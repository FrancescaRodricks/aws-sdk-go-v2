// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Re-enables an EnabledBaseline resource. For example, this API can re-apply the
// existing Baseline after a new member account is moved to the target OU.
func (c *Client) ResetEnabledBaseline(ctx context.Context, params *ResetEnabledBaselineInput, optFns ...func(*Options)) (*ResetEnabledBaselineOutput, error) {
	if params == nil {
		params = &ResetEnabledBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetEnabledBaseline", params, optFns, c.addOperationResetEnabledBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetEnabledBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetEnabledBaselineInput struct {

	// Specifies the ID of the EnabledBaseline resource to be re-enabled, in ARN
	// format.
	//
	// This member is required.
	EnabledBaselineIdentifier *string

	noSmithyDocumentSerde
}

type ResetEnabledBaselineOutput struct {

	// The ID (in UUID format) of the asynchronous ResetEnabledBaseline operation.
	// This operationIdentifier is used to track status through calls to the
	// GetBaselineOperation API.
	//
	// This member is required.
	OperationIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetEnabledBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResetEnabledBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResetEnabledBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ResetEnabledBaseline"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpResetEnabledBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetEnabledBaseline(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opResetEnabledBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ResetEnabledBaseline",
	}
}
