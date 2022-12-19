package result

type ResultChannel[T any] struct {
	res chan T
	err chan error
}

func (r ResultChannel[T]) Ok() <-chan T {
	return r.res
}

func (r ResultChannel[T]) Err() <-chan error {
	return r.err
}

func (r ResultChannel[T]) Push(val Result[T]) {
	if val.IsErr() {
		r.err <- val.Err()
		return
	}

	r.res <- val.Ok()
}

func Channel[T any]() *ResultChannel[T] {
	return &ResultChannel[T]{
		res: make(chan T),
		err: make(chan error),
	}
}
