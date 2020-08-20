package domain

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

// TransferInCheckMailToken invokes the domain.TransferInCheckMailToken API synchronously
// api document: https://help.aliyun.com/api/domain/transferincheckmailtoken.html
func (client *Client) TransferInCheckMailToken(request *TransferInCheckMailTokenRequest) (response *TransferInCheckMailTokenResponse, err error) {
	response = CreateTransferInCheckMailTokenResponse()
	err = client.DoAction(request, response)
	return
}

// TransferInCheckMailTokenWithChan invokes the domain.TransferInCheckMailToken API asynchronously
// api document: https://help.aliyun.com/api/domain/transferincheckmailtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransferInCheckMailTokenWithChan(request *TransferInCheckMailTokenRequest) (<-chan *TransferInCheckMailTokenResponse, <-chan error) {
	responseChan := make(chan *TransferInCheckMailTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferInCheckMailToken(request)
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

// TransferInCheckMailTokenWithCallback invokes the domain.TransferInCheckMailToken API asynchronously
// api document: https://help.aliyun.com/api/domain/transferincheckmailtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransferInCheckMailTokenWithCallback(request *TransferInCheckMailTokenRequest, callback func(response *TransferInCheckMailTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferInCheckMailTokenResponse
		var err error
		defer close(result)
		response, err = client.TransferInCheckMailToken(request)
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

// TransferInCheckMailTokenRequest is the request struct for api TransferInCheckMailToken
type TransferInCheckMailTokenRequest struct {
	*requests.RpcRequest
	Token        string `position:"Query" name:"Token"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// TransferInCheckMailTokenResponse is the response struct for api TransferInCheckMailToken
type TransferInCheckMailTokenResponse struct {
	*responses.BaseResponse
	RequestId   string                                `json:"RequestId" xml:"RequestId"`
	SuccessList SuccessListInTransferInCheckMailToken `json:"SuccessList" xml:"SuccessList"`
	FailList    FailListInTransferInCheckMailToken    `json:"FailList" xml:"FailList"`
}

// CreateTransferInCheckMailTokenRequest creates a request to invoke TransferInCheckMailToken API
func CreateTransferInCheckMailTokenRequest() (request *TransferInCheckMailTokenRequest) {
	request = &TransferInCheckMailTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "TransferInCheckMailToken", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTransferInCheckMailTokenResponse creates a response to parse from TransferInCheckMailToken response
func CreateTransferInCheckMailTokenResponse() (response *TransferInCheckMailTokenResponse) {
	response = &TransferInCheckMailTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
