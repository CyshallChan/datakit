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

//
type UpdateCredentialResult struct {
	// IAM用户ID。
	UserId string `json:"user_id"`
	// 修改的AK。
	Access string `json:"access"`
	// 访问密钥状态。
	Status string `json:"status"`
	// 访问密钥创建时间。
	CreateTime string `json:"create_time"`
	// 访问密钥描述信息。
	Description string `json:"description"`
}

func (o UpdateCredentialResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCredentialResult struct{}"
	}

	return strings.Join([]string{"UpdateCredentialResult", string(data)}, " ")
}
