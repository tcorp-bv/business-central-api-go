/*
 * MIT License
 *
 * Copyright (c) 2020 TCorp BV
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package businesscentral

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tcorp-bv/business-central-api-go/auth"
	"github.com/tcorp-bv/business-central-api-go/swagger"
	"io/ioutil"
	"net/http"
)

const (
	productionEnvironment = "production"
	endpointFmt           = "https://api.businesscentral.dynamics.com/v2.0/%s/api/v1.0"
	// The company ID for TCorp (really not so smart to hardcode this but whatever)
	TCorp = "253a6f6a-d6cf-4ee6-a9c4-86d18e7c397f"
)

// Get a new business central api instance pointing to the "production" environment.
func New(credentials auth.CredentialProvider) (*swagger.APIClient, error) {
	return NewWithEnvironment(credentials, productionEnvironment)
}

// Get a new api instance for the desired environment (usually "production" or "sandbox")
func NewWithEnvironment(credentials auth.CredentialProvider, environment string) (*swagger.APIClient, error) {
	return NewWithEnvironmentDebug(credentials, environment, false)
}

type debugRoundTripper struct {
	Transport http.RoundTripper
}

func (t *debugRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Printf("Executing %s %s:\n", req.Method, req.URL.String())
	res, err := t.Transport.RoundTrip(req)
	if err != nil {
		fmt.Printf("transport error: %v", err)
		return res, err
	}
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	res.Body.Close() //  must close
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Printf("Returned: %s\n", string(bodyBytes))
	return res, err
}

// Debugs logs all http request and responses, debug is currently not used
func NewWithEnvironmentDebug(credentials auth.CredentialProvider, environment string, debug bool) (*swagger.APIClient, error) {
	auth, err := auth.New(credentials)
	if err != nil {
		return nil, err
	}
	httpClient := auth.Client(context.Background())
	if debug {
		httpClient.Transport = &debugRoundTripper{Transport: httpClient.Transport}
	}
	config := swagger.NewConfiguration()
	config.HTTPClient = httpClient
	config.UserAgent = ""
	config.BasePath = fmt.Sprintf(endpointFmt, environment)

	return swagger.NewAPIClient(config), nil
}
