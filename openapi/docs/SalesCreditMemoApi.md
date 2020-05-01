# \SalesCreditMemoApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelActionSalesCreditMemos**](SalesCreditMemoApi.md#CancelActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.cancel | Performs the cancel action for salesCreditMemos entity
[**CancelAndSendActionSalesCreditMemos**](SalesCreditMemoApi.md#CancelAndSendActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.cancelAndSend | Performs the cancelAndSend action for salesCreditMemos entity
[**DeleteSalesCreditMemo**](SalesCreditMemoApi.md#DeleteSalesCreditMemo) | **Delete** /companies({company_id})/salesCreditMemos({salesCreditMemo_id}) | Deletes an object of type salesCreditMemo in Dynamics 365 Business Central
[**GetSalesCreditMemo**](SalesCreditMemoApi.md#GetSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id}) | Retrieve the properties and relationships of an object of type salesCreditMemo for Dynamics 365 Business Central.
[**ListSalesCreditMemos**](SalesCreditMemoApi.md#ListSalesCreditMemos) | **Get** /companies({company_id})/salesCreditMemos | Returns a list of salesCreditMemos
[**PatchSalesCreditMemo**](SalesCreditMemoApi.md#PatchSalesCreditMemo) | **Patch** /companies({company_id})/salesCreditMemos({salesCreditMemo_id}) | Updates an object of type salesCreditMemo in Dynamics 365 Business Central
[**PostActionSalesCreditMemos**](SalesCreditMemoApi.md#PostActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.post | Performs the post action for salesCreditMemos entity
[**PostAndSendActionSalesCreditMemos**](SalesCreditMemoApi.md#PostAndSendActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.postAndSend | Performs the postAndSend action for salesCreditMemos entity
[**PostSalesCreditMemo**](SalesCreditMemoApi.md#PostSalesCreditMemo) | **Post** /companies({company_id})/salesCreditMemos | Creates an object of type salesCreditMemo in Dynamics 365 Business Central
[**SendActionSalesCreditMemos**](SalesCreditMemoApi.md#SendActionSalesCreditMemos) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/Microsoft.NAV.send | Performs the send action for salesCreditMemos entity



## CancelActionSalesCreditMemos

> CancelActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)

Performs the cancel action for salesCreditMemos entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

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


## CancelAndSendActionSalesCreditMemos

> CancelAndSendActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)

Performs the cancelAndSend action for salesCreditMemos entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

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


## DeleteSalesCreditMemo

> DeleteSalesCreditMemo(ctx, companyId, salesCreditMemoId)

Deletes an object of type salesCreditMemo in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

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


## GetSalesCreditMemo

> SalesCreditMemo GetSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)

Retrieve the properties and relationships of an object of type salesCreditMemo for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***GetSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSalesCreditMemoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesCreditMemo**](salesCreditMemo.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSalesCreditMemos

> InlineResponse20043 ListSalesCreditMemos(ctx, companyId, optional)

Returns a list of salesCreditMemos

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListSalesCreditMemosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSalesCreditMemosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSalesCreditMemo

> SalesCreditMemo PatchSalesCreditMemo(ctx, companyId, salesCreditMemoId, contentType, ifMatch, uNKNOWNBASETYPE)

Updates an object of type salesCreditMemo in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
**contentType** | **string**| application/json | 
**ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**SalesCreditMemo**](salesCreditMemo.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostActionSalesCreditMemos

> PostActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)

Performs the post action for salesCreditMemos entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

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


## PostAndSendActionSalesCreditMemos

> PostAndSendActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)

Performs the postAndSend action for salesCreditMemos entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

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


## PostSalesCreditMemo

> SalesCreditMemo PostSalesCreditMemo(ctx, companyId, contentType, uNKNOWNBASETYPE)

Creates an object of type salesCreditMemo in Dynamics 365 Business Central

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**contentType** | **string**| application/json | 
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**SalesCreditMemo**](salesCreditMemo.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendActionSalesCreditMemos

> SendActionSalesCreditMemos(ctx, companyId, salesCreditMemoId)

Performs the send action for salesCreditMemos entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 

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

