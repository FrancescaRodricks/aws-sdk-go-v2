// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new custom medical vocabulary.
//
// Before creating a new custom medical vocabulary, you must first upload a text
// file that contains your vocabulary table into an Amazon S3 bucket. Note that
// this differs from , where you can include a list of terms within your request
// using the Phrases flag; CreateMedicalVocabulary does not support the Phrases
// flag and only accepts vocabularies in table format.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Custom vocabularies].
//
// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
func (c *Client) CreateMedicalVocabulary(ctx context.Context, params *CreateMedicalVocabularyInput, optFns ...func(*Options)) (*CreateMedicalVocabularyOutput, error) {
	if params == nil {
		params = &CreateMedicalVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMedicalVocabulary", params, optFns, c.addOperationCreateMedicalVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMedicalVocabularyInput struct {

	// The language code that represents the language of the entries in your custom
	// vocabulary. US English ( en-US ) is the only language supported with Amazon
	// Transcribe Medical.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The Amazon S3 location (URI) of the text file that contains your custom medical
	// vocabulary. The URI must be in the same Amazon Web Services Region as the
	// resource you're calling.
	//
	// Here's an example URI path: s3://DOC-EXAMPLE-BUCKET/my-vocab-file.txt
	//
	// This member is required.
	VocabularyFileUri *string

	// A unique name, chosen by you, for your new custom medical vocabulary.
	//
	// This name is case sensitive, cannot contain spaces, and must be unique within
	// an Amazon Web Services account. If you try to create a new custom medical
	// vocabulary with the same name as an existing custom medical vocabulary, you get
	// a ConflictException error.
	//
	// This member is required.
	VocabularyName *string

	// Adds one or more custom tags, each in the form of a key:value pair, to a new
	// custom medical vocabulary at the time you create this new custom vocabulary.
	//
	// To learn more about using tags with Amazon Transcribe, refer to [Tagging resources].
	//
	// [Tagging resources]: https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMedicalVocabularyOutput struct {

	// If VocabularyState is FAILED , FailureReason contains information about why the
	// medical transcription job request failed. See also: [Common Errors].
	//
	// [Common Errors]: https://docs.aws.amazon.com/transcribe/latest/APIReference/CommonErrors.html
	FailureReason *string

	// The language code you selected for your custom medical vocabulary. US English (
	// en-US ) is the only language supported with Amazon Transcribe Medical.
	LanguageCode types.LanguageCode

	// The date and time you created your custom medical vocabulary.
	//
	// Timestamps are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name you chose for your custom medical vocabulary.
	VocabularyName *string

	// The processing state of your custom medical vocabulary. If the state is READY ,
	// you can use the custom vocabulary in a StartMedicalTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMedicalVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMedicalVocabulary"); err != nil {
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
	if err = addOpCreateMedicalVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMedicalVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMedicalVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMedicalVocabulary",
	}
}
