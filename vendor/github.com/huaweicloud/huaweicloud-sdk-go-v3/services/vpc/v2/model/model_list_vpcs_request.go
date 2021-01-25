/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVpcsRequest struct {
	Limit               *int32  `json:"limit,omitempty"`
	Marker              *string `json:"marker,omitempty"`
	Id                  *string `json:"id,omitempty"`
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListVpcsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsRequest", string(data)}, " ")
}
