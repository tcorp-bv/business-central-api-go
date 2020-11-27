# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelActionSalesInvoices**](SalesInvoiceApi.md#CancelActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.cancel | Performs the cancel action for salesInvoices entity
[**CancelAndSendActionSalesInvoices**](SalesInvoiceApi.md#CancelAndSendActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.cancelAndSend | Performs the cancelAndSend action for salesInvoices entity
[**DeleteSalesInvoice**](SalesInvoiceApi.md#DeleteSalesInvoice) | **Delete** /companies({company_id})/salesInvoices({salesInvoice_id}) | Deletes an object of type salesInvoice in Dynamics 365 Business Central
[**GetSalesInvoice**](SalesInvoiceApi.md#GetSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id}) | Retrieve the properties and relationships of an object of type salesInvoice for Dynamics 365 Business Central.
[**ListSalesInvoices**](SalesInvoiceApi.md#ListSalesInvoices) | **Get** /companies({company_id})/salesInvoices | Returns a list of salesInvoices
[**MakeCorrectiveCreditMemoActionSalesInvoices**](SalesInvoiceApi.md#MakeCorrectiveCreditMemoActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.makeCorrectiveCreditMemo | Performs the makeCorrectiveCreditMemo action for salesInvoices entity
[**PatchSalesInvoice**](SalesInvoiceApi.md#PatchSalesInvoice) | **Patch** /companies({company_id})/salesInvoices({salesInvoice_id}) | Updates an object of type salesInvoice in Dynamics 365 Business Central
[**PostActionSalesInvoices**](SalesInvoiceApi.md#PostActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.post | Performs the post action for salesInvoices entity
[**PostAndSendActionSalesInvoices**](SalesInvoiceApi.md#PostAndSendActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.postAndSend | Performs the postAndSend action for salesInvoices entity
[**PostSalesInvoice**](SalesInvoiceApi.md#PostSalesInvoice) | **Post** /companies({company_id})/salesInvoices | Creates an object of type salesInvoice in Dynamics 365 Business Central
[**SendActionSalesInvoices**](SalesInvoiceApi.md#SendActionSalesInvoices) | **Post** /companies({company_id})/salesInvoices({salesInvoice_id})/Microsoft.NAV.send | Performs the send action for salesInvoices entity

# **CancelActionSalesInvoices**
> CancelActionSalesInvoices(ctx, companyId, salesInvoiceId)
Performs the cancel action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelAndSendActionSalesInvoices**
> CancelAndSendActionSalesInvoices(ctx, companyId, salesInvoiceId)
Performs the cancelAndSend action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesInvoice**
> DeleteSalesInvoice(ctx, companyId, salesInvoiceId)
Deletes an object of type salesInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesInvoice**
> SalesInvoice GetSalesInvoice(ctx, companyId, salesInvoiceId, optional)
Retrieve the properties and relationships of an object of type salesInvoice for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***SalesInvoiceApiGetSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesInvoiceApiGetSalesInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesInvoice**](salesInvoice.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesInvoices**
> InlineResponse20040 ListSalesInvoices(ctx, companyId, optional)
Returns a list of salesInvoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***SalesInvoiceApiListSalesInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesInvoiceApiListSalesInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeCorrectiveCreditMemoActionSalesInvoices**
> MakeCorrectiveCreditMemoActionSalesInvoices(ctx, companyId, salesInvoiceId)
Performs the makeCorrectiveCreditMemo action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesInvoice**
> SalesInvoice PatchSalesInvoice(ctx, body, companyId, salesInvoiceId, contentType, ifMatch)
Updates an object of type salesInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesInvoice**](salesInvoice.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostActionSalesInvoices**
> PostActionSalesInvoices(ctx, companyId, salesInvoiceId)
Performs the post action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAndSendActionSalesInvoices**
> PostAndSendActionSalesInvoices(ctx, companyId, salesInvoiceId)
Performs the postAndSend action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesInvoice**
> SalesInvoice PostSalesInvoice(ctx, body, companyId, contentType)
Creates an object of type salesInvoice in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesInvoice**](salesInvoice.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendActionSalesInvoices**
> SendActionSalesInvoices(ctx, companyId, salesInvoiceId)
Performs the send action for salesInvoices entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

