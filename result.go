package result

type Result[T any] interface {
	IsOk() bool
	IsErr() bool
	Ok() T
	Err() error
}

type OkOf[T any] struct {
	value T
}

func (OkOf[T]) IsOk() bool {
	return true
}

func (OkOf[T]) IsErr() bool {
	return false
}

func (r OkOf[T]) Ok() T {
	return r.value
}

func (OkOf[T]) Err() error {
	panic("cannot call Err on Ok")
}

func Ok[T any](val T) Result[T] {
	return &OkOf[T]{
		value: val,
	}
}

type ErrorOf[T any] struct {
	err error
}

func (ErrorOf[T]) IsOk() bool {
	return false
}

func (ErrorOf[T]) IsErr() bool {
	return true
}

func (ErrorOf[T]) Ok() T {
	panic("cannot call Ok on Err")
}

func (e ErrorOf[T]) Err() error {
	return e.err
}

func Error[T any](err error) Result[T] {
	return &ErrorOf[T]{
		err,
	}
}
