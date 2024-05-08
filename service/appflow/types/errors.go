// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// AppFlow/Requester has invalid or missing permissions.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//	There was a conflict when processing the request (for example, a flow with the
//
// given name already exists within the account. Check for conflicting resource
// names and try again.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred when authenticating with the connector endpoint.
type ConnectorAuthenticationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConnectorAuthenticationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConnectorAuthenticationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConnectorAuthenticationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConnectorAuthenticationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConnectorAuthenticationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred when retrieving data from the connector endpoint.
type ConnectorServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConnectorServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConnectorServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConnectorServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConnectorServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConnectorServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//	An internal service error occurred during the processing of your request. Try
//
// again later.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

//	The resource specified in the request (such as the source or destination
//
// connector profile) is not found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//	The request would cause a service quota (such as the number of flows) to be
//
// exceeded.
type ServiceQuotaExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceQuotaExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// API calls have exceeded the maximum allowed API request rate per account and
// per Region.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation is not supported for the current flow.
type UnsupportedOperationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedOperationException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request has invalid or missing parameters.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
