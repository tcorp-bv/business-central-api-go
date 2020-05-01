# \SalesInvoiceLineApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesInvoiceLine**](SalesInvoiceLineApi.md#DeleteSalesInvoiceLine) | **Delete** /companies({company_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central
[**DeleteSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#DeleteSalesInvoiceLineForSalesInvoice) | **Delete** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central
[**GetSalesInvoiceLine**](SalesInvoiceLineApi.md#GetSalesInvoiceLine) | **Get** /companies({company_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.
[**GetSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#GetSalesInvoiceLineForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.
[**ListSalesInvoiceLines**](SalesInvoiceLineApi.md#ListSalesInvoiceLines) | **Get** /companies({company_id})/salesInvoiceLines | Returns a list of salesInvoiceLines
[**ListSalesInvoiceLinesForSalesInvoice**](SalesInvoiceLineApi.md#ListSalesInvoiceLinesForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines | Returns a list of salesInvoiceLines
[**PatchSalesInvoiceLine**](SalesInvoiceLineApi.md#PatchSalesInvoiceLine) | **Patch** /companies({company_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | Updates an object of type salesInvoiceLine in Dynamics 365 Business Central
[**PatchSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#PatchSalesInvoiceLineForSalesInvoice) | **Patch** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines(&#39;{salesInvoiceLine_id}&#39;) | Updates an object of type salesInvoiceLine in Dynamics 365 Business Central
[**PostSalesInvoiceLine**](SalesInvoiceLineApi.md#PostSalesInvoiceLine) | **Post** /companies({company_id})/salesInvoiceLines | Creates an object of type salesInvoiceLine in Dynamics 365 Business Central
[**PostSalesInvoiceLineForSalesInvoice**](SalesInvoiceLineApi.md#PostSalesInvoiceLineForSalesInvoice) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines | Creates an object of type salesInvoiceLine in Dynamics 365 Business Central



## DeleteSalesInvoiceLine

> DeleteSalesInvoiceLine(ctx, companyId, salesInvoiceLineId)

Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceLineId** | **string**| id for salesInvoiceLine | 

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


## DeleteSalesInvoiceLineForSalesInvoice

> DeleteSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, salesInvoiceLineId)

Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
**salesInvoiceLineId** | **string**| id for salesInvoiceLine | 

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


## GetSalesInvoiceLine

> SalesInvoiceLine GetSalesInvoiceLine(ctx, companyId, salesInvoiceLineId, optional)

Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
 **optional** | ***GetSalesInvoiceLineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSalesInvoiceLineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesInvoiceLineForSalesInvoice

> SalesInvoiceLine GetSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, salesInvoiceLineId, optional)

Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
**salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
 **optional** | ***GetSalesInvoiceLineForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSalesInvoiceLineForSalesInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSalesInvoiceLines

> InlineResponse2009 ListSalesInvoiceLines(ctx, companyId, optional)

Returns a list of salesInvoiceLines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesInvoiceLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSalesInvoiceLinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSalesInvoiceLinesForSalesInvoice

> InlineResponse2009 ListSalesInvoiceLinesForSalesInvoice(ctx, companyId, salesInvoiceId, optional)

Returns a list of salesInvoiceLines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***ListSalesInvoiceLinesForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSalesInvoiceLinesForSalesInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSalesInvoiceLine

> SalesInvoiceLine PatchSalesInvoiceLine(ctx, companyId, salesInvoiceLineId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSalesInvoiceLineForSalesInvoice

> SalesInvoiceLine PatchSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, salesInvoiceLineId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
**salesInvoiceLineId** | **string**| id for salesInvoiceLine | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesInvoiceLine

> SalesInvoiceLine PostSalesInvoiceLine(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesInvoiceLineForSalesInvoice

> SalesInvoiceLine PostSalesInvoiceLineForSalesInvoice(ctx, companyId, salesInvoiceId, contentType, uNKNOWNBASETYPE)

Creates an object of type salesInvoiceLine in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**SalesInvoiceLine**](salesInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

