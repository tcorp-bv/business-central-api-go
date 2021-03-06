# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCashFlowStatement**](CashFlowStatementApi.md#GetCashFlowStatement) | **Get** /companies({company_id})/cashFlowStatement({cashFlowStatement_lineNumber}) | Retrieve the properties and relationships of an object of type cashFlowStatement for Dynamics 365 Business Central.
[**ListCashFlowStatement**](CashFlowStatementApi.md#ListCashFlowStatement) | **Get** /companies({company_id})/cashFlowStatement | Returns a list of cashFlowStatement

# **GetCashFlowStatement**
> CashFlowStatement GetCashFlowStatement(ctx, companyId, cashFlowStatementLineNumber, optional)
Retrieve the properties and relationships of an object of type cashFlowStatement for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **cashFlowStatementLineNumber** | **int32**| lineNumber for cashFlowStatement | 
 **optional** | ***CashFlowStatementApiGetCashFlowStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CashFlowStatementApiGetCashFlowStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**CashFlowStatement**](cashFlowStatement.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCashFlowStatement**
> InlineResponse2007 ListCashFlowStatement(ctx, companyId, optional)
Returns a list of cashFlowStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***CashFlowStatementApiListCashFlowStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CashFlowStatementApiListCashFlowStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

