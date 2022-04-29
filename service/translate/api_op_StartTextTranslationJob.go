// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous batch translation job. Batch translation jobs can be used
// to translate large volumes of text across multiple documents at once. For more
// information, see async. Batch translation jobs can be described with the
// DescribeTextTranslationJob operation, listed with the ListTextTranslationJobs
// operation, and stopped with the StopTextTranslationJob operation. Amazon
// Translate does not support batch translation of multiple source languages at
// once.
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

	// A unique identifier for the request. This token is auto-generated when using the
	// Amazon Translate SDK.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that grants Amazon Translate read access to your input data. For more
	// information, see identity-and-access-management.
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and S3 location of the input documents for the translation
	// job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// Specifies the S3 folder to which your job output will be saved.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The language code of the input language. For a list of language codes, see
	// what-is-languages. Amazon Translate does not automatically detect a source
	// language during batch translation jobs.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code of the output language.
	//
	// This member is required.
	TargetLanguageCodes []string

	// The name of the batch translation job to be performed.
	JobName *string

	// The name of a parallel data resource to add to the translation job. This
	// resource consists of examples that show how you want segments of text to be
	// translated. When you add parallel data to a translation job, you create an
	// Active Custom Translation job. This parameter accepts only one parallel data
	// resource. Active Custom Translation jobs are priced at a higher rate than other
	// jobs that don't use parallel data. For more information, see Amazon Translate
	// pricing (http://aws.amazon.com/translate/pricing/). For a list of available
	// parallel data resources, use the ListParallelData operation. For more
	// information, see customizing-translations-parallel-data.
	ParallelDataNames []string

	// Settings to configure your translation output, including the option to mask
	// profane words and phrases.
	Settings *types.TranslationSettings

	// The name of a custom terminology resource to add to the translation job. This
	// resource lists examples source terms and the desired translation for each term.
	// This parameter accepts only one custom terminology resource. For a list of
	// available custom terminology resources, use the ListTerminologies operation. For
	// more information, see how-custom-terminology.
	TerminologyNames []string

	noSmithyDocumentSerde
}

type StartTextTranslationJobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this ID
	// with the DescribeTextTranslationJob operation.
	JobId *string

	// The status of the job. Possible values include:
	//
	// * SUBMITTED - The job has been
	// received and is queued for processing.
	//
	// * IN_PROGRESS - Amazon Translate is
	// processing the job.
	//
	// * COMPLETED - The job was successfully completed and the
	// output is available.
	//
	// * COMPLETED_WITH_ERROR - The job was completed with
	// errors. The errors can be analyzed in the job's output.
	//
	// * FAILED - The job did
	// not complete. To get details, use the DescribeTextTranslationJob operation.
	//
	// *
	// STOP_REQUESTED - The user who started the job has requested that it be
	// stopped.
	//
	// * STOPPED - The job has been stopped.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTextTranslationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTextTranslationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTextTranslationJob{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
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
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
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
		SigningName:   "translate",
		OperationName: "StartTextTranslationJob",
	}
}