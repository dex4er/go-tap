# Run Go tests with TAP

## Install

```sh
go get github.com/dex4er/go-tap
```

## t/simple/simple.go

```go
package main

import "github.com/dex4er/go-tap"

func main() {
    tap.Ok(true, "Ok")
    tap.Is("Aaa", "Aaa", "Is")
    tap.Is(123, 123, "Is")
	tap.DoneTesting()
}
```
