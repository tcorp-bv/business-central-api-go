# \CustomerPaymentJournalApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomerPaymentJournal**](CustomerPaymentJournalApi.md#DeleteCustomerPaymentJournal) | **Delete** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id}) | Deletes an object of type customerPaymentJournal in Dynamics 365 Business Central
[**GetCustomerPaymentJournal**](CustomerPaymentJournalApi.md#GetCustomerPaymentJournal) | **Get** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id}) | Retrieve the properties and relationships of an object of type customerPaymentJournal for Dynamics 365 Business Central.
[**ListCustomerPaymentJournals**](CustomerPaymentJournalApi.md#ListCustomerPaymentJournals) | **Get** /companies({company_id})/customerPaymentJournals | Returns a list of customerPaymentJournals
[**PatchCustomerPaymentJournal**](CustomerPaymentJournalApi.md#PatchCustomerPaymentJournal) | **Patch** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id}) | Updates an object of type customerPaymentJournal in Dynamics 365 Business Central
[**PostCustomerPaymentJournal**](CustomerPaymentJournalApi.md#PostCustomerPaymentJournal) | **Post** /companies({company_id})/customerPaymentJournals | Creates an object of type customerPaymentJournal in Dynamics 365 Business Central



## DeleteCustomerPaymentJournal

> DeleteCustomerPaymentJournal(ctx, companyId, customerPaymentJournalId)

Deletes an object of type customerPaymentJournal in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 

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


## GetCustomerPaymentJournal

> CustomerPaymentJournal GetCustomerPaymentJournal(ctx, companyId, customerPaymentJournalId, optional)

Retrieve the properties and relationships of an object of type customerPaymentJournal for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
 **optional** | ***GetCustomerPaymentJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerPaymentJournalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerPaymentJournal**](customerPaymentJournal.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerPaymentJournals

> InlineResponse20011 ListCustomerPaymentJournals(ctx, companyId, optional)

Returns a list of customerPaymentJournals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListCustomerPaymentJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCustomerPaymentJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCustomerPaymentJournal

> CustomerPaymentJournal PatchCustomerPaymentJournal(ctx, companyId, customerPaymentJournalId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type customerPaymentJournal in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**CustomerPaymentJournal**](customerPaymentJournal.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCustomerPaymentJournal

> CustomerPaymentJournal PostCustomerPaymentJournal(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type customerPaymentJournal in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**CustomerPaymentJournal**](customerPaymentJournal.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

