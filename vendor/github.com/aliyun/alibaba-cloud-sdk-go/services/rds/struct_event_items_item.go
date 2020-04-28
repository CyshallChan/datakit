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

// EventItemsItem is a nested struct in rds response
type EventItemsItem struct {
	EventName       string `json:"EventName" xml:"EventName"`
	EventPayload    string `json:"EventPayload" xml:"EventPayload"`
	ResourceType    string `json:"ResourceType" xml:"ResourceType"`
	EventReason     string `json:"EventReason" xml:"EventReason"`
	EventSig        string `json:"EventSig" xml:"EventSig"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
	EventType       string `json:"EventType" xml:"EventType"`
	ResourceName    string `json:"ResourceName" xml:"ResourceName"`
	EventContent    string `json:"EventContent" xml:"EventContent"`
	EventId         int    `json:"EventId" xml:"EventId"`
	EventRcpt       string `json:"EventRcpt" xml:"EventRcpt"`
	EventTime       string `json:"EventTime" xml:"EventTime"`
	EventUserType   string `json:"EventUserType" xml:"EventUserType"`
	EventRecordTime string `json:"EventRecordTime" xml:"EventRecordTime"`
}
