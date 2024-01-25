package users

import (
	"github.com/Masterminds/squirrel"

	consts "github.com/prapsky/sawitpro/repository/consts/users"
	"github.com/prapsky/sawitpro/repository/query_builder"
)

type FindByIDQueryBuilder struct {
	id uint64
}

func NewFindByIDQueryBuilder(id uint64) FindByIDQueryBuilder {
	return FindByIDQueryBuilder{
		id: id,
	}
}

func (qb FindByIDQueryBuilder) Build() query_builder.QueryBuilder {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	builder := sq.Select(
		consts.FullNameColumn,
		consts.PhoneNumberColumn,
	).From(consts.UsersTable).
		Where(squirrel.Eq{consts.IDColumn: qb.id}).
		Limit(1)

	sql, params := builder.MustSql()

	return query_builder.NewQueryBuilderResult(sql, params)
}
