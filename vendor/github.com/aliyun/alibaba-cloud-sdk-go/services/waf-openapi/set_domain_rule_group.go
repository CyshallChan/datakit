package waf_openapi

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

// SetDomainRuleGroup invokes the waf_openapi.SetDomainRuleGroup API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/setdomainrulegroup.html
func (client *Client) SetDomainRuleGroup(request *SetDomainRuleGroupRequest) (response *SetDomainRuleGroupResponse, err error) {
	response = CreateSetDomainRuleGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SetDomainRuleGroupWithChan invokes the waf_openapi.SetDomainRuleGroup API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/setdomainrulegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainRuleGroupWithChan(request *SetDomainRuleGroupRequest) (<-chan *SetDomainRuleGroupResponse, <-chan error) {
	responseChan := make(chan *SetDomainRuleGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomainRuleGroup(request)
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

// SetDomainRuleGroupWithCallback invokes the waf_openapi.SetDomainRuleGroup API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/setdomainrulegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainRuleGroupWithCallback(request *SetDomainRuleGroupRequest, callback func(response *SetDomainRuleGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainRuleGroupResponse
		var err error
		defer close(result)
		response, err = client.SetDomainRuleGroup(request)
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

// SetDomainRuleGroupRequest is the request struct for api SetDomainRuleGroup
type SetDomainRuleGroupRequest struct {
	*requests.RpcRequest
	WafVersion      requests.Integer `position:"Query" name:"WafVersion"`
	RuleGroupId     requests.Integer `position:"Query" name:"RuleGroupId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	Domains         string           `position:"Query" name:"Domains"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
}

// SetDomainRuleGroupResponse is the response struct for api SetDomainRuleGroup
type SetDomainRuleGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDomainRuleGroupRequest creates a request to invoke SetDomainRuleGroup API
func CreateSetDomainRuleGroupRequest() (request *SetDomainRuleGroupRequest) {
	request = &SetDomainRuleGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "SetDomainRuleGroup", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDomainRuleGroupResponse creates a response to parse from SetDomainRuleGroup response
func CreateSetDomainRuleGroupResponse() (response *SetDomainRuleGroupResponse) {
	response = &SetDomainRuleGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
