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

// DescribeDomainRealTimeTrafficData invokes the cdn.DescribeDomainRealTimeTrafficData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimetrafficdata.html
func (client *Client) DescribeDomainRealTimeTrafficData(request *DescribeDomainRealTimeTrafficDataRequest) (response *DescribeDomainRealTimeTrafficDataResponse, err error) {
	response = CreateDescribeDomainRealTimeTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeTrafficDataWithChan invokes the cdn.DescribeDomainRealTimeTrafficData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimetrafficdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRealTimeTrafficDataWithChan(request *DescribeDomainRealTimeTrafficDataRequest) (<-chan *DescribeDomainRealTimeTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeTrafficData(request)
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

// DescribeDomainRealTimeTrafficDataWithCallback invokes the cdn.DescribeDomainRealTimeTrafficData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimetrafficdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRealTimeTrafficDataWithCallback(request *DescribeDomainRealTimeTrafficDataRequest, callback func(response *DescribeDomainRealTimeTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeTrafficData(request)
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

// DescribeDomainRealTimeTrafficDataRequest is the request struct for api DescribeDomainRealTimeTrafficData
type DescribeDomainRealTimeTrafficDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainRealTimeTrafficDataResponse is the response struct for api DescribeDomainRealTimeTrafficData
type DescribeDomainRealTimeTrafficDataResponse struct {
	*responses.BaseResponse
	RequestId                      string                         `json:"RequestId" xml:"RequestId"`
	DomainName                     string                         `json:"DomainName" xml:"DomainName"`
	StartTime                      string                         `json:"StartTime" xml:"StartTime"`
	EndTime                        string                         `json:"EndTime" xml:"EndTime"`
	DataInterval                   string                         `json:"DataInterval" xml:"DataInterval"`
	RealTimeTrafficDataPerInterval RealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval" xml:"RealTimeTrafficDataPerInterval"`
}

// CreateDescribeDomainRealTimeTrafficDataRequest creates a request to invoke DescribeDomainRealTimeTrafficData API
func CreateDescribeDomainRealTimeTrafficDataRequest() (request *DescribeDomainRealTimeTrafficDataRequest) {
	request = &DescribeDomainRealTimeTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeTrafficData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainRealTimeTrafficDataResponse creates a response to parse from DescribeDomainRealTimeTrafficData response
func CreateDescribeDomainRealTimeTrafficDataResponse() (response *DescribeDomainRealTimeTrafficDataResponse) {
	response = &DescribeDomainRealTimeTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
