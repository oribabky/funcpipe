package funcpipe

func NewPipe[T any](list []T) Pipe[T] {
	pipe := pipe[T]{list}
	return pipe
}

type Pipe[T any] interface {
	Filter(fn func(T) bool) pipe[T]
	Collect() []T
}
type pipe[T any] struct {
	list []T
}

func (p pipe[T]) Filter(fn func(T) bool) pipe[T] {
	var res []T
	for _, e := range p.list {
		if fn(e) == true {
			res = append(res, e)
		}
	}
	p.list = res
	return p
}

func (p pipe[T]) Collect() []T {
	return p.list
}
