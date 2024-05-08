// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents the content of a particular audit event.
type AuditEvent struct {

	// Unique identifier of a case audit history event.
	//
	// This member is required.
	EventId *string

	// A list of Case Audit History event fields.
	//
	// This member is required.
	Fields []*AuditEventField

	// Time at which an Audit History event took place.
	//
	// This member is required.
	PerformedTime *time.Time

	// The Type of an audit history event.
	//
	// This member is required.
	Type AuditEventType

	// Information of the user which performed the audit.
	PerformedBy *AuditEventPerformedBy

	// The Type of the related item.
	RelatedItemType RelatedItemType

	noSmithyDocumentSerde
}

// Fields for audit event.
type AuditEventField struct {

	// Unique identifier of field in an Audit History entry.
	//
	// This member is required.
	EventFieldId *string

	// Union of potential field value types.
	//
	// This member is required.
	NewValue AuditEventFieldValueUnion

	// Union of potential field value types.
	OldValue AuditEventFieldValueUnion

	noSmithyDocumentSerde
}

// Object to store union of Field values.
//
// The following types satisfy this interface:
//
//	AuditEventFieldValueUnionMemberBooleanValue
//	AuditEventFieldValueUnionMemberDoubleValue
//	AuditEventFieldValueUnionMemberEmptyValue
//	AuditEventFieldValueUnionMemberStringValue
//	AuditEventFieldValueUnionMemberUserArnValue
type AuditEventFieldValueUnion interface {
	isAuditEventFieldValueUnion()
}

// Can be either null, or have a Boolean value type. Only one value can be
// provided.
type AuditEventFieldValueUnionMemberBooleanValue struct {
	Value bool

	noSmithyDocumentSerde
}

func (*AuditEventFieldValueUnionMemberBooleanValue) isAuditEventFieldValueUnion() {}

// Can be either null, or have a Double value type. Only one value can be provided.
type AuditEventFieldValueUnionMemberDoubleValue struct {
	Value float64

	noSmithyDocumentSerde
}

func (*AuditEventFieldValueUnionMemberDoubleValue) isAuditEventFieldValueUnion() {}

// An empty value. You cannot set EmptyFieldValue on a field that is required on a
// case template.
//
// This structure will never have any data members. It signifies an empty value on
// a case field.
type AuditEventFieldValueUnionMemberEmptyValue struct {
	Value EmptyFieldValue

	noSmithyDocumentSerde
}

func (*AuditEventFieldValueUnionMemberEmptyValue) isAuditEventFieldValueUnion() {}

// Can be either null, or have a String value type. Only one value can be provided.
type AuditEventFieldValueUnionMemberStringValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*AuditEventFieldValueUnionMemberStringValue) isAuditEventFieldValueUnion() {}

// Can be either null, or have a String value type formatted as an ARN. Only one
// value can be provided.
type AuditEventFieldValueUnionMemberUserArnValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*AuditEventFieldValueUnionMemberUserArnValue) isAuditEventFieldValueUnion() {}

// Information of the user which performed the audit.
type AuditEventPerformedBy struct {

	// Unique identifier of an IAM role.
	//
	// This member is required.
	IamPrincipalArn *string

	// Represents the identity of the person who performed the action.
	User UserUnion

	noSmithyDocumentSerde
}

// Content specific to BasicLayout type. It configures fields in the top panel and
// More Info tab of agent application.
type BasicLayout struct {

	// This represents sections in a tab of the page layout.
	MoreInfo *LayoutSections

	// This represents sections in a panel of the page layout.
	TopPanel *LayoutSections

	noSmithyDocumentSerde
}

// Details of what case data is published through the case event stream.
type CaseEventIncludedData struct {

	// List of field identifiers.
	//
	// This member is required.
	Fields []FieldIdentifier

	noSmithyDocumentSerde
}

// A filter for cases. Only one value can be provided.
//
// The following types satisfy this interface:
//
//	CaseFilterMemberAndAll
//	CaseFilterMemberField
//	CaseFilterMemberNot
//	CaseFilterMemberOrAll
type CaseFilter interface {
	isCaseFilter()
}

// Provides "and all" filtering.
type CaseFilterMemberAndAll struct {
	Value []CaseFilter

	noSmithyDocumentSerde
}

func (*CaseFilterMemberAndAll) isCaseFilter() {}

// A list of fields to filter on.
type CaseFilterMemberField struct {
	Value FieldFilter

	noSmithyDocumentSerde
}

func (*CaseFilterMemberField) isCaseFilter() {}

// A filter for cases. Only one value can be provided.
type CaseFilterMemberNot struct {
	Value CaseFilter

	noSmithyDocumentSerde
}

func (*CaseFilterMemberNot) isCaseFilter() {}

// Provides "or all" filtering.
type CaseFilterMemberOrAll struct {
	Value []CaseFilter

	noSmithyDocumentSerde
}

func (*CaseFilterMemberOrAll) isCaseFilter() {}

// Case summary information.
type CaseSummary struct {

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	noSmithyDocumentSerde
}

// Represents the content of a Comment to be returned to agents.
type CommentContent struct {

	// Text in the body of a Comment on a case.
	//
	// This member is required.
	Body *string

	// Type of the text in the box of a Comment on a case.
	//
	// This member is required.
	ContentType CommentBodyTextType

	noSmithyDocumentSerde
}

// A filter for related items of type Comment .
type CommentFilter struct {
	noSmithyDocumentSerde
}

// An object that represents an Amazon Connect contact object.
type Contact struct {

	// A unique identifier of a contact in Amazon Connect.
	//
	// This member is required.
	ContactArn *string

	noSmithyDocumentSerde
}

// An object that represents a content of an Amazon Connect contact object.
type ContactContent struct {

	// A list of channels to filter on for related items of type Contact .
	//
	// This member is required.
	Channel *string

	// The difference between the InitiationTimestamp and the DisconnectTimestamp of
	// the contact.
	//
	// This member is required.
	ConnectedToSystemTime *time.Time

	// A unique identifier of a contact in Amazon Connect.
	//
	// This member is required.
	ContactArn *string

	noSmithyDocumentSerde
}

// A filter for related items of type Contact .
type ContactFilter struct {

	// A list of channels to filter on for related items of type Contact .
	Channel []string

	// A unique identifier of a contact in Amazon Connect.
	ContactArn *string

	noSmithyDocumentSerde
}

// Object for the summarized details of the domain.
type DomainSummary struct {

	// The Amazon Resource Name (ARN) of the domain.
	//
	// This member is required.
	DomainArn *string

	// The unique identifier of the domain.
	//
	// This member is required.
	DomainId *string

	// The name of the domain.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// An empty value. You cannot set EmptyFieldValue on a field that is required on a
// case template.
//
// This structure will never have any data members. It signifies an empty value on
// a case field.
type EmptyFieldValue struct {
	noSmithyDocumentSerde
}

// Configuration to enable EventBridge case event delivery and determine what data
// is delivered.
type EventBridgeConfiguration struct {

	// Indicates whether the to broadcast case event data to the customer.
	//
	// This member is required.
	Enabled *bool

	// Details of what case and related item data is published through the case event
	// stream.
	IncludedData *EventIncludedData

	noSmithyDocumentSerde
}

// Details of what case and related item data is published through the case event
// stream.
type EventIncludedData struct {

	// Details of what case data is published through the case event stream.
	CaseData *CaseEventIncludedData

	// Details of what related item data is published through the case event stream.
	RelatedItemData *RelatedItemEventIncludedData

	noSmithyDocumentSerde
}

// Object for errors on fields.
type FieldError struct {

	// The error code from getting a field.
	//
	// This member is required.
	ErrorCode *string

	// The field identifier that caused the error.
	//
	// This member is required.
	Id *string

	// The error message from getting a field.
	Message *string

	noSmithyDocumentSerde
}

// A filter for fields. Only one value can be provided.
//
// The following types satisfy this interface:
//
//	FieldFilterMemberContains
//	FieldFilterMemberEqualTo
//	FieldFilterMemberGreaterThan
//	FieldFilterMemberGreaterThanOrEqualTo
//	FieldFilterMemberLessThan
//	FieldFilterMemberLessThanOrEqualTo
type FieldFilter interface {
	isFieldFilter()
}

// Object containing field identifier and value information.
type FieldFilterMemberContains struct {
	Value FieldValue

	noSmithyDocumentSerde
}

func (*FieldFilterMemberContains) isFieldFilter() {}

// Object containing field identifier and value information.
type FieldFilterMemberEqualTo struct {
	Value FieldValue

	noSmithyDocumentSerde
}

func (*FieldFilterMemberEqualTo) isFieldFilter() {}

// Object containing field identifier and value information.
type FieldFilterMemberGreaterThan struct {
	Value FieldValue

	noSmithyDocumentSerde
}

func (*FieldFilterMemberGreaterThan) isFieldFilter() {}

// Object containing field identifier and value information.
type FieldFilterMemberGreaterThanOrEqualTo struct {
	Value FieldValue

	noSmithyDocumentSerde
}

func (*FieldFilterMemberGreaterThanOrEqualTo) isFieldFilter() {}

// Object containing field identifier and value information.
type FieldFilterMemberLessThan struct {
	Value FieldValue

	noSmithyDocumentSerde
}

func (*FieldFilterMemberLessThan) isFieldFilter() {}

// Object containing field identifier and value information.
type FieldFilterMemberLessThanOrEqualTo struct {
	Value FieldValue

	noSmithyDocumentSerde
}

func (*FieldFilterMemberLessThanOrEqualTo) isFieldFilter() {}

// Object for a group of fields and associated properties.
type FieldGroup struct {

	// Represents an ordered list containing field related information.
	//
	// This member is required.
	Fields []FieldItem

	// Name of the field group.
	Name *string

	noSmithyDocumentSerde
}

// Object for unique identifier of a field.
type FieldIdentifier struct {

	// Unique identifier of a field.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// Object for field related information.
type FieldItem struct {

	// Unique identifier of a field.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// Object for field Options information.
type FieldOption struct {

	// Describes whether the FieldOption is active (displayed) or inactive.
	//
	// This member is required.
	Active *bool

	// FieldOptionName has max length 100 and disallows trailing spaces.
	//
	// This member is required.
	Name *string

	// FieldOptionValue has max length 100 and must be alphanumeric with hyphens and
	// underscores.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Object for field Options errors.
type FieldOptionError struct {

	// Error code from creating or updating field option.
	//
	// This member is required.
	ErrorCode *string

	// Error message from creating or updating field option.
	//
	// This member is required.
	Message *string

	// The field option value that caused the error.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Object for the summarized details of the field.
type FieldSummary struct {

	// The Amazon Resource Name (ARN) of the field.
	//
	// This member is required.
	FieldArn *string

	// The unique identifier of a field.
	//
	// This member is required.
	FieldId *string

	// Name of the field.
	//
	// This member is required.
	Name *string

	// The namespace of a field.
	//
	// This member is required.
	Namespace FieldNamespace

	// The type of a field.
	//
	// This member is required.
	Type FieldType

	noSmithyDocumentSerde
}

// Object for case field values.
type FieldValue struct {

	// Unique identifier of a field.
	//
	// This member is required.
	Id *string

	// Union of potential field value types.
	//
	// This member is required.
	Value FieldValueUnion

	noSmithyDocumentSerde
}

// Object to store union of Field values.
//
// The Summary system field accepts 1500 characters while all other fields accept
// 500 characters.
//
// The following types satisfy this interface:
//
//	FieldValueUnionMemberBooleanValue
//	FieldValueUnionMemberDoubleValue
//	FieldValueUnionMemberEmptyValue
//	FieldValueUnionMemberStringValue
//	FieldValueUnionMemberUserArnValue
type FieldValueUnion interface {
	isFieldValueUnion()
}

// Can be either null, or have a Boolean value type. Only one value can be
// provided.
type FieldValueUnionMemberBooleanValue struct {
	Value bool

	noSmithyDocumentSerde
}

func (*FieldValueUnionMemberBooleanValue) isFieldValueUnion() {}

// Can be either null, or have a Double number value type. Only one value can be
// provided.
type FieldValueUnionMemberDoubleValue struct {
	Value float64

	noSmithyDocumentSerde
}

func (*FieldValueUnionMemberDoubleValue) isFieldValueUnion() {}

// An empty value.
type FieldValueUnionMemberEmptyValue struct {
	Value EmptyFieldValue

	noSmithyDocumentSerde
}

func (*FieldValueUnionMemberEmptyValue) isFieldValueUnion() {}

// String value type.
type FieldValueUnionMemberStringValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*FieldValueUnionMemberStringValue) isFieldValueUnion() {}

// Represents the user that performed the audit.
type FieldValueUnionMemberUserArnValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*FieldValueUnionMemberUserArnValue) isFieldValueUnion() {}

// An object that represents a content of an Amazon Connect file object.
type FileContent struct {

	// The Amazon Resource Name (ARN) of a File in Amazon Connect.
	//
	// This member is required.
	FileArn *string

	noSmithyDocumentSerde
}

// A filter for related items of type File .
type FileFilter struct {

	// The Amazon Resource Name (ARN) of the file.
	FileArn *string

	noSmithyDocumentSerde
}

// Object to store detailed field information.
type GetFieldResponse struct {

	// The Amazon Resource Name (ARN) of the field.
	//
	// This member is required.
	FieldArn *string

	// Unique identifier of the field.
	//
	// This member is required.
	FieldId *string

	// Name of the field.
	//
	// This member is required.
	Name *string

	// Namespace of the field.
	//
	// This member is required.
	Namespace FieldNamespace

	// Type of the field.
	//
	// This member is required.
	Type FieldType

	// Timestamp at which the resource was created.
	CreatedTime *time.Time

	// Denotes whether or not the resource has been deleted.
	Deleted bool

	// Description of the field.
	Description *string

	// Timestamp at which the resource was created or last modified.
	LastModifiedTime *time.Time

	// A map of of key-value pairs that represent tags on a resource. Tags are used to
	// organize, track, or control access for this resource.
	Tags map[string]*string

	noSmithyDocumentSerde
}

// Object to store configuration of layouts associated to the template.
type LayoutConfiguration struct {

	//  Unique identifier of a layout.
	DefaultLayout *string

	noSmithyDocumentSerde
}

// Object to store union of different versions of layout content.
//
// The following types satisfy this interface:
//
//	LayoutContentMemberBasic
type LayoutContent interface {
	isLayoutContent()
}

// Content specific to BasicLayout type. It configures fields in the top panel and
// More Info tab of Cases user interface.
type LayoutContentMemberBasic struct {
	Value BasicLayout

	noSmithyDocumentSerde
}

func (*LayoutContentMemberBasic) isLayoutContent() {}

// Ordered list containing different kinds of sections that can be added. A
// LayoutSections object can only contain one section.
type LayoutSections struct {

	// Ordered list containing different kinds of sections that can be added.
	Sections []Section

	noSmithyDocumentSerde
}

// Object for the summarized details of the layout.
type LayoutSummary struct {

	// The Amazon Resource Name (ARN) of the layout.
	//
	// This member is required.
	LayoutArn *string

	// The unique identifier for of the layout.
	//
	// This member is required.
	LayoutId *string

	// The name of the layout.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Represents the content of a particular type of related item.
//
// The following types satisfy this interface:
//
//	RelatedItemContentMemberComment
//	RelatedItemContentMemberContact
//	RelatedItemContentMemberFile
type RelatedItemContent interface {
	isRelatedItemContent()
}

// Represents the content of a comment to be returned to agents.
type RelatedItemContentMemberComment struct {
	Value CommentContent

	noSmithyDocumentSerde
}

func (*RelatedItemContentMemberComment) isRelatedItemContent() {}

// Represents the content of a contact to be returned to agents.
type RelatedItemContentMemberContact struct {
	Value ContactContent

	noSmithyDocumentSerde
}

func (*RelatedItemContentMemberContact) isRelatedItemContent() {}

// Represents the content of a File to be returned to agents.
type RelatedItemContentMemberFile struct {
	Value FileContent

	noSmithyDocumentSerde
}

func (*RelatedItemContentMemberFile) isRelatedItemContent() {}

// Details of what related item data is published through the case event stream.
type RelatedItemEventIncludedData struct {

	// Details of what related item data is published through the case event stream.
	//
	// This member is required.
	IncludeContent *bool

	noSmithyDocumentSerde
}

// Represents the content of a related item to be created.
//
// The following types satisfy this interface:
//
//	RelatedItemInputContentMemberComment
//	RelatedItemInputContentMemberContact
//	RelatedItemInputContentMemberFile
type RelatedItemInputContent interface {
	isRelatedItemInputContent()
}

// Represents the content of a comment to be returned to agents.
type RelatedItemInputContentMemberComment struct {
	Value CommentContent

	noSmithyDocumentSerde
}

func (*RelatedItemInputContentMemberComment) isRelatedItemInputContent() {}

// Object representing a contact in Amazon Connect as an API request field.
type RelatedItemInputContentMemberContact struct {
	Value Contact

	noSmithyDocumentSerde
}

func (*RelatedItemInputContentMemberContact) isRelatedItemInputContent() {}

// A file of related items.
type RelatedItemInputContentMemberFile struct {
	Value FileContent

	noSmithyDocumentSerde
}

func (*RelatedItemInputContentMemberFile) isRelatedItemInputContent() {}

// The list of types of related items and their parameters to use for filtering.
//
// The following types satisfy this interface:
//
//	RelatedItemTypeFilterMemberComment
//	RelatedItemTypeFilterMemberContact
//	RelatedItemTypeFilterMemberFile
type RelatedItemTypeFilter interface {
	isRelatedItemTypeFilter()
}

// A filter for related items of type Comment .
type RelatedItemTypeFilterMemberComment struct {
	Value CommentFilter

	noSmithyDocumentSerde
}

func (*RelatedItemTypeFilterMemberComment) isRelatedItemTypeFilter() {}

// A filter for related items of type Contact .
type RelatedItemTypeFilterMemberContact struct {
	Value ContactFilter

	noSmithyDocumentSerde
}

func (*RelatedItemTypeFilterMemberContact) isRelatedItemTypeFilter() {}

// A filter for related items of this type of File .
type RelatedItemTypeFilterMemberFile struct {
	Value FileFilter

	noSmithyDocumentSerde
}

func (*RelatedItemTypeFilterMemberFile) isRelatedItemTypeFilter() {}

// List of fields that must have a value provided to create a case.
type RequiredField struct {

	// Unique identifier of a field.
	//
	// This member is required.
	FieldId *string

	noSmithyDocumentSerde
}

// A list of items that represent cases.
type SearchCasesResponseItem struct {

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// List of case field values.
	//
	// This member is required.
	Fields []FieldValue

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	// A map of of key-value pairs that represent tags on a resource. Tags are used to
	// organize, track, or control access for this resource.
	Tags map[string]*string

	noSmithyDocumentSerde
}

// A list of items that represent RelatedItems.
type SearchRelatedItemsResponseItem struct {

	// Time at which a related item was associated with a case.
	//
	// This member is required.
	AssociationTime *time.Time

	// Represents the content of a particular type of related item.
	//
	// This member is required.
	Content RelatedItemContent

	// Unique identifier of a related item.
	//
	// This member is required.
	RelatedItemId *string

	// Type of a related item.
	//
	// This member is required.
	Type RelatedItemType

	// Represents the creator of the related item.
	PerformedBy UserUnion

	// A map of of key-value pairs that represent tags on a resource. Tags are used to
	// organize, track, or control access for this resource.
	Tags map[string]*string

	noSmithyDocumentSerde
}

// This represents a sections within a panel or tab of the page layout.
//
// The following types satisfy this interface:
//
//	SectionMemberFieldGroup
type Section interface {
	isSection()
}

// Consists of a group of fields and associated properties.
type SectionMemberFieldGroup struct {
	Value FieldGroup

	noSmithyDocumentSerde
}

func (*SectionMemberFieldGroup) isSection() {}

// A structured set of sort terms.
type Sort struct {

	// Unique identifier of a field.
	//
	// This member is required.
	FieldId *string

	// A structured set of sort terms
	//
	// This member is required.
	SortOrder Order

	noSmithyDocumentSerde
}

// Template summary information.
type TemplateSummary struct {

	// The template name.
	//
	// This member is required.
	Name *string

	// The status of the template.
	//
	// This member is required.
	Status TemplateStatus

	// The Amazon Resource Name (ARN) of the template.
	//
	// This member is required.
	TemplateArn *string

	// The unique identifier for the template.
	//
	// This member is required.
	TemplateId *string

	noSmithyDocumentSerde
}

// Represents the identity of the person who performed the action.
//
// The following types satisfy this interface:
//
//	UserUnionMemberUserArn
type UserUnion interface {
	isUserUnion()
}

// Represents the Amazon Connect ARN of the user.
type UserUnionMemberUserArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*UserUnionMemberUserArn) isUserUnion() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAuditEventFieldValueUnion() {}
func (*UnknownUnionMember) isCaseFilter()                {}
func (*UnknownUnionMember) isFieldFilter()               {}
func (*UnknownUnionMember) isFieldValueUnion()           {}
func (*UnknownUnionMember) isLayoutContent()             {}
func (*UnknownUnionMember) isRelatedItemContent()        {}
func (*UnknownUnionMember) isRelatedItemInputContent()   {}
func (*UnknownUnionMember) isRelatedItemTypeFilter()     {}
func (*UnknownUnionMember) isSection()                   {}
func (*UnknownUnionMember) isUserUnion()                 {}
