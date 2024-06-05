// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the Slack workspace authorization that allows channels to be configured
// in that workspace. This requires all configured channels in the workspace to be
// deleted.
func (c *Client) DeleteSlackWorkspaceAuthorization(ctx context.Context, params *DeleteSlackWorkspaceAuthorizationInput, optFns ...func(*Options)) (*DeleteSlackWorkspaceAuthorizationOutput, error) {
	if params == nil {
		params = &DeleteSlackWorkspaceAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSlackWorkspaceAuthorization", params, optFns, c.addOperationDeleteSlackWorkspaceAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSlackWorkspaceAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSlackWorkspaceAuthorizationInput struct {

	// The ID of the Slack workspace authorized with AWS Chatbot.
	//
	// This member is required.
	SlackTeamId *string

	noSmithyDocumentSerde
}

type DeleteSlackWorkspaceAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSlackWorkspaceAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSlackWorkspaceAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSlackWorkspaceAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteSlackWorkspaceAuthorization"); err != nil {
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
	if err = addOpDeleteSlackWorkspaceAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSlackWorkspaceAuthorization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSlackWorkspaceAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteSlackWorkspaceAuthorization",
	}
}
