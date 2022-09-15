# strive
Helper for error propagation a handling in GO


## Installation

```bash
go get github.com/rodriez/strive
```

## Usage Strive with Check

```go
package main

import (
    "fmt"
	"strconv"
    
    "github.com/rodriez/strive"
)


func main() {
	strive.Strive(func() {
        i := strive.Check(strconv.Atoi("XXXXX"))
        fmt.Println(i)
    }, func(err error) any {
        fmt.Println(err)
    })
}

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
	i := strive.Try(func() int {
        return strive.Check(strconv.Atoi("XXXXX"))    
    }, func(err error) int {
        fmt.Println(err)
        return 0
    })

    fmt.Println(i)
}

```

## Usage Strive with CheckFn

```go
package main

import (
    "fmt"
	"strconv"

    "github.com/rodriez/strive"
)


func main() {
    stri := "XXXXX"

	strive.Strive(func() {
        i := strive.CheckFn(func() (int, strive.Exception) {
            return strconv.Atoi(stri)
        })

        fmt.Println(i)
    }, func(err error) {
        fmt.Println(err)
    })
}

```

## Usage Strive with literal panic

```go
package main

import (
    "fmt"

    "github.com/rodriez/strive"
)


func main() {
    strive.Strive(func() {
        err := fmt.Errorf("not implemented")
		strive.CheckError(err)
    }, func(err error) {    
        fmt.Println(err)
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
    }, func(e error) int {
        fmt.Println(err)
        return -1
    })

    fmt.Println(i)
}

```