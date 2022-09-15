package strive

// Function without expected errors
type ExactFn[T any] func() T

// Function with expected errors
type InexactFn[T any] func() (T, error)

// Function with the logic to convert an exception
type ExceptionMapper[T any] func(e error) T

// Function with the logic to handle an exception
type ExceptionHandler func(e error)

// Execute a function in a safe context
// In the try context you can panic exceptional errors safely
func Try[T any](fn ExactFn[T], catch ExceptionMapper[T]) (resp T) {
	defer panicMapper(catch, &resp)

	return fn()
}

// Execute a function in a safe context
// In the Strive context you can panic exceptional errors safely
func Strive(cmd func(), catch ExceptionHandler) {
	defer panicHandler(catch)

	cmd()
}

// If exception is not nil throw(panic) the exception
// Otherwise return the expected response
// Should be called only un Try context
func Check[T any](r T, e error) T {
	CheckError(e)

	return r
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// Run the fn checking for exception
// Should be called only un Try context
func CheckFn[T any](fn InexactFn[T]) T {
	return Check(fn())
}

func panicMapper[T any](mapFn ExceptionMapper[T], output *T) {
	if e := recover(); e != nil {
		*output = mapFn(e.(error))
	}
}

func panicHandler(catch ExceptionHandler) {
	if e := recover(); e != nil {
		catch(e.(error))
	}
}
