package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDBInstancePayType invokes the rds.ModifyDBInstancePayType API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancepaytype.html
func (client *Client) ModifyDBInstancePayType(request *ModifyDBInstancePayTypeRequest) (response *ModifyDBInstancePayTypeResponse, err error) {
	response = CreateModifyDBInstancePayTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstancePayTypeWithChan invokes the rds.ModifyDBInstancePayType API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancepaytype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstancePayTypeWithChan(request *ModifyDBInstancePayTypeRequest) (<-chan *ModifyDBInstancePayTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstancePayTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstancePayType(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDBInstancePayTypeWithCallback invokes the rds.ModifyDBInstancePayType API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancepaytype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstancePayTypeWithCallback(request *ModifyDBInstancePayTypeRequest, callback func(response *ModifyDBInstancePayTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstancePayTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstancePayType(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDBInstancePayTypeRequest is the request struct for api ModifyDBInstancePayType
type ModifyDBInstancePayTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	Period               string           `position:"Query" name:"Period"`
	AgentId              string           `position:"Query" name:"AgentId"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Resource             string           `position:"Query" name:"Resource"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime             requests.Integer `position:"Query" name:"UsedTime"`
	PayType              string           `position:"Query" name:"PayType"`
}

// ModifyDBInstancePayTypeResponse is the response struct for api ModifyDBInstancePayType
type ModifyDBInstancePayTypeResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId      int64  `json:"OrderId" xml:"OrderId"`
}

// CreateModifyDBInstancePayTypeRequest creates a request to invoke ModifyDBInstancePayType API
func CreateModifyDBInstancePayTypeRequest() (request *ModifyDBInstancePayTypeRequest) {
	request = &ModifyDBInstancePayTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstancePayType", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstancePayTypeResponse creates a response to parse from ModifyDBInstancePayType response
func CreateModifyDBInstancePayTypeResponse() (response *ModifyDBInstancePayTypeResponse) {
	response = &ModifyDBInstancePayTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
