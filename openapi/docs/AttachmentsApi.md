# \AttachmentsApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttachments**](AttachmentsApi.md#DeleteAttachments) | **Delete** /companies({company_id})/attachments({attachments_parentId},{attachments_id}) | Deletes an object of type attachments in Dynamics 365 Business Central
[**DeleteAttachmentsForJournalLine**](AttachmentsApi.md#DeleteAttachmentsForJournalLine) | **Delete** /companies({company_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | Deletes an object of type attachments in Dynamics 365 Business Central
[**DeleteAttachmentsForJournalLineForJournal**](AttachmentsApi.md#DeleteAttachmentsForJournalLineForJournal) | **Delete** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | Deletes an object of type attachments in Dynamics 365 Business Central
[**GetAttachments**](AttachmentsApi.md#GetAttachments) | **Get** /companies({company_id})/attachments({attachments_parentId},{attachments_id}) | Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.
[**GetAttachmentsForJournalLine**](AttachmentsApi.md#GetAttachmentsForJournalLine) | **Get** /companies({company_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.
[**GetAttachmentsForJournalLineForJournal**](AttachmentsApi.md#GetAttachmentsForJournalLineForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.
[**ListAttachments**](AttachmentsApi.md#ListAttachments) | **Get** /companies({company_id})/attachments | Returns a list of attachments
[**ListAttachmentsForJournalLine**](AttachmentsApi.md#ListAttachmentsForJournalLine) | **Get** /companies({company_id})/journalLines({journalLine_id})/attachments | Returns a list of attachments
[**ListAttachmentsForJournalLineForJournal**](AttachmentsApi.md#ListAttachmentsForJournalLineForJournal) | **Get** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments | Returns a list of attachments
[**PatchAttachments**](AttachmentsApi.md#PatchAttachments) | **Patch** /companies({company_id})/attachments({attachments_parentId},{attachments_id}) | Updates an object of type attachments in Dynamics 365 Business Central
[**PatchAttachmentsForJournalLine**](AttachmentsApi.md#PatchAttachmentsForJournalLine) | **Patch** /companies({company_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | Updates an object of type attachments in Dynamics 365 Business Central
[**PatchAttachmentsForJournalLineForJournal**](AttachmentsApi.md#PatchAttachmentsForJournalLineForJournal) | **Patch** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments({attachments_parentId},{attachments_id}) | Updates an object of type attachments in Dynamics 365 Business Central
[**PostAttachments**](AttachmentsApi.md#PostAttachments) | **Post** /companies({company_id})/attachments | Creates an object of type attachments in Dynamics 365 Business Central
[**PostAttachmentsForJournalLine**](AttachmentsApi.md#PostAttachmentsForJournalLine) | **Post** /companies({company_id})/journalLines({journalLine_id})/attachments | Creates an object of type attachments in Dynamics 365 Business Central
[**PostAttachmentsForJournalLineForJournal**](AttachmentsApi.md#PostAttachmentsForJournalLineForJournal) | **Post** /companies({company_id})/journals({journal_id})/journalLines({journalLine_id})/attachments | Creates an object of type attachments in Dynamics 365 Business Central



## DeleteAttachments

> DeleteAttachments(ctx, companyId, attachmentsParentId, attachmentsId)

Deletes an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 

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


## DeleteAttachmentsForJournalLine

> DeleteAttachmentsForJournalLine(ctx, companyId, journalLineId, attachmentsParentId, attachmentsId)

Deletes an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 

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


## DeleteAttachmentsForJournalLineForJournal

> DeleteAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, attachmentsParentId, attachmentsId)

Deletes an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalId** | [**string**](.md)| id for journal | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 

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


## GetAttachments

> Attachments GetAttachments(ctx, companyId, attachmentsParentId, attachmentsId, optional)

Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 
 **optional** | ***GetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAttachmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentsForJournalLine

> Attachments GetAttachmentsForJournalLine(ctx, companyId, journalLineId, attachmentsParentId, attachmentsId, optional)

Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 
 **optional** | ***GetAttachmentsForJournalLineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAttachmentsForJournalLineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentsForJournalLineForJournal

> Attachments GetAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, attachmentsParentId, attachmentsId, optional)

Retrieve the properties and relationships of an object of type attachments for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalId** | [**string**](.md)| id for journal | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 
 **optional** | ***GetAttachmentsForJournalLineForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAttachmentsForJournalLineForJournalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachments

> InlineResponse20017 ListAttachments(ctx, companyId, optional)

Returns a list of attachments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAttachmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachmentsForJournalLine

> InlineResponse20017 ListAttachmentsForJournalLine(ctx, companyId, journalLineId, optional)

Returns a list of attachments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***ListAttachmentsForJournalLineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAttachmentsForJournalLineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachmentsForJournalLineForJournal

> InlineResponse20017 ListAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, optional)

Returns a list of attachments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalId** | [**string**](.md)| id for journal | 
**journalLineId** | [**string**](.md)| id for journalLine | 
 **optional** | ***ListAttachmentsForJournalLineForJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAttachmentsForJournalLineForJournalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAttachments

> Attachments PatchAttachments(ctx, companyId, attachmentsParentId, attachmentsId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAttachmentsForJournalLine

> Attachments PatchAttachmentsForJournalLine(ctx, companyId, journalLineId, attachmentsParentId, attachmentsId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAttachmentsForJournalLineForJournal

> Attachments PatchAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, attachmentsParentId, attachmentsId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalId** | [**string**](.md)| id for journal | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**attachmentsParentId** | [**string**](.md)| parentId for attachments | 
**attachmentsId** | [**string**](.md)| id for attachments | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAttachments

> Attachments PostAttachments(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAttachmentsForJournalLine

> Attachments PostAttachmentsForJournalLine(ctx, companyId, journalLineId, contentType, uNKNOWNBASETYPE)

Creates an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAttachmentsForJournalLineForJournal

> Attachments PostAttachmentsForJournalLineForJournal(ctx, companyId, journalId, journalLineId, contentType, uNKNOWNBASETYPE)

Creates an object of type attachments in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**journalId** | [**string**](.md)| id for journal | 
**journalLineId** | [**string**](.md)| id for journalLine | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**Attachments**](attachments.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

