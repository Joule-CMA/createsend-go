createsend-go
=============

createsend-go is a [Go](http://golang.org) library for accessing the [Campaign
Monitor API](http://www.campaignmonitor.com/api/).

This is an unofficial library that is not affiliated with [Campaign
Monitor](http://www.campaignmonitor.com). Official libraries are available at
[github.com/campaignmonitor](https://github.com/campaignmonitor).

This is branched of the work of sourcegraph but now has been archived.

**Documentation:** <https://github.com/sourcegraph/createsend-go>

[![Build Status](https://travis-ci.org/sourcegraph/createsend-go.png?branch=master)](https://travis-ci.org/sourcegraph/createsend-go)

Example usage
-------------

```go
package createsend_test

import (
	"fmt"
	"github.com/sourcegraph/createsend-go/createsend"
	"net/http"
	"os"
)

func Example() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "You must set your Campaign Monitor API key in the API_KEY environment variable to run example_test.go. (Skipping.)\n")
		os.Exit(0)
	}

	authClient := &http.Client{
		Transport: &createsend.APIKeyAuthTransport{APIKey: apiKey},
	}

	c := createsend.NewAPIClient(authClient)
	clients, err := c.ListClients()
	if err != nil {
		fmt.Printf("Error listing clients: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d clients.\n", len(clients))
	for _, client := range clients {
		fmt.Printf(" - %s (ID: [%d-char ID])\n", client.Name, len(client.ClientID))
	}

	// This output will be different for each account.

	// output:
	// Found 1 clients.
	//  - Sourcegraph (ID: [32-char ID])
}
```

See the [createsend/example_test.go](./createsend/example_test.go) file for the full source.


Running the tests
-----------------

To run the tests:

```
go test ./createsend
```

To run the included example (in `createsend/example_test.go`), set your Campaign
Monitor API key in the `API_KEY` environment variable (available in Account
Settings).

```
API_KEY=your-api-key go test ./createsend
```

Acknowledgements
----------------

The library's architecture and testing code are adapted from
[go-github](https://github.com/google/go-github), created by [Will
Norris](https://github.com/willnorris).
