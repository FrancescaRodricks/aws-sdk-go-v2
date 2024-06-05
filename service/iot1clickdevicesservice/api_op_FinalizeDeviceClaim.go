// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Given a device ID, finalizes the claim request for the associated device.
//
// Claiming a device consists of initiating a claim, then publishing a device
// event, and finalizing the claim. For a device of type button, a device event can
// be published by simply clicking the device.
func (c *Client) FinalizeDeviceClaim(ctx context.Context, params *FinalizeDeviceClaimInput, optFns ...func(*Options)) (*FinalizeDeviceClaimOutput, error) {
	if params == nil {
		params = &FinalizeDeviceClaimInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FinalizeDeviceClaim", params, optFns, c.addOperationFinalizeDeviceClaimMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FinalizeDeviceClaimOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FinalizeDeviceClaimInput struct {

	// The unique identifier of the device.
	//
	// This member is required.
	DeviceId *string

	// A collection of key/value pairs defining the resource tags. For example, {
	// "tags": {"key1": "value1", "key2": "value2"} }. For more information, see [AWS Tagging Strategies].
	//
	// [AWS Tagging Strategies]: https://aws.amazon.com/answers/account-management/aws-tagging-strategies/
	Tags map[string]string

	noSmithyDocumentSerde
}

type FinalizeDeviceClaimOutput struct {

	// The device's final claim state.
	State *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFinalizeDeviceClaimMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpFinalizeDeviceClaim{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpFinalizeDeviceClaim{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "FinalizeDeviceClaim"); err != nil {
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
	if err = addOpFinalizeDeviceClaimValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFinalizeDeviceClaim(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFinalizeDeviceClaim(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "FinalizeDeviceClaim",
	}
}
