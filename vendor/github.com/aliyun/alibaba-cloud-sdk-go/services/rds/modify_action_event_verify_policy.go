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

// ModifyActionEventVerifyPolicy invokes the rds.ModifyActionEventVerifyPolicy API synchronously
// api document: https://help.aliyun.com/api/rds/modifyactioneventverifypolicy.html
func (client *Client) ModifyActionEventVerifyPolicy(request *ModifyActionEventVerifyPolicyRequest) (response *ModifyActionEventVerifyPolicyResponse, err error) {
	response = CreateModifyActionEventVerifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyActionEventVerifyPolicyWithChan invokes the rds.ModifyActionEventVerifyPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyactioneventverifypolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyActionEventVerifyPolicyWithChan(request *ModifyActionEventVerifyPolicyRequest) (<-chan *ModifyActionEventVerifyPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyActionEventVerifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyActionEventVerifyPolicy(request)
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

// ModifyActionEventVerifyPolicyWithCallback invokes the rds.ModifyActionEventVerifyPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyactioneventverifypolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyActionEventVerifyPolicyWithCallback(request *ModifyActionEventVerifyPolicyRequest, callback func(response *ModifyActionEventVerifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyActionEventVerifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyActionEventVerifyPolicy(request)
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

// ModifyActionEventVerifyPolicyRequest is the request struct for api ModifyActionEventVerifyPolicy
type ModifyActionEventVerifyPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UserPublicKey        string           `position:"Query" name:"UserPublicKey"`
}

// ModifyActionEventVerifyPolicyResponse is the response struct for api ModifyActionEventVerifyPolicy
type ModifyActionEventVerifyPolicyResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
	ServerPublicKey string `json:"ServerPublicKey" xml:"ServerPublicKey"`
}

// CreateModifyActionEventVerifyPolicyRequest creates a request to invoke ModifyActionEventVerifyPolicy API
func CreateModifyActionEventVerifyPolicyRequest() (request *ModifyActionEventVerifyPolicyRequest) {
	request = &ModifyActionEventVerifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyActionEventVerifyPolicy", "rds", "openAPI")
	return
}

// CreateModifyActionEventVerifyPolicyResponse creates a response to parse from ModifyActionEventVerifyPolicy response
func CreateModifyActionEventVerifyPolicyResponse() (response *ModifyActionEventVerifyPolicyResponse) {
	response = &ModifyActionEventVerifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
