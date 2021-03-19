/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateGroupRequest struct {
	Body *KeystoneCreateGroupRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateGroupRequest", string(data)}, " ")
}
