# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#DeletePurchaseInvoiceLine) | **Delete** /companies({company_id})/purchaseInvoiceLines(&#x27;{purchaseInvoiceLine_id}&#x27;) | Deletes an object of type purchaseInvoiceLine in Dynamics 365 Business Central
[**DeletePurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#DeletePurchaseInvoiceLineForPurchaseInvoice) | **Delete** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines(&#x27;{purchaseInvoiceLine_id}&#x27;) | Deletes an object of type purchaseInvoiceLine in Dynamics 365 Business Central
[**GetPurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#GetPurchaseInvoiceLine) | **Get** /companies({company_id})/purchaseInvoiceLines(&#x27;{purchaseInvoiceLine_id}&#x27;) | Retrieve the properties and relationships of an object of type purchaseInvoiceLine for Dynamics 365 Business Central.
[**GetPurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#GetPurchaseInvoiceLineForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines(&#x27;{purchaseInvoiceLine_id}&#x27;) | Retrieve the properties and relationships of an object of type purchaseInvoiceLine for Dynamics 365 Business Central.
[**ListPurchaseInvoiceLines**](PurchaseInvoiceLineApi.md#ListPurchaseInvoiceLines) | **Get** /companies({company_id})/purchaseInvoiceLines | Returns a list of purchaseInvoiceLines
[**ListPurchaseInvoiceLinesForPurchaseInvoice**](PurchaseInvoiceLineApi.md#ListPurchaseInvoiceLinesForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines | Returns a list of purchaseInvoiceLines
[**PatchPurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#PatchPurchaseInvoiceLine) | **Patch** /companies({company_id})/purchaseInvoiceLines(&#x27;{purchaseInvoiceLine_id}&#x27;) | Updates an object of type purchaseInvoiceLine in Dynamics 365 Business Central
[**PatchPurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#PatchPurchaseInvoiceLineForPurchaseInvoice) | **Patch** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines(&#x27;{purchaseInvoiceLine_id}&#x27;) | Updates an object of type purchaseInvoiceLine in Dynamics 365 Business Central
[**PostPurchaseInvoiceLine**](PurchaseInvoiceLineApi.md#PostPurchaseInvoiceLine) | **Post** /companies({company_id})/purchaseInvoiceLines | Creates an object of type purchaseInvoiceLine in Dynamics 365 Business Central
[**PostPurchaseInvoiceLineForPurchaseInvoice**](PurchaseInvoiceLineApi.md#PostPurchaseInvoiceLineForPurchaseInvoice) | **Post** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/purchaseInvoiceLines | Creates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

# **DeletePurchaseInvoiceLine**
> DeletePurchaseInvoiceLine(ctx, companyId, purchaseInvoiceLineId)
Deletes an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePurchaseInvoiceLineForPurchaseInvoice**
> DeletePurchaseInvoiceLineForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, purchaseInvoiceLineId)
Deletes an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchaseInvoiceLine**
> PurchaseInvoiceLine GetPurchaseInvoiceLine(ctx, companyId, purchaseInvoiceLineId, optional)
Retrieve the properties and relationships of an object of type purchaseInvoiceLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
 **optional** | ***PurchaseInvoiceLineApiGetPurchaseInvoiceLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseInvoiceLineApiGetPurchaseInvoiceLineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchaseInvoiceLineForPurchaseInvoice**
> PurchaseInvoiceLine GetPurchaseInvoiceLineForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, purchaseInvoiceLineId, optional)
Retrieve the properties and relationships of an object of type purchaseInvoiceLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
 **optional** | ***PurchaseInvoiceLineApiGetPurchaseInvoiceLineForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseInvoiceLineApiGetPurchaseInvoiceLineForPurchaseInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPurchaseInvoiceLines**
> InlineResponse20047 ListPurchaseInvoiceLines(ctx, companyId, optional)
Returns a list of purchaseInvoiceLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***PurchaseInvoiceLineApiListPurchaseInvoiceLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseInvoiceLineApiListPurchaseInvoiceLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPurchaseInvoiceLinesForPurchaseInvoice**
> InlineResponse20047 ListPurchaseInvoiceLinesForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)
Returns a list of purchaseInvoiceLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***PurchaseInvoiceLineApiListPurchaseInvoiceLinesForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseInvoiceLineApiListPurchaseInvoiceLinesForPurchaseInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPurchaseInvoiceLine**
> PurchaseInvoiceLine PatchPurchaseInvoiceLine(ctx, body, companyId, purchaseInvoiceLineId, contentType, ifMatch)
Updates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPurchaseInvoiceLineForPurchaseInvoice**
> PurchaseInvoiceLine PatchPurchaseInvoiceLineForPurchaseInvoice(ctx, body, companyId, purchaseInvoiceId, purchaseInvoiceLineId, contentType, ifMatch)
Updates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **purchaseInvoiceLineId** | **string**| id for purchaseInvoiceLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchaseInvoiceLine**
> PurchaseInvoiceLine PostPurchaseInvoiceLine(ctx, body, companyId, contentType)
Creates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchaseInvoiceLineForPurchaseInvoice**
> PurchaseInvoiceLine PostPurchaseInvoiceLineForPurchaseInvoice(ctx, body, companyId, purchaseInvoiceId, contentType)
Creates an object of type purchaseInvoiceLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **contentType** | **string**| application/json | 

### Return type

[**PurchaseInvoiceLine**](purchaseInvoiceLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

