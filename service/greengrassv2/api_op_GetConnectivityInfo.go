// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves connectivity information for a Greengrass core device.
//
// Connectivity information includes endpoints and ports where client devices can
// connect to an MQTT broker on the core device. When a client device calls the [IoT Greengrass discovery API],
// IoT Greengrass returns connectivity information for all of the core devices
// where the client device can connect. For more information, see [Connect client devices to core devices]in the IoT
// Greengrass Version 2 Developer Guide.
//
// [Connect client devices to core devices]: https://docs.aws.amazon.com/greengrass/v2/developerguide/connect-client-devices.html
// [IoT Greengrass discovery API]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-discover-api.html
func (c *Client) GetConnectivityInfo(ctx context.Context, params *GetConnectivityInfoInput, optFns ...func(*Options)) (*GetConnectivityInfoOutput, error) {
	if params == nil {
		params = &GetConnectivityInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnectivityInfo", params, optFns, c.addOperationGetConnectivityInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectivityInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectivityInfoInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	ThingName *string

	noSmithyDocumentSerde
}

type GetConnectivityInfoOutput struct {

	// The connectivity information for the core device.
	ConnectivityInfo []types.ConnectivityInfo

	// A message about the connectivity information request.
	Message *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConnectivityInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConnectivityInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConnectivityInfo{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConnectivityInfo"); err != nil {
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
	if err = addOpGetConnectivityInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnectivityInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConnectivityInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConnectivityInfo",
	}
}
