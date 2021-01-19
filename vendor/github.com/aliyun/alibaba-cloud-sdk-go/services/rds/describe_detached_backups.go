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

// DescribeDetachedBackups invokes the rds.DescribeDetachedBackups API synchronously
// api document: https://help.aliyun.com/api/rds/describedetachedbackups.html
func (client *Client) DescribeDetachedBackups(request *DescribeDetachedBackupsRequest) (response *DescribeDetachedBackupsResponse, err error) {
	response = CreateDescribeDetachedBackupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDetachedBackupsWithChan invokes the rds.DescribeDetachedBackups API asynchronously
// api document: https://help.aliyun.com/api/rds/describedetachedbackups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDetachedBackupsWithChan(request *DescribeDetachedBackupsRequest) (<-chan *DescribeDetachedBackupsResponse, <-chan error) {
	responseChan := make(chan *DescribeDetachedBackupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDetachedBackups(request)
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

// DescribeDetachedBackupsWithCallback invokes the rds.DescribeDetachedBackups API asynchronously
// api document: https://help.aliyun.com/api/rds/describedetachedbackups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDetachedBackupsWithCallback(request *DescribeDetachedBackupsRequest, callback func(response *DescribeDetachedBackupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDetachedBackupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDetachedBackups(request)
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

// DescribeDetachedBackupsRequest is the request struct for api DescribeDetachedBackups
type DescribeDetachedBackupsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	BackupLocation       string           `position:"Query" name:"BackupLocation"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	BackupStatus         string           `position:"Query" name:"BackupStatus"`
	BackupMode           string           `position:"Query" name:"BackupMode"`
	Region               string           `position:"Query" name:"Region"`
}

// DescribeDetachedBackupsResponse is the response struct for api DescribeDetachedBackups
type DescribeDetachedBackupsResponse struct {
	*responses.BaseResponse
	RequestId        string                         `json:"RequestId" xml:"RequestId"`
	TotalRecordCount string                         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       string                         `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  string                         `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalBackupSize  int64                          `json:"TotalBackupSize" xml:"TotalBackupSize"`
	Items            ItemsInDescribeDetachedBackups `json:"Items" xml:"Items"`
}

// CreateDescribeDetachedBackupsRequest creates a request to invoke DescribeDetachedBackups API
func CreateDescribeDetachedBackupsRequest() (request *DescribeDetachedBackupsRequest) {
	request = &DescribeDetachedBackupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDetachedBackups", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDetachedBackupsResponse creates a response to parse from DescribeDetachedBackups response
func CreateDescribeDetachedBackupsResponse() (response *DescribeDetachedBackupsResponse) {
	response = &DescribeDetachedBackupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
