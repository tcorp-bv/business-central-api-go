# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgedAccountsReceivable**](AgedAccountsReceivableApi.md#GetAgedAccountsReceivable) | **Get** /companies({company_id})/agedAccountsReceivable({agedAccountsReceivable_customerId}) | Retrieve the properties and relationships of an object of type agedAccountsReceivable for Dynamics 365 Business Central.
[**ListAgedAccountsReceivable**](AgedAccountsReceivableApi.md#ListAgedAccountsReceivable) | **Get** /companies({company_id})/agedAccountsReceivable | Returns a list of agedAccountsReceivable

# **GetAgedAccountsReceivable**
> AgedAccountsReceivable GetAgedAccountsReceivable(ctx, companyId, agedAccountsReceivableCustomerId, optional)
Retrieve the properties and relationships of an object of type agedAccountsReceivable for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **agedAccountsReceivableCustomerId** | [**string**](.md)| customerId for agedAccountsReceivable | 
 **optional** | ***AgedAccountsReceivableApiGetAgedAccountsReceivableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgedAccountsReceivableApiGetAgedAccountsReceivableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**AgedAccountsReceivable**](agedAccountsReceivable.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAgedAccountsReceivable**
> InlineResponse2003 ListAgedAccountsReceivable(ctx, companyId, optional)
Returns a list of agedAccountsReceivable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***AgedAccountsReceivableApiListAgedAccountsReceivableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgedAccountsReceivableApiListAgedAccountsReceivableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

