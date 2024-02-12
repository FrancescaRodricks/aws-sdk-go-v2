// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Puts a resource launch action.
func (c *Client) PutLaunchAction(ctx context.Context, params *PutLaunchActionInput, optFns ...func(*Options)) (*PutLaunchActionOutput, error) {
	if params == nil {
		params = &PutLaunchActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLaunchAction", params, optFns, c.addOperationPutLaunchActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLaunchActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLaunchActionInput struct {

	// Launch action code.
	//
	// This member is required.
	ActionCode *string

	// Launch action Id.
	//
	// This member is required.
	ActionId *string

	// Launch action version.
	//
	// This member is required.
	ActionVersion *string

	// Whether the launch action is active.
	//
	// This member is required.
	Active *bool

	// Launch action category.
	//
	// This member is required.
	Category types.LaunchActionCategory

	// Launch action description.
	//
	// This member is required.
	Description *string

	// Launch action name.
	//
	// This member is required.
	Name *string

	// Whether the launch will not be marked as failed if this action fails.
	//
	// This member is required.
	Optional *bool

	// Launch action order.
	//
	// This member is required.
	Order *int32

	// Launch configuration template Id or Source Server Id
	//
	// This member is required.
	ResourceId *string

	// Launch action parameters.
	Parameters map[string]types.LaunchActionParameter

	noSmithyDocumentSerde
}

type PutLaunchActionOutput struct {

	// Launch action code.
	ActionCode *string

	// Launch action Id.
	ActionId *string

	// Launch action version.
	ActionVersion *string

	// Whether the launch action is active.
	Active *bool

	// Launch action category.
	Category types.LaunchActionCategory

	// Launch action description.
	Description *string

	// Launch action name.
	Name *string

	// Whether the launch will not be marked as failed if this action fails.
	Optional *bool

	// Launch action order.
	Order *int32

	// Launch action parameters.
	Parameters map[string]types.LaunchActionParameter

	// Launch configuration template Id or Source Server Id
	ResourceId *string

	// Launch action type.
	Type types.LaunchActionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLaunchActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLaunchAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLaunchAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutLaunchAction"); err != nil {
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
	if err = addOpPutLaunchActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLaunchAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLaunchAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutLaunchAction",
	}
}
