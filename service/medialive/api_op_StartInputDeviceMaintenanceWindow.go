// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Start a maintenance window for the specified input device. Starting a
// maintenance window will give the device up to two hours to install software. If
// the device was streaming prior to the maintenance, it will resume streaming when
// the software is fully installed. Devices automatically install updates while
// they are powered on and their MediaLive channels are stopped. A maintenance
// window allows you to update a device without having to stop MediaLive channels
// that use the device. The device must remain powered on and connected to the
// internet for the duration of the maintenance.
func (c *Client) StartInputDeviceMaintenanceWindow(ctx context.Context, params *StartInputDeviceMaintenanceWindowInput, optFns ...func(*Options)) (*StartInputDeviceMaintenanceWindowOutput, error) {
	if params == nil {
		params = &StartInputDeviceMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartInputDeviceMaintenanceWindow", params, optFns, c.addOperationStartInputDeviceMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartInputDeviceMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for StartInputDeviceMaintenanceWindowRequest
type StartInputDeviceMaintenanceWindowInput struct {

	// The unique ID of the input device to start a maintenance window for. For
	// example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for StartInputDeviceMaintenanceWindowResponse
type StartInputDeviceMaintenanceWindowOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartInputDeviceMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartInputDeviceMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartInputDeviceMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartInputDeviceMaintenanceWindow"); err != nil {
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
	if err = addOpStartInputDeviceMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartInputDeviceMaintenanceWindow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartInputDeviceMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartInputDeviceMaintenanceWindow",
	}
}
