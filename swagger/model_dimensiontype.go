/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Dimensiontype struct {
	// The code property for the Dynamics 365 Business Central dimensiontype entity
	Code string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central dimensiontype entity
	DisplayName string `json:"displayName,omitempty"`
	// The valueCode property for the Dynamics 365 Business Central dimensiontype entity
	ValueCode string `json:"valueCode,omitempty"`
	// The valueDisplayName property for the Dynamics 365 Business Central dimensiontype entity
	ValueDisplayName string `json:"valueDisplayName,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
}