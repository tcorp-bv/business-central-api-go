# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJournalLine**](JournalLineApi.md#DeleteJournalLine) | **Delete** /companies({company_id})/journalLines({journalLine_id}) | Deletes an object of type journalLine in Dynamics 365 Business Central
[**DeleteJournalLineForJournal**](JournalLineApi.md#DeleteJournalLineForJournal) | **Delete** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id}) | Deletes an object of type journalLine in Dynamics 365 Business Central
[**GetJournalLine**](JournalLineApi.md#GetJournalLine) | **Get** /companies({company_id})/journalLines({journalLine_id}) | Retrieve the properties and relationships of an object of type journalLine for Dynamics 365 Business Central.
[**GetJournalLineForJournal**](JournalLineApi.md#GetJournalLineForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id}) | Retrieve the properties and relationships of an object of type journalLine for Dynamics 365 Business Central.
[**ListJournalLines**](JournalLineApi.md#ListJournalLines) | **Get** /companies({company_id})/journalLines | Returns a list of journalLines
[**ListJournalLinesForJournal**](JournalLineApi.md#ListJournalLinesForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines | Returns a list of journalLines
[**PatchJournalLine**](JournalLineApi.md#PatchJournalLine) | **Patch** /companies({company_id})/journalLines({journalLine_id}) | Updates an object of type journalLine in Dynamics 365 Business Central
[**PatchJournalLineForJournal**](JournalLineApi.md#PatchJournalLineForJournal) | **Patch** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id}) | Updates an object of type journalLine in Dynamics 365 Business Central
[**PostJournalLine**](JournalLineApi.md#PostJournalLine) | **Post** /companies({company_id})/journalLines | Creates an object of type journalLine in Dynamics 365 Business Central
[**PostJournalLineForJournal**](JournalLineApi.md#PostJournalLineForJournal) | **Post** /companies({company_id})/journals({journal_id})/journalLines | Creates an object of type journalLine in Dynamics 365 Business Central

# **DeleteJournalLine**
> DeleteJournalLine(ctx, companyId, journalLineId)
Deletes an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJournalLineForJournal**
> DeleteJournalLineForJournal(ctx, companyId, journalId, journalLineId)
Deletes an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJournalLine**
> JournalLine GetJournalLine(ctx, companyId, journalLineId, optional)
Retrieve the properties and relationships of an object of type journalLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***JournalLineApiGetJournalLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JournalLineApiGetJournalLineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJournalLineForJournal**
> JournalLine GetJournalLineForJournal(ctx, companyId, journalId, journalLineId, optional)
Retrieve the properties and relationships of an object of type journalLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***JournalLineApiGetJournalLineForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JournalLineApiGetJournalLineForJournalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJournalLines**
> InlineResponse20028 ListJournalLines(ctx, companyId, optional)
Returns a list of journalLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***JournalLineApiListJournalLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JournalLineApiListJournalLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJournalLinesForJournal**
> InlineResponse20028 ListJournalLinesForJournal(ctx, companyId, journalId, optional)
Returns a list of journalLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
 **optional** | ***JournalLineApiListJournalLinesForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JournalLineApiListJournalLinesForJournalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchJournalLine**
> JournalLine PatchJournalLine(ctx, body, companyId, journalLineId, contentType, ifMatch)
Updates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchJournalLineForJournal**
> JournalLine PatchJournalLineForJournal(ctx, body, companyId, journalId, journalLineId, contentType, ifMatch)
Updates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **journalLineId** | [**string**](.md)| id for journalLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostJournalLine**
> JournalLine PostJournalLine(ctx, body, companyId, contentType)
Creates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostJournalLineForJournal**
> JournalLine PostJournalLineForJournal(ctx, body, companyId, journalId, contentType)
Creates an object of type journalLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **journalId** | [**string**](.md)| id for journal | 
  **contentType** | **string**| application/json | 

### Return type

[**JournalLine**](journalLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

