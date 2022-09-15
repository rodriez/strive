# strive
Helper for error propagation a handling in GO


## Installation

```bash
go get github.com/rodriez/strive
```

## Usage Strive with Check
[Try it](https://go.dev/play/p/yqfYOAOSUyj)

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
    }, func(err error) {
        fmt.Println(err)
    })
}

```

## Usage Try with Check
[Try it](https://go.dev/play/p/lvQS5ESdgI3)

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
[Try it](https://go.dev/play/p/lJ142OEN-BO)

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
        i := strive.CheckFn(func() (int, error) {
            return strconv.Atoi(stri)
        })

        fmt.Println(i)
    }, func(err error) {
        fmt.Println(err)
    })
}

```

## Usage Strive with literal panic
[Try it](https://go.dev/play/p/Ulphqmr--Qk)

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
[Try it](https://go.dev/play/p/1Odvi4zPhaF)

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
    }, func(err error) int {
        fmt.Println(err)
        return -1
    })

    fmt.Println(i)
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

[Hey 👋 buy me a beer! ](https://www.buymeacoffee.com/rodriez)