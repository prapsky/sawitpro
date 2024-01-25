package repository

import (
	"context"
	"database/sql"

	"github.com/prapsky/sawitpro/entity"
	qbLoginAttempts "github.com/prapsky/sawitpro/repository/query_builder/login_attempts"
	qbUsers "github.com/prapsky/sawitpro/repository/query_builder/users"
)

func (r *Repository) Insert(ctx context.Context, input entity.User) (uint64, error) {
	builder := qbUsers.NewInsertQueryBuilder(input).Build()

	id := uint64(0)
	err := r.db.QueryRowContext(ctx, builder.GetQuery(), builder.GetValues()...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	builder := qbUsers.NewFindByPhoneNumberQueryBuilder(phoneNumber).Build()

	row := r.db.QueryRowContext(ctx, builder.GetQuery(), builder.GetValues()...)
	user := &entity.User{}
	err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.PasswordHash,
		&user.SuccessfulLogins,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *Repository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	builder := qbUsers.NewFindByIDQueryBuilder(id).Build()

	row := r.db.QueryRowContext(ctx, builder.GetQuery(), builder.GetValues()...)
	user := &entity.User{}
	err := row.Scan(
		&user.FullName,
		&user.PhoneNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *Repository) UpdateSuccessfulLogins(ctx context.Context, input entity.User) error {
	builder := qbUsers.NewUpdateSuccessfulLoginsQueryBuilder(input).Build()

	_, err := r.db.ExecContext(ctx, builder.GetQuery(), builder.GetValues()...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertLoginAttempts(ctx context.Context, input entity.LoginAttempt) error {
	builder := qbLoginAttempts.NewInsertQueryBuilder(input).Build()

	_, err := r.db.ExecContext(ctx, builder.GetQuery(), builder.GetValues()...)
	if err != nil {
		return err
	}

	return nil
}
