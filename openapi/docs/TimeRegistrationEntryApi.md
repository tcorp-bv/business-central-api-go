# \TimeRegistrationEntryApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeRegistrationEntry**](TimeRegistrationEntryApi.md#DeleteTimeRegistrationEntry) | **Delete** /companies({company_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | Deletes an object of type timeRegistrationEntry in Dynamics 365 Business Central
[**DeleteTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#DeleteTimeRegistrationEntryForEmployee) | **Delete** /companies({company_id})/employees({employee_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | Deletes an object of type timeRegistrationEntry in Dynamics 365 Business Central
[**GetTimeRegistrationEntry**](TimeRegistrationEntryApi.md#GetTimeRegistrationEntry) | **Get** /companies({company_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | Retrieve the properties and relationships of an object of type timeRegistrationEntry for Dynamics 365 Business Central.
[**GetTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#GetTimeRegistrationEntryForEmployee) | **Get** /companies({company_id})/employees({employee_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | Retrieve the properties and relationships of an object of type timeRegistrationEntry for Dynamics 365 Business Central.
[**ListTimeRegistrationEntries**](TimeRegistrationEntryApi.md#ListTimeRegistrationEntries) | **Get** /companies({company_id})/timeRegistrationEntries | Returns a list of timeRegistrationEntries
[**ListTimeRegistrationEntriesForEmployee**](TimeRegistrationEntryApi.md#ListTimeRegistrationEntriesForEmployee) | **Get** /companies({company_id})/employees({employee_id})/timeRegistrationEntries | Returns a list of timeRegistrationEntries
[**PatchTimeRegistrationEntry**](TimeRegistrationEntryApi.md#PatchTimeRegistrationEntry) | **Patch** /companies({company_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | Updates an object of type timeRegistrationEntry in Dynamics 365 Business Central
[**PatchTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#PatchTimeRegistrationEntryForEmployee) | **Patch** /companies({company_id})/employees({employee_id})/timeRegistrationEntries({timeRegistrationEntry_id}) | Updates an object of type timeRegistrationEntry in Dynamics 365 Business Central
[**PostTimeRegistrationEntry**](TimeRegistrationEntryApi.md#PostTimeRegistrationEntry) | **Post** /companies({company_id})/timeRegistrationEntries | Creates an object of type timeRegistrationEntry in Dynamics 365 Business Central
[**PostTimeRegistrationEntryForEmployee**](TimeRegistrationEntryApi.md#PostTimeRegistrationEntryForEmployee) | **Post** /companies({company_id})/employees({employee_id})/timeRegistrationEntries | Creates an object of type timeRegistrationEntry in Dynamics 365 Business Central



## DeleteTimeRegistrationEntry

> DeleteTimeRegistrationEntry(ctx, companyId, timeRegistrationEntryId)

Deletes an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTimeRegistrationEntryForEmployee

> DeleteTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, timeRegistrationEntryId)

Deletes an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**employeeId** | [**string**](.md)| id for employee | 
**timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeRegistrationEntry

> TimeRegistrationEntry GetTimeRegistrationEntry(ctx, companyId, timeRegistrationEntryId, optional)

Retrieve the properties and relationships of an object of type timeRegistrationEntry for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
 **optional** | ***GetTimeRegistrationEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTimeRegistrationEntryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeRegistrationEntryForEmployee

> TimeRegistrationEntry GetTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, timeRegistrationEntryId, optional)

Retrieve the properties and relationships of an object of type timeRegistrationEntry for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**employeeId** | [**string**](.md)| id for employee | 
**timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
 **optional** | ***GetTimeRegistrationEntryForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTimeRegistrationEntryForEmployeeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimeRegistrationEntries

> InlineResponse20019 ListTimeRegistrationEntries(ctx, companyId, optional)

Returns a list of timeRegistrationEntries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListTimeRegistrationEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTimeRegistrationEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimeRegistrationEntriesForEmployee

> InlineResponse20019 ListTimeRegistrationEntriesForEmployee(ctx, companyId, employeeId, optional)

Returns a list of timeRegistrationEntries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**employeeId** | [**string**](.md)| id for employee | 
 **optional** | ***ListTimeRegistrationEntriesForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTimeRegistrationEntriesForEmployeeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTimeRegistrationEntry

> TimeRegistrationEntry PatchTimeRegistrationEntry(ctx, companyId, timeRegistrationEntryId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTimeRegistrationEntryForEmployee

> TimeRegistrationEntry PatchTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, timeRegistrationEntryId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**employeeId** | [**string**](.md)| id for employee | 
**timeRegistrationEntryId** | [**string**](.md)| id for timeRegistrationEntry | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeRegistrationEntry

> TimeRegistrationEntry PostTimeRegistrationEntry(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeRegistrationEntryForEmployee

> TimeRegistrationEntry PostTimeRegistrationEntryForEmployee(ctx, companyId, employeeId, contentType, uNKNOWNBASETYPE)

Creates an object of type timeRegistrationEntry in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**employeeId** | [**string**](.md)| id for employee | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**TimeRegistrationEntry**](timeRegistrationEntry.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

