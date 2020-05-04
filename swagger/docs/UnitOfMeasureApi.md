# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUnitOfMeasure**](UnitOfMeasureApi.md#DeleteUnitOfMeasure) | **Delete** /companies({company_id})/unitsOfMeasure({unitOfMeasure_id}) | Deletes an object of type unitOfMeasure in Dynamics 365 Business Central
[**GetUnitOfMeasure**](UnitOfMeasureApi.md#GetUnitOfMeasure) | **Get** /companies({company_id})/unitsOfMeasure({unitOfMeasure_id}) | Retrieve the properties and relationships of an object of type unitOfMeasure for Dynamics 365 Business Central.
[**ListUnitsOfMeasure**](UnitOfMeasureApi.md#ListUnitsOfMeasure) | **Get** /companies({company_id})/unitsOfMeasure | Returns a list of unitsOfMeasure
[**PatchUnitOfMeasure**](UnitOfMeasureApi.md#PatchUnitOfMeasure) | **Patch** /companies({company_id})/unitsOfMeasure({unitOfMeasure_id}) | Updates an object of type unitOfMeasure in Dynamics 365 Business Central
[**PostUnitOfMeasure**](UnitOfMeasureApi.md#PostUnitOfMeasure) | **Post** /companies({company_id})/unitsOfMeasure | Creates an object of type unitOfMeasure in Dynamics 365 Business Central

# **DeleteUnitOfMeasure**
> DeleteUnitOfMeasure(ctx, companyId, unitOfMeasureId)
Deletes an object of type unitOfMeasure in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **unitOfMeasureId** | [**string**](.md)| id for unitOfMeasure | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnitOfMeasure**
> UnitOfMeasure GetUnitOfMeasure(ctx, companyId, unitOfMeasureId, optional)
Retrieve the properties and relationships of an object of type unitOfMeasure for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **unitOfMeasureId** | [**string**](.md)| id for unitOfMeasure | 
 **optional** | ***UnitOfMeasureApiGetUnitOfMeasureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnitOfMeasureApiGetUnitOfMeasureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**UnitOfMeasure**](unitOfMeasure.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUnitsOfMeasure**
> InlineResponse20049 ListUnitsOfMeasure(ctx, companyId, optional)
Returns a list of unitsOfMeasure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***UnitOfMeasureApiListUnitsOfMeasureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnitOfMeasureApiListUnitsOfMeasureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUnitOfMeasure**
> UnitOfMeasure PatchUnitOfMeasure(ctx, body, companyId, unitOfMeasureId, contentType, ifMatch)
Updates an object of type unitOfMeasure in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **unitOfMeasureId** | [**string**](.md)| id for unitOfMeasure | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**UnitOfMeasure**](unitOfMeasure.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUnitOfMeasure**
> UnitOfMeasure PostUnitOfMeasure(ctx, body, companyId, contentType)
Creates an object of type unitOfMeasure in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**UnitOfMeasure**](unitOfMeasure.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

