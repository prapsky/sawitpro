package users

import (
	"github.com/Masterminds/squirrel"

	"github.com/prapsky/sawitpro/entity"
	consts "github.com/prapsky/sawitpro/repository/consts/users"
	"github.com/prapsky/sawitpro/repository/query_builder"
)

type UpdateSuccessfulLoginsQueryBuilder struct {
	user entity.User
}

func NewUpdateSuccessfulLoginsQueryBuilder(user entity.User) UpdateSuccessfulLoginsQueryBuilder {
	return UpdateSuccessfulLoginsQueryBuilder{
		user: user,
	}
}

func (qb UpdateSuccessfulLoginsQueryBuilder) Build() query_builder.QueryBuilder {
	user := qb.user
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	builder := sq.Update(consts.UsersTable).
		Set(consts.SuccessfulLoginsColumn, user.SuccessfulLogins).
		Set(consts.LastLoginAtColumn, user.LastLoginAt).
		Where(squirrel.Eq{consts.IDColumn: user.ID})

	sql, params := builder.MustSql()

	return query_builder.NewQueryBuilderResult(sql, params)
}
