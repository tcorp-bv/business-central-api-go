# \AgedAccountsPayableApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgedAccountsPayable**](AgedAccountsPayableApi.md#GetAgedAccountsPayable) | **Get** /companies({company_id})/agedAccountsPayable({agedAccountsPayable_vendorId}) | Retrieve the properties and relationships of an object of type agedAccountsPayable for Dynamics 365 Business Central.
[**ListAgedAccountsPayable**](AgedAccountsPayableApi.md#ListAgedAccountsPayable) | **Get** /companies({company_id})/agedAccountsPayable | Returns a list of agedAccountsPayable



## GetAgedAccountsPayable

> AgedAccountsPayable GetAgedAccountsPayable(ctx, companyId, agedAccountsPayableVendorId, optional)

Retrieve the properties and relationships of an object of type agedAccountsPayable for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**agedAccountsPayableVendorId** | [**string**](.md)| vendorId for agedAccountsPayable | 
 **optional** | ***GetAgedAccountsPayableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAgedAccountsPayableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**AgedAccountsPayable**](agedAccountsPayable.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgedAccountsPayable

> InlineResponse20036 ListAgedAccountsPayable(ctx, companyId, optional)

Returns a list of agedAccountsPayable

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListAgedAccountsPayableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAgedAccountsPayableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

