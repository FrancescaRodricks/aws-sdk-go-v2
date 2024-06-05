// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you export your virtual machine (VM) from its virtualization environment,
// that process creates a set of one or more disk container files that act as
// snapshots of your VM’s environment, settings, and data. The Amazon EC2 API [ImportImage]
// action uses those files to import your VM and create an AMI. To import using the
// CLI command, see [import-image]
//
// You can reference the task ID from the VM import to pull in the AMI that the
// import created as the base image for your Image Builder recipe.
//
// [ImportImage]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportImage.html
// [import-image]: https://docs.aws.amazon.com/cli/latest/reference/ec2/import-image.html
func (c *Client) ImportVmImage(ctx context.Context, params *ImportVmImageInput, optFns ...func(*Options)) (*ImportVmImageOutput, error) {
	if params == nil {
		params = &ImportVmImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportVmImage", params, optFns, c.addOperationImportVmImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportVmImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportVmImageInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// This member is required.
	ClientToken *string

	// The name of the base image that is created by the import process.
	//
	// This member is required.
	Name *string

	// The operating system platform for the imported VM.
	//
	// This member is required.
	Platform types.Platform

	// The semantic version to attach to the base image that was created during the
	// import process. This version follows the semantic version syntax.
	//
	// The semantic version has four nodes: ../. You can assign values for the first
	// three, and can filter on all of them.
	//
	// Assignment: For the first three nodes you can assign any positive integer
	// value, including zero, with an upper limit of 2^30-1, or 1073741823 for each
	// node. Image Builder automatically assigns the build number to the fourth node.
	//
	// Patterns: You can use any numeric pattern that adheres to the assignment
	// requirements for the nodes that you can assign. For example, you might choose a
	// software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	//
	// This member is required.
	SemanticVersion *string

	// The importTaskId (API) or ImportTaskId (CLI) from the Amazon EC2 VM import
	// process. Image Builder retrieves information from the import process to pull in
	// the AMI that is created from the VM source as the base image for your recipe.
	//
	// This member is required.
	VmImportTaskId *string

	// The description for the base image that is created by the import process.
	Description *string

	// The operating system version for the imported VM.
	OsVersion *string

	// Tags that are attached to the import resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ImportVmImageOutput struct {

	// The client token that uniquely identifies the request.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the AMI that was created during the VM import
	// process. This AMI is used as the base image for the recipe that imported the VM.
	ImageArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportVmImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportVmImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportVmImage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportVmImage"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addIdempotencyToken_opImportVmImageMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportVmImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportVmImage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpImportVmImage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpImportVmImage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpImportVmImage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ImportVmImageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ImportVmImageInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opImportVmImageMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpImportVmImage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opImportVmImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportVmImage",
	}
}
