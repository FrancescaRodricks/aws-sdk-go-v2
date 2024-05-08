// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/controltower/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Updates the configuration of an already enabled control.
//
// If the enabled control shows an EnablementStatus of SUCCEEDED, supply
// parameters that are different from the currently configured parameters.
// Otherwise, Amazon Web Services Control Tower will not accept the request.
//
// If the enabled control shows an EnablementStatus of FAILED, Amazon Web Services
// Control Tower will update the control to match any valid parameters that you
// supply.
//
// If the DriftSummary status for the control shows as DRIFTED, you cannot call
// this API. Instead, you can update the control by calling DisableControl and
// again calling EnableControl , or you can run an extending governance operation.
// For usage examples, see [the Amazon Web Services Control Tower User Guide]
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/control-api-examples-short.html
func (c *Client) UpdateEnabledControl(ctx context.Context, params *UpdateEnabledControlInput, optFns ...func(*Options)) (*UpdateEnabledControlOutput, error) {
	if params == nil {
		params = &UpdateEnabledControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnabledControl", params, optFns, c.addOperationUpdateEnabledControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnabledControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnabledControlInput struct {

	//  The ARN of the enabled control that will be updated.
	//
	// This member is required.
	EnabledControlIdentifier *string

	// A key/value pair, where Key is of type String and Value is of type Document .
	//
	// This member is required.
	Parameters []types.EnabledControlParameter

	noSmithyDocumentSerde
}

type UpdateEnabledControlOutput struct {

	//  The operation identifier for this UpdateEnabledControl operation.
	//
	// This member is required.
	OperationIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnabledControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEnabledControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEnabledControl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnabledControl"); err != nil {
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
	if err = addOpUpdateEnabledControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnabledControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEnabledControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnabledControl",
	}
}
