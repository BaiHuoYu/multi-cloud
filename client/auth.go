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
	"log"
	"strings"
)

const (
	// The values of OS_AUTH_AUTHSTRATEGY
	Keystone = "keystone"
	Noauth   = "noauth"

	// Api environment variable name in docker-compose.yml
	MicroServerAddress = "MICRO_SERVER_ADDRESS"
	OsAuthAuthstrategy = "OS_AUTH_AUTHSTRATEGY"
	OsAuthURL          = "OS_AUTH_URL"
	OsUserName         = "OS_USERNAME"
	OsPassword         = "OS_PASSWORD"
	OsTenantName       = "OS_TENANT_NAME"
	OsProjectName      = "OS_PROJECT_NAME"
	OsUserDominID      = "OS_USER_DOMIN_ID"
)

type AuthOptions interface {
	GetTenantId() string
}

func NewKeystoneAuthOptions() *KeystoneAuthOptions {
	return &KeystoneAuthOptions{}
}

type KeystoneAuthOptions struct {
	IdentityEndpoint string
	Username         string
	UserID           string
	Password         string
	DomainID         string
	DomainName       string
	TenantID         string
	TenantName       string
	AllowReauth      bool
	TokenID          string
}

func (k *KeystoneAuthOptions) GetTenantId() string {
	return k.TenantID
}

func NewNoauthOptions(tenantId string) *NoAuthOptions {
	return &NoAuthOptions{TenantID: tenantId}
}

type NoAuthOptions struct {
	TenantID string
}

func (n *NoAuthOptions) GetTenantId() string {
	return n.TenantID
}

func GetValueFromStrArray(strArray []string, key string) string {
	value := ""

	for _, str := range strArray {
		if strings.HasPrefix(str, key+"=") {
			authArray := strings.Split(str, "=")

			if len(authArray) > 1 {
				value = authArray[1]
			} else {
				log.Printf("There is no value in %+v ", key)
			}

			break
		}
	}

	log.Printf("There is no %+v in %+v ", key, strArray)
	return value
}

func LoadKeystoneAuthOptions(envs []string) *KeystoneAuthOptions {
	opt := NewKeystoneAuthOptions()
	opt.IdentityEndpoint = GetValueFromStrArray(envs, OsAuthURL)
	opt.Username = GetValueFromStrArray(envs, OsUserName)
	opt.Password = GetValueFromStrArray(envs, OsPassword)
	opt.TenantName = GetValueFromStrArray(envs, OsTenantName)
	projectName := GetValueFromStrArray(envs, OsProjectName)
	opt.DomainID = GetValueFromStrArray(envs, OsUserDominID)
	if opt.TenantName == "" {
		opt.TenantName = projectName
	}

	return opt
}
