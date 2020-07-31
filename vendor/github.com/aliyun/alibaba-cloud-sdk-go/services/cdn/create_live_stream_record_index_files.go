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

// CreateLiveStreamRecordIndexFiles invokes the cdn.CreateLiveStreamRecordIndexFiles API synchronously
// api document: https://help.aliyun.com/api/cdn/createlivestreamrecordindexfiles.html
func (client *Client) CreateLiveStreamRecordIndexFiles(request *CreateLiveStreamRecordIndexFilesRequest) (response *CreateLiveStreamRecordIndexFilesResponse, err error) {
	response = CreateCreateLiveStreamRecordIndexFilesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLiveStreamRecordIndexFilesWithChan invokes the cdn.CreateLiveStreamRecordIndexFiles API asynchronously
// api document: https://help.aliyun.com/api/cdn/createlivestreamrecordindexfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLiveStreamRecordIndexFilesWithChan(request *CreateLiveStreamRecordIndexFilesRequest) (<-chan *CreateLiveStreamRecordIndexFilesResponse, <-chan error) {
	responseChan := make(chan *CreateLiveStreamRecordIndexFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLiveStreamRecordIndexFiles(request)
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

// CreateLiveStreamRecordIndexFilesWithCallback invokes the cdn.CreateLiveStreamRecordIndexFiles API asynchronously
// api document: https://help.aliyun.com/api/cdn/createlivestreamrecordindexfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLiveStreamRecordIndexFilesWithCallback(request *CreateLiveStreamRecordIndexFilesRequest, callback func(response *CreateLiveStreamRecordIndexFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLiveStreamRecordIndexFilesResponse
		var err error
		defer close(result)
		response, err = client.CreateLiveStreamRecordIndexFiles(request)
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

// CreateLiveStreamRecordIndexFilesRequest is the request struct for api CreateLiveStreamRecordIndexFiles
type CreateLiveStreamRecordIndexFilesRequest struct {
	*requests.RpcRequest
	OssEndpoint   string           `position:"Query" name:"OssEndpoint"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OssObject     string           `position:"Query" name:"OssObject"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	StreamName    string           `position:"Query" name:"StreamName"`
	OssBucket     string           `position:"Query" name:"OssBucket"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateLiveStreamRecordIndexFilesResponse is the response struct for api CreateLiveStreamRecordIndexFiles
type CreateLiveStreamRecordIndexFilesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	RecordInfo RecordInfo `json:"RecordInfo" xml:"RecordInfo"`
}

// CreateCreateLiveStreamRecordIndexFilesRequest creates a request to invoke CreateLiveStreamRecordIndexFiles API
func CreateCreateLiveStreamRecordIndexFilesRequest() (request *CreateLiveStreamRecordIndexFilesRequest) {
	request = &CreateLiveStreamRecordIndexFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "CreateLiveStreamRecordIndexFiles", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateLiveStreamRecordIndexFilesResponse creates a response to parse from CreateLiveStreamRecordIndexFiles response
func CreateCreateLiveStreamRecordIndexFilesResponse() (response *CreateLiveStreamRecordIndexFilesResponse) {
	response = &CreateLiveStreamRecordIndexFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
