// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Creates a new session or modifies an existing session with an Amazon Lex bot.
// Use this operation to enable your application to set the state of the bot.
//
// For more information, see [Managing Sessions].
//
// [Managing Sessions]: https://docs.aws.amazon.com/lex/latest/dg/how-session-api.html
func (c *Client) PutSession(ctx context.Context, params *PutSessionInput, optFns ...func(*Options)) (*PutSessionOutput, error) {
	if params == nil {
		params = &PutSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSession", params, optFns, c.addOperationPutSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSessionInput struct {

	// The alias in use for the bot that contains the session data.
	//
	// This member is required.
	BotAlias *string

	// The name of the bot that contains the session data.
	//
	// This member is required.
	BotName *string

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot.
	//
	// This member is required.
	UserId *string

	// The message that Amazon Lex returns in the response can be either text or
	// speech based depending on the value of this field.
	//
	//   - If the value is text/plain; charset=utf-8 , Amazon Lex returns text in the
	//   response.
	//
	//   - If the value begins with audio/ , Amazon Lex returns speech in the response.
	//   Amazon Lex uses Amazon Polly to generate the speech in the configuration that
	//   you specify. For example, if you specify audio/mpeg as the value, Amazon Lex
	//   returns speech in the MPEG format.
	//
	//   - If the value is audio/pcm , the speech is returned as audio/pcm in 16-bit,
	//   little endian format.
	//
	//   - The following are the accepted values:
	//
	//   - audio/mpeg
	//
	//   - audio/ogg
	//
	//   - audio/pcm
	//
	//   - audio/* (defaults to mpeg)
	//
	//   - text/plain; charset=utf-8
	Accept *string

	// A list of contexts active for the request. A context can be activated when a
	// previous intent is fulfilled, or by including the context in the request,
	//
	// If you don't specify a list of contexts, Amazon Lex will use the current list
	// of contexts for the session. If you specify an empty list, all contexts for the
	// session are cleared.
	ActiveContexts []types.ActiveContext

	// Sets the next action that the bot should take to fulfill the conversation.
	DialogAction *types.DialogAction

	// A summary of the recent intents for the bot. You can use the intent summary
	// view to set a checkpoint label on an intent and modify attributes of intents.
	// You can also use it to remove or add intent summary objects to the list.
	//
	// An intent that you modify or add to the list must make sense for the bot. For
	// example, the intent name must be valid for the bot. You must provide valid
	// values for:
	//
	//   - intentName
	//
	//   - slot names
	//
	//   - slotToElict
	//
	// If you send the recentIntentSummaryView parameter in a PutSession request, the
	// contents of the new summary view replaces the old summary view. For example, if
	// a GetSession request returns three intents in the summary view and you call
	// PutSession with one intent in the summary view, the next call to GetSession
	// will only return one intent.
	RecentIntentSummaryView []types.IntentSummary

	// Map of key/value pairs representing the session-specific context information.
	// It contains application information passed between Amazon Lex and a client
	// application.
	SessionAttributes map[string]string

	noSmithyDocumentSerde
}

type PutSessionOutput struct {

	// A list of active contexts for the session.
	//
	// This value conforms to the media type: application/json
	ActiveContexts *string

	// The audio version of the message to convey to the user.
	AudioStream io.ReadCloser

	// Content type as specified in the Accept HTTP header in the request.
	ContentType *string

	//   - ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response to confirm
	//   the intent before fulfilling an intent.
	//
	//   - ElicitIntent - Amazon Lex wants to elicit the user's intent.
	//
	//   - ElicitSlot - Amazon Lex is expecting the value of a slot for the current
	//   intent.
	//
	//   - Failed - Conveys that the conversation with the user has failed. This can
	//   happen for various reasons, including the user does not provide an appropriate
	//   response to prompts from the service, or if the Lambda function fails to fulfill
	//   the intent.
	//
	//   - Fulfilled - Conveys that the Lambda function has sucessfully fulfilled the
	//   intent.
	//
	//   - ReadyForFulfillment - Conveys that the client has to fulfill the intent.
	DialogState types.DialogState

	// The next message that should be presented to the user.
	//
	// The encodedMessage field is base-64 encoded. You must decode the field before
	// you can use the value.
	EncodedMessage *string

	// The name of the current intent.
	IntentName *string

	// The next message that should be presented to the user.
	//
	// You can only use this field in the de-DE, en-AU, en-GB, en-US, es-419, es-ES,
	// es-US, fr-CA, fr-FR, and it-IT locales. In all other locales, the message field
	// is null. You should use the encodedMessage field instead.
	//
	// Deprecated: The message field is deprecated, use the encodedMessage field
	// instead. The message field is available only in the de-DE, en-AU, en-GB, en-US,
	// es-419, es-ES, es-US, fr-CA, fr-FR and it-IT locales.
	Message *string

	// The format of the response message. One of the following values:
	//
	//   - PlainText - The message contains plain UTF-8 text.
	//
	//   - CustomPayload - The message is a custom format for the client.
	//
	//   - SSML - The message contains text formatted for voice output.
	//
	//   - Composite - The message contains an escaped JSON object containing one or
	//   more messages from the groups that messages were assigned to when the intent was
	//   created.
	MessageFormat types.MessageFormatType

	// Map of key/value pairs representing session-specific context information.
	//
	// This value conforms to the media type: application/json
	SessionAttributes *string

	// A unique identifier for the session.
	SessionId *string

	// If the dialogState is ElicitSlot , returns the name of the slot for which Amazon
	// Lex is eliciting a value.
	SlotToElicit *string

	// Map of zero or more intent slots Amazon Lex detected from the user input during
	// the conversation.
	//
	// Amazon Lex creates a resolution list containing likely values for a slot. The
	// value that it returns is determined by the valueSelectionStrategy selected when
	// the slot type was created or updated. If valueSelectionStrategy is set to
	// ORIGINAL_VALUE , the value provided by the user is returned, if the user value
	// is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list or, if
	// there is no resolution list, null. If you don't specify a valueSelectionStrategy
	// the default is ORIGINAL_VALUE .
	//
	// This value conforms to the media type: application/json
	Slots *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutSession"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpPutSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutSession",
	}
}
