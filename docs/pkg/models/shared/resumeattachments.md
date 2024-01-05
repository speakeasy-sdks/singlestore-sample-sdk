# ResumeAttachments

Represents information related to database attachments


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Attachment`                                                  | [shared.Attachment](../../../pkg/models/shared/attachment.md) | :heavy_check_mark:                                            | The type of attachment                                        |
| `Database`                                                    | *string*                                                      | :heavy_check_mark:                                            | Name of the database                                          |
| `Error`                                                       | **string*                                                     | :heavy_minus_sign:                                            | The error if the attachment was not successful                |
| `Success`                                                     | *bool*                                                        | :heavy_check_mark:                                            | Whether the attachment was successful or not                  |