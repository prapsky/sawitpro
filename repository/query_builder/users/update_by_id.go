package users

import (
	"github.com/Masterminds/squirrel"

	"github.com/prapsky/sawitpro/entity"
	consts "github.com/prapsky/sawitpro/repository/consts/users"
	"github.com/prapsky/sawitpro/repository/query_builder"
)

type UpdateByIDQueryBuilder struct {
	user entity.User
}

func NewUpdateByIDQueryBuilder(user entity.User) UpdateByIDQueryBuilder {
	return UpdateByIDQueryBuilder{
		user: user,
	}
}

func (qb UpdateByIDQueryBuilder) Build() query_builder.QueryBuilder {
	user := qb.user
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	builder := sq.Update(consts.UsersTable).
		Set(consts.UpdatedAtColumn, user.UpdatedAt)

	if user.FullName != "" {
		builder = builder.Set(consts.FullNameColumn, user.FullName)
	}

	if user.PhoneNumber != "" {
		builder = builder.Set(consts.PhoneNumberColumn, user.PhoneNumber)
	}

	builder = builder.Where(squirrel.Eq{consts.IDColumn: user.ID})

	sql, params := builder.MustSql()

	return query_builder.NewQueryBuilderResult(sql, params)
}
