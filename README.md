# ApiDQ API GoLang Client

---

This is the GoLang ApiDQ API client. This library allows using of the actual API version. You can find more info in
the [documentation](https://docs.apidq.io).

## Installation

Follow those steps to install the library:

1. Import the library our code:

```shell
go get github.com/nikitaksv/apidq-client-go
```

## Usage

The client is separated into several resource groups, all of which are accessible through the Client's public
properties. You can call API methods from those groups like this:

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nikitaksv/apidq-client-go"
	"github.com/nikitaksv/apidq-client-go/dto/address"
)

func main() {
	client, err := apidq.NewClient(http.DefaultClient, apidq.BaseURL)
	if err != nil {
		panic(err)
	}

	client.WithAuth("your token here")
	// Or set a individual ApiKey for a specific service
	// client.WithAuthService("you_token_here", "address")

	cleanRsp, _, err := client.Address.Clean(context.TODO(), &address.CleanRequest{
		Query:       "москва спартаковская 10с12",
		CountryCode: "RU",
	})
	if err != nil {
		panic(err)
    }

	fmt.Println(cleanRsp.Address.Address) // -- print: г Москва, пл Спартаковская
}
```

To handle errors you must use one type of errors:

* `apidq.ErrorResponse` for the api service error.
