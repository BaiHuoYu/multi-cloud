// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
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

	"github.com/opensds/multi-cloud/backend/proto"
)

// NewBackendMgr
func NewBackendMgr(r Receiver, edp string, tenantId string) *BackendMgr {
	return &BackendMgr{
		Receiver: r,
		Endpoint: edp,
		TenantId: tenantId,
	}
}

// BackendMgr
type BackendMgr struct {
	Receiver
	Endpoint string
	TenantId string
}

func CurrentVersion() string {
	return "v1"
}

func generateURL(resource string, tenantId string, in ...string) string {
	// If project id is not specified, ignore it.
	if tenantId == "" {
		value := []string{CurrentVersion(), resource}
		value = append(value, in...)
		return strings.Join(value, "/")
	}

	value := []string{CurrentVersion(), tenantId, resource}
	value = append(value, in...)

	return strings.Join(value, "/")
}

// GetVolume
func (b *BackendMgr) GetBackend(Id string) (*backend.GetBackendResponse, error) {
	var res backend.GetBackendResponse
	url := strings.Join([]string{
		b.Endpoint,
		generateURL("backends", b.TenantId, Id)}, "/")

	if err := b.Recv(url, "GET", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListBackends
func (b *BackendMgr) ListBackends() (*backend.GetBackendResponse, error) {
	var res backend.GetBackendResponse
	url := strings.Join([]string{
		b.Endpoint,
		generateURL("backends", b.TenantId)}, "/")

	if err := b.Recv(url, "GET", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
