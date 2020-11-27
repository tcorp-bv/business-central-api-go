# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomerPayment**](CustomerPaymentApi.md#DeleteCustomerPayment) | **Delete** /companies({company_id})/customerPayments({customerPayment_id}) | Deletes an object of type customerPayment in Dynamics 365 Business Central
[**DeleteCustomerPaymentForCustomerPaymentJournal**](CustomerPaymentApi.md#DeleteCustomerPaymentForCustomerPaymentJournal) | **Delete** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id})/customerPayments({customerPayment_id}) | Deletes an object of type customerPayment in Dynamics 365 Business Central
[**GetCustomerPayment**](CustomerPaymentApi.md#GetCustomerPayment) | **Get** /companies({company_id})/customerPayments({customerPayment_id}) | Retrieve the properties and relationships of an object of type customerPayment for Dynamics 365 Business Central.
[**GetCustomerPaymentForCustomerPaymentJournal**](CustomerPaymentApi.md#GetCustomerPaymentForCustomerPaymentJournal) | **Get** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id})/customerPayments({customerPayment_id}) | Retrieve the properties and relationships of an object of type customerPayment for Dynamics 365 Business Central.
[**ListCustomerPayments**](CustomerPaymentApi.md#ListCustomerPayments) | **Get** /companies({company_id})/customerPayments | Returns a list of customerPayments
[**ListCustomerPaymentsForCustomerPaymentJournal**](CustomerPaymentApi.md#ListCustomerPaymentsForCustomerPaymentJournal) | **Get** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id})/customerPayments | Returns a list of customerPayments
[**PatchCustomerPayment**](CustomerPaymentApi.md#PatchCustomerPayment) | **Patch** /companies({company_id})/customerPayments({customerPayment_id}) | Updates an object of type customerPayment in Dynamics 365 Business Central
[**PatchCustomerPaymentForCustomerPaymentJournal**](CustomerPaymentApi.md#PatchCustomerPaymentForCustomerPaymentJournal) | **Patch** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id})/customerPayments({customerPayment_id}) | Updates an object of type customerPayment in Dynamics 365 Business Central
[**PostCustomerPayment**](CustomerPaymentApi.md#PostCustomerPayment) | **Post** /companies({company_id})/customerPayments | Creates an object of type customerPayment in Dynamics 365 Business Central
[**PostCustomerPaymentForCustomerPaymentJournal**](CustomerPaymentApi.md#PostCustomerPaymentForCustomerPaymentJournal) | **Post** /companies({company_id})/customerPaymentJournals({customerPaymentJournal_id})/customerPayments | Creates an object of type customerPayment in Dynamics 365 Business Central

# **DeleteCustomerPayment**
> DeleteCustomerPayment(ctx, companyId, customerPaymentId)
Deletes an object of type customerPayment in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentId** | [**string**](.md)| id for customerPayment | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerPaymentForCustomerPaymentJournal**
> DeleteCustomerPaymentForCustomerPaymentJournal(ctx, companyId, customerPaymentJournalId, customerPaymentId)
Deletes an object of type customerPayment in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
  **customerPaymentId** | [**string**](.md)| id for customerPayment | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerPayment**
> CustomerPayment GetCustomerPayment(ctx, companyId, customerPaymentId, optional)
Retrieve the properties and relationships of an object of type customerPayment for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentId** | [**string**](.md)| id for customerPayment | 
 **optional** | ***CustomerPaymentApiGetCustomerPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerPaymentApiGetCustomerPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerPayment**](customerPayment.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerPaymentForCustomerPaymentJournal**
> CustomerPayment GetCustomerPaymentForCustomerPaymentJournal(ctx, companyId, customerPaymentJournalId, customerPaymentId, optional)
Retrieve the properties and relationships of an object of type customerPayment for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
  **customerPaymentId** | [**string**](.md)| id for customerPayment | 
 **optional** | ***CustomerPaymentApiGetCustomerPaymentForCustomerPaymentJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerPaymentApiGetCustomerPaymentForCustomerPaymentJournalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CustomerPayment**](customerPayment.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerPayments**
> InlineResponse20013 ListCustomerPayments(ctx, companyId, optional)
Returns a list of customerPayments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***CustomerPaymentApiListCustomerPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerPaymentApiListCustomerPaymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerPaymentsForCustomerPaymentJournal**
> InlineResponse20013 ListCustomerPaymentsForCustomerPaymentJournal(ctx, companyId, customerPaymentJournalId, optional)
Returns a list of customerPayments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
 **optional** | ***CustomerPaymentApiListCustomerPaymentsForCustomerPaymentJournalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerPaymentApiListCustomerPaymentsForCustomerPaymentJournalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCustomerPayment**
> CustomerPayment PatchCustomerPayment(ctx, body, companyId, customerPaymentId, contentType, ifMatch)
Updates an object of type customerPayment in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentId** | [**string**](.md)| id for customerPayment | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**CustomerPayment**](customerPayment.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCustomerPaymentForCustomerPaymentJournal**
> CustomerPayment PatchCustomerPaymentForCustomerPaymentJournal(ctx, body, companyId, customerPaymentJournalId, customerPaymentId, contentType, ifMatch)
Updates an object of type customerPayment in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
  **customerPaymentId** | [**string**](.md)| id for customerPayment | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**CustomerPayment**](customerPayment.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomerPayment**
> CustomerPayment PostCustomerPayment(ctx, body, companyId, contentType)
Creates an object of type customerPayment in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **contentType** | **string**| application/json | 

### Return type

[**CustomerPayment**](customerPayment.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomerPaymentForCustomerPaymentJournal**
> CustomerPayment PostCustomerPaymentForCustomerPaymentJournal(ctx, body, companyId, customerPaymentJournalId, contentType)
Creates an object of type customerPayment in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **customerPaymentJournalId** | [**string**](.md)| id for customerPaymentJournal | 
  **contentType** | **string**| application/json | 

### Return type

[**CustomerPayment**](customerPayment.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

