/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

type SalesOrder struct {
	// The billToCustomerId property for the Dynamics 365 Business Central salesOrder entity
	BillToCustomerId string `json:"billToCustomerId,omitempty"`
	// The billToCustomerNumber property for the Dynamics 365 Business Central salesOrder entity
	BillToCustomerNumber string `json:"billToCustomerNumber,omitempty"`
	// The billToName property for the Dynamics 365 Business Central salesOrder entity
	BillToName           string             `json:"billToName,omitempty"`
	BillingPostalAddress *Postaladdresstype `json:"billingPostalAddress,omitempty"`
	// The contactId property for the Dynamics 365 Business Central salesOrder entity
	ContactId string    `json:"contactId,omitempty"`
	Currency  *Currency `json:"currency,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central salesOrder entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central salesOrder entity
	CurrencyId string    `json:"currencyId,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`
	// The customerId property for the Dynamics 365 Business Central salesOrder entity
	CustomerId string `json:"customerId,omitempty"`
	// The customerName property for the Dynamics 365 Business Central salesOrder entity
	CustomerName string `json:"customerName,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central salesOrder entity
	CustomerNumber string `json:"customerNumber,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesOrder entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesOrder entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The email property for the Dynamics 365 Business Central salesOrder entity
	Email string `json:"email,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central salesOrder entity
	ExternalDocumentNumber string `json:"externalDocumentNumber,omitempty"`
	// The fullyShipped property for the Dynamics 365 Business Central salesOrder entity
	FullyShipped bool `json:"fullyShipped,omitempty"`
	// The id property for the Dynamics 365 Business Central salesOrder entity
	Id string `json:"id,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central salesOrder entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The number property for the Dynamics 365 Business Central salesOrder entity
	Number string `json:"number,omitempty"`
	// The orderDate property for the Dynamics 365 Business Central salesOrder entity
	OrderDate string `json:"orderDate,omitempty"`
	// The partialShipping property for the Dynamics 365 Business Central salesOrder entity
	PartialShipping bool         `json:"partialShipping,omitempty"`
	PaymentTerm     *PaymentTerm `json:"paymentTerm,omitempty"`
	// The paymentTermsId property for the Dynamics 365 Business Central salesOrder entity
	PaymentTermsId string `json:"paymentTermsId,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central salesOrder entity
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// The pricesIncludeTax property for the Dynamics 365 Business Central salesOrder entity
	PricesIncludeTax bool `json:"pricesIncludeTax,omitempty"`
	// The requestedDeliveryDate property for the Dynamics 365 Business Central salesOrder entity
	RequestedDeliveryDate string           `json:"requestedDeliveryDate,omitempty"`
	SalesOrderLines       []SalesOrderLine `json:"salesOrderLines,omitempty"`
	// The salesperson property for the Dynamics 365 Business Central salesOrder entity
	Salesperson          string             `json:"salesperson,omitempty"`
	SellingPostalAddress *Postaladdresstype `json:"sellingPostalAddress,omitempty"`
	// The shipToContact property for the Dynamics 365 Business Central salesOrder entity
	ShipToContact string `json:"shipToContact,omitempty"`
	// The shipToName property for the Dynamics 365 Business Central salesOrder entity
	ShipToName     string          `json:"shipToName,omitempty"`
	ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
	// The shipmentMethodId property for the Dynamics 365 Business Central salesOrder entity
	ShipmentMethodId      string             `json:"shipmentMethodId,omitempty"`
	ShippingPostalAddress *Postaladdresstype `json:"shippingPostalAddress,omitempty"`
	// The status property for the Dynamics 365 Business Central salesOrder entity
	Status string `json:"status,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central salesOrder entity
	TotalAmountExcludingTax float64 `json:"totalAmountExcludingTax,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central salesOrder entity
	TotalAmountIncludingTax float64 `json:"totalAmountIncludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesOrder entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
}
