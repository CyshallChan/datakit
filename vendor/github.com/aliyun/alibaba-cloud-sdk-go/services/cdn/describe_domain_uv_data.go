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

// DescribeDomainUvData invokes the cdn.DescribeDomainUvData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainuvdata.html
func (client *Client) DescribeDomainUvData(request *DescribeDomainUvDataRequest) (response *DescribeDomainUvDataResponse, err error) {
	response = CreateDescribeDomainUvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainUvDataWithChan invokes the cdn.DescribeDomainUvData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainuvdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainUvDataWithChan(request *DescribeDomainUvDataRequest) (<-chan *DescribeDomainUvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainUvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainUvData(request)
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

// DescribeDomainUvDataWithCallback invokes the cdn.DescribeDomainUvData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainuvdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainUvDataWithCallback(request *DescribeDomainUvDataRequest, callback func(response *DescribeDomainUvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainUvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainUvData(request)
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

// DescribeDomainUvDataRequest is the request struct for api DescribeDomainUvData
type DescribeDomainUvDataRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDomainUvDataResponse is the response struct for api DescribeDomainUvData
type DescribeDomainUvDataResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	DomainName     string         `json:"DomainName" xml:"DomainName"`
	DataInterval   string         `json:"DataInterval" xml:"DataInterval"`
	StartTime      string         `json:"StartTime" xml:"StartTime"`
	EndTime        string         `json:"EndTime" xml:"EndTime"`
	UvDataInterval UvDataInterval `json:"UvDataInterval" xml:"UvDataInterval"`
}

// CreateDescribeDomainUvDataRequest creates a request to invoke DescribeDomainUvData API
func CreateDescribeDomainUvDataRequest() (request *DescribeDomainUvDataRequest) {
	request = &DescribeDomainUvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainUvData", "", "")
	return
}

// CreateDescribeDomainUvDataResponse creates a response to parse from DescribeDomainUvData response
func CreateDescribeDomainUvDataResponse() (response *DescribeDomainUvDataResponse) {
	response = &DescribeDomainUvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
