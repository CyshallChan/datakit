package ecs

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

// DescribeEipMonitorData invokes the ecs.DescribeEipMonitorData API synchronously
// api document: https://help.aliyun.com/api/ecs/describeeipmonitordata.html
func (client *Client) DescribeEipMonitorData(request *DescribeEipMonitorDataRequest) (response *DescribeEipMonitorDataResponse, err error) {
	response = CreateDescribeEipMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEipMonitorDataWithChan invokes the ecs.DescribeEipMonitorData API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeipmonitordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEipMonitorDataWithChan(request *DescribeEipMonitorDataRequest) (<-chan *DescribeEipMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeEipMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEipMonitorData(request)
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

// DescribeEipMonitorDataWithCallback invokes the ecs.DescribeEipMonitorData API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeeipmonitordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEipMonitorDataWithCallback(request *DescribeEipMonitorDataRequest, callback func(response *DescribeEipMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEipMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeEipMonitorData(request)
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

// DescribeEipMonitorDataRequest is the request struct for api DescribeEipMonitorData
type DescribeEipMonitorDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeEipMonitorDataResponse is the response struct for api DescribeEipMonitorData
type DescribeEipMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId       string                                  `json:"RequestId" xml:"RequestId"`
	EipMonitorDatas EipMonitorDatasInDescribeEipMonitorData `json:"EipMonitorDatas" xml:"EipMonitorDatas"`
}

// CreateDescribeEipMonitorDataRequest creates a request to invoke DescribeEipMonitorData API
func CreateDescribeEipMonitorDataRequest() (request *DescribeEipMonitorDataRequest) {
	request = &DescribeEipMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEipMonitorData", "ecs", "openAPI")
	return
}

// CreateDescribeEipMonitorDataResponse creates a response to parse from DescribeEipMonitorData response
func CreateDescribeEipMonitorDataResponse() (response *DescribeEipMonitorDataResponse) {
	response = &DescribeEipMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
