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
	strive.Try(func() any {
        i := strive.Check(strconv.Atoi("XXXXX"))

        return i
    },
    //Catch
    func(e strive.Exception) any {
        err := e.(error)

        fmt.Println(err)
        return nil
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
    },
    //Catch
    func(e strive.Exception) any {
        err := e.(error)

        fmt.Println(err)
        return nil
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
    },
    //Catch
    func(e strive.Exception) any {
        err := e.(error)

        fmt.Println(err)
        return nil
    })
}

```

## Usage Try in success case

```go
package main

import (
    "fmt"
	"strconv"

    "github.com/rodriez/strive"
)


func main() {
    stri := "12345"

	i := strive.Try(func() int {
        return strive.Check(strconv.Atoi(stri))
    },
    //Catch
    func(e strive.Exception) int {
        err := e.(error)

        fmt.Println(err)
        return 0
    })

    fmt.Println(i)
}

```