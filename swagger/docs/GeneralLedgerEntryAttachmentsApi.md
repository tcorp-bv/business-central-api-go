# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGeneralLedgerEntryAttachments**](GeneralLedgerEntryAttachmentsApi.md#DeleteGeneralLedgerEntryAttachments) | **Delete** /companies({company_id})/generalLedgerEntryAttachments({generalLedgerEntryAttachments_generalLedgerEntryNumber},{generalLedgerEntryAttachments_id}) | Deletes an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central
[**GetGeneralLedgerEntryAttachments**](GeneralLedgerEntryAttachmentsApi.md#GetGeneralLedgerEntryAttachments) | **Get** /companies({company_id})/generalLedgerEntryAttachments({generalLedgerEntryAttachments_generalLedgerEntryNumber},{generalLedgerEntryAttachments_id}) | Retrieve the properties and relationships of an object of type generalLedgerEntryAttachments for Dynamics 365 Business Central.
[**ListGeneralLedgerEntryAttachments**](GeneralLedgerEntryAttachmentsApi.md#ListGeneralLedgerEntryAttachments) | **Get** /companies({company_id})/generalLedgerEntryAttachments | Returns a list of generalLedgerEntryAttachments
[**PatchGeneralLedgerEntryAttachments**](GeneralLedgerEntryAttachmentsApi.md#PatchGeneralLedgerEntryAttachments) | **Patch** /companies({company_id})/generalLedgerEntryAttachments({generalLedgerEntryAttachments_generalLedgerEntryNumber},{generalLedgerEntryAttachments_id}) | Updates an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central
[**PostGeneralLedgerEntryAttachments**](GeneralLedgerEntryAttachmentsApi.md#PostGeneralLedgerEntryAttachments) | **Post** /companies({company_id})/generalLedgerEntryAttachments | Creates an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central

# **DeleteGeneralLedgerEntryAttachments**
> DeleteGeneralLedgerEntryAttachments(ctx, companyId, generalLedgerEntryAttachmentsGeneralLedgerEntryNumber, generalLedgerEntryAttachmentsId)
Deletes an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **generalLedgerEntryAttachmentsGeneralLedgerEntryNumber** | **int32**| generalLedgerEntryNumber for generalLedgerEntryAttachments | 
  **generalLedgerEntryAttachmentsId** | [**string**](.md)| id for generalLedgerEntryAttachments | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGeneralLedgerEntryAttachments**
> GeneralLedgerEntryAttachments GetGeneralLedgerEntryAttachments(ctx, companyId, generalLedgerEntryAttachmentsGeneralLedgerEntryNumber, generalLedgerEntryAttachmentsId, optional)
Retrieve the properties and relationships of an object of type generalLedgerEntryAttachments for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **generalLedgerEntryAttachmentsGeneralLedgerEntryNumber** | **int32**| generalLedgerEntryNumber for generalLedgerEntryAttachments | 
  **generalLedgerEntryAttachmentsId** | [**string**](.md)| id for generalLedgerEntryAttachments | 
 **optional** | ***GeneralLedgerEntryAttachmentsApiGetGeneralLedgerEntryAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GeneralLedgerEntryAttachmentsApiGetGeneralLedgerEntryAttachmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**GeneralLedgerEntryAttachments**](generalLedgerEntryAttachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGeneralLedgerEntryAttachments**
> InlineResponse20024 ListGeneralLedgerEntryAttachments(ctx, companyId, optional)
Returns a list of generalLedgerEntryAttachments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***GeneralLedgerEntryAttachmentsApiListGeneralLedgerEntryAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GeneralLedgerEntryAttachmentsApiListGeneralLedgerEntryAttachmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGeneralLedgerEntryAttachments**
> GeneralLedgerEntryAttachments PatchGeneralLedgerEntryAttachments(ctx, body, companyId, generalLedgerEntryAttachmentsGeneralLedgerEntryNumber, generalLedgerEntryAttachmentsId, contentType, ifMatch)
Updates an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **generalLedgerEntryAttachmentsGeneralLedgerEntryNumber** | **int32**| generalLedgerEntryNumber for generalLedgerEntryAttachments | 
  **generalLedgerEntryAttachmentsId** | [**string**](.md)| id for generalLedgerEntryAttachments | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**GeneralLedgerEntryAttachments**](generalLedgerEntryAttachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGeneralLedgerEntryAttachments**
> GeneralLedgerEntryAttachments PostGeneralLedgerEntryAttachments(ctx, body, companyId, contentType)
Creates an object of type generalLedgerEntryAttachments in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**GeneralLedgerEntryAttachments**](generalLedgerEntryAttachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

