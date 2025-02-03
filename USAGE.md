<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := esvsdkgo.New(
		esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
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