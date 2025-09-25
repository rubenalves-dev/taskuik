package common

type Result[T any] struct {
	Err        error
	StatusCode int
	Data       T
}

type StatusErr struct {
	Err        error
	StatusCode int
}

type ListResult[T any] struct {
	Items []T
	Total int
	StatusErr
}

func OkListResult[T any](items []T, total, statusCode int) ListResult[T] {
	return ListResult[T]{
		Items:     items,
		Total:     total,
		StatusErr: StatusErr{StatusCode: statusCode},
	}
}

func ErrListResult[T any](err error, statusCode int) ListResult[T] {
	return ListResult[T]{
		StatusErr: StatusErr{
			Err:        err,
			StatusCode: statusCode,
		},
	}
}

func OkResult[T any](data T, statusCode int) Result[T] {
	return Result[T]{
		StatusCode: statusCode,
		Data:       data,
	}
}

func ErrResult[T any](err error, statusCode int) Result[T] {
	return Result[T]{
		Err:        err,
		StatusCode: statusCode,
	}
}

func OkStatus(statusCode int) StatusErr {
	return StatusErr{
		StatusCode: statusCode,
	}
}

func ErrStatus(err error, statusCode int) StatusErr {
	return StatusErr{
		Err:        err,
		StatusCode: statusCode,
	}
}
