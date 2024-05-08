// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthorizationStrategy string

// Enum values for AuthorizationStrategy
const (
	AuthorizationStrategySmartv1 AuthorizationStrategy = "SMART_ON_FHIR_V1"
	AuthorizationStrategyAwsAuth AuthorizationStrategy = "AWS_AUTH"
)

// Values returns all known values for AuthorizationStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthorizationStrategy) Values() []AuthorizationStrategy {
	return []AuthorizationStrategy{
		"SMART_ON_FHIR_V1",
		"AWS_AUTH",
	}
}

type CmkType string

// Enum values for CmkType
const (
	CmkTypeCmCmk CmkType = "CUSTOMER_MANAGED_KMS_KEY"
	CmkTypeAoCmk CmkType = "AWS_OWNED_KMS_KEY"
)

// Values returns all known values for CmkType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CmkType) Values() []CmkType {
	return []CmkType{
		"CUSTOMER_MANAGED_KMS_KEY",
		"AWS_OWNED_KMS_KEY",
	}
}

type DatastoreStatus string

// Enum values for DatastoreStatus
const (
	DatastoreStatusCreating     DatastoreStatus = "CREATING"
	DatastoreStatusActive       DatastoreStatus = "ACTIVE"
	DatastoreStatusDeleting     DatastoreStatus = "DELETING"
	DatastoreStatusDeleted      DatastoreStatus = "DELETED"
	DatastoreStatusCreateFailed DatastoreStatus = "CREATE_FAILED"
)

// Values returns all known values for DatastoreStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DatastoreStatus) Values() []DatastoreStatus {
	return []DatastoreStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"CREATE_FAILED",
	}
}

type ErrorCategory string

// Enum values for ErrorCategory
const (
	ErrorCategoryRetryableError    ErrorCategory = "RETRYABLE_ERROR"
	ErrorCategoryNonRetryableError ErrorCategory = "NON_RETRYABLE_ERROR"
)

// Values returns all known values for ErrorCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCategory) Values() []ErrorCategory {
	return []ErrorCategory{
		"RETRYABLE_ERROR",
		"NON_RETRYABLE_ERROR",
	}
}

type FHIRVersion string

// Enum values for FHIRVersion
const (
	FHIRVersionR4 FHIRVersion = "R4"
)

// Values returns all known values for FHIRVersion. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FHIRVersion) Values() []FHIRVersion {
	return []FHIRVersion{
		"R4",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted           JobStatus = "SUBMITTED"
	JobStatusInProgress          JobStatus = "IN_PROGRESS"
	JobStatusCompletedWithErrors JobStatus = "COMPLETED_WITH_ERRORS"
	JobStatusCompleted           JobStatus = "COMPLETED"
	JobStatusFailed              JobStatus = "FAILED"
	JobStatusCancelSubmitted     JobStatus = "CANCEL_SUBMITTED"
	JobStatusCancelInProgress    JobStatus = "CANCEL_IN_PROGRESS"
	JobStatusCancelCompleted     JobStatus = "CANCEL_COMPLETED"
	JobStatusCancelFailed        JobStatus = "CANCEL_FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED_WITH_ERRORS",
		"COMPLETED",
		"FAILED",
		"CANCEL_SUBMITTED",
		"CANCEL_IN_PROGRESS",
		"CANCEL_COMPLETED",
		"CANCEL_FAILED",
	}
}

type PreloadDataType string

// Enum values for PreloadDataType
const (
	PreloadDataTypeSynthea PreloadDataType = "SYNTHEA"
)

// Values returns all known values for PreloadDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PreloadDataType) Values() []PreloadDataType {
	return []PreloadDataType{
		"SYNTHEA",
	}
}
