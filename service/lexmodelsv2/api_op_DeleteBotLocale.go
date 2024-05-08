// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a locale from a bot.
//
// When you delete a locale, all intents, slots, and slot types defined for the
// locale are also deleted.
func (c *Client) DeleteBotLocale(ctx context.Context, params *DeleteBotLocaleInput, optFns ...func(*Options)) (*DeleteBotLocaleOutput, error) {
	if params == nil {
		params = &DeleteBotLocaleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBotLocale", params, optFns, c.addOperationDeleteBotLocaleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotLocaleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotLocaleInput struct {

	// The unique identifier of the bot that contains the locale.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the locale.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale that will be deleted. The string must
	// match one of the supported locales. For more information, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type DeleteBotLocaleOutput struct {

	// The identifier of the bot that contained the deleted locale.
	BotId *string

	// The status of deleting the bot locale. The locale first enters the Deleting
	// status. Once the locale is deleted it no longer appears in the list of locales
	// for the bot.
	BotLocaleStatus types.BotLocaleStatus

	// The version of the bot that contained the deleted locale.
	BotVersion *string

	// The language and locale of the deleted locale.
	LocaleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBotLocaleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBotLocale{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBotLocale{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteBotLocale"); err != nil {
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
	if err = addOpDeleteBotLocaleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBotLocale(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBotLocale(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteBotLocale",
	}
}
