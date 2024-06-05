// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deactivates the specified device identifier.
func (c *Client) DeactivateDeviceIdentifier(ctx context.Context, params *DeactivateDeviceIdentifierInput, optFns ...func(*Options)) (*DeactivateDeviceIdentifierOutput, error) {
	if params == nil {
		params = &DeactivateDeviceIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeactivateDeviceIdentifier", params, optFns, c.addOperationDeactivateDeviceIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeactivateDeviceIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeactivateDeviceIdentifierInput struct {

	// The Amazon Resource Name (ARN) of the device identifier.
	//
	// This member is required.
	DeviceIdentifierArn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [How to ensure idempotency].
	//
	// [How to ensure idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html
	ClientToken *string

	noSmithyDocumentSerde
}

type DeactivateDeviceIdentifierOutput struct {

	// Information about the device identifier.
	//
	// This member is required.
	DeviceIdentifier *types.DeviceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeactivateDeviceIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeactivateDeviceIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeactivateDeviceIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeactivateDeviceIdentifier"); err != nil {
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
	if err = addOpDeactivateDeviceIdentifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivateDeviceIdentifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeactivateDeviceIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeactivateDeviceIdentifier",
	}
}
