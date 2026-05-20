# Passages

## Overview

### Available Operations

* [GetHTML](#gethtml) - Get Bible passage HTML
* [Search](#search) - Search Bible passages
* [GetAudio](#getaudio) - Get Bible passage audio
* [GetText](#gettext) - Get Bible passage text

## GetHTML

Returns Bible passage text with HTML formatting

Esv.org API Docs for `/v3/passages/html`
<https://api.esv.org/docs/passage-html/>

### Example Usage: ChapterRange

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="ChapterRange" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
        Query: "Genesis 1-3",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: CompactNotation

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="CompactNotation" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
        Query: "jn11.35",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: DigitalRange

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="DigitalRange" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
        Query: "01001001-01011032",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: MultiReference

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="MultiReference" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
        Query: "John1.1;Genesis1.1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: NumericalEncoding

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="NumericalEncoding" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
        Query: "43011035",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: StandardReference

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="StandardReference" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
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
### Example Usage: StructuredDigital

<!-- UsageSnippet language="go" operationID="getPassageHtml" method="get" path="/passage/html/" example="StructuredDigital" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetHTML(ctx, operations.GetPassageHTMLRequest{
        Query: "19001001-19001006,19003001-19003008",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPassageHTMLRequest](../../models/operations/getpassagehtmlrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetPassageHTMLResponse](../../models/operations/getpassagehtmlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Search

Returns search results for Bible passages based on the provided query

Esv.org API Docs for `/v3/passage/search`
<https://api.esv.org/docs/passage-search/>

### Example Usage

<!-- UsageSnippet language="go" operationID="searchPassages" method="get" path="/passage/search/" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.Search(ctx, "<value>", esvsdkgo.Pointer[int64](20), esvsdkgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `query`                                                  | `string`                                                 | :heavy_check_mark:                                       | The text to search for                                   |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Number of results to return per page                     |
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | Page number to return                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SearchPassagesResponse](../../models/operations/searchpassagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAudio

Returns audio file for Bible passages based on the provided query

Esv.org API Docs for `/v3/passage/audio`
<https://api.esv.org/docs/passage-audio/>

### Example Usage: ChapterRange

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="ChapterRange" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "Genesis 1-3")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```
### Example Usage: CompactNotation

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="CompactNotation" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "jn11.35")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```
### Example Usage: DigitalRange

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="DigitalRange" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "01001001-01011032")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```
### Example Usage: MultiReference

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="MultiReference" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "John1.1;Genesis1.1")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```
### Example Usage: NumericalEncoding

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="NumericalEncoding" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "43011035")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```
### Example Usage: StandardReference

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="StandardReference" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "John 1:1")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```
### Example Usage: StructuredDigital

<!-- UsageSnippet language="go" operationID="getPassageAudio" method="get" path="/passage/audio/" example="StructuredDigital" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetAudio(ctx, "19001001-19001006,19003001-19003008")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `query`                                                   | `string`                                                  | :heavy_check_mark:                                        | Bible passage reference (e.g., "John 3:16" or "43011016") |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.GetPassageAudioResponse](../../models/operations/getpassageaudioresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetText

Returns Bible passage text based on the provided query parameters

Esv.org API Docs for `/v3/passages/text`
<https://api.esv.org/docs/passage-text/>

### Example Usage: ChapterRange

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="ChapterRange" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
        Query: "Genesis 1-3",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: CompactNotation

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="CompactNotation" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
        Query: "jn11.35",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: DigitalRange

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="DigitalRange" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
        Query: "01001001-01011032",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: MultiReference

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="MultiReference" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
        Query: "John1.1;Genesis1.1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: NumericalEncoding

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="NumericalEncoding" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
        Query: "43011035",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```
### Example Usage: StandardReference

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="StandardReference" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
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
### Example Usage: StructuredDigital

<!-- UsageSnippet language="go" operationID="getPassageText" method="get" path="/passage/text/" example="StructuredDigital" -->
```go
package main

import(
	"context"
	"os"
	esvsdkgo "github.com/walker-tx/esv-sdk-go"
	"github.com/walker-tx/esv-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := esvsdkgo.New(
        esvsdkgo.WithSecurity(os.Getenv("ESV_API_KEY")),
    )

    res, err := s.Passages.GetText(ctx, operations.GetPassageTextRequest{
        Query: "19001001-19001006,19003001-19003008",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PassageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetPassageTextRequest](../../models/operations/getpassagetextrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetPassageTextResponse](../../models/operations/getpassagetextresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |