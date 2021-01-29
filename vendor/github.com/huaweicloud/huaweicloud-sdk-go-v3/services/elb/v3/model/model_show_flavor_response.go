/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowFlavorResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	Flavor         *Flavor `json:"flavor,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFlavorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFlavorResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorResponse", string(data)}, " ")
}
