# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesCreditMemoLine**](SalesCreditMemoLineApi.md#DeleteSalesCreditMemoLine) | **Delete** /companies({company_id})/salesCreditMemoLines(&#x27;{salesCreditMemoLine_id}&#x27;) | Deletes an object of type salesCreditMemoLine in Dynamics 365 Business Central
[**DeleteSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#DeleteSalesCreditMemoLineForSalesCreditMemo) | **Delete** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines(&#x27;{salesCreditMemoLine_id}&#x27;) | Deletes an object of type salesCreditMemoLine in Dynamics 365 Business Central
[**GetSalesCreditMemoLine**](SalesCreditMemoLineApi.md#GetSalesCreditMemoLine) | **Get** /companies({company_id})/salesCreditMemoLines(&#x27;{salesCreditMemoLine_id}&#x27;) | Retrieve the properties and relationships of an object of type salesCreditMemoLine for Dynamics 365 Business Central.
[**GetSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#GetSalesCreditMemoLineForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines(&#x27;{salesCreditMemoLine_id}&#x27;) | Retrieve the properties and relationships of an object of type salesCreditMemoLine for Dynamics 365 Business Central.
[**ListSalesCreditMemoLines**](SalesCreditMemoLineApi.md#ListSalesCreditMemoLines) | **Get** /companies({company_id})/salesCreditMemoLines | Returns a list of salesCreditMemoLines
[**ListSalesCreditMemoLinesForSalesCreditMemo**](SalesCreditMemoLineApi.md#ListSalesCreditMemoLinesForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines | Returns a list of salesCreditMemoLines
[**PatchSalesCreditMemoLine**](SalesCreditMemoLineApi.md#PatchSalesCreditMemoLine) | **Patch** /companies({company_id})/salesCreditMemoLines(&#x27;{salesCreditMemoLine_id}&#x27;) | Updates an object of type salesCreditMemoLine in Dynamics 365 Business Central
[**PatchSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#PatchSalesCreditMemoLineForSalesCreditMemo) | **Patch** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines(&#x27;{salesCreditMemoLine_id}&#x27;) | Updates an object of type salesCreditMemoLine in Dynamics 365 Business Central
[**PostSalesCreditMemoLine**](SalesCreditMemoLineApi.md#PostSalesCreditMemoLine) | **Post** /companies({company_id})/salesCreditMemoLines | Creates an object of type salesCreditMemoLine in Dynamics 365 Business Central
[**PostSalesCreditMemoLineForSalesCreditMemo**](SalesCreditMemoLineApi.md#PostSalesCreditMemoLineForSalesCreditMemo) | **Post** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/salesCreditMemoLines | Creates an object of type salesCreditMemoLine in Dynamics 365 Business Central

# **DeleteSalesCreditMemoLine**
> DeleteSalesCreditMemoLine(ctx, companyId, salesCreditMemoLineId)
Deletes an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesCreditMemoLineForSalesCreditMemo**
> DeleteSalesCreditMemoLineForSalesCreditMemo(ctx, companyId, salesCreditMemoId, salesCreditMemoLineId)
Deletes an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesCreditMemoLine**
> SalesCreditMemoLine GetSalesCreditMemoLine(ctx, companyId, salesCreditMemoLineId, optional)
Retrieve the properties and relationships of an object of type salesCreditMemoLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
 **optional** | ***SalesCreditMemoLineApiGetSalesCreditMemoLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesCreditMemoLineApiGetSalesCreditMemoLineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesCreditMemoLineForSalesCreditMemo**
> SalesCreditMemoLine GetSalesCreditMemoLineForSalesCreditMemo(ctx, companyId, salesCreditMemoId, salesCreditMemoLineId, optional)
Retrieve the properties and relationships of an object of type salesCreditMemoLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
 **optional** | ***SalesCreditMemoLineApiGetSalesCreditMemoLineForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesCreditMemoLineApiGetSalesCreditMemoLineForSalesCreditMemoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesCreditMemoLines**
> InlineResponse20037 ListSalesCreditMemoLines(ctx, companyId, optional)
Returns a list of salesCreditMemoLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***SalesCreditMemoLineApiListSalesCreditMemoLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesCreditMemoLineApiListSalesCreditMemoLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesCreditMemoLinesForSalesCreditMemo**
> InlineResponse20037 ListSalesCreditMemoLinesForSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)
Returns a list of salesCreditMemoLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***SalesCreditMemoLineApiListSalesCreditMemoLinesForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesCreditMemoLineApiListSalesCreditMemoLinesForSalesCreditMemoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesCreditMemoLine**
> SalesCreditMemoLine PatchSalesCreditMemoLine(ctx, body, companyId, salesCreditMemoLineId, contentType, ifMatch)
Updates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesCreditMemoLineForSalesCreditMemo**
> SalesCreditMemoLine PatchSalesCreditMemoLineForSalesCreditMemo(ctx, body, companyId, salesCreditMemoId, salesCreditMemoLineId, contentType, ifMatch)
Updates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **salesCreditMemoLineId** | **string**| id for salesCreditMemoLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesCreditMemoLine**
> SalesCreditMemoLine PostSalesCreditMemoLine(ctx, body, companyId, contentType)
Creates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesCreditMemoLineForSalesCreditMemo**
> SalesCreditMemoLine PostSalesCreditMemoLineForSalesCreditMemo(ctx, body, companyId, salesCreditMemoId, contentType)
Creates an object of type salesCreditMemoLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesCreditMemoLine**](salesCreditMemoLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

