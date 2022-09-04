package strive

// Error throwed with panic
type Exception any

// Function without expected errors
type ExactFunc[T any] func() T

// Function with expected errors
type InexactFunc[T any] func() (T, Exception)

// Function with the logic to handle an exception
type ExceptionHandler func(e Exception)

// Execute a function in a safe context
// In the try context you can panic exceptional errors safely
func Try[T any](fn ExactFunc[T], catch ExceptionHandler) T {
	defer panicListener(catch)

	return fn()
}

// If exception is not nil throw(panic) the exception
// Otherwise return the expected response
// Should be called only un Try context
func Check[T any](r T, e Exception) T {
	if e != nil {
		panic(e)
	} else {
		return r
	}
}

// Run the fn checking for exception
// Should be called only un Try context
func CheckFn[T any](fn InexactFunc[T]) T {
	return Check(fn())
}

func panicListener(catch ExceptionHandler) {
	if e := recover(); e != nil {
		catch(e)
	}
}
