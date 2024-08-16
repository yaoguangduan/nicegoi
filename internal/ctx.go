package nice

type Query map[string]any

func (q Query) Has(key string) bool {
	_, exist := q[key]
	return exist
}

func (q Query) Get(key string) any {
	return q[key]
}
func (q Query) GetOr(key string, def any) any {
	v, exist := q[key]
	if exist {
		return v
	}
	return def
}

type GoiContext interface {
	Query() Query
}

type pageCtx struct {
	q Query
}

func (p *pageCtx) Query() Query {
	return p.q
}
