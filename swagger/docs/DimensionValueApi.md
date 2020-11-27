# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDimensionValue**](DimensionValueApi.md#GetDimensionValue) | **Get** /companies({company_id})/dimensionValues({dimensionValue_id}) | Retrieve the properties and relationships of an object of type dimensionValue for Dynamics 365 Business Central.
[**GetDimensionValueForDimension**](DimensionValueApi.md#GetDimensionValueForDimension) | **Get** /companies({company_id})/dimensions({dimension_id})/dimensionValues({dimensionValue_id}) | Retrieve the properties and relationships of an object of type dimensionValue for Dynamics 365 Business Central.
[**ListDimensionValues**](DimensionValueApi.md#ListDimensionValues) | **Get** /companies({company_id})/dimensionValues | Returns a list of dimensionValues
[**ListDimensionValuesForDimension**](DimensionValueApi.md#ListDimensionValuesForDimension) | **Get** /companies({company_id})/dimensions({dimension_id})/dimensionValues | Returns a list of dimensionValues

# **GetDimensionValue**
> DimensionValue GetDimensionValue(ctx, companyId, dimensionValueId, optional)
Retrieve the properties and relationships of an object of type dimensionValue for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionValueId** | [**string**](.md)| id for dimensionValue | 
 **optional** | ***DimensionValueApiGetDimensionValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DimensionValueApiGetDimensionValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DimensionValue**](dimensionValue.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDimensionValueForDimension**
> DimensionValue GetDimensionValueForDimension(ctx, companyId, dimensionId, dimensionValueId, optional)
Retrieve the properties and relationships of an object of type dimensionValue for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionId** | [**string**](.md)| id for dimension | 
  **dimensionValueId** | [**string**](.md)| id for dimensionValue | 
 **optional** | ***DimensionValueApiGetDimensionValueForDimensionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DimensionValueApiGetDimensionValueForDimensionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**DimensionValue**](dimensionValue.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDimensionValues**
> InlineResponse20024 ListDimensionValues(ctx, companyId, optional)
Returns a list of dimensionValues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***DimensionValueApiListDimensionValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DimensionValueApiListDimensionValuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDimensionValuesForDimension**
> InlineResponse20024 ListDimensionValuesForDimension(ctx, companyId, dimensionId, optional)
Returns a list of dimensionValues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **dimensionId** | [**string**](.md)| id for dimension | 
 **optional** | ***DimensionValueApiListDimensionValuesForDimensionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DimensionValueApiListDimensionValuesForDimensionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

