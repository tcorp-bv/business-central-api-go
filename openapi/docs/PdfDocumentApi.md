# \PdfDocumentApi

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPdfDocument**](PdfDocumentApi.md#GetPdfDocument) | **Get** /companies({company_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForPurchaseInvoice**](PdfDocumentApi.md#GetPdfDocumentForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForSalesCreditMemo**](PdfDocumentApi.md#GetPdfDocumentForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForSalesInvoice**](PdfDocumentApi.md#GetPdfDocumentForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForSalesQuote**](PdfDocumentApi.md#GetPdfDocumentForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**ListPdfDocument**](PdfDocumentApi.md#ListPdfDocument) | **Get** /companies({company_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForPurchaseInvoice**](PdfDocumentApi.md#ListPdfDocumentForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForSalesCreditMemo**](PdfDocumentApi.md#ListPdfDocumentForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForSalesInvoice**](PdfDocumentApi.md#ListPdfDocumentForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForSalesQuote**](PdfDocumentApi.md#ListPdfDocumentForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/pdfDocument | Returns a list of pdfDocument



## GetPdfDocument

> PdfDocument GetPdfDocument(ctx, companyId, pdfDocumentId, optional)

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPdfDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdfDocumentForPurchaseInvoice

> PdfDocument GetPdfDocumentForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, pdfDocumentId, optional)

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
**pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPdfDocumentForPurchaseInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdfDocumentForSalesCreditMemo

> PdfDocument GetPdfDocumentForSalesCreditMemo(ctx, companyId, salesCreditMemoId, pdfDocumentId, optional)

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
**pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPdfDocumentForSalesCreditMemoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdfDocumentForSalesInvoice

> PdfDocument GetPdfDocumentForSalesInvoice(ctx, companyId, salesInvoiceId, pdfDocumentId, optional)

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
**pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPdfDocumentForSalesInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdfDocumentForSalesQuote

> PdfDocument GetPdfDocumentForSalesQuote(ctx, companyId, salesQuoteId, pdfDocumentId, optional)

Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesQuoteId** | [**string**](.md)| id for salesQuote | 
**pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***GetPdfDocumentForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPdfDocumentForSalesQuoteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPdfDocument

> InlineResponse20010 ListPdfDocument(ctx, companyId, optional)

Returns a list of pdfDocument

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
 **optional** | ***ListPdfDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPdfDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPdfDocumentForPurchaseInvoice

> InlineResponse20010 ListPdfDocumentForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)

Returns a list of pdfDocument

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***ListPdfDocumentForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPdfDocumentForPurchaseInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPdfDocumentForSalesCreditMemo

> InlineResponse20010 ListPdfDocumentForSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)

Returns a list of pdfDocument

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***ListPdfDocumentForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPdfDocumentForSalesCreditMemoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPdfDocumentForSalesInvoice

> InlineResponse20010 ListPdfDocumentForSalesInvoice(ctx, companyId, salesInvoiceId, optional)

Returns a list of pdfDocument

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***ListPdfDocumentForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPdfDocumentForSalesInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPdfDocumentForSalesQuote

> InlineResponse20010 ListPdfDocumentForSalesQuote(ctx, companyId, salesQuoteId, optional)

Returns a list of pdfDocument

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | [**string**](.md)| id for company | 
**salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***ListPdfDocumentForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPdfDocumentForSalesQuoteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

