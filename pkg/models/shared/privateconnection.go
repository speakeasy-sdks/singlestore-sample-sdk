// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Status - The status of the private connection
type Status string

const (
	StatusPending Status = "PENDING"
	StatusActive  Status = "ACTIVE"
	StatusDeleted Status = "DELETED"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING":
		fallthrough
	case "ACTIVE":
		fallthrough
	case "DELETED":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// Type - The private connection type
type Type string

const (
	TypeInbound  Type = "INBOUND"
	TypeOutbound Type = "OUTBOUND"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INBOUND":
		fallthrough
	case "OUTBOUND":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// PrivateConnection - Represents information related to a private link connection
type PrivateConnection struct {
	// The timestamp of when the private connection became active
	ActiveAt *string `json:"activeAt,omitempty"`
	// The private connection allow list. This is the account ID for AWS,  subscription ID for Azure, and the project name GCP
	AllowList *string `json:"allowList,omitempty"`
	// The timestamp of when the private connection was created
	CreatedAt *string `json:"createdAt,omitempty"`
	// The timestamp of when the private connection was deleted
	DeletedAt *string `json:"deletedAt,omitempty"`
	// The account ID which must be allowed for outbound connections
	OutboundAllowList *string `json:"outboundAllowList,omitempty"`
	// The ID of the private connection
	PrivateConnectionID string `json:"privateConnectionID"`
	// The name of the private connection service
	ServiceName *string `json:"serviceName,omitempty"`
	// The status of the private connection
	Status *Status `json:"status,omitempty"`
	// The private connection type
	Type *Type `json:"type,omitempty"`
	// The timestamp of when the private connection was last updated
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The ID of the workspace group containing the private connection
	WorkspaceGroupID string `json:"workspaceGroupID"`
	// The ID of the workspace to connect with
	WorkspaceID *string `json:"workspaceID,omitempty"`
}

func (o *PrivateConnection) GetActiveAt() *string {
	if o == nil {
		return nil
	}
	return o.ActiveAt
}

func (o *PrivateConnection) GetAllowList() *string {
	if o == nil {
		return nil
	}
	return o.AllowList
}

func (o *PrivateConnection) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PrivateConnection) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *PrivateConnection) GetOutboundAllowList() *string {
	if o == nil {
		return nil
	}
	return o.OutboundAllowList
}

func (o *PrivateConnection) GetPrivateConnectionID() string {
	if o == nil {
		return ""
	}
	return o.PrivateConnectionID
}

func (o *PrivateConnection) GetServiceName() *string {
	if o == nil {
		return nil
	}
	return o.ServiceName
}

func (o *PrivateConnection) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PrivateConnection) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PrivateConnection) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PrivateConnection) GetWorkspaceGroupID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceGroupID
}

func (o *PrivateConnection) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}
