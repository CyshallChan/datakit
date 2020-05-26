package sas

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

// DescribeAlarmEventList invokes the sas.DescribeAlarmEventList API synchronously
// api document: https://help.aliyun.com/api/sas/describealarmeventlist.html
func (client *Client) DescribeAlarmEventList(request *DescribeAlarmEventListRequest) (response *DescribeAlarmEventListResponse, err error) {
	response = CreateDescribeAlarmEventListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlarmEventListWithChan invokes the sas.DescribeAlarmEventList API asynchronously
// api document: https://help.aliyun.com/api/sas/describealarmeventlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmEventListWithChan(request *DescribeAlarmEventListRequest) (<-chan *DescribeAlarmEventListResponse, <-chan error) {
	responseChan := make(chan *DescribeAlarmEventListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlarmEventList(request)
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

// DescribeAlarmEventListWithCallback invokes the sas.DescribeAlarmEventList API asynchronously
// api document: https://help.aliyun.com/api/sas/describealarmeventlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmEventListWithCallback(request *DescribeAlarmEventListRequest, callback func(response *DescribeAlarmEventListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlarmEventListResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlarmEventList(request)
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

// DescribeAlarmEventListRequest is the request struct for api DescribeAlarmEventList
type DescribeAlarmEventListRequest struct {
	*requests.RpcRequest
	AlarmEventName string           `position:"Query" name:"AlarmEventName"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	PageSize       string           `position:"Query" name:"PageSize"`
	AlarmEventType string           `position:"Query" name:"AlarmEventType"`
	Dealed         string           `position:"Query" name:"Dealed"`
	From           string           `position:"Query" name:"From"`
	Remark         string           `position:"Query" name:"Remark"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	Lang           string           `position:"Query" name:"Lang"`
	Levels         string           `position:"Query" name:"Levels"`
}

// DescribeAlarmEventListResponse is the response struct for api DescribeAlarmEventList
type DescribeAlarmEventListResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	PageInfo   PageInfo         `json:"PageInfo" xml:"PageInfo"`
	SuspEvents []SuspEventsItem `json:"SuspEvents" xml:"SuspEvents"`
}

// CreateDescribeAlarmEventListRequest creates a request to invoke DescribeAlarmEventList API
func CreateDescribeAlarmEventListRequest() (request *DescribeAlarmEventListRequest) {
	request = &DescribeAlarmEventListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeAlarmEventList", "sas", "openAPI")
	return
}

// CreateDescribeAlarmEventListResponse creates a response to parse from DescribeAlarmEventList response
func CreateDescribeAlarmEventListResponse() (response *DescribeAlarmEventListResponse) {
	response = &DescribeAlarmEventListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
