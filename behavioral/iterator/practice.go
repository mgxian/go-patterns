package iterator

type paginator interface {
	first() []interface{}
	prev() []interface{}
	next() []interface{}
	last() []interface{}
}

type simplePaginator struct {
	pageSize    int
	currentPage int
	totalPage   int
	data        []interface{}
}

func newSimplePaginator(d []interface{}, pageSize int) *simplePaginator {
	totalPage := len(d) / pageSize
	if len(d)%pageSize != 0 {
		totalPage++
	}

	return &simplePaginator{
		pageSize:    pageSize,
		currentPage: 0,
		totalPage:   totalPage,
		data:        d,
	}
}

func (s *simplePaginator) getPage(n int) []interface{} {
	start := s.pageSize * (n - 1)
	if n < 1 {
		start = 0
	}

	end := start + s.pageSize
	if end > len(s.data) {
		end = len(s.data)
	}

	s.currentPage = n
	return s.data[start:end]
}

func (s *simplePaginator) first() []interface{} {
	return s.getPage(0)
}

func (s *simplePaginator) prev() []interface{} {
	return s.getPage(s.currentPage - 1)
}

func (s *simplePaginator) next() []interface{} {
	return s.getPage(s.currentPage + 1)
}

func (s *simplePaginator) last() []interface{} {
	return s.getPage(s.totalPage)
}
