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

	S3model "github.com/opensds/multi-cloud/s3/pkg/model"
	"github.com/opensds/multi-cloud/s3/proto"
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
func (b *BucketMgr) CreateBucket(name string, body *S3model.CreateBucketConfiguration) (*CBaseResponse, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID), name}, "/")

	res := CBaseResponse{}
	if err := b.Recv(url, "PUT", XmlHeaders, body, &res, "", ""); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteBucket implementation
func (b *BucketMgr) DeleteBucket(name string) (*CBaseResponse, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID), name}, "/")

	res := CBaseResponse{}
	if err := b.Recv(url, "DELETE", XmlHeaders, nil, &res, "", ""); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListBuckets implementation
func (b *BucketMgr) ListBuckets() ([]S3model.Bucket, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID)}, "/")

	res := S3model.ListAllMyBucketsResult{}
	if err := b.Recv(url, "GET", XmlHeaders, nil, &res, "", ""); err != nil {
		return nil, err
	}

	return res.Buckets, nil
}

// ListObjects implementation
func (b *BucketMgr) ListObjects(BucketName string) ([]*s3.Object, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID), BucketName}, "/")

	res := s3.ListObjectResponse{}
	if err := b.Recv(url, "GET", XmlHeaders, nil, &res, "", ""); err != nil {
		return nil, err
	}

	return res.ListObjects, nil
}

// UploadObject implementation
func (b *BucketMgr) UploadObject(BucketName, ObjectKey, Object string) (*CBaseResponse, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID), BucketName, ObjectKey}, "/")

	res := CBaseResponse{}
	//buf, err := ioutil.ReadFile(Object)
	//if err != nil {
	//	return &res, err
	//}

	//Headers := HeaderOption{obs.HEADER_CONTENT_TYPE: "application/xml",
	//	obs.HEADER_CONTENT_LENGTH: strconv.Itoa(len(buf)),
	//}

	if err := b.Recv(url, "PUT", XmlHeaders, nil, &res, ObjectKey, Object); err != nil {
		return nil, err
	}

	return &res, nil
}

// DownloadObject implementation
func (b *BucketMgr) DownloadObject(BucketName string, Object string) (*CBaseResponse, error) {
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID), BucketName, Object}, "/")

	res := CBaseResponse{}
	if err := b.Recv(url, "GET", XmlHeaders, nil, &res, "", ""); err != nil {
		return nil, err
	}

	return &res, nil
}
