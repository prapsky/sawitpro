package users

import (
	"github.com/Masterminds/squirrel"

	consts "github.com/prapsky/sawitpro/repository/consts/users"
	"github.com/prapsky/sawitpro/repository/query_builder"
)

type FindByPhoneNumberQueryBuilder struct {
	phoneNumber string
}

func NewFindByPhoneNumberQueryBuilder(phoneNumber string) FindByPhoneNumberQueryBuilder {
	return FindByPhoneNumberQueryBuilder{
		phoneNumber: phoneNumber,
	}
}

func (qb FindByPhoneNumberQueryBuilder) Build() query_builder.QueryBuilder {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	builder := sq.Select(
		consts.IDColumn,
		consts.FullNameColumn,
		consts.PhoneNumberColumn,
		consts.PasswordHashColumn,
		consts.SuccessfulLoginsColumn,
	).From(consts.UsersTable).
		Where(squirrel.Eq{consts.PhoneNumberColumn: qb.phoneNumber}).
		Limit(1)

	sql, params := builder.MustSql()

	return query_builder.NewQueryBuilderResult(sql, params)
}
