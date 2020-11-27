# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchaseInvoice**](PurchaseInvoiceApi.md#DeletePurchaseInvoice) | **Delete** /companies({company_id})/purchaseInvoices({purchaseInvoice_id}) | Deletes an object of type purchaseInvoice in Dynamics 365 Business Central
[**GetPurchaseInvoice**](PurchaseInvoiceApi.md#GetPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id}) | Retrieve the properties and relationships of an object of type purchaseInvoice for Dynamics 365 Business Central.
[**ListPurchaseInvoices**](PurchaseInvoiceApi.md#ListPurchaseInvoices) | **Get** /companies({company_id})/purchaseInvoices | Returns a list of purchaseInvoices
[**PatchPurchaseInvoice**](PurchaseInvoiceApi.md#PatchPurchaseInvoice) | **Patch** /companies({company_id})/purchaseInvoices({purchaseInvoice_id}) | Updates an object of type purchaseInvoice in Dynamics 365 Business Central
[**PostActionPurchaseInvoices**](PurchaseInvoiceApi.md#PostActionPurchaseInvoices) | **Post** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/Microsoft.NAV.post | Performs the post action for purchaseInvoices entity
[**PostPurchaseInvoice**](PurchaseInvoiceApi.md#PostPurchaseInvoice) | **Post** /companies({company_id})/purchaseInvoices | Creates an object of type purchaseInvoice in Dynamics 365 Business Central

# **DeletePurchaseInvoice**
> DeletePurchaseInvoice(ctx, companyId, purchaseInvoiceId)
Deletes an object of type purchaseInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurchaseInvoice**
> PurchaseInvoice GetPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)
Retrieve the properties and relationships of an object of type purchaseInvoice for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***PurchaseInvoiceApiGetPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseInvoiceApiGetPurchaseInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PurchaseInvoice**](purchaseInvoice.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPurchaseInvoices**
> InlineResponse20035 ListPurchaseInvoices(ctx, companyId, optional)
Returns a list of purchaseInvoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***PurchaseInvoiceApiListPurchaseInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseInvoiceApiListPurchaseInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPurchaseInvoice**
> PurchaseInvoice PatchPurchaseInvoice(ctx, body, companyId, purchaseInvoiceId, contentType, ifMatch)
Updates an object of type purchaseInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**PurchaseInvoice**](purchaseInvoice.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostActionPurchaseInvoices**
> PostActionPurchaseInvoices(ctx, companyId, purchaseInvoiceId)
Performs the post action for purchaseInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPurchaseInvoice**
> PurchaseInvoice PostPurchaseInvoice(ctx, body, companyId, contentType)
Creates an object of type purchaseInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**PurchaseInvoice**](purchaseInvoice.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

