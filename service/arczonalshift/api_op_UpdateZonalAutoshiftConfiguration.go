// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You can update the zonal autoshift status for a resource, to enable or disable
// zonal autoshift. When zonal autoshift is ENABLED , Amazon Web Services shifts
// away resource traffic from an Availability Zone, on your behalf, when Amazon Web
// Services determines that there's an issue in the Availability Zone that could
// potentially affect customers.
func (c *Client) UpdateZonalAutoshiftConfiguration(ctx context.Context, params *UpdateZonalAutoshiftConfigurationInput, optFns ...func(*Options)) (*UpdateZonalAutoshiftConfigurationOutput, error) {
	if params == nil {
		params = &UpdateZonalAutoshiftConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateZonalAutoshiftConfiguration", params, optFns, c.addOperationUpdateZonalAutoshiftConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateZonalAutoshiftConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateZonalAutoshiftConfigurationInput struct {

	// The identifier for the resource that you want to update the zonal autoshift
	// configuration for. The identifier is the Amazon Resource Name (ARN) for the
	// resource.
	//
	// This member is required.
	ResourceIdentifier *string

	// The zonal autoshift status for the resource that you want to update the zonal
	// autoshift configuration for.
	//
	// This member is required.
	ZonalAutoshiftStatus types.ZonalAutoshiftStatus

	noSmithyDocumentSerde
}

type UpdateZonalAutoshiftConfigurationOutput struct {

	// The identifier for the resource that you updated the zonal autoshift
	// configuration for. The identifier is the Amazon Resource Name (ARN) for the
	// resource.
	//
	// This member is required.
	ResourceIdentifier *string

	// The zonal autoshift status for the resource that you updated the zonal
	// autoshift configuration for.
	//
	// This member is required.
	ZonalAutoshiftStatus types.ZonalAutoshiftStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateZonalAutoshiftConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateZonalAutoshiftConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateZonalAutoshiftConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateZonalAutoshiftConfiguration"); err != nil {
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
	if err = addOpUpdateZonalAutoshiftConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateZonalAutoshiftConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateZonalAutoshiftConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateZonalAutoshiftConfiguration",
	}
}
