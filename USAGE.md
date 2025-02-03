<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	esvsdk "github.com/walker-tx/esv-sdk"
	"github.com/walker-tx/esv-sdk/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := esvsdk.New(
		esvsdk.WithSecurity(os.Getenv("ESV_API_KEY")),
	)

	res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
		Query: "John 1:1",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PassageResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->