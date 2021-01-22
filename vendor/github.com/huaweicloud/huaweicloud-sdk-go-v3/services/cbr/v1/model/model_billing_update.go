/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type BillingUpdate struct {
	// 存储库规格
	ConsistentLevel *BillingUpdateConsistentLevel `json:"consistent_level,omitempty"`
	// 存储库大小，单位为GB
	Size *int32 `json:"size,omitempty"`
}

func (o BillingUpdate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BillingUpdate struct{}"
	}

	return strings.Join([]string{"BillingUpdate", string(data)}, " ")
}

type BillingUpdateConsistentLevel struct {
	value string
}

type BillingUpdateConsistentLevelEnum struct {
	APP_CONSISTENT   BillingUpdateConsistentLevel
	CRASH_CONSISTENT BillingUpdateConsistentLevel
}

func GetBillingUpdateConsistentLevelEnum() BillingUpdateConsistentLevelEnum {
	return BillingUpdateConsistentLevelEnum{
		APP_CONSISTENT: BillingUpdateConsistentLevel{
			value: "app_consistent",
		},
		CRASH_CONSISTENT: BillingUpdateConsistentLevel{
			value: "crash_consistent",
		},
	}
}

func (c BillingUpdateConsistentLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BillingUpdateConsistentLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
