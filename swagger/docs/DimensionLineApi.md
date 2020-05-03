# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDimensionLine**](DimensionLineApi.md#DeleteDimensionLine) | **Delete** /companies({company_id})/dimensionLines({dimensionLine_parentId},{dimensionLine_id}) | Deletes an object of type dimensionLine in Dynamics 365 Business Central
[**GetDimensionLine**](DimensionLineApi.md#GetDimensionLine) | **Get** /companies({company_id})/dimensionLines({dimensionLine_parentId},{dimensionLine_id}) | Retrieve the properties and relationships of an object of type dimensionLine for Dynamics 365 Business Central.
[**ListDimensionLines**](DimensionLineApi.md#ListDimensionLines) | **Get** /companies({company_id})/dimensionLines | Returns a list of dimensionLines
[**PatchDimensionLine**](DimensionLineApi.md#PatchDimensionLine) | **Patch** /companies({company_id})/dimensionLines({dimensionLine_parentId},{dimensionLine_id}) | Updates an object of type dimensionLine in Dynamics 365 Business Central
[**PostDimensionLine**](DimensionLineApi.md#PostDimensionLine) | **Post** /companies({company_id})/dimensionLines | Creates an object of type dimensionLine in Dynamics 365 Business Central

# **DeleteDimensionLine**
> DeleteDimensionLine(ctx, companyId, dimensionLineParentId, dimensionLineId)
Deletes an object of type dimensionLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionLineParentId** | [**string**](.md)| parentId for dimensionLine | 
  **dimensionLineId** | [**string**](.md)| id for dimensionLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDimensionLine**
> DimensionLine GetDimensionLine(ctx, companyId, dimensionLineParentId, dimensionLineId, optional)
Retrieve the properties and relationships of an object of type dimensionLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionLineParentId** | [**string**](.md)| parentId for dimensionLine | 
  **dimensionLineId** | [**string**](.md)| id for dimensionLine | 
 **optional** | ***DimensionLineApiGetDimensionLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DimensionLineApiGetDimensionLineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DimensionLine**](dimensionLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDimensionLines**
> InlineResponse20025 ListDimensionLines(ctx, companyId, optional)
Returns a list of dimensionLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***DimensionLineApiListDimensionLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DimensionLineApiListDimensionLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDimensionLine**
> DimensionLine PatchDimensionLine(ctx, body, companyId, dimensionLineParentId, dimensionLineId, contentType, ifMatch)
Updates an object of type dimensionLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **dimensionLineParentId** | [**string**](.md)| parentId for dimensionLine | 
  **dimensionLineId** | [**string**](.md)| id for dimensionLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**DimensionLine**](dimensionLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDimensionLine**
> DimensionLine PostDimensionLine(ctx, body, companyId, contentType)
Creates an object of type dimensionLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**DimensionLine**](dimensionLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

