# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesQuoteLine**](SalesQuoteLineApi.md#DeleteSalesQuoteLine) | **Delete** /companies({company_id})/salesQuoteLines(&#x27;{salesQuoteLine_id}&#x27;) | Deletes an object of type salesQuoteLine in Dynamics 365 Business Central
[**DeleteSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#DeleteSalesQuoteLineForSalesQuote) | **Delete** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines(&#x27;{salesQuoteLine_id}&#x27;) | Deletes an object of type salesQuoteLine in Dynamics 365 Business Central
[**GetSalesQuoteLine**](SalesQuoteLineApi.md#GetSalesQuoteLine) | **Get** /companies({company_id})/salesQuoteLines(&#x27;{salesQuoteLine_id}&#x27;) | Retrieve the properties and relationships of an object of type salesQuoteLine for Dynamics 365 Business Central.
[**GetSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#GetSalesQuoteLineForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines(&#x27;{salesQuoteLine_id}&#x27;) | Retrieve the properties and relationships of an object of type salesQuoteLine for Dynamics 365 Business Central.
[**ListSalesQuoteLines**](SalesQuoteLineApi.md#ListSalesQuoteLines) | **Get** /companies({company_id})/salesQuoteLines | Returns a list of salesQuoteLines
[**ListSalesQuoteLinesForSalesQuote**](SalesQuoteLineApi.md#ListSalesQuoteLinesForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines | Returns a list of salesQuoteLines
[**PatchSalesQuoteLine**](SalesQuoteLineApi.md#PatchSalesQuoteLine) | **Patch** /companies({company_id})/salesQuoteLines(&#x27;{salesQuoteLine_id}&#x27;) | Updates an object of type salesQuoteLine in Dynamics 365 Business Central
[**PatchSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#PatchSalesQuoteLineForSalesQuote) | **Patch** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines(&#x27;{salesQuoteLine_id}&#x27;) | Updates an object of type salesQuoteLine in Dynamics 365 Business Central
[**PostSalesQuoteLine**](SalesQuoteLineApi.md#PostSalesQuoteLine) | **Post** /companies({company_id})/salesQuoteLines | Creates an object of type salesQuoteLine in Dynamics 365 Business Central
[**PostSalesQuoteLineForSalesQuote**](SalesQuoteLineApi.md#PostSalesQuoteLineForSalesQuote) | **Post** /companies({company_id})/salesQuotes({salesQuote_id})/salesQuoteLines | Creates an object of type salesQuoteLine in Dynamics 365 Business Central

# **DeleteSalesQuoteLine**
> DeleteSalesQuoteLine(ctx, companyId, salesQuoteLineId)
Deletes an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesQuoteLineForSalesQuote**
> DeleteSalesQuoteLineForSalesQuote(ctx, companyId, salesQuoteId, salesQuoteLineId)
Deletes an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesQuoteLine**
> SalesQuoteLine GetSalesQuoteLine(ctx, companyId, salesQuoteLineId, optional)
Retrieve the properties and relationships of an object of type salesQuoteLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
 **optional** | ***SalesQuoteLineApiGetSalesQuoteLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesQuoteLineApiGetSalesQuoteLineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesQuoteLineForSalesQuote**
> SalesQuoteLine GetSalesQuoteLineForSalesQuote(ctx, companyId, salesQuoteId, salesQuoteLineId, optional)
Retrieve the properties and relationships of an object of type salesQuoteLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
 **optional** | ***SalesQuoteLineApiGetSalesQuoteLineForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesQuoteLineApiGetSalesQuoteLineForSalesQuoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesQuoteLines**
> InlineResponse20043 ListSalesQuoteLines(ctx, companyId, optional)
Returns a list of salesQuoteLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***SalesQuoteLineApiListSalesQuoteLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesQuoteLineApiListSalesQuoteLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesQuoteLinesForSalesQuote**
> InlineResponse20043 ListSalesQuoteLinesForSalesQuote(ctx, companyId, salesQuoteId, optional)
Returns a list of salesQuoteLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***SalesQuoteLineApiListSalesQuoteLinesForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesQuoteLineApiListSalesQuoteLinesForSalesQuoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesQuoteLine**
> SalesQuoteLine PatchSalesQuoteLine(ctx, body, companyId, salesQuoteLineId, contentType, ifMatch)
Updates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesQuoteLineForSalesQuote**
> SalesQuoteLine PatchSalesQuoteLineForSalesQuote(ctx, body, companyId, salesQuoteId, salesQuoteLineId, contentType, ifMatch)
Updates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **salesQuoteLineId** | **string**| id for salesQuoteLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesQuoteLine**
> SalesQuoteLine PostSalesQuoteLine(ctx, body, companyId, contentType)
Creates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesQuoteLineForSalesQuote**
> SalesQuoteLine PostSalesQuoteLineForSalesQuote(ctx, body, companyId, salesQuoteId, contentType)
Creates an object of type salesQuoteLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesQuoteLine**](salesQuoteLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

