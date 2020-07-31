package bssopenapi

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

// QueryBill invokes the bssopenapi.QueryBill API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybill.html
func (client *Client) QueryBill(request *QueryBillRequest) (response *QueryBillResponse, err error) {
	response = CreateQueryBillResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBillWithChan invokes the bssopenapi.QueryBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBillWithChan(request *QueryBillRequest) (<-chan *QueryBillResponse, <-chan error) {
	responseChan := make(chan *QueryBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBill(request)
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

// QueryBillWithCallback invokes the bssopenapi.QueryBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBillWithCallback(request *QueryBillRequest, callback func(response *QueryBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBillResponse
		var err error
		defer close(result)
		response, err = client.QueryBill(request)
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

// QueryBillRequest is the request struct for api QueryBill
type QueryBillRequest struct {
	*requests.RpcRequest
	ProductCode            string           `position:"Query" name:"ProductCode"`
	IsHideZeroCharge       requests.Boolean `position:"Query" name:"IsHideZeroCharge"`
	IsDisplayLocalCurrency requests.Boolean `position:"Query" name:"IsDisplayLocalCurrency"`
	SubscriptionType       string           `position:"Query" name:"SubscriptionType"`
	BillingCycle           string           `position:"Query" name:"BillingCycle"`
	Type                   string           `position:"Query" name:"Type"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	PageNum                requests.Integer `position:"Query" name:"PageNum"`
	ProductType            string           `position:"Query" name:"ProductType"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
}

// QueryBillResponse is the response struct for api QueryBill
type QueryBillResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryBillRequest creates a request to invoke QueryBill API
func CreateQueryBillRequest() (request *QueryBillRequest) {
	request = &QueryBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryBill", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryBillResponse creates a response to parse from QueryBill response
func CreateQueryBillResponse() (response *QueryBillResponse) {
	response = &QueryBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
