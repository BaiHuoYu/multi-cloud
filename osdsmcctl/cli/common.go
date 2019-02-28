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

package cli

import (
	"fmt"
	"os"

	c "github.com/opensds/multi-cloud/client"
	"github.com/spf13/cobra"
)

const (
	errorPrefix = "ERROR:"
	debugPrefix = "DEBUG:"
	warnPrefix  = "WARNING:"
)

// Printf implementation
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

// Debugf implementation
func Debugf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		return fmt.Fprintf(os.Stdout, debugPrefix+" "+format, a...)
	}
	return 0, nil
}

// Warnf implementation
func Warnf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, warnPrefix+" "+format, a...)
}

// Errorf implementation
func Errorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, errorPrefix+" "+format, a...)
}

// Fatalf implementation
func Fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, errorPrefix+" "+format, a...)
	os.Exit(-1)
}

// Println implementation
func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

// Debugln implementation
func Debugln(a ...interface{}) (n int, err error) {
	if Debug {
		a = append([]interface{}{debugPrefix}, a...)
		return fmt.Fprintln(os.Stdout, a...)
	}
	return 0, nil
}

// Warnln implementation
func Warnln(a ...interface{}) (n int, err error) {
	a = append([]interface{}{warnPrefix}, a...)
	return fmt.Fprintln(os.Stdout, a...)
}

// Errorln implementation
func Errorln(a ...interface{}) (n int, err error) {
	a = append([]interface{}{errorPrefix}, a...)
	return fmt.Fprintln(os.Stderr, a...)
}

// Fatalln implementation
func Fatalln(a ...interface{}) {
	a = append([]interface{}{errorPrefix}, a...)
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(-1)
}

// HTTPErrStrip Strip some redundant message from client http error.
func HTTPErrStrip(err error) error {
	fmt.Printf("HTTPErrStrip 95 %+v/n", err)
	if httpErr, ok := err.(*c.HTTPError); ok {
		fmt.Printf("HTTPErrStrip 97 %+v-----------%+v/n", httpErr.Msg, httpErr.Desc)
		httpErr.Decode()
		if "" != httpErr.Msg {
			return fmt.Errorf(httpErr.Msg)
		}

		if "" != httpErr.Desc {
			return fmt.Errorf(httpErr.Desc)
		}
	}

	return err
}

// ArgsNumCheck implementation
func ArgsNumCheck(cmd *cobra.Command, args []string, invalidNum int) {
	if len(args) != invalidNum {
		Errorln("The number of args is not correct!")
		cmd.Usage()
		os.Exit(1)
	}
}
