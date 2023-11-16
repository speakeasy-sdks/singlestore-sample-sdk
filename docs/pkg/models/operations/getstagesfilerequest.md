# GetStagesFileRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Path`                                            | *string*                                          | :heavy_check_mark:                                | Path in Stages to a file or folder                |
| `WorkspaceGroupID`                                | *string*                                          | :heavy_check_mark:                                | ID of the Stages-enabled workspace group          |
| `Metadata`                                        | **bool*                                           | :heavy_minus_sign:                                | If enabled, the API request returns only metadata |