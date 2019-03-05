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
	"os"

	c "github.com/opensds/multi-cloud/client"
	s3 "github.com/opensds/multi-cloud/s3/pkg/model"
	"github.com/spf13/cobra"
)

var (
	xmlns              string
	locationconstraint string
)

type S3BaseResp struct {
	ErrorCode string
	Message   string
}

var bucketCommand = &cobra.Command{
	Use:   "bucket",
	Short: "manage buckets",
	Run:   bucketAction,
}

var objectCommand = &cobra.Command{
	Use:   "object",
	Short: "manage objects",
	Run:   objectAction,
}

var bucketCreateCommand = &cobra.Command{
	Use:   "create <bucket info>",
	Short: "create a bucket",
	Run:   bucketCreateAction,
}

var bucketDeleteCommand = &cobra.Command{
	Use:   "delete <name>",
	Short: "delete a bucket",
	Run:   bucketDeleteAction,
}

var objectListCommand = &cobra.Command{
	Use:   "list <BucketName>",
	Short: "list objects in a bucket",
	Run:   objectListAction,
}

func init() {
	bucketCommand.AddCommand(bucketCreateCommand)
	bucketCreateCommand.Flags().StringVarP(&xmlns, "xmlns", "x", "", "the xmlns of updated bucket")
	bucketCreateCommand.Flags().StringVarP(&locationconstraint, "locationconstraint", "l", "", "the location constraint of updated bucket")

	bucketCommand.AddCommand(bucketDeleteCommand)

	objectCommand.AddCommand(objectListCommand)
}

func bucketAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func objectAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func PrintS3BaseResp(resp *c.CBaseResponse) {
	S3Resp := S3BaseResp{
		ErrorCode: resp.CErrorCode.Value,
		Message:   resp.CMsg.Value,
	}

	keys := KeyList{"ErrorCode", "Message"}
	PrintDict(S3Resp, keys, FormatterList{})
}

func bucketCreateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	bucket := &s3.CreateBucketConfiguration{
		Xmlns:              xmlns,
		LocationConstraint: locationconstraint,
	}

	resp, err := client.CreateBucket(args[0], bucket)
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}

	PrintS3BaseResp(resp)
}

func bucketDeleteAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)

	resp, err := client.DeleteBucket(args[0])
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}

	PrintS3BaseResp(resp)
}

func objectListAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)

	resp, err := client.ListObjects(args[0])
	if err != nil {
		Fatalln(HTTPErrStrip(err))
	}

	keys := KeyList{"ObjectKey", "BucketName", "Size", "Backend"}
	PrintList(resp.ListObjects, keys, FormatterList{})
}
