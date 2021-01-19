package resourcemanager

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

// MoveAccount invokes the resourcemanager.MoveAccount API synchronously
// api document: https://help.aliyun.com/api/resourcemanager/moveaccount.html
func (client *Client) MoveAccount(request *MoveAccountRequest) (response *MoveAccountResponse, err error) {
	response = CreateMoveAccountResponse()
	err = client.DoAction(request, response)
	return
}

// MoveAccountWithChan invokes the resourcemanager.MoveAccount API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/moveaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoveAccountWithChan(request *MoveAccountRequest) (<-chan *MoveAccountResponse, <-chan error) {
	responseChan := make(chan *MoveAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoveAccount(request)
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

// MoveAccountWithCallback invokes the resourcemanager.MoveAccount API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/moveaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoveAccountWithCallback(request *MoveAccountRequest, callback func(response *MoveAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoveAccountResponse
		var err error
		defer close(result)
		response, err = client.MoveAccount(request)
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

// MoveAccountRequest is the request struct for api MoveAccount
type MoveAccountRequest struct {
	*requests.RpcRequest
	AccountId           string `position:"Query" name:"AccountId"`
	DestinationFolderId string `position:"Query" name:"DestinationFolderId"`
}

// MoveAccountResponse is the response struct for api MoveAccount
type MoveAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMoveAccountRequest creates a request to invoke MoveAccount API
func CreateMoveAccountRequest() (request *MoveAccountRequest) {
	request = &MoveAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "MoveAccount", "resourcemanager", "openAPI")
	return
}

// CreateMoveAccountResponse creates a response to parse from MoveAccount response
func CreateMoveAccountResponse() (response *MoveAccountResponse) {
	response = &MoveAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
