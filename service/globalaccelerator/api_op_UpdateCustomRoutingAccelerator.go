// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a custom routing accelerator.
func (c *Client) UpdateCustomRoutingAccelerator(ctx context.Context, params *UpdateCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*UpdateCustomRoutingAcceleratorOutput, error) {
	if params == nil {
		params = &UpdateCustomRoutingAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCustomRoutingAccelerator", params, optFns, c.addOperationUpdateCustomRoutingAcceleratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCustomRoutingAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCustomRoutingAcceleratorInput struct {

	// The Amazon Resource Name (ARN) of the accelerator to update.
	//
	// This member is required.
	AcceleratorArn *string

	// Indicates whether an accelerator is enabled. The value is true or false. The
	// default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to
	// false, the accelerator can be deleted.
	Enabled *bool

	// The IP address type that an accelerator supports. For a custom routing
	// accelerator, the value must be IPV4.
	IpAddressType types.IpAddressType

	// The name of the accelerator. The name can have a maximum of 64 characters, must
	// contain only alphanumeric characters, periods (.), or hyphens (-), and must not
	// begin or end with a hyphen or period.
	Name *string

	noSmithyDocumentSerde
}

type UpdateCustomRoutingAcceleratorOutput struct {

	// Information about the updated custom routing accelerator.
	Accelerator *types.CustomRoutingAccelerator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCustomRoutingAcceleratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCustomRoutingAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCustomRoutingAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCustomRoutingAccelerator"); err != nil {
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
	if err = addOpUpdateCustomRoutingAcceleratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCustomRoutingAccelerator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCustomRoutingAccelerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCustomRoutingAccelerator",
	}
}
