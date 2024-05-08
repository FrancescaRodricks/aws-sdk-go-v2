// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChangeDirectionEnum string

// Enum values for ChangeDirectionEnum
const (
	ChangeDirectionEnumIncrease ChangeDirectionEnum = "INCREASE"
	ChangeDirectionEnumDecrease ChangeDirectionEnum = "DECREASE"
)

// Values returns all known values for ChangeDirectionEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChangeDirectionEnum) Values() []ChangeDirectionEnum {
	return []ChangeDirectionEnum{
		"INCREASE",
		"DECREASE",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeEvaluation EventType = "aws.evidently.evaluation"
	EventTypeCustom     EventType = "aws.evidently.custom"
)

// Values returns all known values for EventType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"aws.evidently.evaluation",
		"aws.evidently.custom",
	}
}

type ExperimentBaseStat string

// Enum values for ExperimentBaseStat
const (
	ExperimentBaseStatMean ExperimentBaseStat = "Mean"
)

// Values returns all known values for ExperimentBaseStat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentBaseStat) Values() []ExperimentBaseStat {
	return []ExperimentBaseStat{
		"Mean",
	}
}

type ExperimentReportName string

// Enum values for ExperimentReportName
const (
	ExperimentReportNameBayesianInference ExperimentReportName = "BayesianInference"
)

// Values returns all known values for ExperimentReportName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentReportName) Values() []ExperimentReportName {
	return []ExperimentReportName{
		"BayesianInference",
	}
}

type ExperimentResultRequestType string

// Enum values for ExperimentResultRequestType
const (
	ExperimentResultRequestTypeBaseStat           ExperimentResultRequestType = "BaseStat"
	ExperimentResultRequestTypeTreatmentEffect    ExperimentResultRequestType = "TreatmentEffect"
	ExperimentResultRequestTypeConfidenceInterval ExperimentResultRequestType = "ConfidenceInterval"
	ExperimentResultRequestTypePValue             ExperimentResultRequestType = "PValue"
)

// Values returns all known values for ExperimentResultRequestType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentResultRequestType) Values() []ExperimentResultRequestType {
	return []ExperimentResultRequestType{
		"BaseStat",
		"TreatmentEffect",
		"ConfidenceInterval",
		"PValue",
	}
}

type ExperimentResultResponseType string

// Enum values for ExperimentResultResponseType
const (
	ExperimentResultResponseTypeMean                         ExperimentResultResponseType = "Mean"
	ExperimentResultResponseTypeTreatmentEffect              ExperimentResultResponseType = "TreatmentEffect"
	ExperimentResultResponseTypeConfidenceIntervalUpperbound ExperimentResultResponseType = "ConfidenceIntervalUpperBound"
	ExperimentResultResponseTypeConfidenceIntervalLowerbound ExperimentResultResponseType = "ConfidenceIntervalLowerBound"
	ExperimentResultResponseTypePValue                       ExperimentResultResponseType = "PValue"
)

// Values returns all known values for ExperimentResultResponseType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentResultResponseType) Values() []ExperimentResultResponseType {
	return []ExperimentResultResponseType{
		"Mean",
		"TreatmentEffect",
		"ConfidenceIntervalUpperBound",
		"ConfidenceIntervalLowerBound",
		"PValue",
	}
}

type ExperimentStatus string

// Enum values for ExperimentStatus
const (
	ExperimentStatusCreated   ExperimentStatus = "CREATED"
	ExperimentStatusUpdating  ExperimentStatus = "UPDATING"
	ExperimentStatusRunning   ExperimentStatus = "RUNNING"
	ExperimentStatusCompleted ExperimentStatus = "COMPLETED"
	ExperimentStatusCancelled ExperimentStatus = "CANCELLED"
)

// Values returns all known values for ExperimentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentStatus) Values() []ExperimentStatus {
	return []ExperimentStatus{
		"CREATED",
		"UPDATING",
		"RUNNING",
		"COMPLETED",
		"CANCELLED",
	}
}

type ExperimentStopDesiredState string

// Enum values for ExperimentStopDesiredState
const (
	ExperimentStopDesiredStateCompleted ExperimentStopDesiredState = "COMPLETED"
	ExperimentStopDesiredStateCancelled ExperimentStopDesiredState = "CANCELLED"
)

// Values returns all known values for ExperimentStopDesiredState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentStopDesiredState) Values() []ExperimentStopDesiredState {
	return []ExperimentStopDesiredState{
		"COMPLETED",
		"CANCELLED",
	}
}

type ExperimentType string

// Enum values for ExperimentType
const (
	ExperimentTypeOnlineAbExperiment ExperimentType = "aws.evidently.onlineab"
)

// Values returns all known values for ExperimentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentType) Values() []ExperimentType {
	return []ExperimentType{
		"aws.evidently.onlineab",
	}
}

type FeatureEvaluationStrategy string

// Enum values for FeatureEvaluationStrategy
const (
	FeatureEvaluationStrategyAllRules         FeatureEvaluationStrategy = "ALL_RULES"
	FeatureEvaluationStrategyDefaultVariation FeatureEvaluationStrategy = "DEFAULT_VARIATION"
)

// Values returns all known values for FeatureEvaluationStrategy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FeatureEvaluationStrategy) Values() []FeatureEvaluationStrategy {
	return []FeatureEvaluationStrategy{
		"ALL_RULES",
		"DEFAULT_VARIATION",
	}
}

type FeatureStatus string

// Enum values for FeatureStatus
const (
	FeatureStatusAvailable FeatureStatus = "AVAILABLE"
	FeatureStatusUpdating  FeatureStatus = "UPDATING"
)

// Values returns all known values for FeatureStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FeatureStatus) Values() []FeatureStatus {
	return []FeatureStatus{
		"AVAILABLE",
		"UPDATING",
	}
}

type LaunchStatus string

// Enum values for LaunchStatus
const (
	LaunchStatusCreated   LaunchStatus = "CREATED"
	LaunchStatusUpdating  LaunchStatus = "UPDATING"
	LaunchStatusRunning   LaunchStatus = "RUNNING"
	LaunchStatusCompleted LaunchStatus = "COMPLETED"
	LaunchStatusCancelled LaunchStatus = "CANCELLED"
)

// Values returns all known values for LaunchStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LaunchStatus) Values() []LaunchStatus {
	return []LaunchStatus{
		"CREATED",
		"UPDATING",
		"RUNNING",
		"COMPLETED",
		"CANCELLED",
	}
}

type LaunchStopDesiredState string

// Enum values for LaunchStopDesiredState
const (
	LaunchStopDesiredStateCompleted LaunchStopDesiredState = "COMPLETED"
	LaunchStopDesiredStateCancelled LaunchStopDesiredState = "CANCELLED"
)

// Values returns all known values for LaunchStopDesiredState. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LaunchStopDesiredState) Values() []LaunchStopDesiredState {
	return []LaunchStopDesiredState{
		"COMPLETED",
		"CANCELLED",
	}
}

type LaunchType string

// Enum values for LaunchType
const (
	LaunchTypeScheduledSplitsLaunch LaunchType = "aws.evidently.splits"
)

// Values returns all known values for LaunchType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LaunchType) Values() []LaunchType {
	return []LaunchType{
		"aws.evidently.splits",
	}
}

type ProjectStatus string

// Enum values for ProjectStatus
const (
	ProjectStatusAvailable ProjectStatus = "AVAILABLE"
	ProjectStatusUpdating  ProjectStatus = "UPDATING"
)

// Values returns all known values for ProjectStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProjectStatus) Values() []ProjectStatus {
	return []ProjectStatus{
		"AVAILABLE",
		"UPDATING",
	}
}

type SegmentReferenceResourceType string

// Enum values for SegmentReferenceResourceType
const (
	SegmentReferenceResourceTypeExperiment SegmentReferenceResourceType = "EXPERIMENT"
	SegmentReferenceResourceTypeLaunch     SegmentReferenceResourceType = "LAUNCH"
)

// Values returns all known values for SegmentReferenceResourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SegmentReferenceResourceType) Values() []SegmentReferenceResourceType {
	return []SegmentReferenceResourceType{
		"EXPERIMENT",
		"LAUNCH",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}

type VariationValueType string

// Enum values for VariationValueType
const (
	VariationValueTypeString  VariationValueType = "STRING"
	VariationValueTypeLong    VariationValueType = "LONG"
	VariationValueTypeDouble  VariationValueType = "DOUBLE"
	VariationValueTypeBoolean VariationValueType = "BOOLEAN"
)

// Values returns all known values for VariationValueType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VariationValueType) Values() []VariationValueType {
	return []VariationValueType{
		"STRING",
		"LONG",
		"DOUBLE",
		"BOOLEAN",
	}
}
