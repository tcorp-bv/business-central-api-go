# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesQuote**](SalesQuoteApi.md#DeleteSalesQuote) | **Delete** /companies({company_id})/salesQuotes({salesQuote_id}) | Deletes an object of type salesQuote in Dynamics 365 Business Central
[**GetSalesQuote**](SalesQuoteApi.md#GetSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id}) | Retrieve the properties and relationships of an object of type salesQuote for Dynamics 365 Business Central.
[**ListSalesQuotes**](SalesQuoteApi.md#ListSalesQuotes) | **Get** /companies({company_id})/salesQuotes | Returns a list of salesQuotes
[**MakeInvoiceActionSalesQuotes**](SalesQuoteApi.md#MakeInvoiceActionSalesQuotes) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/Microsoft.NAV.makeInvoice | Performs the makeInvoice action for salesQuotes entity
[**MakeOrderActionSalesQuotes**](SalesQuoteApi.md#MakeOrderActionSalesQuotes) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/Microsoft.NAV.makeOrder | Performs the makeOrder action for salesQuotes entity
[**PatchSalesQuote**](SalesQuoteApi.md#PatchSalesQuote) | **Patch** /companies({company_id})/salesQuotes({salesQuote_id}) | Updates an object of type salesQuote in Dynamics 365 Business Central
[**PostSalesQuote**](SalesQuoteApi.md#PostSalesQuote) | **Post** /companies({company_id})/salesQuotes | Creates an object of type salesQuote in Dynamics 365 Business Central
[**SendActionSalesQuotes**](SalesQuoteApi.md#SendActionSalesQuotes) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/Microsoft.NAV.send | Performs the send action for salesQuotes entity

# **DeleteSalesQuote**
> DeleteSalesQuote(ctx, companyId, salesQuoteId)
Deletes an object of type salesQuote in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesQuote**
> SalesQuote GetSalesQuote(ctx, companyId, salesQuoteId, optional)
Retrieve the properties and relationships of an object of type salesQuote for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***SalesQuoteApiGetSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesQuoteApiGetSalesQuoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesQuote**](salesQuote.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesQuotes**
> InlineResponse20041 ListSalesQuotes(ctx, companyId, optional)
Returns a list of salesQuotes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***SalesQuoteApiListSalesQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesQuoteApiListSalesQuotesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeInvoiceActionSalesQuotes**
> MakeInvoiceActionSalesQuotes(ctx, companyId, salesQuoteId)
Performs the makeInvoice action for salesQuotes entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeOrderActionSalesQuotes**
> MakeOrderActionSalesQuotes(ctx, companyId, salesQuoteId)
Performs the makeOrder action for salesQuotes entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesQuote**
> SalesQuote PatchSalesQuote(ctx, body, companyId, salesQuoteId, contentType, ifMatch)
Updates an object of type salesQuote in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesQuote**](salesQuote.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesQuote**
> SalesQuote PostSalesQuote(ctx, body, companyId, contentType)
Creates an object of type salesQuote in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesQuote**](salesQuote.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendActionSalesQuotes**
> SendActionSalesQuotes(ctx, companyId, salesQuoteId)
Performs the send action for salesQuotes entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

