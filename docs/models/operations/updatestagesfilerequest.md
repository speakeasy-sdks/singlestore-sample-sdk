# UpdateStagesFileRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `StagesPatch`                                             | [*shared.StagesPatch](../../models/shared/stagespatch.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `Path`                                                    | *string*                                                  | :heavy_check_mark:                                        | Path in Stages to modify                                  |
| `WorkspaceGroupID`                                        | *string*                                                  | :heavy_check_mark:                                        | ID of the Stages-enabled workspace group                  |