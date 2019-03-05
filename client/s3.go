// Copyright (c) 2019 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"encoding/xml"
	"strings"

	bucket "github.com/opensds/multi-cloud/s3/pkg/model"
)

type CBaseResponse struct {
	XMLName               xml.Name               `xml:"BaseResponse,omitempty" json:"BaseResponse,omitempty"`
	CErrorCode            *CErrorCode            `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`
	CMsg                  *CMsg                  `xml:"Msg,omitempty" json:"Msg,omitempty"`
	CXXX_NoUnkeyedLiteral *CXXX_NoUnkeyedLiteral `xml:"XXX_NoUnkeyedLiteral,omitempty" json:"XXX_NoUnkeyedLiteral,omitempty"`
	CXXX_sizecache        *CXXX_sizecache        `xml:"XXX_sizecache,omitempty" json:"XXX_sizecache,omitempty"`
	CXXX_unrecognized     *CXXX_unrecognized     `xml:"XXX_unrecognized,omitempty" json:"XXX_unrecognized,omitempty"`
}

type CErrorCode struct {
	XMLName xml.Name `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CMsg struct {
	XMLName xml.Name `xml:"Msg,omitempty" json:"Msg,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CXXX_NoUnkeyedLiteral struct {
	XMLName xml.Name `xml:"XXX_NoUnkeyedLiteral,omitempty" json:"XXX_NoUnkeyedLiteral,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CXXX_sizecache struct {
	XMLName xml.Name `xml:"XXX_sizecache,omitempty" json:"XXX_sizecache,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CXXX_unrecognized struct {
	XMLName xml.Name `xml:"XXX_unrecognized,omitempty" json:"XXX_unrecognized,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

// NewBucketMgr implementation
func NewBucketMgr(r Receiver, edp string, tenantID string) *BucketMgr {
	return &BucketMgr{
		Receiver: r,
		Endpoint: edp,
		TenantID: tenantID,
	}
}

// BucketMgr implementation
type BucketMgr struct {
	Receiver
	Endpoint string
	TenantID string
}

// CreateBucket implementation
func (b *BucketMgr) CreateBucket(name string, body *bucket.CreateBucketConfiguration) (*CBaseResponse, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID), name}, "/")

	res := CBaseResponse{}
	if err := b.Recv(url, "PUT", XmlHeaders, body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
