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
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/credentials"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/opensds/multi-cloud/api/pkg/filters/signature/credentials/keystonecredentials"
	"github.com/opensds/multi-cloud/api/pkg/filters/signature/signer"
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
		reqBody interface{}, respBody interface{}, needMarshal bool, outFileName string) error
}

// NewReceiver implementation
func NewReceiver() Receiver {
	return &receiver{}
}

func CalculateSignature(headers HeaderOption, req *http.Request) string {
	authorization, ok := headers[constants.AuthorizationHeader]
	if !ok {
		log.Printf("no ", constants.AuthorizationHeader)
		return ""
	}

	//Get the X-Auth-Date Header from the request
	requestDateTime, ok := headers[constants.SignDateHeader]
	if !ok {
		log.Printf("no ", constants.SignDateHeader)
		return ""
	}

	log.Printf("authorization %+v, requestDateTime %+v", authorization, requestDateTime)
	//Get the Authorization parameters from the Authorization String
	authorizationParts := strings.Split(authorization, ",")
	credential, _ := strings.TrimSpace(authorizationParts[0]), strings.TrimSpace(authorizationParts[2])
	credentialParts := strings.Split(credential, " ")
	creds := credentialParts[1]
	credsParts := strings.Split(creds, "=")
	credentialStr := credsParts[1]
	credentialStrParts := strings.Split(credentialStr, "/")
	accessKeyID, requestDate, region, service := credentialStrParts[0], credentialStrParts[1], credentialStrParts[2], credentialStrParts[3]
	log.Printf("accessKeyID:%+v, requestDate:%+v, region:%+v, service:%+v", accessKeyID, requestDate, region, service)
	//TODO Get Request Body
	body := ""

	//Create a keystone credentials Provider client for retrieving credentials
	credentials := keystonecredentials.NewCredentialsClient(accessKeyID)
	log.Printf("credentials %+v", credentials)
	//Create a Signer and the calculate the signature based on the Header parameters passed in request
	Signer := signer.NewSigner(credentials)
	calculatedSignature, err := Signer.Sign(req, body, service, region, requestDateTime, requestDate, credentialStr)
	log.Printf("req:%+v, body:%+v, service:%+v, region:%+v, requestDateTime:%+v, requestDate:%+v, credentialStr:%+v", req, body, service, region, requestDateTime, requestDate, credentialStr)
	if err != nil {
		log.Printf("signer.Sign err:%+v", err)
		return ""
	}

	log.Printf("calculatedSignature:%+v", calculatedSignature)
	return calculatedSignature
}

// request implementation
func request(url string, method string, headers HeaderOption,
	reqBody interface{}, respBody interface{}, needMarshal bool, outFileName string) error {
	log.Printf("\nurl=%+v\nmethod=%+v\nheaders=%+v\nreqBody=%+v\nrespBody=%+v\nneedMarshal=%+v\noutFileName=%+v\n",
		url, method, headers, reqBody, respBody, needMarshal, outFileName)
	var err error
	req := httplib.NewBeegoRequest(url, strings.ToUpper(method))
	req.SetTimeout(time.Minute*6, time.Minute*6)
	contentType, ok := headers[obs.HEADER_CONTENT_TYPE]
	if !ok {
		log.Printf("Content-Type was not be configured in the request header")
	}

	if reqBody != nil {
		var body []byte

		if needMarshal {
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
		}

		log.Printf("Request body:\n%s\n", string(body))
		if needMarshal {
			req.Body(body)
		} else {
			req.Body(reqBody)
		}
	}

	calculatedSignature := CalculateSignature(headers, req.GetRequest())
	log.Printf("calculatedSignature:%v", calculatedSignature)
	//init header
	if headers != nil {
		for k, v := range headers {
			req.Header(k, v)
		}
	}

	// Get http response.
	resp, err := req.Response()
	if err != nil {
		return err
	}

	if 400 <= resp.StatusCode && resp.StatusCode <= 599 {
		log.Printf("Response statusCode: %+v\n", resp.StatusCode)
		return NewHTTPError(resp.StatusCode, "")
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll failed, err: %+v\n", err)
		return err
	}

	defer resp.Body.Close()
	log.Printf("Response Status: %s\nResponse Body:\n%s\n", resp.Status, string(rbody))
	if (nil == rbody) || ("" == string(rbody)) {
		return nil
	}

	if "" == outFileName {
		if nil == respBody {
			return nil
		}

		var respContentType string
		respContentTypes, ok := resp.Header["Content-Type"]
		log.Printf("ok=%+v, respContentTypes=%+v, len=%v\n", ok, respContentTypes, len(respContentTypes))

		if ok && len(respContentTypes) > 0 {
			respContentType = respContentTypes[0]
		}

		switch respContentType {
		case constants.HeaderValueJson:
			if err = json.Unmarshal(rbody, respBody); err != nil {
				return fmt.Errorf("failed to unmarshal result message: %v", err)
			}
			log.Printf("application/json, respBody=%+v\n", respBody)
			break
		case constants.HeaderValueXml, "text/xml; charset=utf-8":
			if err = xml.Unmarshal(rbody, respBody); err != nil {
				return fmt.Errorf("failed to unmarshal result message: %v", err)
			}
			log.Printf("application/xml, respBody=%+v\n", respBody)
			break
		default:
			log.Printf("Failure to process the response body!")
		}
	} else {
		path := fmt.Sprintf("./%s", outFileName)
		file, err := os.Create(path)
		if err != nil {
			log.Printf("Failed to create file:%+v\n", err)
		}
		defer file.Close()

		n, err := file.Write(rbody)
		if err != nil {
			log.Printf("Failed to Write file,err:%+v\n, n:%+v\n", err, n)
		}
		log.Printf("Save file successfully, n:%+v\n", n)
	}

	return nil
}

type receiver struct{}

func (*receiver) Recv(url string, method string, headers HeaderOption,
	reqBody interface{}, respBody interface{}, needMarshal bool, outFileName string) error {
	return request(url, method, headers, reqBody, respBody, needMarshal, outFileName)
}

// NewKeystoneReciver implementation
func NewKeystoneReciver(auth *KeystoneAuthOptions) Receiver {
	k := &KeystoneReciver{Auth: auth}
	err := k.GetTokenAndCredential()
	if err != nil {
		log.Printf("Failed to get token: %v", err)
	}
	return k
}

// KeystoneReciver implementation
type KeystoneReciver struct {
	Auth *KeystoneAuthOptions
}

// GetToken implementation
func (k *KeystoneReciver) GetTokenAndCredential() error {
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

	credentialOpts := credentials.CreateOpts{
		Blob:      "{\"access\":\"181920\",\"secret\":\"secretKey\"}",
		ProjectID: project.ID,
		Type:      "ec2",
		UserID:    k.Auth.UserID,
	}

	credential, err := credentials.Create(identity, &credentialOpts).Extract()
	if err != nil {
		log.Printf("cannot create credential, err:%v, CreateOpts:%+v\n", err, credentialOpts)
		return err
	}
	log.Printf("credential:%v", err, credential)

	return nil
}

// Recv implementation
func (k *KeystoneReciver) Recv(url string, method string, headers HeaderOption,
	reqBody interface{}, respBody interface{}, needMarshal bool, outFileName string) error {
	desc := fmt.Sprintf("%s %s", method, url)
	return utils.Retry(2, desc, true, func(retryIdx int, lastErr error) error {
		if retryIdx > 0 {
			err, ok := lastErr.(*HTTPError)
			if ok && err.Code == http.StatusUnauthorized {
				err := k.GetTokenAndCredential()
				if err != nil {
					log.Printf("Failed to get token: %v", err)
				}
			} else {
				return lastErr
			}
		}
		log.Printf("Recv----url: %v", url)
		headers[constants.AuthTokenHeader] = k.Auth.TokenID
		if strings.Contains(url, "/s3") {
			headers[constants.AuthorizationHeader] = "OPENSDS-HMAC-SHA256 Credential=access_key/20190301/us-east-1/s3/sign_request,SignedHeaders=authorization;host;x-auth-date,Signature=472f0a1b7815974847620da53fcdd2fdd53203b5d8d08e7ce81943b260560e26"
			headers[constants.SignDateHeader] = time.Now().Format("20060102T150405Z")
		}
		log.Printf("Recv----headers: %v", url)
		return request(url, method, headers, reqBody, respBody, needMarshal, outFileName)
	})
}

func checkHTTPResponseStatusCode(resp *http.Response) error {
	if 400 <= resp.StatusCode && resp.StatusCode <= 599 {
		return fmt.Errorf("response == %d, %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	return nil
}
