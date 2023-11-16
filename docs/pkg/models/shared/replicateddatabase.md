# ReplicatedDatabase

Represents information related to a database's replication status


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `DatabaseName`                                                            | *string*                                                                  | :heavy_check_mark:                                                        | Name of the database                                                      |
| `DuplicationState`                                                        | [shared.DuplicationState](../../../pkg/models/shared/duplicationstate.md) | :heavy_check_mark:                                                        | Duplication state of the database                                         |
| `Region`                                                                  | *string*                                                                  | :heavy_check_mark:                                                        | Name of the region                                                        |