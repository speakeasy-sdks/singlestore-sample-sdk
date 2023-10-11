# ReplicatedDatabase

Represents information related to a database's replication status


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `DatabaseName`                                                                                  | *string*                                                                                        | :heavy_check_mark:                                                                              | Name of the database                                                                            |
| `DuplicationState`                                                                              | [ReplicatedDatabaseDuplicationState](../../models/shared/replicateddatabaseduplicationstate.md) | :heavy_check_mark:                                                                              | Duplication state of the database                                                               |
| `Region`                                                                                        | *string*                                                                                        | :heavy_check_mark:                                                                              | Name of the region                                                                              |