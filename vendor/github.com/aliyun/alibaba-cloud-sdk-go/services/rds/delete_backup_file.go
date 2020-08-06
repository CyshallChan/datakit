package rds

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

// DeleteBackupFile invokes the rds.DeleteBackupFile API synchronously
// api document: https://help.aliyun.com/api/rds/deletebackupfile.html
func (client *Client) DeleteBackupFile(request *DeleteBackupFileRequest) (response *DeleteBackupFileResponse, err error) {
	response = CreateDeleteBackupFileResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBackupFileWithChan invokes the rds.DeleteBackupFile API asynchronously
// api document: https://help.aliyun.com/api/rds/deletebackupfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBackupFileWithChan(request *DeleteBackupFileRequest) (<-chan *DeleteBackupFileResponse, <-chan error) {
	responseChan := make(chan *DeleteBackupFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBackupFile(request)
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

// DeleteBackupFileWithCallback invokes the rds.DeleteBackupFile API asynchronously
// api document: https://help.aliyun.com/api/rds/deletebackupfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBackupFileWithCallback(request *DeleteBackupFileRequest, callback func(response *DeleteBackupFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBackupFileResponse
		var err error
		defer close(result)
		response, err = client.DeleteBackupFile(request)
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

// DeleteBackupFileRequest is the request struct for api DeleteBackupFile
type DeleteBackupFileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	BackupTime           string           `position:"Query" name:"BackupTime"`
	DBName               string           `position:"Query" name:"DBName"`
}

// DeleteBackupFileResponse is the response struct for api DeleteBackupFile
type DeleteBackupFileResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	DeletedBaksetIds DeletedBaksetIds `json:"DeletedBaksetIds" xml:"DeletedBaksetIds"`
}

// CreateDeleteBackupFileRequest creates a request to invoke DeleteBackupFile API
func CreateDeleteBackupFileRequest() (request *DeleteBackupFileRequest) {
	request = &DeleteBackupFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteBackupFile", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteBackupFileResponse creates a response to parse from DeleteBackupFile response
func CreateDeleteBackupFileResponse() (response *DeleteBackupFileResponse) {
	response = &DeleteBackupFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
