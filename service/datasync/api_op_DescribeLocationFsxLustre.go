// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about how an DataSync transfer location for an Amazon FSx for
// Lustre file system is configured.
func (c *Client) DescribeLocationFsxLustre(ctx context.Context, params *DescribeLocationFsxLustreInput, optFns ...func(*Options)) (*DescribeLocationFsxLustreOutput, error) {
	if params == nil {
		params = &DescribeLocationFsxLustreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationFsxLustre", params, optFns, c.addOperationDescribeLocationFsxLustreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationFsxLustreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocationFsxLustreInput struct {

	// The Amazon Resource Name (ARN) of the FSx for Lustre location to describe.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

type DescribeLocationFsxLustreOutput struct {

	// The time that the FSx for Lustre location was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the FSx for Lustre location that was
	// described.
	LocationArn *string

	// The URI of the FSx for Lustre location that was described.
	LocationUri *string

	// The Amazon Resource Names (ARNs) of the security groups that are configured for
	// the FSx for Lustre file system.
	SecurityGroupArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationFsxLustreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationFsxLustre{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationFsxLustre{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLocationFsxLustre"); err != nil {
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
	if err = addOpDescribeLocationFsxLustreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationFsxLustre(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationFsxLustre(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLocationFsxLustre",
	}
}
