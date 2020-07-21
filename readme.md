# Business central go api
Generated client for the Microsoft Dynamics [Business Central API](https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/).

## Getting started
Using `auth.NewEnvProvider()` requires the following environment variables:
* `BC_CLIENT_ID`: The Azure active directory app ClientId (app has delegated user BC permission)
* `BC_USERNAME`: Your microsoft username (email), this user should be a BC admin
* `BC_PASSWORD`: Your microsoft password
* `BC_TENANT_ID`: The azure active directory app TenantId (should be same as your bc tenantId)


Setup the api once:
```go
import (
	businesscentral "github.com/tcorp-bv/business-central-api-go"
	"github.com/tcorp-bv/business-central-api-go/auth"
)

func main() {
    api, err := businesscentral.NewWithEnvironmentDebug(auth.NewEnvProvider())
    if err != nil {
	    // handle error (eg environment variable missing)
    }
}
```

Example of calling the api:
```go
ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
res, _,  err := api.CompanyApi.ListCompanies(ctx, &openapi.ListCompaniesOpts{})
if err != nil {
	// handle err
}
for _, c := range res.Value {
	fmt.Printf("Company %s.\n", c*c.Name)
}
```

and in your `go.mod`:
```shell script
github.com/tcorp-bv/business-central-api-go v1.6.0
```

## Documentation
To see many examples of how to use the API, please see our [integration tests](api_integration_test.go).


### Creating a new entity
Because the API is generated, any POST operation with a body takes a `body interface{}` parameter.
You can simply pass the expected resource in here.

To create a customer for example:
```go
customer := swagger.Customer{
	Number: testCustomerNumber,
	DisplayName: "Test Customer",
	Email: "asd@asd.be",
	PhoneNumber: "0412320312",
	Type_: "Person",
}
// You also need to provide "application/json" as the Content-Type
c, _, err := api.CustomerApi.PostCustomer(ctx, customer, "application/json", compId)
```
### Date format
String dates should be formatted as `YYYY-MM-DD`.

### Filtering
The API resources (invoices, customers, companies...) use their own unique identifier (different from the number identifier).
This identifier is called **Id**. 

Usually you want to find a resource by some other identifier (customer number...).
To do that, you will need to use the list operation with a filter.

Here is an example:
```go
filter := fmt.Sprintf("number eq '%s'", testCustomerNumber)
res, _, err := api.CustomerApi.ListCustomers(ctx, compId, &swagger.CustomerApiListCustomersOpts{Filter: optional.NewString(filter)})
for _, c := range res.Value {
	api.CustomerApi.DeleteCustomer(ctx, compId, c.Id) // we ignore the result of this, as it is just a cleanup
}
```

For more information on all the possible filters, see the [Microsoft API guidelines](https://github.com/Microsoft/api-guidelines/blob/master/Guidelines.md#97-filtering).

## Generating the client 

To regenerate the swagger client ([swagger/](./swagger)) execute the following command:

```shell script
make generate
```


## Business Central authentication
As we've wasted a lot of time getting to authenticate with the oauth for business central,
I'll document the process here.

The API uses delegated AAD permissions. This means you must actually authenticate as a user.
https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/Because microsoft does not allow the creation of client secrets for this, you need to use the oauth "password" grant_type.
This means the application will actually manage your username and password.

The app registry requires the `https://dynamics.microsoft.com/business-central/overview/user_impersonation` permission.


## Copyright notice
Files in the [swagger](./swagger) directory are generated from the [business central api specification](https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/contracts/bcoas1.0.yaml).
TCorp BV does not hold any rights over this specification and is not associated with Business Central.