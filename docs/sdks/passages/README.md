# Passages
(*Passages*)

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

### Example Usage

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

    res, err := s.Passages.Search(ctx, "<value>", nil, nil)
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
| `query`                                                  | *string*                                                 | :heavy_check_mark:                                       | The text to search for                                   |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Number of results to return per page                     |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number to return                                    |
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

### Example Usage

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

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `query`                                                   | *string*                                                  | :heavy_check_mark:                                        | Bible passage reference (e.g., "John 3:16" or "43011016") |
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

### Example Usage

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