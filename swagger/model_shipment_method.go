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

type ShipmentMethod struct {
	// The id property for the Dynamics 365 Business Central shipmentMethod entity
	Id string `json:"id,omitempty"`
	// The code property for the Dynamics 365 Business Central shipmentMethod entity
	Code string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central shipmentMethod entity
	DisplayName string `json:"displayName,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central shipmentMethod entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
}
