/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchRebootServersRequest struct {
	Body *BatchRebootServersRequestBody `json:"body,omitempty"`
}

func (o BatchRebootServersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchRebootServersRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootServersRequest", string(data)}, " ")
}
