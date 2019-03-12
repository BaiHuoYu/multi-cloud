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

	dataflow "github.com/opensds/multi-cloud/dataflow/proto"
	"github.com/spf13/cobra"
)

var planCommand = &cobra.Command{
	Use:   "plan",
	Short: "manage plans in the multi-cloud",
	Run:   backendAction,
}

var planCreateCommand = &cobra.Command{
	Use:   "create <plan info>",
	Short: "create a plan",
	Run:   planCreateAction,
}

var policyCommand = &cobra.Command{
	Use:   "policy",
	Short: "manage policies in the multi-cloud",
	Run:   policyAction,
}

var policyCreateCommand = &cobra.Command{
	Use:   "create <policy info>",
	Short: "create a policy",
	Run:   policyCreateAction,
}

func init() {
	planCommand.AddCommand(planCreateCommand)

	policyCommand.AddCommand(policyCreateCommand)
}

func planAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func policyAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func planCreateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	plan := &dataflow.Plan{}
	if err := json.Unmarshal([]byte(args[0]), plan); err != nil {
		Errorln(err)
		cmd.Usage()
		os.Exit(1)
	}

	resp, err := client.CreatePlan(plan)
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}
	keys := KeyList{"Id", "Name", "Description", "Type", "PolicyId", "PolicyName",
		"SourceConn", "DestConn", "Filter", "RemainSource", "TenantId", "UserId",
		"PolicyEnabled"}
	PrintDict(resp, keys, FormatterList{})
}

func policyCreateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	policy := &dataflow.Policy{}
	if err := json.Unmarshal([]byte(args[0]), policy); err != nil {
		Errorln(err)
		cmd.Usage()
		os.Exit(1)
	}

	resp, err := client.CreatePolicy(policy)
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}
	keys := KeyList{"Id", "Name", "Tenant", "Description", "Schedule"}
	PrintDict(resp, keys, FormatterList{})
}
