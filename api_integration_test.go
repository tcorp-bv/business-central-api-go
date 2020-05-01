//+build integration

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
	"github.com/tcorp-bv/business-central-api-go/auth"
	"github.com/tcorp-bv/business-central-api-go/openapi"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	credentials := auth.NewEnvProvider()
	api, err := NewWithEnvironmentDebug(&credentials, "production", true)
	if err != nil {
		t.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, _, err := api.CompanyApi.ListCompanies(ctx, &openapi.ListCompaniesOpts{})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Value) == 0 {
		t.Error("received empty company list")
	}
	for _, c := range res.Value {
		if c.Id == "" {
			t.Error("received empty company id")
		}
		if c.Name == nil || *c.Name == "" {
			t.Error("received empty company name")
		}
	}
}
