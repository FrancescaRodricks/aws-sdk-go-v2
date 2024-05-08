// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the space.
func (c *Client) DescribeSpace(ctx context.Context, params *DescribeSpaceInput, optFns ...func(*Options)) (*DescribeSpaceOutput, error) {
	if params == nil {
		params = &DescribeSpaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpace", params, optFns, c.addOperationDescribeSpaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSpaceInput struct {

	// The ID of the associated domain.
	//
	// This member is required.
	DomainId *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	noSmithyDocumentSerde
}

type DescribeSpaceOutput struct {

	// The creation time.
	CreationTime *time.Time

	// The ID of the associated domain.
	DomainId *string

	// The failure reason.
	FailureReason *string

	// The ID of the space's profile in the Amazon EFS volume.
	HomeEfsFileSystemUid *string

	// The last modified time.
	LastModifiedTime *time.Time

	// The collection of ownership settings for a space.
	OwnershipSettings *types.OwnershipSettings

	// The space's Amazon Resource Name (ARN).
	SpaceArn *string

	// The name of the space that appears in the Amazon SageMaker Studio UI.
	SpaceDisplayName *string

	// The name of the space.
	SpaceName *string

	// A collection of space settings.
	SpaceSettings *types.SpaceSettings

	// The collection of space sharing settings for a space.
	SpaceSharingSettings *types.SpaceSharingSettings

	// The status.
	Status types.SpaceStatus

	// Returns the URL of the space. If the space is created with Amazon Web Services
	// IAM Identity Center (Successor to Amazon Web Services Single Sign-On)
	// authentication, users can navigate to the URL after appending the respective
	// redirect parameter for the application type to be federated through Amazon Web
	// Services IAM Identity Center.
	//
	// The following application types are supported:
	//
	//   - Studio Classic: &redirect=JupyterServer
	//
	//   - JupyterLab: &redirect=JupyterLab
	//
	//   - Code Editor, based on Code-OSS, Visual Studio Code - Open Source:
	//   &redirect=CodeEditor
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSpaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSpace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSpace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSpace"); err != nil {
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
	if err = addOpDescribeSpaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSpace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSpace",
	}
}
