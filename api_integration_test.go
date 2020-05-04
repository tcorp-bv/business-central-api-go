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
	"fmt"
	"github.com/antihax/optional"
	"github.com/matryer/is"
	"github.com/tcorp-bv/business-central-api-go/auth"
	"github.com/tcorp-bv/business-central-api-go/swagger"
	"testing"
	"time"
)

const (
	testEmployeeNumber = "TESTING"
	testCustomerNumber = "BOLTEST123456"
)

func TestGetCompanies(t *testing.T) {
	is := is.New(t)
	api := getAPI(is)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, _, err := api.CompanyApi.ListCompanies(ctx, &swagger.CompanyApiListCompaniesOpts{})
	is.NoErr(err)
	is.True(len(res.Value) > 0)
	for _, c := range res.Value {
		is.True(c.Id != "")
		is.True(c.Name != "")
	}
}
func TestPostCustomer(t *testing.T) {
	is := is.New(t)
	api := getAPI(is)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	compId := getSomeCompanyId(is, api)

	// Clean up the customer if they already exist
	filter := fmt.Sprintf("number eq '%s'", testCustomerNumber)
	res, _, err := api.CustomerApi.ListCustomers(ctx, compId, &swagger.CustomerApiListCustomersOpts{Filter: optional.NewString(filter)})
	for _, c := range res.Value {
		api.CustomerApi.DeleteCustomer(ctx, compId, c.Id) // we ignore the result of this, as it is just a cleanup
	}
	customer := swagger.Customer{
		Number:      testCustomerNumber,
		DisplayName: "Test Customer",
		Email:       "asd@asd.be",
		PhoneNumber: "0412320312",
		Type_:       "Person",
	}
	// Post the customer again
	c, _, err := api.CustomerApi.PostCustomer(ctx, customer, "application/json", compId)
	customerId := c.Id
	defer func() { // cleanup
		_, err := api.CustomerApi.DeleteCustomer(ctx, compId, customerId)
		is.NoErr(err)
	}()
	is.NoErr(err)
	fetchedCustomer, _, err := api.CustomerApi.GetCustomer(ctx, compId, customerId, &swagger.CustomerApiGetCustomerOpts{})
	is.NoErr(err)
	is.Equal(customer.Number, fetchedCustomer.Number)
	is.Equal(customer.DisplayName, fetchedCustomer.DisplayName)
	is.Equal(customer.Email, fetchedCustomer.Email)
	is.Equal(customer.PhoneNumber, fetchedCustomer.PhoneNumber)
	is.Equal(customer.Type_, fetchedCustomer.Type_)
}

type EmployeeWrapper struct {
	swagger.Employee
	EmploymentDate  string `json:"employmentDate,omitempty"`
	BirthDate       string `json:"birthDate,omitempty"`
	TerminationDate string `json:"terminationDate,omitempty"`
}

func TestEmploymentDateField(t *testing.T) {
	is := is.New(t)
	api := getAPI(is)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	compId := getSomeCompanyId(is, api)
	item := swagger.Employee{Number: testEmployeeNumber, GivenName: "Test", Surname: "Employee", EmploymentDate: "2020-05-01"}
	cleanupEmployee(is, ctx, api, compId)
	emp, _, err := api.EmployeeApi.PostEmployee(ctx, item, DefaultContentType, compId)
	is.NoErr(err) // we want to make sure that employmentdate is postable
	defer cleanupEmployee(is, ctx, api, compId)

	resEmp, _, err := api.EmployeeApi.GetEmployee(ctx, compId, emp.Id, &swagger.EmployeeApiGetEmployeeOpts{})
	is.NoErr(err)
	is.Equal(resEmp.Number, testEmployeeNumber)
	is.Equal(resEmp.GivenName, "Test")
	is.Equal(resEmp.Surname, "Employee")
	is.Equal(resEmp.EmploymentDate, "2020-05-01")

}
func cleanupEmployee(is *is.I, ctx context.Context, api *swagger.APIClient, companyId string) {
	filter := optional.NewString(fmt.Sprintf("number eq '%s'", testEmployeeNumber))
	res, _, err := api.EmployeeApi.ListEmployees(ctx, companyId, &swagger.EmployeeApiListEmployeesOpts{Filter: filter})
	is.NoErr(err)
	for _, emp := range res.Value {
		api.EmployeeApi.DeleteEmployee(ctx, companyId, emp.Id) // Ignore result as this is just a cleanup
	}
}
func getAPI(is *is.I) *swagger.APIClient {
	api, err := NewWithEnvironmentDebug(auth.NewEnvProvider(), "sandbox", true)
	is.NoErr(err)
	return api
}
func getSomeCompanyId(is *is.I, api *swagger.APIClient) string {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, _, err := api.CompanyApi.ListCompanies(ctx, &swagger.CompanyApiListCompaniesOpts{})
	is.NoErr(err)
	is.True(len(res.Value) > 0)
	id := res.Value[0].Id
	is.True(id != TCorp)
	return id
}
