// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous batch translation job. Use batch translation jobs to
// translate large volumes of text across multiple documents at once. For batch
// translation, you can input documents with different source languages (specify
// auto as the source language). You can specify one or more target languages.
// Batch translation translates each input document into each of the target
// languages. For more information, see [Asynchronous batch processing].
//
// Batch translation jobs can be described with the DescribeTextTranslationJob operation, listed with the ListTextTranslationJobs
// operation, and stopped with the StopTextTranslationJoboperation.
//
// [Asynchronous batch processing]: https://docs.aws.amazon.com/translate/latest/dg/async.html
func (c *Client) StartTextTranslationJob(ctx context.Context, params *StartTextTranslationJobInput, optFns ...func(*Options)) (*StartTextTranslationJobOutput, error) {
	if params == nil {
		params = &StartTextTranslationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTextTranslationJob", params, optFns, c.addOperationStartTextTranslationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTextTranslationJobInput struct {

	// A unique identifier for the request. This token is generated for you when using
	// the Amazon Translate SDK.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that grants Amazon Translate read access to your input data. For more
	// information, see [Identity and access management].
	//
	// [Identity and access management]: https://docs.aws.amazon.com/translate/latest/dg/identity-and-access-management.html
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input documents for the translation
	// job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// Specifies the S3 folder to which your job output will be saved.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The language code of the input language. Specify the language if all input
	// documents share the same language. If you don't know the language of the source
	// files, or your input documents contains different source languages, select auto
	// . Amazon Translate auto detects the source language for each input document. For
	// a list of supported language codes, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
	//
	// This member is required.
	SourceLanguageCode *string

	// The target languages of the translation job. Enter up to 10 language codes.
	// Each input file is translated into each target language.
	//
	// Each language code is 2 or 5 characters long. For a list of language codes, see [Supported languages]
	// .
	//
	// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
	//
	// This member is required.
	TargetLanguageCodes []string

	// The name of the batch translation job to be performed.
	JobName *string

	// The name of a parallel data resource to add to the translation job. This
	// resource consists of examples that show how you want segments of text to be
	// translated. If you specify multiple target languages for the job, the parallel
	// data file must include translations for all the target languages.
	//
	// When you add parallel data to a translation job, you create an Active Custom
	// Translation job.
	//
	// This parameter accepts only one parallel data resource.
	//
	// Active Custom Translation jobs are priced at a higher rate than other jobs that
	// don't use parallel data. For more information, see [Amazon Translate pricing].
	//
	// For a list of available parallel data resources, use the ListParallelData operation.
	//
	// For more information, see [Customizing your translations with parallel data].
	//
	// [Customizing your translations with parallel data]: https://docs.aws.amazon.com/translate/latest/dg/customizing-translations-parallel-data.html
	// [Amazon Translate pricing]: http://aws.amazon.com/translate/pricing/
	ParallelDataNames []string

	// Settings to configure your translation output. You can configure the following
	// options:
	//
	//   - Brevity: not supported.
	//
	//   - Formality: sets the formality level of the output text.
	//
	//   - Profanity: masks profane words and phrases in your translation output.
	Settings *types.TranslationSettings

	// The name of a custom terminology resource to add to the translation job. This
	// resource lists examples source terms and the desired translation for each term.
	//
	// This parameter accepts only one custom terminology resource.
	//
	// If you specify multiple target languages for the job, translate uses the
	// designated terminology for each requested target language that has an entry for
	// the source term in the terminology file.
	//
	// For a list of available custom terminology resources, use the ListTerminologies operation.
	//
	// For more information, see [Custom terminology].
	//
	// [Custom terminology]: https://docs.aws.amazon.com/translate/latest/dg/how-custom-terminology.html
	TerminologyNames []string

	noSmithyDocumentSerde
}

type StartTextTranslationJobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this ID
	// with the DescribeTextTranslationJoboperation.
	JobId *string

	// The status of the job. Possible values include:
	//
	//   - SUBMITTED - The job has been received and is queued for processing.
	//
	//   - IN_PROGRESS - Amazon Translate is processing the job.
	//
	//   - COMPLETED - The job was successfully completed and the output is available.
	//
	//   - COMPLETED_WITH_ERROR - The job was completed with errors. The errors can be
	//   analyzed in the job's output.
	//
	//   - FAILED - The job did not complete. To get details, use the DescribeTextTranslationJoboperation.
	//
	//   - STOP_REQUESTED - The user who started the job has requested that it be
	//   stopped.
	//
	//   - STOPPED - The job has been stopped.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTextTranslationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTextTranslationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTextTranslationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartTextTranslationJob"); err != nil {
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
	if err = addIdempotencyToken_opStartTextTranslationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartTextTranslationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTextTranslationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartTextTranslationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartTextTranslationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartTextTranslationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartTextTranslationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartTextTranslationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartTextTranslationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartTextTranslationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartTextTranslationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartTextTranslationJob",
	}
}
