// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An object representing a device for a placement template (see PlacementTemplate).
type DeviceTemplate struct {

	// An optional Lambda function to invoke instead of the default Lambda function
	// provided by the placement template.
	CallbackOverrides map[string]string

	// The device type, which currently must be "button" .
	DeviceType *string

	noSmithyDocumentSerde
}

// An object describing a project's placement.
type PlacementDescription struct {

	// The user-defined attributes associated with the placement.
	//
	// This member is required.
	Attributes map[string]string

	// The date when the placement was initially created, in UNIX epoch time format.
	//
	// This member is required.
	CreatedDate *time.Time

	// The name of the placement.
	//
	// This member is required.
	PlacementName *string

	// The name of the project containing the placement.
	//
	// This member is required.
	ProjectName *string

	// The date when the placement was last updated, in UNIX epoch time format. If the
	// placement was not updated, then createdDate and updatedDate are the same.
	//
	// This member is required.
	UpdatedDate *time.Time

	noSmithyDocumentSerde
}

// An object providing summary information for a particular placement.
type PlacementSummary struct {

	// The date when the placement was originally created, in UNIX epoch time format.
	//
	// This member is required.
	CreatedDate *time.Time

	// The name of the placement being summarized.
	//
	// This member is required.
	PlacementName *string

	// The name of the project containing the placement.
	//
	// This member is required.
	ProjectName *string

	// The date when the placement was last updated, in UNIX epoch time format. If the
	// placement was not updated, then createdDate and updatedDate are the same.
	//
	// This member is required.
	UpdatedDate *time.Time

	noSmithyDocumentSerde
}

// An object defining the template for a placement.
type PlacementTemplate struct {

	// The default attributes (key/value pairs) to be applied to all placements using
	// this template.
	DefaultAttributes map[string]string

	// An object specifying the DeviceTemplate for all placements using this (PlacementTemplate ) template.
	DeviceTemplates map[string]DeviceTemplate

	noSmithyDocumentSerde
}

// An object providing detailed information for a particular project associated
// with an AWS account and region.
type ProjectDescription struct {

	// The date when the project was originally created, in UNIX epoch time format.
	//
	// This member is required.
	CreatedDate *time.Time

	// The name of the project for which to obtain information from.
	//
	// This member is required.
	ProjectName *string

	// The date when the project was last updated, in UNIX epoch time format. If the
	// project was not updated, then createdDate and updatedDate are the same.
	//
	// This member is required.
	UpdatedDate *time.Time

	// The ARN of the project.
	Arn *string

	// The description of the project.
	Description *string

	// An object describing the project's placement specifications.
	PlacementTemplate *PlacementTemplate

	// The tags (metadata key/value pairs) associated with the project.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object providing summary information for a particular project for an
// associated AWS account and region.
type ProjectSummary struct {

	// The date when the project was originally created, in UNIX epoch time format.
	//
	// This member is required.
	CreatedDate *time.Time

	// The name of the project being summarized.
	//
	// This member is required.
	ProjectName *string

	// The date when the project was last updated, in UNIX epoch time format. If the
	// project was not updated, then createdDate and updatedDate are the same.
	//
	// This member is required.
	UpdatedDate *time.Time

	// The ARN of the project.
	Arn *string

	// The tags (metadata key/value pairs) associated with the project.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
