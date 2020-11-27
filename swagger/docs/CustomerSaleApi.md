# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerSale**](CustomerSaleApi.md#GetCustomerSale) | **Get** /companies({company_id})/customerSales({customerSale_customerId},&#x27;{customerSale_customerNumber}&#x27;,&#x27;{customerSale_name}&#x27;) | Retrieve the properties and relationships of an object of type customerSale for Dynamics 365 Business Central.
[**ListCustomerSales**](CustomerSaleApi.md#ListCustomerSales) | **Get** /companies({company_id})/customerSales | Returns a list of customerSales

# **GetCustomerSale**
> CustomerSale GetCustomerSale(ctx, companyId, customerSaleCustomerId, customerSaleCustomerNumber, customerSaleName, optional)
Retrieve the properties and relationships of an object of type customerSale for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerSaleCustomerId** | [**string**](.md)| customerId for customerSale | 
  **customerSaleCustomerNumber** | **string**| customerNumber for customerSale | 
  **customerSaleName** | **string**| name for customerSale | 
 **optional** | ***CustomerSaleApiGetCustomerSaleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerSaleApiGetCustomerSaleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerSale**](customerSale.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerSales**
> InlineResponse20050 ListCustomerSales(ctx, companyId, optional)
Returns a list of customerSales

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***CustomerSaleApiListCustomerSalesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerSaleApiListCustomerSalesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

