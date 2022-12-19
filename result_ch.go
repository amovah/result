package result

type ResultCh[T any] struct {
	res chan T
	err chan error
}

func (r ResultCh[T]) Ok() <-chan T {
	return r.res
}

func (r ResultCh[T]) Err() <-chan error {
	return r.err
}

func (r ResultCh[T]) Push(val Result[T]) {
	if val.IsErr() {
		r.err <- val.Err()
		return
	}

	r.res <- val.Ok()
}

func Channel[T any]() *ResultCh[T] {
	return &ResultCh[T]{
		res: make(chan T),
		err: make(chan error),
	}
}
