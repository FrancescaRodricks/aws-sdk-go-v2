// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action sends a request to sync metadata across the specified virtual
// machines.
func (c *Client) StartVirtualMachinesMetadataSync(ctx context.Context, params *StartVirtualMachinesMetadataSyncInput, optFns ...func(*Options)) (*StartVirtualMachinesMetadataSyncOutput, error) {
	if params == nil {
		params = &StartVirtualMachinesMetadataSyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartVirtualMachinesMetadataSync", params, optFns, c.addOperationStartVirtualMachinesMetadataSyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartVirtualMachinesMetadataSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartVirtualMachinesMetadataSyncInput struct {

	// The Amazon Resource Name (ARN) of the hypervisor.
	//
	// This member is required.
	HypervisorArn *string

	noSmithyDocumentSerde
}

type StartVirtualMachinesMetadataSyncOutput struct {

	// The Amazon Resource Name (ARN) of the hypervisor.
	HypervisorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartVirtualMachinesMetadataSyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartVirtualMachinesMetadataSync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartVirtualMachinesMetadataSync{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartVirtualMachinesMetadataSync"); err != nil {
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
	if err = addOpStartVirtualMachinesMetadataSyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartVirtualMachinesMetadataSync(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartVirtualMachinesMetadataSync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartVirtualMachinesMetadataSync",
	}
}
