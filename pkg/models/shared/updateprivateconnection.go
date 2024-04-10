// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdatePrivateConnection - Represents the information specfied when updating a private connection
type UpdatePrivateConnection struct {
	// The private connection allow list
	AllowList *string `json:"allowList,omitempty"`
}

func (o *UpdatePrivateConnection) GetAllowList() *string {
	if o == nil {
		return nil
	}
	return o.AllowList
}
