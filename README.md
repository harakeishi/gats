# gats
Returns the value declared in go's basic type as a `string`.  
gats is named as an abbreviation of `go any to string`.

## Importing
```go
import (
    "github.com/harakeishi/gats"
)
```

## Documentation
Visit the docs on [GoDoc](https://pkg.go.dev/github.com/harakeishi/gats)

## usage

```go
package main

import (
	"fmt"
	"github.com/harakeishi/gats"
)

func main() {
	value := []byte("test")
	result, err := gats.ToString(value)
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(result)
    // Output: test
}
```

## License
Copyright (c) 2023 harakeishi

Licensed under [MIT](LICENSE)

