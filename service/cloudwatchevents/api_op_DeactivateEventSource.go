// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You can use this operation to temporarily stop receiving events from the
// specified partner event source. The matching event bus is not deleted.
//
// When you deactivate a partner event source, the source goes into PENDING state.
// If it remains in PENDING state for more than two weeks, it is deleted.
//
// To activate a deactivated partner event source, use [ActivateEventSource].
//
// [ActivateEventSource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ActivateEventSource.html
func (c *Client) DeactivateEventSource(ctx context.Context, params *DeactivateEventSourceInput, optFns ...func(*Options)) (*DeactivateEventSourceOutput, error) {
	if params == nil {
		params = &DeactivateEventSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeactivateEventSource", params, optFns, c.addOperationDeactivateEventSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeactivateEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeactivateEventSourceInput struct {

	// The name of the partner event source to deactivate.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeactivateEventSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeactivateEventSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeactivateEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeactivateEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeactivateEventSource"); err != nil {
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
	if err = addOpDeactivateEventSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivateEventSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeactivateEventSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeactivateEventSource",
	}
}
