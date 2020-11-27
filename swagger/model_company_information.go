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
	"os"
	"time"
)

type CompanyInformation struct {
	// The id property for the Dynamics 365 Business Central companyInformation entity
	Id string `json:"id,omitempty"`
	// The displayName property for the Dynamics 365 Business Central companyInformation entity
	DisplayName string             `json:"displayName,omitempty"`
	Address     *Postaladdresstype `json:"address,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central companyInformation entity
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// The faxNumber property for the Dynamics 365 Business Central companyInformation entity
	FaxNumber string `json:"faxNumber,omitempty"`
	// The email property for the Dynamics 365 Business Central companyInformation entity
	Email string `json:"email,omitempty"`
	// The website property for the Dynamics 365 Business Central companyInformation entity
	Website string `json:"website,omitempty"`
	// The taxRegistrationNumber property for the Dynamics 365 Business Central companyInformation entity
	TaxRegistrationNumber string `json:"taxRegistrationNumber,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central companyInformation entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The currentFiscalYearStartDate property for the Dynamics 365 Business Central companyInformation entity
	CurrentFiscalYearStartDate time.Time `json:"currentFiscalYearStartDate,omitempty"`
	// The industry property for the Dynamics 365 Business Central companyInformation entity
	Industry string `json:"industry,omitempty"`
	// The picture property for the Dynamics 365 Business Central companyInformation entity
	Picture **os.File `json:"picture,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central companyInformation entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
}
