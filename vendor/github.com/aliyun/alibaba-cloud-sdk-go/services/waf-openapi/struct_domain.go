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

// Domain is a nested struct in waf_openapi response
type Domain struct {
	HttpToUserIp    int         `json:"HttpToUserIp" xml:"HttpToUserIp"`
	HttpsRedirect   int         `json:"HttpsRedirect" xml:"HttpsRedirect"`
	LoadBalancing   int         `json:"LoadBalancing" xml:"LoadBalancing"`
	Cname           string      `json:"Cname" xml:"Cname"`
	IsAccessProduct int         `json:"IsAccessProduct" xml:"IsAccessProduct"`
	Version         int64       `json:"Version" xml:"Version"`
	ClusterType     int         `json:"ClusterType" xml:"ClusterType"`
	ConnectionTime  int         `json:"ConnectionTime" xml:"ConnectionTime"`
	ReadTime        int         `json:"ReadTime" xml:"ReadTime"`
	WriteTime       int         `json:"WriteTime" xml:"WriteTime"`
	ResourceGroupId string      `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SourceIps       []string    `json:"SourceIps" xml:"SourceIps"`
	Http2Port       []string    `json:"Http2Port" xml:"Http2Port"`
	HttpPort        []string    `json:"HttpPort" xml:"HttpPort"`
	HttpsPort       []string    `json:"HttpsPort" xml:"HttpsPort"`
	LogHeaders      []LogHeader `json:"LogHeaders" xml:"LogHeaders"`
}
