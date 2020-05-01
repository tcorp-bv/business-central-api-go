/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// SalesInvoice struct for SalesInvoice
type SalesInvoice struct {
	// The id property for the Dynamics 365 Business Central salesInvoice entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central salesInvoice entity
	Number *string `json:"number,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central salesInvoice entity
	ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
	// The invoiceDate property for the Dynamics 365 Business Central salesInvoice entity
	InvoiceDate *time.Time `json:"invoiceDate,omitempty"`
	// The dueDate property for the Dynamics 365 Business Central salesInvoice entity
	DueDate *time.Time `json:"dueDate,omitempty"`
	// The customerPurchaseOrderReference property for the Dynamics 365 Business Central salesInvoice entity
	CustomerPurchaseOrderReference *string `json:"customerPurchaseOrderReference,omitempty"`
	// The customerId property for the Dynamics 365 Business Central salesInvoice entity
	CustomerId *string `json:"customerId,omitempty"`
	// The contactId property for the Dynamics 365 Business Central salesInvoice entity
	ContactId *string `json:"contactId,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central salesInvoice entity
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// The customerName property for the Dynamics 365 Business Central salesInvoice entity
	CustomerName *string `json:"customerName,omitempty"`
	// The billToName property for the Dynamics 365 Business Central salesInvoice entity
	BillToName *string `json:"billToName,omitempty"`
	// The billToCustomerId property for the Dynamics 365 Business Central salesInvoice entity
	BillToCustomerId *string `json:"billToCustomerId,omitempty"`
	// The billToCustomerNumber property for the Dynamics 365 Business Central salesInvoice entity
	BillToCustomerNumber *string `json:"billToCustomerNumber,omitempty"`
	// The shipToName property for the Dynamics 365 Business Central salesInvoice entity
	ShipToName *string `json:"shipToName,omitempty"`
	// The shipToContact property for the Dynamics 365 Business Central salesInvoice entity
	ShipToContact         *string           `json:"shipToContact,omitempty"`
	SellingPostalAddress  Postaladdresstype `json:"sellingPostalAddress,omitempty"`
	BillingPostalAddress  Postaladdresstype `json:"billingPostalAddress,omitempty"`
	ShippingPostalAddress Postaladdresstype `json:"shippingPostalAddress,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central salesInvoice entity
	CurrencyId *string `json:"currencyId,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central salesInvoice entity
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The orderId property for the Dynamics 365 Business Central salesInvoice entity
	OrderId *string `json:"orderId,omitempty"`
	// The orderNumber property for the Dynamics 365 Business Central salesInvoice entity
	OrderNumber *string `json:"orderNumber,omitempty"`
	// The paymentTermsId property for the Dynamics 365 Business Central salesInvoice entity
	PaymentTermsId *string `json:"paymentTermsId,omitempty"`
	// The shipmentMethodId property for the Dynamics 365 Business Central salesInvoice entity
	ShipmentMethodId *string `json:"shipmentMethodId,omitempty"`
	// The salesperson property for the Dynamics 365 Business Central salesInvoice entity
	Salesperson *string `json:"salesperson,omitempty"`
	// The pricesIncludeTax property for the Dynamics 365 Business Central salesInvoice entity
	PricesIncludeTax *bool `json:"pricesIncludeTax,omitempty"`
	// The remainingAmount property for the Dynamics 365 Business Central salesInvoice entity
	RemainingAmount *float32 `json:"remainingAmount,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesInvoice entity
	DiscountAmount *float32 `json:"discountAmount,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesInvoice entity
	DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central salesInvoice entity
	TotalAmountExcludingTax *float32 `json:"totalAmountExcludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesInvoice entity
	TotalTaxAmount *float32 `json:"totalTaxAmount,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central salesInvoice entity
	TotalAmountIncludingTax *float32 `json:"totalAmountIncludingTax,omitempty"`
	// The status property for the Dynamics 365 Business Central salesInvoice entity
	Status *string `json:"status,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central salesInvoice entity
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central salesInvoice entity
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The email property for the Dynamics 365 Business Central salesInvoice entity
	Email             *string             `json:"email,omitempty"`
	SalesInvoiceLines *[]SalesInvoiceLine `json:"salesInvoiceLines,omitempty"`
	PdfDocument       *[]PdfDocument      `json:"pdfDocument,omitempty"`
	Customer          Customer            `json:"customer,omitempty"`
	Currency          Currency            `json:"currency,omitempty"`
	PaymentTerm       PaymentTerm         `json:"paymentTerm,omitempty"`
	ShipmentMethod    ShipmentMethod      `json:"shipmentMethod,omitempty"`
}
