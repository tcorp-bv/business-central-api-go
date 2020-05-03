# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaxGroup**](TaxGroupApi.md#DeleteTaxGroup) | **Delete** /companies({company_id})/taxGroups({taxGroup_id}) | Deletes an object of type taxGroup in Dynamics 365 Business Central
[**GetTaxGroup**](TaxGroupApi.md#GetTaxGroup) | **Get** /companies({company_id})/taxGroups({taxGroup_id}) | Retrieve the properties and relationships of an object of type taxGroup for Dynamics 365 Business Central.
[**ListTaxGroups**](TaxGroupApi.md#ListTaxGroups) | **Get** /companies({company_id})/taxGroups | Returns a list of taxGroups
[**PatchTaxGroup**](TaxGroupApi.md#PatchTaxGroup) | **Patch** /companies({company_id})/taxGroups({taxGroup_id}) | Updates an object of type taxGroup in Dynamics 365 Business Central
[**PostTaxGroup**](TaxGroupApi.md#PostTaxGroup) | **Post** /companies({company_id})/taxGroups | Creates an object of type taxGroup in Dynamics 365 Business Central

# **DeleteTaxGroup**
> DeleteTaxGroup(ctx, companyId, taxGroupId)
Deletes an object of type taxGroup in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxGroupId** | [**string**](.md)| id for taxGroup | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxGroup**
> TaxGroup GetTaxGroup(ctx, companyId, taxGroupId, optional)
Retrieve the properties and relationships of an object of type taxGroup for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **taxGroupId** | [**string**](.md)| id for taxGroup | 
 **optional** | ***TaxGroupApiGetTaxGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaxGroupApiGetTaxGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TaxGroup**](taxGroup.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaxGroups**
> InlineResponse20014 ListTaxGroups(ctx, companyId, optional)
Returns a list of taxGroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***TaxGroupApiListTaxGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaxGroupApiListTaxGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTaxGroup**
> TaxGroup PatchTaxGroup(ctx, body, companyId, taxGroupId, contentType, ifMatch)
Updates an object of type taxGroup in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **taxGroupId** | [**string**](.md)| id for taxGroup | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**TaxGroup**](taxGroup.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTaxGroup**
> TaxGroup PostTaxGroup(ctx, body, companyId, contentType)
Creates an object of type taxGroup in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**TaxGroup**](taxGroup.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

