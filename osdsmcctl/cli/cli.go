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
/*
This module implements a entry into the OpenSDS CLI service.

*/

package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/opensds/multi-cloud/api/pkg/filters/context"
	c "github.com/opensds/multi-cloud/client"
	"github.com/opensds/opensds/pkg/utils"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	client      *c.Client
	rootCommand = &cobra.Command{
		Use:   "osdsmcctl",
		Short: "Administer the OpenSDS multi-cloud",
		Long:  `Admin utility for the OpenSDS multi-cloud.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
			os.Exit(1)
		},
	}
	Debug bool

	DockerComposePath        = "DOCKER_COMPOSE_PATH"
	DefaultDockerComposePath = "/root/gopath/src/github.com/opensds/multi-cloud/docker-compose.yml"
)

type DockerComposeAPI struct {
	Image       string   `yaml:"image"`
	Volumes     []string `yaml:"volumes"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
}

func init() {
	rootCommand.AddCommand(backendCommand)
	flags := rootCommand.PersistentFlags()
	flags.BoolVar(&Debug, "debug", false, "shows debugging output.")
}

type DummyWriter struct{}

// do nothing
func (writer DummyWriter) Write(data []byte) (n int, err error) {
	return len(data), nil
}

type DebugWriter struct{}

// do nothing
func (writer DebugWriter) Write(data []byte) (n int, err error) {
	Debugf("%s", string(data))
	return len(data), nil
}

func GetAPIEnvs() []string {
	path, ok := os.LookupEnv(DockerComposePath)
	if !ok {
		path = DefaultDockerComposePath
	}

	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Read config yaml file (%s) failed, reason:(%v)", path, err)
		return nil
	}

	apiConf := &DockerComposeAPI{}
	if err = yaml.Unmarshal(ymlFile, apiConf); err != nil {
		log.Printf("Parse error: %v", err)

		return nil
	}

	return apiConf.Environment
}

// Run method indicates how to start a cli tool through cobra.
func Run() error {
	if !utils.Contained("--debug", os.Args) {
		log.SetOutput(DummyWriter{})
	} else {
		log.SetOutput(DebugWriter{})
	}

	ep, ok := os.LookupEnv(c.MultiCloudEndpoint)
	if !ok {
		return fmt.Errorf("ERROR: You must provide the endpoint by setting " +
			"the environment variable MULTI_CLOUD_ENDPOINT")
	}

	cfg := &c.Config{Endpoint: ep}
	APIEnvs := GetAPIEnvs()
	authStrategy := c.GetValueFromStrArray(APIEnvs, "OS_AUTH_AUTHSTRATEGY")

	switch authStrategy {
	case c.Keystone:
		cfg.AuthOptions = c.LoadKeystoneAuthOptions(APIEnvs)
	case c.Noauth:
		cfg.AuthOptions = c.NewNoauthOptions(context.NoAuthAdminTenantId)
	default:
		cfg.AuthOptions = c.NewNoauthOptions(context.DefaultTenantId)
	}

	client = c.NewClient(cfg)

	return rootCommand.Execute()
}
