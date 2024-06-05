// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Teams user identity
func (c *Client) DeleteMicrosoftTeamsUserIdentity(ctx context.Context, params *DeleteMicrosoftTeamsUserIdentityInput, optFns ...func(*Options)) (*DeleteMicrosoftTeamsUserIdentityOutput, error) {
	if params == nil {
		params = &DeleteMicrosoftTeamsUserIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMicrosoftTeamsUserIdentity", params, optFns, c.addOperationDeleteMicrosoftTeamsUserIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMicrosoftTeamsUserIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMicrosoftTeamsUserIdentityInput struct {

	// The ARN of the MicrosoftTeamsChannelConfiguration associated with the user
	// identity to delete.
	//
	// This member is required.
	ChatConfigurationArn *string

	// Id from Microsoft Teams for user.
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type DeleteMicrosoftTeamsUserIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMicrosoftTeamsUserIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMicrosoftTeamsUserIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMicrosoftTeamsUserIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteMicrosoftTeamsUserIdentity"); err != nil {
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
	if err = addOpDeleteMicrosoftTeamsUserIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMicrosoftTeamsUserIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMicrosoftTeamsUserIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteMicrosoftTeamsUserIdentity",
	}
}
