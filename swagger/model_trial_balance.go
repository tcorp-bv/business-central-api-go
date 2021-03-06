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

type TrialBalance struct {
	Account *Account `json:"account,omitempty"`
	// The accountId property for the Dynamics 365 Business Central trialBalance entity
	AccountId string `json:"accountId,omitempty"`
	// The accountType property for the Dynamics 365 Business Central trialBalance entity
	AccountType string `json:"accountType,omitempty"`
	// The balanceAtDateCredit property for the Dynamics 365 Business Central trialBalance entity
	BalanceAtDateCredit string `json:"balanceAtDateCredit,omitempty"`
	// The balanceAtDateDebit property for the Dynamics 365 Business Central trialBalance entity
	BalanceAtDateDebit string `json:"balanceAtDateDebit,omitempty"`
	// The dateFilter property for the Dynamics 365 Business Central trialBalance entity
	DateFilter time.Time `json:"dateFilter,omitempty"`
	// The display property for the Dynamics 365 Business Central trialBalance entity
	Display string `json:"display,omitempty"`
	// The number property for the Dynamics 365 Business Central trialBalance entity
	Number string `json:"number,omitempty"`
	// The totalCredit property for the Dynamics 365 Business Central trialBalance entity
	TotalCredit string `json:"totalCredit,omitempty"`
	// The totalDebit property for the Dynamics 365 Business Central trialBalance entity
	TotalDebit string `json:"totalDebit,omitempty"`
}
