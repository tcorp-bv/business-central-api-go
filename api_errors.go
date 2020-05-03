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
	"encoding/json"
	"fmt"
	"github.com/tcorp-bv/business-central-api-go/swagger"
)

type APIErrorJson struct {
	Error *APIError `json:"error,omitempty"`
}
type APIError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	ErrorCodeInternal_EntityWithSameKeyExists = "Internal_EntityWithSameKeyExists"
)

// Check if a golang error is an Internal_EntityWithSameKeyExists error.
func AlreadyExists(toParse error) bool {
	apiErr, err := AsAPIError(toParse)
	if err != nil {
		return false
	}
	return apiErr.Code == "ErrorCodeInternal_EntityWithSameKeyExists"
}

// Parse a golang error to an APIError
func AsAPIError(toParse error) (APIError, error) {
	parsed := toParse.(swagger.GenericSwaggerError)
	var res APIErrorJson
	err := json.Unmarshal(parsed.Body(), &res)
	if err != nil {
		return APIError{}, err
	}
	if res.Error == nil {
		return APIError{}, fmt.Errorf("error did not contain an 'error' field")
	}
	return *res.Error, nil
}
