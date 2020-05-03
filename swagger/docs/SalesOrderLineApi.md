# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrderLine**](SalesOrderLineApi.md#DeleteSalesOrderLine) | **Delete** /companies({company_id})/salesOrderLines(&#x27;{salesOrderLine_id}&#x27;) | Deletes an object of type salesOrderLine in Dynamics 365 Business Central
[**DeleteSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#DeleteSalesOrderLineForSalesOrder) | **Delete** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines(&#x27;{salesOrderLine_id}&#x27;) | Deletes an object of type salesOrderLine in Dynamics 365 Business Central
[**GetSalesOrderLine**](SalesOrderLineApi.md#GetSalesOrderLine) | **Get** /companies({company_id})/salesOrderLines(&#x27;{salesOrderLine_id}&#x27;) | Retrieve the properties and relationships of an object of type salesOrderLine for Dynamics 365 Business Central.
[**GetSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#GetSalesOrderLineForSalesOrder) | **Get** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines(&#x27;{salesOrderLine_id}&#x27;) | Retrieve the properties and relationships of an object of type salesOrderLine for Dynamics 365 Business Central.
[**ListSalesOrderLines**](SalesOrderLineApi.md#ListSalesOrderLines) | **Get** /companies({company_id})/salesOrderLines | Returns a list of salesOrderLines
[**ListSalesOrderLinesForSalesOrder**](SalesOrderLineApi.md#ListSalesOrderLinesForSalesOrder) | **Get** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines | Returns a list of salesOrderLines
[**PatchSalesOrderLine**](SalesOrderLineApi.md#PatchSalesOrderLine) | **Patch** /companies({company_id})/salesOrderLines(&#x27;{salesOrderLine_id}&#x27;) | Updates an object of type salesOrderLine in Dynamics 365 Business Central
[**PatchSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#PatchSalesOrderLineForSalesOrder) | **Patch** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines(&#x27;{salesOrderLine_id}&#x27;) | Updates an object of type salesOrderLine in Dynamics 365 Business Central
[**PostSalesOrderLine**](SalesOrderLineApi.md#PostSalesOrderLine) | **Post** /companies({company_id})/salesOrderLines | Creates an object of type salesOrderLine in Dynamics 365 Business Central
[**PostSalesOrderLineForSalesOrder**](SalesOrderLineApi.md#PostSalesOrderLineForSalesOrder) | **Post** /companies({company_id})/salesOrders({salesOrder_id})/salesOrderLines | Creates an object of type salesOrderLine in Dynamics 365 Business Central

# **DeleteSalesOrderLine**
> DeleteSalesOrderLine(ctx, companyId, salesOrderLineId)
Deletes an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSalesOrderLineForSalesOrder**
> DeleteSalesOrderLineForSalesOrder(ctx, companyId, salesOrderId, salesOrderLineId)
Deletes an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesOrderLine**
> SalesOrderLine GetSalesOrderLine(ctx, companyId, salesOrderLineId, optional)
Retrieve the properties and relationships of an object of type salesOrderLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
 **optional** | ***SalesOrderLineApiGetSalesOrderLineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesOrderLineApiGetSalesOrderLineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesOrderLineForSalesOrder**
> SalesOrderLine GetSalesOrderLineForSalesOrder(ctx, companyId, salesOrderId, salesOrderLineId, optional)
Retrieve the properties and relationships of an object of type salesOrderLine for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
 **optional** | ***SalesOrderLineApiGetSalesOrderLineForSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesOrderLineApiGetSalesOrderLineForSalesOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesOrderLines**
> InlineResponse20032 ListSalesOrderLines(ctx, companyId, optional)
Returns a list of salesOrderLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***SalesOrderLineApiListSalesOrderLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesOrderLineApiListSalesOrderLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesOrderLinesForSalesOrder**
> InlineResponse20032 ListSalesOrderLinesForSalesOrder(ctx, companyId, salesOrderId, optional)
Returns a list of salesOrderLines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
 **optional** | ***SalesOrderLineApiListSalesOrderLinesForSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesOrderLineApiListSalesOrderLinesForSalesOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesOrderLine**
> SalesOrderLine PatchSalesOrderLine(ctx, body, companyId, salesOrderLineId, contentType, ifMatch)
Updates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesOrderLineForSalesOrder**
> SalesOrderLine PatchSalesOrderLineForSalesOrder(ctx, body, companyId, salesOrderId, salesOrderLineId, contentType, ifMatch)
Updates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **salesOrderLineId** | **string**| id for salesOrderLine | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesOrderLine**
> SalesOrderLine PostSalesOrderLine(ctx, body, companyId, contentType)
Creates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesOrderLineForSalesOrder**
> SalesOrderLine PostSalesOrderLineForSalesOrder(ctx, body, companyId, salesOrderId, contentType)
Creates an object of type salesOrderLine in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesOrderLine**](salesOrderLine.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

