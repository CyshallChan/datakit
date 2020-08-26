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

// CurrentPageCursor is a nested struct in domain response
type CurrentPageCursor struct {
	ErrorMsg            string `json:"ErrorMsg" xml:"ErrorMsg"`
	TaskDetailNo        string `json:"TaskDetailNo" xml:"TaskDetailNo"`
	UpdateTime          string `json:"UpdateTime" xml:"UpdateTime"`
	CreateTimeLong      int64  `json:"CreateTimeLong" xml:"CreateTimeLong"`
	DomainName          string `json:"DomainName" xml:"DomainName"`
	TaskStatusCode      int    `json:"TaskStatusCode" xml:"TaskStatusCode"`
	CreateTime          string `json:"CreateTime" xml:"CreateTime"`
	TaskTypeDescription string `json:"TaskTypeDescription" xml:"TaskTypeDescription"`
	TaskStatus          string `json:"TaskStatus" xml:"TaskStatus"`
	TaskNum             int    `json:"TaskNum" xml:"TaskNum"`
	TaskNo              string `json:"TaskNo" xml:"TaskNo"`
	InstanceId          string `json:"InstanceId" xml:"InstanceId"`
	Clientip            string `json:"Clientip" xml:"Clientip"`
	TryCount            int    `json:"TryCount" xml:"TryCount"`
	TaskType            string `json:"TaskType" xml:"TaskType"`
}
