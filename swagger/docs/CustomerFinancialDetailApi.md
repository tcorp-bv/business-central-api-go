# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerFinancialDetail**](CustomerFinancialDetailApi.md#GetCustomerFinancialDetail) | **Get** /companies({company_id})/customerFinancialDetails({customerFinancialDetail_id}) | Retrieve the properties and relationships of an object of type customerFinancialDetail for Dynamics 365 Business Central.
[**GetCustomerFinancialDetailForCustomer**](CustomerFinancialDetailApi.md#GetCustomerFinancialDetailForCustomer) | **Get** /companies({company_id})/customers({customer_id})/customerFinancialDetails({customerFinancialDetail_id}) | Retrieve the properties and relationships of an object of type customerFinancialDetail for Dynamics 365 Business Central.
[**ListCustomerFinancialDetails**](CustomerFinancialDetailApi.md#ListCustomerFinancialDetails) | **Get** /companies({company_id})/customerFinancialDetails | Returns a list of customerFinancialDetails
[**ListCustomerFinancialDetailsForCustomer**](CustomerFinancialDetailApi.md#ListCustomerFinancialDetailsForCustomer) | **Get** /companies({company_id})/customers({customer_id})/customerFinancialDetails | Returns a list of customerFinancialDetails

# **GetCustomerFinancialDetail**
> CustomerFinancialDetail GetCustomerFinancialDetail(ctx, companyId, customerFinancialDetailId, optional)
Retrieve the properties and relationships of an object of type customerFinancialDetail for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerFinancialDetailId** | [**string**](.md)| id for customerFinancialDetail | 
 **optional** | ***CustomerFinancialDetailApiGetCustomerFinancialDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerFinancialDetailApiGetCustomerFinancialDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerFinancialDetail**](customerFinancialDetail.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerFinancialDetailForCustomer**
> CustomerFinancialDetail GetCustomerFinancialDetailForCustomer(ctx, companyId, customerId, customerFinancialDetailId, optional)
Retrieve the properties and relationships of an object of type customerFinancialDetail for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **customerFinancialDetailId** | [**string**](.md)| id for customerFinancialDetail | 
 **optional** | ***CustomerFinancialDetailApiGetCustomerFinancialDetailForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerFinancialDetailApiGetCustomerFinancialDetailForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerFinancialDetail**](customerFinancialDetail.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerFinancialDetails**
> InlineResponse20011 ListCustomerFinancialDetails(ctx, companyId, optional)
Returns a list of customerFinancialDetails

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***CustomerFinancialDetailApiListCustomerFinancialDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerFinancialDetailApiListCustomerFinancialDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerFinancialDetailsForCustomer**
> InlineResponse20011 ListCustomerFinancialDetailsForCustomer(ctx, companyId, customerId, optional)
Returns a list of customerFinancialDetails

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
 **optional** | ***CustomerFinancialDetailApiListCustomerFinancialDetailsForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerFinancialDetailApiListCustomerFinancialDetailsForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

