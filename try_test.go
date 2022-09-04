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
	strive.Try(func() int {
		var d, n int

		i := int(d / n)

		return i
	}, func(e strive.Exception) {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "runtime error: integer divide by zero")
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryCatchingLiteralPanic(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() int {
		err := fmt.Errorf("an error happened")
		panic(err)
	}, func(e strive.Exception) {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "an error happened")
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryCatchingPanicThrowedInOtherFunction(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() int {
		producePanic()

		return 1
	}, func(e strive.Exception) {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "a panic has been throwed")
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
	strive.Try(func() int {
		return strive.Check(strconv.Atoi("XXXXX"))
	}, func(e strive.Exception) {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "strconv.Atoi: parsing \"XXXXX\": invalid syntax")
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryWithCheckFunction(t *testing.T) {
	errorWasThrown := false
	strive.Try(func() int {
		return strive.CheckFn(func() (int, strive.Exception) {
			return strconv.Atoi("XXXXX")
		})
	}, func(e strive.Exception) {
		errorWasThrown = true
		err := e.(error)

		assert.EqualError(t, err, "strconv.Atoi: parsing \"XXXXX\": invalid syntax")
	})

	if !errorWasThrown {
		assert.Fail(t, "An exception was expected but nothing was throwed")
	}
}

func TestTryWithCheckOK(t *testing.T) {
	i := strive.Try(func() int {
		return strive.CheckFn(func() (int, strive.Exception) {
			return strconv.Atoi("111")
		})
	}, func(e strive.Exception) {
		assert.Fail(t, fmt.Sprintf("Unexpected exception throwed %+v", e))
	})

	assert.Equal(t, i, 111)
}
