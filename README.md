# strive
Helper for error propagation a handling in GO


## Installation

```bash
go get github.com/rodriez/strive
```

## Usage Try with Check

```go
package main

import (
    "fmt"
	"strconv"
    
    "github.com/rodriez/strive"
)


func main() {
	strive.Try(func() int {
        i := strive.Check(strconv.Atoi("XXXXX"))

        return i
    }, func(e strive.Exception) {
        err := e.(error)

        fmt.Printn(err)
    })
}

```

## Usage Try with CheckFn

```go
package main

import (
    "fmt"
	"strconv"

    "github.com/rodriez/strive"
)


func main() {
    stri := "XXXXX"

	strive.Try(func() any {
        i := strive.CheckFn(func() (int, strive.Exception) {
            return strconv.Atoi(stri)
        })

        fmt.Println(i)

        return nil
    }, func(e strive.Exception) {
        err := e.(error)

        fmt.Printn(err)
    })
}

```

## Usage Try with literal panic

```go
package main

import (
    "fmt"

    "github.com/rodriez/strive"
)


func main() {
    strive.Try(func() any {
        err := fmt.Errorf("not implemented")
		panic(err)
    }, func(e strive.Exception) {
        err := e.(error)

        fmt.Printn(err)
    })
}

```