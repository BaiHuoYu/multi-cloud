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
/*
This module implements a entry into the OpenSDS service.

*/

package cli

import (
	"encoding/json"
	"os"

	backend "github.com/opensds/multi-cloud/backend/proto"
	"github.com/spf13/cobra"
)

var backendCommand = &cobra.Command{
	Use:   "backend",
	Short: "manage backends in the multi-cloud",
	Run:   backendAction,
}

var backendCreateCommand = &cobra.Command{
	Use:   "create <backend info>",
	Short: "create a backend in the multi-cloud",
	Run:   backendShowAction,
}

var backendShowCommand = &cobra.Command{
	Use:   "show <id>",
	Short: "show a backend in the multi-cloud",
	Run:   backendShowAction,
}

var backendListCommand = &cobra.Command{
	Use:   "list",
	Short: "list all backends in the multi-cloud",
	Run:   backendListAction,
}

func init() {
	backendCommand.AddCommand(backendCreateCommand)
	backendCommand.AddCommand(backendShowCommand)
	backendCommand.AddCommand(backendListCommand)
}

func backendAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func backendCreateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	backend := &backend.BackendDetail{}
	if err := json.Unmarshal([]byte(args[0]), backend); err != nil {
		Errorln(err)
		cmd.Usage()
		os.Exit(1)
	}

	resp, err := client.CreateBackend(backend)
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}
	keys := KeyList{"Id", "TenantId", "UserId", "Name", "Type", "Region",
		"Endpoint", "BucketName", "Access", "Security"}
	PrintDict(resp.Backend, keys, FormatterList{})
}

func backendShowAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	resp, err := client.GetBackend(args[0])
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}
	keys := KeyList{"Id", "TenantId", "UserId", "Name", "Type", "Region",
		"Endpoint", "BucketName", "Access", "Security"}
	PrintDict(resp.Backend, keys, FormatterList{})
}

func backendListAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 0)
	resp, err := client.ListBackends()
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}
	keys := KeyList{"Id", "TenantId", "UserId", "Name", "Type", "Region",
		"Endpoint", "BucketName", "Access", "Security"}
	PrintList(resp.Backends, keys, FormatterList{})
}
