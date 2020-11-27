# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrder**](SalesOrderApi.md#DeleteSalesOrder) | **Delete** /companies({company_id})/salesOrders({salesOrder_id}) | Deletes an object of type salesOrder in Dynamics 365 Business Central
[**GetSalesOrder**](SalesOrderApi.md#GetSalesOrder) | **Get** /companies({company_id})/salesOrders({salesOrder_id}) | Retrieve the properties and relationships of an object of type salesOrder for Dynamics 365 Business Central.
[**ListSalesOrders**](SalesOrderApi.md#ListSalesOrders) | **Get** /companies({company_id})/salesOrders | Returns a list of salesOrders
[**PatchSalesOrder**](SalesOrderApi.md#PatchSalesOrder) | **Patch** /companies({company_id})/salesOrders({salesOrder_id}) | Updates an object of type salesOrder in Dynamics 365 Business Central
[**PostSalesOrder**](SalesOrderApi.md#PostSalesOrder) | **Post** /companies({company_id})/salesOrders | Creates an object of type salesOrder in Dynamics 365 Business Central
[**ShipAndInvoiceActionSalesOrders**](SalesOrderApi.md#ShipAndInvoiceActionSalesOrders) | **Post** /companies({company_id})/salesOrders({salesOrder_id})/Microsoft.NAV.shipAndInvoice | Performs the shipAndInvoice action for salesOrders entity

# **DeleteSalesOrder**
> DeleteSalesOrder(ctx, companyId, salesOrderId)
Deletes an object of type salesOrder in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSalesOrder**
> SalesOrder GetSalesOrder(ctx, companyId, salesOrderId, optional)
Retrieve the properties and relationships of an object of type salesOrder for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
 **optional** | ***SalesOrderApiGetSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesOrderApiGetSalesOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**SalesOrder**](salesOrder.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSalesOrders**
> InlineResponse20042 ListSalesOrders(ctx, companyId, optional)
Returns a list of salesOrders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***SalesOrderApiListSalesOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesOrderApiListSalesOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSalesOrder**
> SalesOrder PatchSalesOrder(ctx, body, companyId, salesOrderId, contentType, ifMatch)
Updates an object of type salesOrder in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**SalesOrder**](salesOrder.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSalesOrder**
> SalesOrder PostSalesOrder(ctx, body, companyId, contentType)
Creates an object of type salesOrder in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**SalesOrder**](salesOrder.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShipAndInvoiceActionSalesOrders**
> ShipAndInvoiceActionSalesOrders(ctx, companyId, salesOrderId)
Performs the shipAndInvoice action for salesOrders entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesOrderId** | [**string**](.md)| id for salesOrder | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

