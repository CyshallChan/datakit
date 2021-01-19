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

// RenewInstance invokes the rds.RenewInstance API synchronously
// api document: https://help.aliyun.com/api/rds/renewinstance.html
func (client *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
	response = CreateRenewInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RenewInstanceWithChan invokes the rds.RenewInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/renewinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewInstanceWithChan(request *RenewInstanceRequest) (<-chan *RenewInstanceResponse, <-chan error) {
	responseChan := make(chan *RenewInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewInstance(request)
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

// RenewInstanceWithCallback invokes the rds.RenewInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/renewinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewInstanceWithCallback(request *RenewInstanceRequest, callback func(response *RenewInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewInstanceResponse
		var err error
		defer close(result)
		response, err = client.RenewInstance(request)
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

// RenewInstanceRequest is the request struct for api RenewInstance
type RenewInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	Period               requests.Integer `position:"Query" name:"Period"`
	AutoPay              string           `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RenewInstanceResponse is the response struct for api RenewInstance
type RenewInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
}

// CreateRenewInstanceRequest creates a request to invoke RenewInstance API
func CreateRenewInstanceRequest() (request *RenewInstanceRequest) {
	request = &RenewInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RenewInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewInstanceResponse creates a response to parse from RenewInstance response
func CreateRenewInstanceResponse() (response *RenewInstanceResponse) {
	response = &RenewInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
