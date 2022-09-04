package strive_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/rodriez/strive"
	"github.com/stretchr/testify/assert"
)

func TestTryCatchingPanic(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() any {
		var d, n int

		i := int(d / n)

		return i
	}, func(e strive.Exception) any {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "runtime error: integer divide by zero")
		return nil
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryCatchingLiteralPanic(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() any {
		err := fmt.Errorf("an error happened")
		panic(err)
	}, func(e strive.Exception) any {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "an error happened")
		return nil
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryCatchingPanicThrowedInOtherFunction(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() any {
		producePanic()

		return nil
	}, func(e strive.Exception) any {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "a panic has been throwed")
		return nil
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func producePanic() {
	panic(fmt.Errorf("a panic has been throwed"))
}

func TestTryWithCheck(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() any {
		return strive.Check(strconv.Atoi("XXXXX"))
	}, func(e strive.Exception) any {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "strconv.Atoi: parsing \"XXXXX\": invalid syntax")
		return nil
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryWithCheckFunction(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() any {
		return strive.CheckFn(func() (int, strive.Exception) {
			return strconv.Atoi("XXXXX")
		})
	}, func(e strive.Exception) any {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "strconv.Atoi: parsing \"XXXXX\": invalid syntax")
		return nil
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryWithCheckFnFailRecover(t *testing.T) {
	i := strive.Try(func() int {
		return strive.CheckFn(func() (int, strive.Exception) {
			return strconv.Atoi("XXXX")
		})
	}, func(e strive.Exception) int {
		return 111
	})

	assert.Equal(t, i, 111)
}

func TestTryWithCheckFnOK(t *testing.T) {
	i := strive.Try(func() int {
		return strive.CheckFn(func() (int, strive.Exception) {
			return strconv.Atoi("111")
		})
	}, func(e strive.Exception) int {
		assert.Fail(t, fmt.Sprintf("Unexpected exception throwed %+v", e))
		return 0
	})

	assert.Equal(t, i, 111)
}
