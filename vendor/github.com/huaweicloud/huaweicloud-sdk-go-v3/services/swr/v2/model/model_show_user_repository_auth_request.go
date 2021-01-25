/*
 * SWR
 *
 * SWR API
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ShowUserRepositoryAuthRequest struct {
	ContentType ShowUserRepositoryAuthRequestContentType `json:"Content-Type"`
	Namespace   string                                   `json:"namespace"`
	Repository  string                                   `json:"repository"`
}

func (o ShowUserRepositoryAuthRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserRepositoryAuthRequest struct{}"
	}

	return strings.Join([]string{"ShowUserRepositoryAuthRequest", string(data)}, " ")
}

type ShowUserRepositoryAuthRequestContentType struct {
	value string
}

type ShowUserRepositoryAuthRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowUserRepositoryAuthRequestContentType
	APPLICATION_JSON             ShowUserRepositoryAuthRequestContentType
}

func GetShowUserRepositoryAuthRequestContentTypeEnum() ShowUserRepositoryAuthRequestContentTypeEnum {
	return ShowUserRepositoryAuthRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowUserRepositoryAuthRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowUserRepositoryAuthRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowUserRepositoryAuthRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowUserRepositoryAuthRequestContentType) UnmarshalJSON(b []byte) error {
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
