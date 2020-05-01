# \TaxAreaApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaxArea**](TaxAreaApi.md#DeleteTaxArea) | **Delete** /companies({company_id})/taxAreas({taxArea_id}) | Deletes an object of type taxArea in Dynamics 365 Business Central
[**GetTaxArea**](TaxAreaApi.md#GetTaxArea) | **Get** /companies({company_id})/taxAreas({taxArea_id}) | Retrieve the properties and relationships of an object of type taxArea for Dynamics 365 Business Central.
[**ListTaxAreas**](TaxAreaApi.md#ListTaxAreas) | **Get** /companies({company_id})/taxAreas | Returns a list of taxAreas
[**PatchTaxArea**](TaxAreaApi.md#PatchTaxArea) | **Patch** /companies({company_id})/taxAreas({taxArea_id}) | Updates an object of type taxArea in Dynamics 365 Business Central
[**PostTaxArea**](TaxAreaApi.md#PostTaxArea) | **Post** /companies({company_id})/taxAreas | Creates an object of type taxArea in Dynamics 365 Business Central



## DeleteTaxArea

> DeleteTaxArea(ctx, companyId, taxAreaId)

Deletes an object of type taxArea in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**taxAreaId** | [**string**](.md)| id for taxArea | 

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


## GetTaxArea

> TaxArea GetTaxArea(ctx, companyId, taxAreaId, optional)

Retrieve the properties and relationships of an object of type taxArea for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**taxAreaId** | [**string**](.md)| id for taxArea | 
 **optional** | ***GetTaxAreaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTaxAreaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TaxArea**](taxArea.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaxAreas

> InlineResponse20040 ListTaxAreas(ctx, companyId, optional)

Returns a list of taxAreas

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListTaxAreasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaxAreasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTaxArea

> TaxArea PatchTaxArea(ctx, companyId, taxAreaId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type taxArea in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**taxAreaId** | [**string**](.md)| id for taxArea | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**TaxArea**](taxArea.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaxArea

> TaxArea PostTaxArea(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type taxArea in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**TaxArea**](taxArea.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

