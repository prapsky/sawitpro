package user

import (
	"github.com/Masterminds/squirrel"

	"github.com/prapsky/sawitpro/entity"
	consts "github.com/prapsky/sawitpro/repository/consts/user"
	"github.com/prapsky/sawitpro/repository/query_builder"
)

type InsertQueryBuilder struct {
	user entity.User
}

func NewInsertQueryBuilder(user entity.User) InsertQueryBuilder {
	return InsertQueryBuilder{
		user: user,
	}
}

func (qb InsertQueryBuilder) Build() query_builder.QueryBuilder {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	builder := sq.Insert(consts.UsersTable)

	insertClause := map[string]interface{}{}
	insertClause[consts.PhoneNumberColumn] = qb.user.PhoneNumber
	insertClause[consts.FullNameColumn] = qb.user.FullName
	insertClause[consts.PasswordHashColumn] = qb.user.PasswordHash
	insertClause[consts.CreatedAtColumn] = qb.user.CreatedAt

	builder = builder.SetMap(insertClause).Suffix("RETURNING id")
	sql, params := builder.MustSql()

	return query_builder.NewQueryBuilderResult(sql, params)
}
