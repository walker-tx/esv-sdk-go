# PassageResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Query`                                                            | **string*                                                          | :heavy_minus_sign:                                                 | The passage reference that was requested                           |
| `Canonical`                                                        | **string*                                                          | :heavy_minus_sign:                                                 | The canonical version of the passage reference                     |
| `Parsed`                                                           | [][]*int64*                                                        | :heavy_minus_sign:                                                 | Array of parsed passage references                                 |
| `PassageMeta`                                                      | [][components.PassageMeta](../../models/components/passagemeta.md) | :heavy_minus_sign:                                                 | N/A                                                                |
| `Passages`                                                         | []*string*                                                         | :heavy_minus_sign:                                                 | Array of formatted passage text                                    |