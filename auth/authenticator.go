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
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
)

const (
	defaultAuthUrl = "https://login.microsoft.net/{tenant}/oauth2/v2.0/token"
)

// The authenticator manages the oauth authentication to the bol API
type Authenticator interface {
	// Returns the bearer token or an error
	Token() (*oauth2.Token, error)
	Client(ctx context.Context) *http.Client
}

// Creates a new instance of the Authenticator using the credentials from the provider.
func New(provider CredentialProvider) (Authenticator, error) {
	tenant, err := provider.Tenant()
	if err != nil {
		return nil, err
	}
	tokenEndpoint := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", tenant)
	return NewWithTokenEndpoint(provider, tokenEndpoint)
}

func NewWithTokenEndpoint(provider CredentialProvider, tokenEndpoint string) (Authenticator, error) {
	config, err := getConfig(provider, tokenEndpoint)
	if err != nil {
		return nil, err
	}
	return &authenticator{config: config}, nil
}

// Authenticator implementation using golang's oauth2 client.
type authenticator struct {
	config clientcredentials.Config
}

// Fetch Token from the golang oauth2 client
func (a *authenticator) Token() (*oauth2.Token, error) {
	return a.config.Token(context.Background())
}
func (a *authenticator) Client(ctx context.Context) *http.Client {
	return a.config.Client(ctx)
}

// Get config for the golang oauth2 client.
func getConfig(provider CredentialProvider, tokenEndpoint string) (clientcredentials.Config, error) {
	// Get the credentials from the provider
	clientId, err := provider.ClientId()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	username, err := provider.Username()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	password, err := provider.Password()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	// Setup the config
	return clientcredentials.Config{
		ClientID: clientId,
		TokenURL: tokenEndpoint,
		EndpointParams: url.Values{
			"grant_type": {"password"},
			"username":   {username},
			"password":   {password},
			"resource":   []string{"https://api.businesscentral.dynamics.com"},
		},
	}, nil
}
