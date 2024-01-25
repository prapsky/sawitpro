package login_attempts

import (
	"github.com/Masterminds/squirrel"

	"github.com/prapsky/sawitpro/entity"
	consts "github.com/prapsky/sawitpro/repository/consts/login_attempts"
	"github.com/prapsky/sawitpro/repository/query_builder"
)

type InsertQueryBuilder struct {
	loginAttempt entity.LoginAttempt
}

func NewInsertQueryBuilder(loginAttempt entity.LoginAttempt) InsertQueryBuilder {
	return InsertQueryBuilder{
		loginAttempt: loginAttempt,
	}
}

func (qb InsertQueryBuilder) Build() query_builder.QueryBuilder {
	loginAttempt := qb.loginAttempt
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	builder := sq.Insert(consts.LoginAttemptsTable)

	insertClause := map[string]interface{}{}
	insertClause[consts.UserIDColumn] = loginAttempt.UserID
	insertClause[consts.SuccessColumn] = loginAttempt.Success
	insertClause[consts.AttemptedAtColumn] = loginAttempt.AttemptedAt

	builder = builder.SetMap(insertClause)
	sql, params := builder.MustSql()

	return query_builder.NewQueryBuilderResult(sql, params)
}
