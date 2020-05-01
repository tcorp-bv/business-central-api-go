# SalesOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**Number** | Pointer to **string** | The number property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**ExternalDocumentNumber** | Pointer to **string** | The externalDocumentNumber property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**OrderDate** | Pointer to [**time.Time**](time.Time.md) | The orderDate property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**CustomerId** | Pointer to **string** | The customerId property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**ContactId** | Pointer to **string** | The contactId property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**CustomerNumber** | Pointer to **string** | The customerNumber property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**CustomerName** | Pointer to **string** | The customerName property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**BillToName** | Pointer to **string** | The billToName property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**BillToCustomerId** | Pointer to **string** | The billToCustomerId property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**BillToCustomerNumber** | Pointer to **string** | The billToCustomerNumber property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**ShipToName** | Pointer to **string** | The shipToName property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**ShipToContact** | Pointer to **string** | The shipToContact property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**SellingPostalAddress** | [**Postaladdresstype**](postaladdresstype.md) |  | [optional] 
**BillingPostalAddress** | [**Postaladdresstype**](postaladdresstype.md) |  | [optional] 
**ShippingPostalAddress** | [**Postaladdresstype**](postaladdresstype.md) |  | [optional] 
**CurrencyId** | Pointer to **string** | The currencyId property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**CurrencyCode** | Pointer to **string** | The currencyCode property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**PricesIncludeTax** | Pointer to **bool** | The pricesIncludeTax property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**PaymentTermsId** | Pointer to **string** | The paymentTermsId property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**ShipmentMethodId** | Pointer to **string** | The shipmentMethodId property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**Salesperson** | Pointer to **string** | The salesperson property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**PartialShipping** | Pointer to **bool** | The partialShipping property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**RequestedDeliveryDate** | Pointer to [**time.Time**](time.Time.md) | The requestedDeliveryDate property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**DiscountAmount** | Pointer to **float32** | The discountAmount property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**DiscountAppliedBeforeTax** | Pointer to **bool** | The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**TotalAmountExcludingTax** | Pointer to **float32** | The totalAmountExcludingTax property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**TotalTaxAmount** | Pointer to **float32** | The totalTaxAmount property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**TotalAmountIncludingTax** | Pointer to **float32** | The totalAmountIncludingTax property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**FullyShipped** | Pointer to **bool** | The fullyShipped property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**Status** | Pointer to **string** | The status property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | The lastModifiedDateTime property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**PhoneNumber** | Pointer to **string** | The phoneNumber property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**Email** | Pointer to **string** | The email property for the Dynamics 365 Business Central salesOrder entity | [optional] 
**SalesOrderLines** | Pointer to [**[]SalesOrderLine**](salesOrderLine.md) |  | [optional] 
**Customer** | [**Customer**](customer.md) |  | [optional] 
**Currency** | [**Currency**](currency.md) |  | [optional] 
**PaymentTerm** | [**PaymentTerm**](paymentTerm.md) |  | [optional] 
**ShipmentMethod** | [**ShipmentMethod**](shipmentMethod.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


