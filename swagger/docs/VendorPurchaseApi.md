# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendorPurchase**](VendorPurchaseApi.md#GetVendorPurchase) | **Get** /companies({company_id})/vendorPurchases({vendorPurchase_vendorId},&#x27;{vendorPurchase_vendorNumber}&#x27;,&#x27;{vendorPurchase_name}&#x27;) | Retrieve the properties and relationships of an object of type vendorPurchase for Dynamics 365 Business Central.
[**ListVendorPurchases**](VendorPurchaseApi.md#ListVendorPurchases) | **Get** /companies({company_id})/vendorPurchases | Returns a list of vendorPurchases

# **GetVendorPurchase**
> VendorPurchase GetVendorPurchase(ctx, companyId, vendorPurchaseVendorId, vendorPurchaseVendorNumber, vendorPurchaseName, optional)
Retrieve the properties and relationships of an object of type vendorPurchase for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorPurchaseVendorId** | [**string**](.md)| vendorId for vendorPurchase | 
  **vendorPurchaseVendorNumber** | **string**| vendorNumber for vendorPurchase | 
  **vendorPurchaseName** | **string**| name for vendorPurchase | 
 **optional** | ***VendorPurchaseApiGetVendorPurchaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorPurchaseApiGetVendorPurchaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**VendorPurchase**](vendorPurchase.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVendorPurchases**
> InlineResponse20050 ListVendorPurchases(ctx, companyId, optional)
Returns a list of vendorPurchases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***VendorPurchaseApiListVendorPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorPurchaseApiListVendorPurchasesOpts struct
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

