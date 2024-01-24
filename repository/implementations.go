package repository

import (
	"context"

	"github.com/prapsky/sawitpro/entity"
	queryBuilder "github.com/prapsky/sawitpro/repository/query_builder/user"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) Insert(ctx context.Context, input entity.User) (uint64, error) {
	builder := queryBuilder.NewInsertQueryBuilder(input).Build()

	id := uint64(0)
	err := r.db.QueryRowContext(ctx, builder.GetQuery(), builder.GetValues()...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
