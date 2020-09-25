package ons

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

// OnsMessageGetByMsgId invokes the ons.OnsMessageGetByMsgId API synchronously
// api document: https://help.aliyun.com/api/ons/onsmessagegetbymsgid.html
func (client *Client) OnsMessageGetByMsgId(request *OnsMessageGetByMsgIdRequest) (response *OnsMessageGetByMsgIdResponse, err error) {
	response = CreateOnsMessageGetByMsgIdResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMessageGetByMsgIdWithChan invokes the ons.OnsMessageGetByMsgId API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmessagegetbymsgid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMessageGetByMsgIdWithChan(request *OnsMessageGetByMsgIdRequest) (<-chan *OnsMessageGetByMsgIdResponse, <-chan error) {
	responseChan := make(chan *OnsMessageGetByMsgIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMessageGetByMsgId(request)
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

// OnsMessageGetByMsgIdWithCallback invokes the ons.OnsMessageGetByMsgId API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmessagegetbymsgid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMessageGetByMsgIdWithCallback(request *OnsMessageGetByMsgIdRequest, callback func(response *OnsMessageGetByMsgIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMessageGetByMsgIdResponse
		var err error
		defer close(result)
		response, err = client.OnsMessageGetByMsgId(request)
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

// OnsMessageGetByMsgIdRequest is the request struct for api OnsMessageGetByMsgId
type OnsMessageGetByMsgIdRequest struct {
	*requests.RpcRequest
	MsgId      string `position:"Query" name:"MsgId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// OnsMessageGetByMsgIdResponse is the response struct for api OnsMessageGetByMsgId
type OnsMessageGetByMsgIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsMessageGetByMsgIdRequest creates a request to invoke OnsMessageGetByMsgId API
func CreateOnsMessageGetByMsgIdRequest() (request *OnsMessageGetByMsgIdRequest) {
	request = &OnsMessageGetByMsgIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMessageGetByMsgId", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMessageGetByMsgIdResponse creates a response to parse from OnsMessageGetByMsgId response
func CreateOnsMessageGetByMsgIdResponse() (response *OnsMessageGetByMsgIdResponse) {
	response = &OnsMessageGetByMsgIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
