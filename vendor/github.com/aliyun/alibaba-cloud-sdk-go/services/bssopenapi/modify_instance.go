package bssopenapi

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

// ModifyInstance invokes the bssopenapi.ModifyInstance API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/modifyinstance.html
func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
	response = CreateModifyInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceWithChan invokes the bssopenapi.ModifyInstance API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/modifyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceWithChan(request *ModifyInstanceRequest) (<-chan *ModifyInstanceResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstance(request)
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

// ModifyInstanceWithCallback invokes the bssopenapi.ModifyInstance API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/modifyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceWithCallback(request *ModifyInstanceRequest, callback func(response *ModifyInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstance(request)
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

// ModifyInstanceRequest is the request struct for api ModifyInstance
type ModifyInstanceRequest struct {
	*requests.RpcRequest
	ProductCode      string                     `position:"Query" name:"ProductCode"`
	ClientToken      string                     `position:"Query" name:"ClientToken"`
	SubscriptionType string                     `position:"Query" name:"SubscriptionType"`
	OwnerId          requests.Integer           `position:"Query" name:"OwnerId"`
	ProductType      string                     `position:"Query" name:"ProductType"`
	InstanceId       string                     `position:"Query" name:"InstanceId"`
	ModifyType       string                     `position:"Query" name:"ModifyType"`
	Parameter        *[]ModifyInstanceParameter `position:"Query" name:"Parameter"  type:"Repeated"`
}

// ModifyInstanceParameter is a repeated param struct in ModifyInstanceRequest
type ModifyInstanceParameter struct {
	Code  string `name:"Code"`
	Value string `name:"Value"`
}

// ModifyInstanceResponse is the response struct for api ModifyInstance
type ModifyInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateModifyInstanceRequest creates a request to invoke ModifyInstance API
func CreateModifyInstanceRequest() (request *ModifyInstanceRequest) {
	request = &ModifyInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "ModifyInstance", "", "")
	return
}

// CreateModifyInstanceResponse creates a response to parse from ModifyInstance response
func CreateModifyInstanceResponse() (response *ModifyInstanceResponse) {
	response = &ModifyInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
