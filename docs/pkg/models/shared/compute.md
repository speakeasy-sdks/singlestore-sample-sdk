# Compute

Represents information related to a workspace group's latest storage DR operation


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CompletedAttachments`                                                | **int64*                                                              | :heavy_minus_sign:                                                    | The number of database attachments that have been setup               |
| `CompletedWorkspaces`                                                 | **int64*                                                              | :heavy_minus_sign:                                                    | The number of workspaces that have been setup                         |
| `StorageDRState`                                                      | [shared.StorageDRState](../../../pkg/models/shared/storagedrstate.md) | :heavy_check_mark:                                                    | Status of Storage DR operation                                        |
| `StorageDRType`                                                       | [shared.StorageDRType](../../../pkg/models/shared/storagedrtype.md)   | :heavy_check_mark:                                                    | Name of Storage DR operation                                          |
| `TotalAttachments`                                                    | **int64*                                                              | :heavy_minus_sign:                                                    | The total number of database attachments to setup                     |
| `TotalWorkspaces`                                                     | **int64*                                                              | :heavy_minus_sign:                                                    | The total number of workspaces to setup                               |