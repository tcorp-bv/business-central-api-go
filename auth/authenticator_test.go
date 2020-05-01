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

package auth

import (
	"testing"
)

const (
	testClientId = "oRNWbHFXtAECmhnZmEndcjLIaSKbRMVE"
	testUsername = "someUsername@onmicrosoft.com"
	testPassword = "somepassword"
	testEndpoint = "https://login.onmicrosoft.com/mytentant/oauth2/v2.0/token"
)

func TestGetConfig(t *testing.T) {
	c, err := getConfig(&BasicProvider{ClientIdVal: testClientId, UsernameVal: testUsername, PasswordVal: testPassword}, testEndpoint)
	if err != nil {
		t.Error(err)
	}
	if c.ClientID != testClientId || c.EndpointParams["username"][0] != testUsername || c.EndpointParams["password"][0] != testPassword || c.TokenURL != testEndpoint {
		t.Error("Client credentials not properly parsed")
	}
}
