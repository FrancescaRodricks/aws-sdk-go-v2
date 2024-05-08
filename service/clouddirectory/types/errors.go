// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Access denied or directory not found. Either you don't have permissions for
// this directory or the directory does not exist. Try calling ListDirectoriesand check your
// permissions.
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

// A BatchWrite exception has occurred.
type BatchWriteException struct {
	Message *string

	ErrorCodeOverride *string

	Index int32
	Type  BatchWriteExceptionType

	noSmithyDocumentSerde
}

func (e *BatchWriteException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchWriteException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchWriteException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BatchWriteException"
	}
	return *e.ErrorCodeOverride
}
func (e *BatchWriteException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Cannot list the parents of a Directory root.
type CannotListParentOfRootException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CannotListParentOfRootException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CannotListParentOfRootException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CannotListParentOfRootException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CannotListParentOfRootException"
	}
	return *e.ErrorCodeOverride
}
func (e *CannotListParentOfRootException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a Directory could not be created due to a naming conflict. Choose a
// different name and try again.
type DirectoryAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DirectoryAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DirectoryAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *DirectoryAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A directory that has been deleted and to which access has been attempted. Note:
// The requested resource will eventually cease to exist.
type DirectoryDeletedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DirectoryDeletedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryDeletedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryDeletedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DirectoryDeletedException"
	}
	return *e.ErrorCodeOverride
}
func (e *DirectoryDeletedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An operation can only operate on a disabled directory.
type DirectoryNotDisabledException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DirectoryNotDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryNotDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryNotDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DirectoryNotDisabledException"
	}
	return *e.ErrorCodeOverride
}
func (e *DirectoryNotDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Operations are only permitted on enabled directories.
type DirectoryNotEnabledException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DirectoryNotEnabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryNotEnabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryNotEnabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DirectoryNotEnabledException"
	}
	return *e.ErrorCodeOverride
}
func (e *DirectoryNotEnabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A facet with the same name already exists.
type FacetAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *FacetAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FacetAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *FacetAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Occurs when deleting a facet that contains an attribute that is a target to an
// attribute reference in a different facet.
type FacetInUseException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *FacetInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FacetInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *FacetInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified Facet could not be found.
type FacetNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *FacetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FacetNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *FacetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Facet that you provided was not well formed or could not be validated with the
// schema.
type FacetValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *FacetValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FacetValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *FacetValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates a failure occurred while performing a check for backward
// compatibility between the specified schema and the schema that is currently
// applied to the directory.
type IncompatibleSchemaException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *IncompatibleSchemaException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleSchemaException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleSchemaException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IncompatibleSchemaException"
	}
	return *e.ErrorCodeOverride
}
func (e *IncompatibleSchemaException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An object has been attempted to be attached to an object that does not have the
// appropriate attribute value.
type IndexedAttributeMissingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *IndexedAttributeMissingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IndexedAttributeMissingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IndexedAttributeMissingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IndexedAttributeMissingException"
	}
	return *e.ErrorCodeOverride
}
func (e *IndexedAttributeMissingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates a problem that must be resolved by Amazon Web Services. This might be
// a transient error in which case you can retry your request until it succeeds.
// Otherwise, go to the [AWS Service Health Dashboard]site to see if there are any operational issues with the
// service.
//
// [AWS Service Health Dashboard]: http://status.aws.amazon.com/
type InternalServiceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Indicates that the provided ARN value is not valid.
type InvalidArnException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArnException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidArnException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArnException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that an attempt to make an attachment was invalid. For example,
// attaching two nodes with a link type that is not applicable to the nodes or
// attempting to apply a schema to a directory a second time.
type InvalidAttachmentException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidAttachmentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAttachmentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAttachmentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidAttachmentException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidAttachmentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An attempt to modify a Facet resulted in an invalid schema exception.
type InvalidFacetUpdateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidFacetUpdateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFacetUpdateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFacetUpdateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidFacetUpdateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidFacetUpdateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the NextToken value is not valid.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Occurs when any of the rule parameter keys or values are invalid.
type InvalidRuleException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidRuleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRuleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRuleException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRuleException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRuleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided SchemaDoc value is not valid.
type InvalidSchemaDocException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSchemaDocException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSchemaDocException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSchemaDocException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSchemaDocException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSchemaDocException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Can occur for multiple reasons such as when you tag a resource that doesn’t
// exist or if you specify a higher number of tags for a resource than the allowed
// limit. Allowed limit is 50 tags per resource.
type InvalidTaggingRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidTaggingRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTaggingRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTaggingRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTaggingRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTaggingRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that limits are exceeded. See [Limits] for more information.
//
// [Limits]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a link could not be created due to a naming conflict. Choose a
// different name and then try again.
type LinkNameAlreadyInUseException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LinkNameAlreadyInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LinkNameAlreadyInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LinkNameAlreadyInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LinkNameAlreadyInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *LinkNameAlreadyInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the requested operation can only operate on index objects.
type NotIndexException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotIndexException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotIndexException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotIndexException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotIndexException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotIndexException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Occurs when any invalid operations are performed on an object that is not a
// node, such as calling ListObjectChildren for a leaf node object.
type NotNodeException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotNodeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotNodeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotNodeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotNodeException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotNodeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the requested operation can only operate on policy objects.
type NotPolicyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotPolicyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotPolicyException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the object is not attached to the index.
type ObjectAlreadyDetachedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ObjectAlreadyDetachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectAlreadyDetachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectAlreadyDetachedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ObjectAlreadyDetachedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ObjectAlreadyDetachedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the requested operation cannot be completed because the object
// has not been detached from the tree.
type ObjectNotDetachedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ObjectNotDetachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectNotDetachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectNotDetachedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ObjectNotDetachedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ObjectNotDetachedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource could not be found.
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

// Occurs when a conflict with a previous successful write is detected. For
// example, if a write operation occurs on an object and then an attempt is made to
// read the object using “SERIALIZABLE” consistency, this exception may result.
// This generally occurs when the previous write did not have time to propagate to
// the host serving the current request. A retry (with appropriate backoff logic)
// is the recommended response to this exception.
type RetryableConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RetryableConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RetryableConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RetryableConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RetryableConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *RetryableConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a schema could not be created due to a naming conflict. Please
// select a different name and then try again.
type SchemaAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SchemaAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchemaAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchemaAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SchemaAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *SchemaAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a schema is already published.
type SchemaAlreadyPublishedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SchemaAlreadyPublishedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchemaAlreadyPublishedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchemaAlreadyPublishedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SchemaAlreadyPublishedException"
	}
	return *e.ErrorCodeOverride
}
func (e *SchemaAlreadyPublishedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The object could not be deleted because links still exist. Remove the links and
// then try the operation again.
type StillContainsLinksException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *StillContainsLinksException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StillContainsLinksException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StillContainsLinksException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "StillContainsLinksException"
	}
	return *e.ErrorCodeOverride
}
func (e *StillContainsLinksException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the requested index type is not supported.
type UnsupportedIndexTypeException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedIndexTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedIndexTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedIndexTypeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedIndexTypeException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedIndexTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that your request is malformed in some manner. See the exception
// message.
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
