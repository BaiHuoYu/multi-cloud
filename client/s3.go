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
	"strings"

	bucket "github.com/opensds/multi-cloud/s3/pkg/model"
)

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
func (b *BucketMgr) CreateBucket(name string, body *bucket.CreateBucketConfiguration) error {
	var res bucket.CreateBucketConfiguration
	url := strings.Join([]string{
		b.Endpoint,
		GenerateS3URL(b.TenantID)}, "/")

	if err := b.Recv(url, "POST", XmlHeaders, body, &res); err != nil {
		return err
	}

	return nil
}
