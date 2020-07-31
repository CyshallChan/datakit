package cdn

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

// UntagResources invokes the cdn.UntagResources API synchronously
// api document: https://help.aliyun.com/api/cdn/untagresources.html
func (client *Client) UntagResources(request *UntagResourcesRequest) (response *UntagResourcesResponse, err error) {
	response = CreateUntagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// UntagResourcesWithChan invokes the cdn.UntagResources API asynchronously
// api document: https://help.aliyun.com/api/cdn/untagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourcesWithChan(request *UntagResourcesRequest) (<-chan *UntagResourcesResponse, <-chan error) {
	responseChan := make(chan *UntagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UntagResources(request)
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

// UntagResourcesWithCallback invokes the cdn.UntagResources API asynchronously
// api document: https://help.aliyun.com/api/cdn/untagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourcesWithCallback(request *UntagResourcesRequest, callback func(response *UntagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UntagResourcesResponse
		var err error
		defer close(result)
		response, err = client.UntagResources(request)
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

// UntagResourcesRequest is the request struct for api UntagResources
type UntagResourcesRequest struct {
	*requests.RpcRequest
	ResourceId   *[]string        `position:"Query" name:"ResourceId"  type:"Repeated"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	TagKey       *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
}

// UntagResourcesResponse is the response struct for api UntagResources
type UntagResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUntagResourcesRequest creates a request to invoke UntagResources API
func CreateUntagResourcesRequest() (request *UntagResourcesRequest) {
	request = &UntagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "UntagResources", "", "")
	request.Method = requests.POST
	return
}

// CreateUntagResourcesResponse creates a response to parse from UntagResources response
func CreateUntagResourcesResponse() (response *UntagResourcesResponse) {
	response = &UntagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
