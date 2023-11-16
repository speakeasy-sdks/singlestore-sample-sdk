# UpdateStagesFileRequest


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Path`                                                           | *string*                                                         | :heavy_check_mark:                                               | Path in Stages to modify                                         |
| `WorkspaceGroupID`                                               | *string*                                                         | :heavy_check_mark:                                               | ID of the Stages-enabled workspace group                         |
| `StagesPatch`                                                    | [*shared.StagesPatch](../../../pkg/models/shared/stagespatch.md) | :heavy_minus_sign:                                               | N/A                                                              |