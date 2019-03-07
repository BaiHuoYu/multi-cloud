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
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/opensds/multi-cloud/api/pkg/model"
	"github.com/opensds/multi-cloud/api/pkg/utils"
	"github.com/opensds/multi-cloud/api/pkg/utils/constants"
	"github.com/opensds/multi-cloud/api/pkg/utils/obs"
)

// NewHTTPError implementation
func NewHTTPError(code int, msg string) error {
	return &HTTPError{Code: code, Msg: msg}
}

// HTTPError implementation
type HTTPError struct {
	Code int
	Desc string
	Msg  string
}

// Decode implementation
func (e *HTTPError) Decode() {
	errSpec := model.ErrorSpec{}
	err := json.Unmarshal([]byte(e.Msg), &errSpec)
	if err == nil {
		e.Msg = errSpec.Message
	}

	err = json.Unmarshal([]byte(e.Desc), &errSpec)
	if err == nil {
		e.Desc = errSpec.Desc
	}
}

// Error implementation
func (e *HTTPError) Error() string {
	e.Decode()
	return fmt.Sprintf("Code: %v, Desc: %s, Msg: %v", e.Code, http.StatusText(e.Code), e.Msg)
}

// HeaderOption implementation
type HeaderOption map[string]string

// Receiver implementation
type Receiver interface {
	Recv(url string, method string, headers HeaderOption,
		reqBody interface{}, respBody interface{}, ObjectKey string, Object string) error
}

// NewReceiver implementation
func NewReceiver() Receiver {
	return &receiver{}
}

// request implementation
func request(url string, method string, headers HeaderOption,
	reqBody interface{}, respBody interface{}, ObjectKey string, Object string) error {
	req := httplib.NewBeegoRequest(url, strings.ToUpper(method))
	// Set the request timeout a little bit longer upload snapshot to cloud temporarily.
	req.SetTimeout(time.Minute*6, time.Minute*6)
	// init body
	log.Printf("%s %s\n", strings.ToUpper(method), url)
	contentType, ok := headers[obs.HEADER_CONTENT_TYPE]
	if !ok {
		return NewHTTPError(http.StatusInternalServerError,
			"Content-Type must be configured in the header")
	}

	if reqBody != nil {
		var body []byte
		var err error

		switch contentType {
		case constants.HeaderValueJson:
			body, err = json.MarshalIndent(reqBody, "", "  ")
			if err != nil {
				return err
			}
			break
		case constants.HeaderValueXml:
			body, err = xml.Marshal(reqBody)
			if err != nil {
				return err
			}
			break
		default:
			log.Printf("Content-Type is not application/json nor application/xml\n")
		}

		log.Printf("Request body:\n%s\n", string(body))
		req.Body(body)
	}

	if "" != ObjectKey && "" != Object {
		req.PostFile(ObjectKey, Object)
	}

	//init header
	if headers != nil {
		for k, v := range headers {
			req.Header(k, v)
		}
	}

	log.Printf("req=%+v, headers=%v\n", req, headers)
	// Get http response.
	resp, err := req.Response()
	if err != nil {
		return err
	}

	log.Printf("resp=%+v\n", resp)
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("\nStatusCode: %s\nResponse Body:\n%s\n", resp.Status, string(rbody))
	if 400 <= resp.StatusCode && resp.StatusCode <= 599 {
		return NewHTTPError(resp.StatusCode, string(rbody))
	}

	if (respBody == nil) || (nil == rbody) || ("" == string(rbody)) {
		return nil
	}

	switch contentType {
	case constants.HeaderValueJson:
		if err = json.Unmarshal(rbody, respBody); err != nil {
			return fmt.Errorf("failed to unmarshal result message: %v", err)
		}
		break
	case constants.HeaderValueXml:
		if err = xml.Unmarshal(rbody, respBody); err != nil {
			return fmt.Errorf("failed to unmarshal result message: %v", err)
		}
		break
	default:
		log.Printf("Content-Type is not application/json nor application/xml\n")
	}

	return nil
}

type receiver struct{}

func (*receiver) Recv(url string, method string, headers HeaderOption,
	reqBody interface{}, respBody interface{}, ObjectKey string, Object string) error {
	return request(url, method, headers, reqBody, respBody, ObjectKey, Object)
}

// NewKeystoneReciver implementation
func NewKeystoneReciver(auth *KeystoneAuthOptions) Receiver {
	k := &KeystoneReciver{Auth: auth}
	k.GetToken()
	return k
}

// KeystoneReciver implementation
type KeystoneReciver struct {
	Auth *KeystoneAuthOptions
}

// GetToken implementation
func (k *KeystoneReciver) GetToken() error {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: k.Auth.IdentityEndpoint,
		Username:         k.Auth.Username,
		UserID:           k.Auth.UserID,
		Password:         k.Auth.Password,
		DomainID:         k.Auth.DomainID,
		DomainName:       k.Auth.DomainName,
		TenantID:         k.Auth.TenantID,
		TenantName:       k.Auth.TenantName,
		AllowReauth:      k.Auth.AllowReauth,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return fmt.Errorf("When get auth client: %v", err)
	}

	// Only support keystone v3
	identity, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return fmt.Errorf("When get identity session: %v", err)
	}
	r := tokens.Create(identity, &opts)
	token, err := r.ExtractToken()
	if err != nil {
		return fmt.Errorf("When get extract token session: %v", err)
	}
	project, err := r.ExtractProject()
	if err != nil {
		return fmt.Errorf("When get extract project session: %v", err)
	}
	k.Auth.TenantID = project.ID
	k.Auth.TokenID = token.ID
	return nil
}

// Recv implementation
func (k *KeystoneReciver) Recv(url string, method string, headers HeaderOption,
	reqBody interface{}, respBody interface{}, ObjectKey string, Object string) error {
	desc := fmt.Sprintf("%s %s", method, url)
	return utils.Retry(2, desc, true, func(retryIdx int, lastErr error) error {
		if retryIdx > 0 {
			err, ok := lastErr.(*HTTPError)
			if ok && err.Code == http.StatusUnauthorized {
				k.GetToken()
			} else {
				return lastErr
			}
		}

		headers[constants.AuthTokenHeader] = k.Auth.TokenID
		headers[obs.HEADER_CONTENT_TYPE] = constants.HeaderValueJson

		return request(url, method, headers, reqBody, respBody, ObjectKey, Object)
	})
}

func checkHTTPResponseStatusCode(resp *http.Response) error {
	if 400 <= resp.StatusCode && resp.StatusCode <= 599 {
		return fmt.Errorf("response == %d, %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	return nil
}
