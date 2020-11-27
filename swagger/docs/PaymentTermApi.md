# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePaymentTerm**](PaymentTermApi.md#DeletePaymentTerm) | **Delete** /companies({company_id})/paymentTerms({paymentTerm_id}) | Deletes an object of type paymentTerm in Dynamics 365 Business Central
[**GetPaymentTerm**](PaymentTermApi.md#GetPaymentTerm) | **Get** /companies({company_id})/paymentTerms({paymentTerm_id}) | Retrieve the properties and relationships of an object of type paymentTerm for Dynamics 365 Business Central.
[**ListPaymentTerms**](PaymentTermApi.md#ListPaymentTerms) | **Get** /companies({company_id})/paymentTerms | Returns a list of paymentTerms
[**PatchPaymentTerm**](PaymentTermApi.md#PatchPaymentTerm) | **Patch** /companies({company_id})/paymentTerms({paymentTerm_id}) | Updates an object of type paymentTerm in Dynamics 365 Business Central
[**PostPaymentTerm**](PaymentTermApi.md#PostPaymentTerm) | **Post** /companies({company_id})/paymentTerms | Creates an object of type paymentTerm in Dynamics 365 Business Central

# **DeletePaymentTerm**
> DeletePaymentTerm(ctx, companyId, paymentTermId)
Deletes an object of type paymentTerm in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentTermId** | [**string**](.md)| id for paymentTerm | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentTerm**
> PaymentTerm GetPaymentTerm(ctx, companyId, paymentTermId, optional)
Retrieve the properties and relationships of an object of type paymentTerm for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **paymentTermId** | [**string**](.md)| id for paymentTerm | 
 **optional** | ***PaymentTermApiGetPaymentTermOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentTermApiGetPaymentTermOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PaymentTerm**](paymentTerm.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentTerms**
> InlineResponse20026 ListPaymentTerms(ctx, companyId, optional)
Returns a list of paymentTerms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***PaymentTermApiListPaymentTermsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentTermApiListPaymentTermsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPaymentTerm**
> PaymentTerm PatchPaymentTerm(ctx, body, companyId, paymentTermId, contentType, ifMatch)
Updates an object of type paymentTerm in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **paymentTermId** | [**string**](.md)| id for paymentTerm | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**PaymentTerm**](paymentTerm.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPaymentTerm**
> PaymentTerm PostPaymentTerm(ctx, body, companyId, contentType)
Creates an object of type paymentTerm in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**PaymentTerm**](paymentTerm.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

