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

// BatchFuzzyMatchDomainSensitiveWord invokes the domain.BatchFuzzyMatchDomainSensitiveWord API synchronously
// api document: https://help.aliyun.com/api/domain/batchfuzzymatchdomainsensitiveword.html
func (client *Client) BatchFuzzyMatchDomainSensitiveWord(request *BatchFuzzyMatchDomainSensitiveWordRequest) (response *BatchFuzzyMatchDomainSensitiveWordResponse, err error) {
	response = CreateBatchFuzzyMatchDomainSensitiveWordResponse()
	err = client.DoAction(request, response)
	return
}

// BatchFuzzyMatchDomainSensitiveWordWithChan invokes the domain.BatchFuzzyMatchDomainSensitiveWord API asynchronously
// api document: https://help.aliyun.com/api/domain/batchfuzzymatchdomainsensitiveword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchFuzzyMatchDomainSensitiveWordWithChan(request *BatchFuzzyMatchDomainSensitiveWordRequest) (<-chan *BatchFuzzyMatchDomainSensitiveWordResponse, <-chan error) {
	responseChan := make(chan *BatchFuzzyMatchDomainSensitiveWordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchFuzzyMatchDomainSensitiveWord(request)
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

// BatchFuzzyMatchDomainSensitiveWordWithCallback invokes the domain.BatchFuzzyMatchDomainSensitiveWord API asynchronously
// api document: https://help.aliyun.com/api/domain/batchfuzzymatchdomainsensitiveword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchFuzzyMatchDomainSensitiveWordWithCallback(request *BatchFuzzyMatchDomainSensitiveWordRequest, callback func(response *BatchFuzzyMatchDomainSensitiveWordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchFuzzyMatchDomainSensitiveWordResponse
		var err error
		defer close(result)
		response, err = client.BatchFuzzyMatchDomainSensitiveWord(request)
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

// BatchFuzzyMatchDomainSensitiveWordRequest is the request struct for api BatchFuzzyMatchDomainSensitiveWord
type BatchFuzzyMatchDomainSensitiveWordRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Keyword      string `position:"Query" name:"Keyword"`
	Lang         string `position:"Query" name:"Lang"`
}

// BatchFuzzyMatchDomainSensitiveWordResponse is the response struct for api BatchFuzzyMatchDomainSensitiveWord
type BatchFuzzyMatchDomainSensitiveWordResponse struct {
	*responses.BaseResponse
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	SensitiveWordMatchResultList SensitiveWordMatchResultList `json:"SensitiveWordMatchResultList" xml:"SensitiveWordMatchResultList"`
}

// CreateBatchFuzzyMatchDomainSensitiveWordRequest creates a request to invoke BatchFuzzyMatchDomainSensitiveWord API
func CreateBatchFuzzyMatchDomainSensitiveWordRequest() (request *BatchFuzzyMatchDomainSensitiveWordRequest) {
	request = &BatchFuzzyMatchDomainSensitiveWordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "BatchFuzzyMatchDomainSensitiveWord", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchFuzzyMatchDomainSensitiveWordResponse creates a response to parse from BatchFuzzyMatchDomainSensitiveWord response
func CreateBatchFuzzyMatchDomainSensitiveWordResponse() (response *BatchFuzzyMatchDomainSensitiveWordResponse) {
	response = &BatchFuzzyMatchDomainSensitiveWordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
