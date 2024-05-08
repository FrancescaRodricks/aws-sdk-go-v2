// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides metadata information about a bot recommendation. This information will
// enable you to get a description on the request inputs, to download associated
// transcripts after processing is complete, and to download intents and slot-types
// generated by the bot recommendation.
func (c *Client) DescribeBotRecommendation(ctx context.Context, params *DescribeBotRecommendationInput, optFns ...func(*Options)) (*DescribeBotRecommendationOutput, error) {
	if params == nil {
		params = &DescribeBotRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBotRecommendation", params, optFns, c.addOperationDescribeBotRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotRecommendationInput struct {

	// The unique identifier of the bot associated with the bot recommendation.
	//
	// This member is required.
	BotId *string

	// The identifier of the bot recommendation to describe.
	//
	// This member is required.
	BotRecommendationId *string

	// The version of the bot associated with the bot recommendation.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale of the bot recommendation to
	// describe. The string must match one of the supported locales. For more
	// information, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type DescribeBotRecommendationOutput struct {

	// The identifier of the bot associated with the bot recommendation.
	BotId *string

	// The identifier of the bot recommendation being described.
	BotRecommendationId *string

	// The object representing the URL of the bot definition, the URL of the
	// associated transcript and a statistical summary of the bot recommendation
	// results.
	BotRecommendationResults *types.BotRecommendationResults

	// The status of the bot recommendation. If the status is Failed, then the reasons
	// for the failure are listed in the failureReasons field.
	BotRecommendationStatus types.BotRecommendationStatus

	// The version of the bot associated with the bot recommendation.
	BotVersion *string

	// The date and time that the bot recommendation was created.
	CreationDateTime *time.Time

	// The object representing the passwords that were used to encrypt the data
	// related to the bot recommendation results, as well as the KMS key ARN used to
	// encrypt the associated metadata.
	EncryptionSetting *types.EncryptionSetting

	// If botRecommendationStatus is Failed, Amazon Lex explains why.
	FailureReasons []string

	// The date and time that the bot recommendation was last updated.
	LastUpdatedDateTime *time.Time

	// The identifier of the language and locale of the bot recommendation to describe.
	LocaleId *string

	// The object representing the Amazon S3 bucket containing the transcript, as well
	// as the associated metadata.
	TranscriptSourceSetting *types.TranscriptSourceSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBotRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBotRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBotRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBotRecommendation"); err != nil {
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
	if err = addOpDescribeBotRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBotRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBotRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBotRecommendation",
	}
}
