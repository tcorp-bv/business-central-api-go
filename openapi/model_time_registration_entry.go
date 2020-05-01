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

// TimeRegistrationEntry struct for TimeRegistrationEntry
type TimeRegistrationEntry struct {
	// The id property for the Dynamics 365 Business Central timeRegistrationEntry entity
	Id string `json:"id,omitempty"`
	// The employeeId property for the Dynamics 365 Business Central timeRegistrationEntry entity
	EmployeeId *string `json:"employeeId,omitempty"`
	// The employeeNumber property for the Dynamics 365 Business Central timeRegistrationEntry entity
	EmployeeNumber *string `json:"employeeNumber,omitempty"`
	// The jobId property for the Dynamics 365 Business Central timeRegistrationEntry entity
	JobId *string `json:"jobId,omitempty"`
	// The jobNumber property for the Dynamics 365 Business Central timeRegistrationEntry entity
	JobNumber *string `json:"jobNumber,omitempty"`
	// The absence property for the Dynamics 365 Business Central timeRegistrationEntry entity
	Absence *string `json:"absence,omitempty"`
	// The lineNumber property for the Dynamics 365 Business Central timeRegistrationEntry entity
	LineNumber *int32 `json:"lineNumber,omitempty"`
	// The date property for the Dynamics 365 Business Central timeRegistrationEntry entity
	Date *time.Time `json:"date,omitempty"`
	// The quantity property for the Dynamics 365 Business Central timeRegistrationEntry entity
	Quantity *float32 `json:"quantity,omitempty"`
	// The status property for the Dynamics 365 Business Central timeRegistrationEntry entity
	Status *string `json:"status,omitempty"`
	// The unitOfMeasureId property for the Dynamics 365 Business Central timeRegistrationEntry entity
	UnitOfMeasureId *string           `json:"unitOfMeasureId,omitempty"`
	UnitOfMeasure   Unitofmeasuretype `json:"unitOfMeasure,omitempty"`
	Dimensions      []Dimensiontype   `json:"dimensions,omitempty"`
	// The lastModfiedDateTime property for the Dynamics 365 Business Central timeRegistrationEntry entity
	LastModfiedDateTime *time.Time `json:"lastModfiedDateTime,omitempty"`
	Project             Project    `json:"project,omitempty"`
}
