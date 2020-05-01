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
	"context"
	"fmt"
	"github.com/tcorp-bv/business-central-api-go/auth"
	"github.com/tcorp-bv/business-central-api-go/openapi"
)

const (
	productionEnvironment = "production"
	endpointFmt           = "https://api.businesscentral.dynamics.com/v2.0/%s/api/v1.0"
	// The company ID for TCorp (really not so smart to hardcode this but whatever)
	TCorp = "253a6f6a-d6cf-4ee6-a9c4-86d18e7c397f"
)

// Get a new business central api instance pointing to the "production" environment.
func New(credentials auth.CredentialProvider) (*openapi.APIClient, error) {
	return NewWithEnvironment(credentials, productionEnvironment)
}

// Get a new api instance for the desired environment (usually "production" or "sandbox")
func NewWithEnvironment(credentials auth.CredentialProvider, environment string) (*openapi.APIClient, error) {
	return NewWithEnvironmentDebug(credentials, environment, false)
}

// Debugs logs all http request and responses
func NewWithEnvironmentDebug(credentials auth.CredentialProvider, environment string, debug bool) (*openapi.APIClient, error) {
	auth, err := auth.New(credentials)
	if err != nil {
		return nil, err
	}
	httpClient := auth.Client(context.Background())
	config := openapi.NewConfiguration()
	config.HTTPClient = httpClient
	config.UserAgent = ""
	config.BasePath = fmt.Sprintf(endpointFmt, environment)
	config.Debug = debug
	return openapi.NewAPIClient(config), nil
}
