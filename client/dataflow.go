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

	"github.com/opensds/multi-cloud/dataflow/proto"
)

// NewDataflowMgr implementation
func NewDataflowMgr(r Receiver, edp string, tenantID string) *DataflowMgr {
	return &DataflowMgr{
		Receiver: r,
		Endpoint: edp,
		TenantID: tenantID,
	}
}

// DataflowMgr implementation
type DataflowMgr struct {
	Receiver
	Endpoint string
	TenantID string
}

// CreatePlan implementation
func (b *BackendMgr) CreatePlan(body *dataflow.Plan) (*dataflow.Plan, error) {
	res := dataflow.Plan{}
	url := strings.Join([]string{
		b.Endpoint,
		GeneratePlanURL(b.TenantID)}, "/")

	if err := b.Recv(url, "POST", JsonHeaders, body, &res, "", ""); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreatePolicy implementation
func (b *BackendMgr) CreatePolicy(body *dataflow.Policy) (*dataflow.Policy, error) {
	res := dataflow.Policy{}
	url := strings.Join([]string{
		b.Endpoint,
		GeneratePolicyURL(b.TenantID)}, "/")

	if err := b.Recv(url, "POST", JsonHeaders, body, &res, "", ""); err != nil {
		return nil, err
	}

	return &res, nil
}
