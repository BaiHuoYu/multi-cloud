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

	"github.com/opensds/multi-cloud/api/pkg/utils/constants"
)

func GenerateBackendURL(tenantId string, in ...string) string {
	return generateURL("backends", tenantId, in...)
}

func GenerateTypeURL(tenantId string, in ...string) string {
	return generateURL("types", tenantId, in...)
}

func GeneratePlanURL(tenantId string, in ...string) string {
	return generateURL("plans", tenantId, in...)
}

func GenerateJobURL(tenantId string, in ...string) string {
	return generateURL("jobs", tenantId, in...)
}

func GeneratePolicyURL(tenantId string, in ...string) string {
	return generateURL("policies", tenantId, in...)
}

func GenerateS3URL(tenantId string, in ...string) string {
	return generateURL("s3", tenantId, in...)
}

func CurrentVersion() string {
	return constants.APIVersion
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
