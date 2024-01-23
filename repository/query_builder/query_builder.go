package query_builder

type QueryBuilder interface {
	GetQuery() string
	GetValues() []interface{}
}

type QueryBuilderResult struct {
	query  string
	values []interface{}
}

func NewQueryBuilderResult(query string, values []interface{}) QueryBuilderResult {
	return QueryBuilderResult{
		query:  query,
		values: values,
	}
}

func (b QueryBuilderResult) GetQuery() string {
	return b.query
}

func (b QueryBuilderResult) GetValues() []interface{} {
	return b.values
}
