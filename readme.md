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
api, err := NewWithEnvironmentDebug(auth.NewEnvProvider(), "production", true)
if err != nil {
	// handle error (eg environment variable missing)
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
github.com/tcorp-bv/business-central-api-go v1.0.1
```

## Generating the client 

To regenerate the swagger client ([openapi/](./openapi)) execute the following command:

```shell script
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/contracts/bcoas1.0.yaml -g go -o /local/openapi
```


## Business Central authentication
As we've wasted a lot of time getting to authenticate with the oauth for business central,
I'll document the process here.

The API uses delegated AAD permissions. This means you must actually authenticate as a user.
https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/Because microsoft does not allow the creation of client secrets for this, you need to use the oauth "password" grant_type.
This means the application will actually manage your username and password.

The app registry requires the `https://dynamics.microsoft.com/business-central/overview/user_impersonation` permission.


## Copyright notice
Files in the [openapi](./openapi) directory are generated from the [business central api specification](https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/contracts/bcoas1.0.yaml).
TCorp BV does not hold any rights over this specification and is not associated with Business Central.