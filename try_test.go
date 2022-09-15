package strive_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/rodriez/strive"
	"github.com/stretchr/testify/assert"
)

func TestTryCatchingPanic(t *testing.T) {
	thrown := strive.Try(func() bool {
		var d, n int

		i := int(d / n)
		fmt.Println(i)

		return false
	}, func(e error) bool {
		assert.EqualError(t, e, "runtime error: integer divide by zero")
		return true
	})

	if !thrown {
		assert.Fail(t, "An exception was expected but nothing was thrown")
	}
}

func TestTryCatchingLiteralPanic(t *testing.T) {
	thrown := strive.Try(func() bool {
		err := fmt.Errorf("an error happened")
		panic(err)
	}, func(e error) bool {

		assert.EqualError(t, e, "an error happened")
		return true
	})

	if !thrown {
		assert.Fail(t, "An exception was expected but nothing was thrown")
	}
}

func TestTryCatchingPanicThrownInOtherFunction(t *testing.T) {
	thrown := strive.Try(func() bool {
		producePanic()
		return false
	}, func(e error) bool {
		assert.EqualError(t, e, "a panic has been thrown")
		return true
	})

	if !thrown {
		assert.Fail(t, "An exception was expected but nothing was thrown")
	}
}

func producePanic() {
	panic(fmt.Errorf("a panic has been thrown"))
}

func TestTryWithCheck(t *testing.T) {
	thrown := strive.Try(func() bool {
		i := strive.Check(strconv.Atoi("XXXXX"))

		fmt.Println(i)

		return false
	}, func(e error) bool {
		assert.EqualError(t, e, "strconv.Atoi: parsing \"XXXXX\": invalid syntax")
		return true
	})

	if !thrown {
		assert.Fail(t, "An exception was expected but nothing was thrown")
	}
}

func TestTryWithCheckFunction(t *testing.T) {
	thrown := strive.Try(func() bool {
		i := strive.CheckFn(func() (int, error) {
			return strconv.Atoi("XXXXX")
		})

		fmt.Println(i)

		return false
	}, func(e error) bool {
		assert.EqualError(t, e, "strconv.Atoi: parsing \"XXXXX\": invalid syntax")
		return true
	})

	if !thrown {
		assert.Fail(t, "An exception was expected but nothing was thrown")
	}
}

func TestTryWithCheckFnFailRecover(t *testing.T) {
	i := strive.Try(func() int {
		return strive.CheckFn(func() (int, error) {
			return strconv.Atoi("XXXX")
		})
	}, func(e error) int {
		return 111
	})

	assert.Equal(t, i, 111)
}

func TestTryWithCheckFnOK(t *testing.T) {
	i := strive.Try(func() int {
		return strive.CheckFn(func() (int, error) {
			return strconv.Atoi("111")
		})
	}, func(e error) int {
		assert.Fail(t, fmt.Sprintf("Unexpected exception thrown %+v", e))
		return 0
	})

	assert.Equal(t, i, 111)
}

func TestStriveCatchingPanic(t *testing.T) {
	strive.Strive(func() {
		var d, n int

		i := int(d / n)
		fmt.Println(i)

		assert.Fail(t, "An exception was expected but nothing was thrown")
	}, func(e error) {
		assert.EqualError(t, e, "runtime error: integer divide by zero")
	})
}

func TestStriveCatchingLiteralPanic(t *testing.T) {
	strive.Strive(func() {
		err := fmt.Errorf("an error happened")
		panic(err)
	}, func(e error) {
		assert.EqualError(t, e, "an error happened")
	})
}
