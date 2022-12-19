package result

type ResultChannel[T any] interface {
	Ok() <-chan T
	Err() <-chan error
	Push(Result[T])
	Close()
}

type resultChannel[T any] struct {
	res chan T
	err chan error
}

func (r resultChannel[T]) Ok() <-chan T {
	return r.res
}

func (r resultChannel[T]) Err() <-chan error {
	return r.err
}

func (r resultChannel[T]) Push(val Result[T]) {
	if val.IsErr() {
		r.err <- val.Err()
		return
	}

	r.res <- val.Ok()
}

func (r resultChannel[T]) Close() {
	close(r.err)
	close(r.res)
}

func Channel[T any]() ResultChannel[T] {
	return &resultChannel[T]{
		res: make(chan T),
		err: make(chan error),
	}
}
