// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Launch Configuration Template.
func (c *Client) CreateLaunchConfigurationTemplate(ctx context.Context, params *CreateLaunchConfigurationTemplateInput, optFns ...func(*Options)) (*CreateLaunchConfigurationTemplateOutput, error) {
	if params == nil {
		params = &CreateLaunchConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchConfigurationTemplate", params, optFns, c.addOperationCreateLaunchConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationTemplateInput struct {

	// Associate public Ip address.
	AssociatePublicIpAddress *bool

	// Launch configuration template boot mode.
	BootMode types.BootMode

	// Copy private Ip.
	CopyPrivateIp *bool

	// Copy tags.
	CopyTags *bool

	// Enable map auto tagging.
	EnableMapAutoTagging *bool

	// Large volume config.
	LargeVolumeConf *types.LaunchTemplateDiskConf

	// Launch disposition.
	LaunchDisposition types.LaunchDisposition

	// Configure Licensing.
	Licensing *types.Licensing

	// Launch configuration template map auto tagging MPE ID.
	MapAutoTaggingMpeID *string

	// Launch configuration template post launch actions.
	PostLaunchActions *types.PostLaunchActions

	// Small volume config.
	SmallVolumeConf *types.LaunchTemplateDiskConf

	// Small volume maximum size.
	SmallVolumeMaxSize int64

	// Request to associate tags during creation of a Launch Configuration Template.
	Tags map[string]string

	// Target instance type right-sizing method.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	noSmithyDocumentSerde
}

type CreateLaunchConfigurationTemplateOutput struct {

	// ID of the Launch Configuration Template.
	//
	// This member is required.
	LaunchConfigurationTemplateID *string

	// ARN of the Launch Configuration Template.
	Arn *string

	// Associate public Ip address.
	AssociatePublicIpAddress *bool

	// Launch configuration template boot mode.
	BootMode types.BootMode

	// Copy private Ip.
	CopyPrivateIp *bool

	// Copy tags.
	CopyTags *bool

	// EC2 launch template ID.
	Ec2LaunchTemplateID *string

	// Enable map auto tagging.
	EnableMapAutoTagging *bool

	// Large volume config.
	LargeVolumeConf *types.LaunchTemplateDiskConf

	// Launch disposition.
	LaunchDisposition types.LaunchDisposition

	// Configure Licensing.
	Licensing *types.Licensing

	// Launch configuration template map auto tagging MPE ID.
	MapAutoTaggingMpeID *string

	// Post Launch Actions of the Launch Configuration Template.
	PostLaunchActions *types.PostLaunchActions

	// Small volume config.
	SmallVolumeConf *types.LaunchTemplateDiskConf

	// Small volume maximum size.
	SmallVolumeMaxSize int64

	// Tags of the Launch Configuration Template.
	Tags map[string]string

	// Target instance type right-sizing method.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLaunchConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLaunchConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateLaunchConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfigurationTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateLaunchConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "CreateLaunchConfigurationTemplate",
	}
}
