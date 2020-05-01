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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	defaultDir          = "/etc/secrets"
	defaultClientIdFile = "clientId"
	defaultUsernameFile = "username"
	defaultPasswordFile = "password"
	defaultTenantFile   = "password"
)

// The credentials to communicate with the bc API.
type CredentialProvider interface {
	// oauth clientID (see AAD app registry)
	ClientId() (string, error)
	// username of admin
	Username() (string, error)
	// password of admin
	Password() (string, error)
	// The tenant id (see AAD app registry)
	Tenant() (string, error)
}

/*
	By default we use a FileProvider. This allows us to mount a kubernetes secret into the container.
	The mounted volume will create two files containing the clientId and clientSecret
	The FileProvider is responsible for reading them.
*/
// CredentialProvider implementation for a dir containing two files, each containing the secret as a string.
type FileProvider struct {
	// The directory with the files, this is usually the place where you mount the secret volume
	Dir string // Should equal /etc/secrets

	// The filename within Dir containing the ClientId
	ClientIdFile string // Should equal "clientId"

	// The filename within Dir containing the username
	UsernameFile string // Should equal "username"

	// The filename within Dir containing the password
	PasswordFile string // Should equal "password"

	// The filename within Dir containing the AAD tenantId
	TenantFile string
}

func NewEnvProvider() *EnvironmentProvider {
	return &EnvironmentProvider{
		ClientIdKey: "BC_CLIENT_ID",
		UsernameKey: "BC_USERNAME",
		PasswordKey: "BC_PASSWORD",
		TenantKey:   "BC_TENANT_ID",
	}
}

func NewFileProvider() *FileProvider {
	return &FileProvider{Dir: defaultDir, ClientIdFile: defaultClientIdFile, UsernameFile: defaultUsernameFile, PasswordFile: defaultPasswordFile, TenantFile: defaultTenantFile}
}

func (p *FileProvider) ClientId() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.ClientIdFile))
	return string(data), err
}

func (p *FileProvider) Username() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.UsernameFile))
	return string(data), err
}

func (p *FileProvider) Password() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.PasswordFile))
	return string(data), err
}

func (p *FileProvider) Tenant() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.TenantFile))
	return string(data), err
}

// Simple raw value credential provider
type BasicProvider struct {
	ClientIdVal string
	UsernameVal string
	PasswordVal string
	TenantVal   string
}

func (p *BasicProvider) ClientId() (string, error) {
	return p.ClientIdVal, nil
}

func (p *BasicProvider) Username() (string, error) {
	return p.UsernameVal, nil
}
func (p *BasicProvider) Password() (string, error) {
	return p.PasswordVal, nil
}

func (p *BasicProvider) Tenant() (string, error) {
	return p.TenantVal, nil
}

// Fetches the credentials from the environment
type EnvironmentProvider struct {
	ClientIdKey string
	UsernameKey string
	PasswordKey string
	TenantKey   string
}

func (p *EnvironmentProvider) ClientId() (string, error) {
	clientId, present := os.LookupEnv(p.ClientIdKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.ClientIdKey)
	}
	return clientId, nil
}

func (p *EnvironmentProvider) Username() (string, error) {
	username, present := os.LookupEnv(p.UsernameKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.UsernameKey)
	}
	return username, nil
}

func (p *EnvironmentProvider) Password() (string, error) {
	password, present := os.LookupEnv(p.PasswordKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.PasswordKey)
	}
	return password, nil
}

func (p *EnvironmentProvider) Tenant() (string, error) {
	tenant, present := os.LookupEnv(p.TenantKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.TenantKey)
	}
	return tenant, nil
}
