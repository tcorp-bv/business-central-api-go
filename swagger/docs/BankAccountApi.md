# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBankAccount**](BankAccountApi.md#DeleteBankAccount) | **Delete** /companies({company_id})/bankAccounts({bankAccount_id}) | Deletes an object of type bankAccount in Dynamics 365 Business Central
[**GetBankAccount**](BankAccountApi.md#GetBankAccount) | **Get** /companies({company_id})/bankAccounts({bankAccount_id}) | Retrieve the properties and relationships of an object of type bankAccount for Dynamics 365 Business Central.
[**ListBankAccounts**](BankAccountApi.md#ListBankAccounts) | **Get** /companies({company_id})/bankAccounts | Returns a list of bankAccounts
[**PatchBankAccount**](BankAccountApi.md#PatchBankAccount) | **Patch** /companies({company_id})/bankAccounts({bankAccount_id}) | Updates an object of type bankAccount in Dynamics 365 Business Central
[**PostBankAccount**](BankAccountApi.md#PostBankAccount) | **Post** /companies({company_id})/bankAccounts | Creates an object of type bankAccount in Dynamics 365 Business Central

# **DeleteBankAccount**
> DeleteBankAccount(ctx, companyId, bankAccountId)
Deletes an object of type bankAccount in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **bankAccountId** | [**string**](.md)| id for bankAccount | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAccount**
> BankAccount GetBankAccount(ctx, companyId, bankAccountId, optional)
Retrieve the properties and relationships of an object of type bankAccount for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **bankAccountId** | [**string**](.md)| id for bankAccount | 
 **optional** | ***BankAccountApiGetBankAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BankAccountApiGetBankAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**BankAccount**](bankAccount.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBankAccounts**
> InlineResponse2006 ListBankAccounts(ctx, companyId, optional)
Returns a list of bankAccounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***BankAccountApiListBankAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BankAccountApiListBankAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchBankAccount**
> BankAccount PatchBankAccount(ctx, body, companyId, bankAccountId, contentType, ifMatch)
Updates an object of type bankAccount in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **bankAccountId** | [**string**](.md)| id for bankAccount | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**BankAccount**](bankAccount.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBankAccount**
> BankAccount PostBankAccount(ctx, body, companyId, contentType)
Creates an object of type bankAccount in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**BankAccount**](bankAccount.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

