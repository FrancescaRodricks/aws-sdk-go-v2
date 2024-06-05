// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the bandwidth rate limits of a gateway. You can update both the upload
// and download bandwidth rate limit or specify only one of the two. If you don't
// set a bandwidth rate limit, the existing rate limit remains. This operation is
// supported only for the stored volume, cached volume, and tape gateway types. To
// update bandwidth rate limits for S3 file gateways, use UpdateBandwidthRateLimitSchedule.
//
// By default, a gateway's bandwidth rate limits are not set. If you don't set any
// limit, the gateway does not have any limitations on its bandwidth usage and
// could potentially use the maximum available bandwidth.
//
// To specify which gateway to update, use the Amazon Resource Name (ARN) of the
// gateway in your request.
func (c *Client) UpdateBandwidthRateLimit(ctx context.Context, params *UpdateBandwidthRateLimitInput, optFns ...func(*Options)) (*UpdateBandwidthRateLimitOutput, error) {
	if params == nil {
		params = &UpdateBandwidthRateLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBandwidthRateLimit", params, optFns, c.addOperationUpdateBandwidthRateLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBandwidthRateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//
// # UpdateBandwidthRateLimitInput$AverageDownloadRateLimitInBitsPerSec
//
// UpdateBandwidthRateLimitInput$AverageUploadRateLimitInBitsPerSec
type UpdateBandwidthRateLimitInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The average download bandwidth rate limit in bits per second.
	AverageDownloadRateLimitInBitsPerSec *int64

	// The average upload bandwidth rate limit in bits per second.
	AverageUploadRateLimitInBitsPerSec *int64

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// throttle information was updated.
type UpdateBandwidthRateLimitOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBandwidthRateLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBandwidthRateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBandwidthRateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBandwidthRateLimit"); err != nil {
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
	if err = addOpUpdateBandwidthRateLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBandwidthRateLimit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBandwidthRateLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBandwidthRateLimit",
	}
}
