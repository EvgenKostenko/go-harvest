#go-harvest

[Go](https://golang.org/) client library for [Harvest](http://help.getharvest.com/api/).

## Features

* Authentication (HTTP Basic)


### Authenticate with session cookie

Some actions require an authenticated user.

Here is an example with HTTP Basic authentification.

```go
package main


import (
	"github.com/EvgenKostenko/go-harvest"
	"fmt"
)

func main() {
	harvestClient, err := harvest.NewClient(nil, "https://test.harvestapp.com/")
	if err != nil {
		panic(err)
	}

	res, err := harvestClient.Authentication.Acquire("test@test.com", "foobar")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	if harvestClient.Authentication.Authenticated() {
		fmt.Printf("Done")
	}

}
```