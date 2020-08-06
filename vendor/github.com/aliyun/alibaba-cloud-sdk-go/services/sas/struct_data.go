package sas

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

// Data is a nested struct in sas response
type Data struct {
	DataSource          string         `json:"DataSource" xml:"DataSource"`
	Account             int            `json:"Account" xml:"Account"`
	EndTime             int64          `json:"EndTime" xml:"EndTime"`
	AlarmUniqueInfo     string         `json:"AlarmUniqueInfo" xml:"AlarmUniqueInfo"`
	Uuid                string         `json:"Uuid" xml:"Uuid"`
	StartTime           int64          `json:"StartTime" xml:"StartTime"`
	Vul                 int            `json:"Vul" xml:"Vul"`
	IntranetIp          string         `json:"IntranetIp" xml:"IntranetIp"`
	CanCancelFault      bool           `json:"CanCancelFault" xml:"CanCancelFault"`
	Health              int            `json:"Health" xml:"Health"`
	Type                string         `json:"Type" xml:"Type"`
	Solution            string         `json:"Solution" xml:"Solution"`
	InternetIp          string         `json:"InternetIp" xml:"InternetIp"`
	Level               string         `json:"Level" xml:"Level"`
	InstanceName        string         `json:"InstanceName" xml:"InstanceName"`
	Suspicious          int            `json:"Suspicious" xml:"Suspicious"`
	NewSuspicious       int            `json:"NewSuspicious" xml:"NewSuspicious"`
	AlarmEventDesc      string         `json:"AlarmEventDesc" xml:"AlarmEventDesc"`
	AlarmEventAliasName string         `json:"AlarmEventAliasName" xml:"AlarmEventAliasName"`
	CanBeDealOnLine     bool           `json:"CanBeDealOnLine" xml:"CanBeDealOnLine"`
	Trojan              int            `json:"Trojan" xml:"Trojan"`
	EntityTypeList      []EntityType   `json:"EntityTypeList" xml:"EntityTypeList"`
	CauseDetails        []CauseDetail  `json:"CauseDetails" xml:"CauseDetails"`
	VertexList          []Vertex       `json:"VertexList" xml:"VertexList"`
	RelationTypeList    []RelationType `json:"RelationTypeList" xml:"RelationTypeList"`
	EdgeList            []Edge         `json:"EdgeList" xml:"EdgeList"`
}
