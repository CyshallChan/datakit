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

// ListTagResources invokes the ons.ListTagResources API synchronously
// api document: https://help.aliyun.com/api/ons/listtagresources.html
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (response *ListTagResourcesResponse, err error) {
	response = CreateListTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagResourcesWithChan invokes the ons.ListTagResources API asynchronously
// api document: https://help.aliyun.com/api/ons/listtagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagResourcesWithChan(request *ListTagResourcesRequest) (<-chan *ListTagResourcesResponse, <-chan error) {
	responseChan := make(chan *ListTagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagResources(request)
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

// ListTagResourcesWithCallback invokes the ons.ListTagResources API asynchronously
// api document: https://help.aliyun.com/api/ons/listtagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagResourcesWithCallback(request *ListTagResourcesRequest, callback func(response *ListTagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagResourcesResponse
		var err error
		defer close(result)
		response, err = client.ListTagResources(request)
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

// ListTagResourcesRequest is the request struct for api ListTagResources
type ListTagResourcesRequest struct {
	*requests.RpcRequest
	ResourceId   *[]string              `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceType string                 `position:"Query" name:"ResourceType"`
	InstanceId   string                 `position:"Query" name:"InstanceId"`
	NextToken    string                 `position:"Query" name:"NextToken"`
	Tag          *[]ListTagResourcesTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// ListTagResourcesTag is a repeated param struct in ListTagResourcesRequest
type ListTagResourcesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListTagResourcesResponse is the response struct for api ListTagResources
type ListTagResourcesResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	NextToken    string        `json:"NextToken" xml:"NextToken"`
	TagResources []TagResource `json:"TagResources" xml:"TagResources"`
}

// CreateListTagResourcesRequest creates a request to invoke ListTagResources API
func CreateListTagResourcesRequest() (request *ListTagResourcesRequest) {
	request = &ListTagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "ListTagResources", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTagResourcesResponse creates a response to parse from ListTagResources response
func CreateListTagResourcesResponse() (response *ListTagResourcesResponse) {
	response = &ListTagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
