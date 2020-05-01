# \ItemCategoryApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItemCategory**](ItemCategoryApi.md#DeleteItemCategory) | **Delete** /companies({company_id})/itemCategories({itemCategory_id}) | Deletes an object of type itemCategory in Dynamics 365 Business Central
[**GetItemCategory**](ItemCategoryApi.md#GetItemCategory) | **Get** /companies({company_id})/itemCategories({itemCategory_id}) | Retrieve the properties and relationships of an object of type itemCategory for Dynamics 365 Business Central.
[**ListItemCategories**](ItemCategoryApi.md#ListItemCategories) | **Get** /companies({company_id})/itemCategories | Returns a list of itemCategories
[**PatchItemCategory**](ItemCategoryApi.md#PatchItemCategory) | **Patch** /companies({company_id})/itemCategories({itemCategory_id}) | Updates an object of type itemCategory in Dynamics 365 Business Central
[**PostItemCategory**](ItemCategoryApi.md#PostItemCategory) | **Post** /companies({company_id})/itemCategories | Creates an object of type itemCategory in Dynamics 365 Business Central



## DeleteItemCategory

> DeleteItemCategory(ctx, companyId, itemCategoryId)

Deletes an object of type itemCategory in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**itemCategoryId** | [**string**](.md)| id for itemCategory | 

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


## GetItemCategory

> ItemCategory GetItemCategory(ctx, companyId, itemCategoryId, optional)

Retrieve the properties and relationships of an object of type itemCategory for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**itemCategoryId** | [**string**](.md)| id for itemCategory | 
 **optional** | ***GetItemCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemCategoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**ItemCategory**](itemCategory.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItemCategories

> InlineResponse20028 ListItemCategories(ctx, companyId, optional)

Returns a list of itemCategories

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListItemCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListItemCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchItemCategory

> ItemCategory PatchItemCategory(ctx, companyId, itemCategoryId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type itemCategory in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**itemCategoryId** | [**string**](.md)| id for itemCategory | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**ItemCategory**](itemCategory.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemCategory

> ItemCategory PostItemCategory(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type itemCategory in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**ItemCategory**](itemCategory.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

