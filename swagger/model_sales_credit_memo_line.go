/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SalesCreditMemoLine struct {
	Account *Account `json:"account,omitempty"`
	// The accountId property for the Dynamics 365 Business Central salesCreditMemoLine entity
	AccountId string `json:"accountId,omitempty"`
	// The amountExcludingTax property for the Dynamics 365 Business Central salesCreditMemoLine entity
	AmountExcludingTax float64 `json:"amountExcludingTax,omitempty"`
	// The amountIncludingTax property for the Dynamics 365 Business Central salesCreditMemoLine entity
	AmountIncludingTax float64 `json:"amountIncludingTax,omitempty"`
	// The description property for the Dynamics 365 Business Central salesCreditMemoLine entity
	Description string `json:"description,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesCreditMemoLine entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central salesCreditMemoLine entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The discountPercent property for the Dynamics 365 Business Central salesCreditMemoLine entity
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	// The documentId property for the Dynamics 365 Business Central salesCreditMemoLine entity
	DocumentId string `json:"documentId,omitempty"`
	// The id property for the Dynamics 365 Business Central salesCreditMemoLine entity
	Id string `json:"id,omitempty"`
	// The invoiceDiscountAllocation property for the Dynamics 365 Business Central salesCreditMemoLine entity
	InvoiceDiscountAllocation float64 `json:"invoiceDiscountAllocation,omitempty"`
	Item                      *Item   `json:"item,omitempty"`
	// The itemId property for the Dynamics 365 Business Central salesCreditMemoLine entity
	ItemId      string                         `json:"itemId,omitempty"`
	LineDetails *Documentlineobjectdetailstype `json:"lineDetails,omitempty"`
	// The lineType property for the Dynamics 365 Business Central salesCreditMemoLine entity
	LineType string `json:"lineType,omitempty"`
	// The netAmount property for the Dynamics 365 Business Central salesCreditMemoLine entity
	NetAmount float64 `json:"netAmount,omitempty"`
	// The netAmountIncludingTax property for the Dynamics 365 Business Central salesCreditMemoLine entity
	NetAmountIncludingTax float64 `json:"netAmountIncludingTax,omitempty"`
	// The netTaxAmount property for the Dynamics 365 Business Central salesCreditMemoLine entity
	NetTaxAmount float64 `json:"netTaxAmount,omitempty"`
	// The quantity property for the Dynamics 365 Business Central salesCreditMemoLine entity
	Quantity float64 `json:"quantity,omitempty"`
	// The sequence property for the Dynamics 365 Business Central salesCreditMemoLine entity
	Sequence int32 `json:"sequence,omitempty"`
	// The shipmentDate property for the Dynamics 365 Business Central salesCreditMemoLine entity
	ShipmentDate string `json:"shipmentDate,omitempty"`
	// The taxCode property for the Dynamics 365 Business Central salesCreditMemoLine entity
	TaxCode string `json:"taxCode,omitempty"`
	// The taxPercent property for the Dynamics 365 Business Central salesCreditMemoLine entity
	TaxPercent float64 `json:"taxPercent,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesCreditMemoLine entity
	TotalTaxAmount float64            `json:"totalTaxAmount,omitempty"`
	UnitOfMeasure  *Unitofmeasuretype `json:"unitOfMeasure,omitempty"`
	// The unitOfMeasureId property for the Dynamics 365 Business Central salesCreditMemoLine entity
	UnitOfMeasureId string `json:"unitOfMeasureId,omitempty"`
	// The unitPrice property for the Dynamics 365 Business Central salesCreditMemoLine entity
	UnitPrice float64 `json:"unitPrice,omitempty"`
}
