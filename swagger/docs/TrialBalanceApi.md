# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrialBalance**](TrialBalanceApi.md#GetTrialBalance) | **Get** /companies({company_id})/trialBalance(&#x27;{trialBalance_number}&#x27;) | Retrieve the properties and relationships of an object of type trialBalance for Dynamics 365 Business Central.
[**ListTrialBalance**](TrialBalanceApi.md#ListTrialBalance) | **Get** /companies({company_id})/trialBalance | Returns a list of trialBalance

# **GetTrialBalance**
> TrialBalance GetTrialBalance(ctx, companyId, trialBalanceNumber, optional)
Retrieve the properties and relationships of an object of type trialBalance for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **trialBalanceNumber** | **string**| number for trialBalance | 
 **optional** | ***TrialBalanceApiGetTrialBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrialBalanceApiGetTrialBalanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**TrialBalance**](trialBalance.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrialBalance**
> InlineResponse20048 ListTrialBalance(ctx, companyId, optional)
Returns a list of trialBalance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***TrialBalanceApiListTrialBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrialBalanceApiListTrialBalanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **expand** | [**optional.Interface of []string**](string.md)| Entities to expand | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

