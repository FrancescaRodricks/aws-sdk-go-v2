// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set a protect configuration as your account default. You can only have one
// account default protect configuration at a time. The current account default
// protect configuration is replaced with the provided protect configuration.
func (c *Client) SetAccountDefaultProtectConfiguration(ctx context.Context, params *SetAccountDefaultProtectConfigurationInput, optFns ...func(*Options)) (*SetAccountDefaultProtectConfigurationOutput, error) {
	if params == nil {
		params = &SetAccountDefaultProtectConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetAccountDefaultProtectConfiguration", params, optFns, c.addOperationSetAccountDefaultProtectConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetAccountDefaultProtectConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetAccountDefaultProtectConfigurationInput struct {

	// The unique identifier for the protect configuration.
	//
	// This member is required.
	ProtectConfigurationId *string

	noSmithyDocumentSerde
}

type SetAccountDefaultProtectConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the account default protect configuration.
	//
	// This member is required.
	DefaultProtectConfigurationArn *string

	// The unique identifier of the account default protect configuration.
	//
	// This member is required.
	DefaultProtectConfigurationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetAccountDefaultProtectConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSetAccountDefaultProtectConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSetAccountDefaultProtectConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetAccountDefaultProtectConfiguration"); err != nil {
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
	if err = addOpSetAccountDefaultProtectConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetAccountDefaultProtectConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetAccountDefaultProtectConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetAccountDefaultProtectConfiguration",
	}
}
