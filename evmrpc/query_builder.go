package evmrpc

type QueryBuilder struct {
	conditions []string
}

func NewHeadQueryBuilder() *QueryBuilder { _ = "STUB: not implemented"; return nil }

func NewBlockQueryBuilder() *QueryBuilder { _ = "STUB: not implemented"; return nil }

func (q *QueryBuilder) Build() string { _ = "STUB: not implemented"; return "" }
