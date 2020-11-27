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

type GeneralLedgerEntry struct {
	Account *Account `json:"account,omitempty"`
	// The accountId property for the Dynamics 365 Business Central generalLedgerEntry entity
	AccountId string `json:"accountId,omitempty"`
	// The accountNumber property for the Dynamics 365 Business Central generalLedgerEntry entity
	AccountNumber string `json:"accountNumber,omitempty"`
	// The creditAmount property for the Dynamics 365 Business Central generalLedgerEntry entity
	CreditAmount float64 `json:"creditAmount,omitempty"`
	// The debitAmount property for the Dynamics 365 Business Central generalLedgerEntry entity
	DebitAmount float64 `json:"debitAmount,omitempty"`
	// The description property for the Dynamics 365 Business Central generalLedgerEntry entity
	Description string          `json:"description,omitempty"`
	Dimensions  []Dimensiontype `json:"dimensions,omitempty"`
	// The documentNumber property for the Dynamics 365 Business Central generalLedgerEntry entity
	DocumentNumber string `json:"documentNumber,omitempty"`
	// The documentType property for the Dynamics 365 Business Central generalLedgerEntry entity
	DocumentType string `json:"documentType,omitempty"`
	// The id property for the Dynamics 365 Business Central generalLedgerEntry entity
	Id int32 `json:"id,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central generalLedgerEntry entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The postingDate property for the Dynamics 365 Business Central generalLedgerEntry entity
	PostingDate string `json:"postingDate,omitempty"`
}
