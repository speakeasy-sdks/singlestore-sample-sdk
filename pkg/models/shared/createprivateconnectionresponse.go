// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreatePrivateConnectionResponse struct {
	PrivateConnectionID string `json:"privateConnectionID"`
}

func (o *CreatePrivateConnectionResponse) GetPrivateConnectionID() string {
	if o == nil {
		return ""
	}
	return o.PrivateConnectionID
}
