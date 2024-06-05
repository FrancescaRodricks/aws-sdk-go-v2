// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transcribes patient-clinician conversations and generates clinical notes.
//
// Amazon Web Services HealthScribe automatically provides rich conversation
// transcripts, identifies speaker roles, classifies dialogues, extracts medical
// terms, and generates preliminary clinical notes. To learn more about these
// features, refer to [Amazon Web Services HealthScribe].
//
// To make a StartMedicalScribeJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// You must include the following parameters in your StartMedicalTranscriptionJob
// request:
//
//   - DataAccessRoleArn : The ARN of an IAM role with the these minimum
//     permissions: read permission on input file Amazon S3 bucket specified in Media
//     , write permission on the Amazon S3 bucket specified in OutputBucketName , and
//     full permissions on the KMS key specified in OutputEncryptionKMSKeyId (if
//     set). The role should also allow transcribe.amazonaws.com to assume it.
//
//   - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
//   - MedicalScribeJobName : A custom name you create for your MedicalScribe job
//     that is unique within your Amazon Web Services account.
//
//   - OutputBucketName : The Amazon S3 bucket where you want your output files
//     stored.
//
//   - Settings : A MedicalScribeSettings obect that must set exactly one of
//     ShowSpeakerLabels or ChannelIdentification to true. If ShowSpeakerLabels is
//     true, MaxSpeakerLabels must also be set.
//
//   - ChannelDefinitions : A MedicalScribeChannelDefinitions array should be set
//     if and only if the ChannelIdentification value of Settings is set to true.
//
// [Amazon Web Services HealthScribe]: https://docs.aws.amazon.com/transcribe/latest/dg/health-scribe.html
func (c *Client) StartMedicalScribeJob(ctx context.Context, params *StartMedicalScribeJobInput, optFns ...func(*Options)) (*StartMedicalScribeJobOutput, error) {
	if params == nil {
		params = &StartMedicalScribeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMedicalScribeJob", params, optFns, c.addOperationStartMedicalScribeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMedicalScribeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMedicalScribeJobInput struct {

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access
	// the Amazon S3 bucket that contains your input files, write to the output bucket,
	// and use your KMS key if supplied. If the role that you specify doesn’t have the
	// appropriate permissions your request fails.
	//
	// IAM role ARNs have the format
	// arn:partition:iam::account:role/role-name-with-path . For example:
	// arn:aws:iam::111122223333:role/Admin .
	//
	// For more information, see [IAM ARNs].
	//
	// [IAM ARNs]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns
	//
	// This member is required.
	DataAccessRoleArn *string

	// Describes the Amazon S3 location of the media file you want to use in your
	// request.
	//
	// For information on supported media formats, refer to the MediaFormat parameter
	// or the [Media formats]section in the Amazon S3 Developer Guide.
	//
	// [Media formats]: https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio
	//
	// This member is required.
	Media *types.Media

	// A unique name, chosen by you, for your Medical Scribe job.
	//
	// This name is case sensitive, cannot contain spaces, and must be unique within
	// an Amazon Web Services account. If you try to create a new job with the same
	// name as an existing job, you get a ConflictException error.
	//
	// This member is required.
	MedicalScribeJobName *string

	// The name of the Amazon S3 bucket where you want your Medical Scribe output
	// stored. Do not include the S3:// prefix of the specified bucket.
	//
	// Note that the role specified in the DataAccessRoleArn request parameter must
	// have permission to use the specified location. You can change Amazon S3
	// permissions using the [Amazon Web Services Management Console]. See also [Permissions Required for IAM User Roles].
	//
	// [Amazon Web Services Management Console]: https://console.aws.amazon.com/s3
	// [Permissions Required for IAM User Roles]: https://docs.aws.amazon.com/transcribe/latest/dg/security_iam_id-based-policy-examples.html#auth-role-iam-user
	//
	// This member is required.
	OutputBucketName *string

	// Makes it possible to control how your Medical Scribe job is processed using a
	// MedicalScribeSettings object. Specify ChannelIdentification if
	// ChannelDefinitions are set. Enabled ShowSpeakerLabels if ChannelIdentification
	// and ChannelDefinitions are not set. One and only one of ChannelIdentification
	// and ShowSpeakerLabels must be set. If ShowSpeakerLabels is set, MaxSpeakerLabels
	// must also be set. Use Settings to specify a vocabulary or vocabulary filter or
	// both using VocabularyName , VocabularyFilterName . VocabularyFilterMethod must
	// be specified if VocabularyFilterName is set.
	//
	// This member is required.
	Settings *types.MedicalScribeSettings

	// Makes it possible to specify which speaker is on which channel. For example, if
	// the clinician is the first participant to speak, you would set ChannelId of the
	// first ChannelDefinition in the list to 0 (to indicate the first channel) and
	// ParticipantRole to CLINICIAN (to indicate that it's the clinician speaking).
	// Then you would set the ChannelId of the second ChannelDefinition in the list to
	// 1 (to indicate the second channel) and ParticipantRole to PATIENT (to indicate
	// that it's the patient speaking).
	ChannelDefinitions []types.MedicalScribeChannelDefinition

	// A map of plain text, non-secret key:value pairs, known as encryption context
	// pairs, that provide an added layer of security for your data. For more
	// information, see [KMS encryption context]and [Asymmetric keys in KMS].
	//
	// [Asymmetric keys in KMS]: https://docs.aws.amazon.com/transcribe/latest/dg/symmetric-asymmetric.html
	// [KMS encryption context]: https://docs.aws.amazon.com/transcribe/latest/dg/key-management.html#kms-context
	KMSEncryptionContext map[string]string

	// The KMS key you want to use to encrypt your Medical Scribe output.
	//
	// If using a key located in the current Amazon Web Services account, you can
	// specify your KMS key in one of four ways:
	//
	//   - Use the KMS key ID itself. For example, 1234abcd-12ab-34cd-56ef-1234567890ab
	//   .
	//
	//   - Use an alias for the KMS key ID. For example, alias/ExampleAlias .
	//
	//   - Use the Amazon Resource Name (ARN) for the KMS key ID. For example,
	//   arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab .
	//
	//   - Use the ARN for the KMS key alias. For example,
	//   arn:aws:kms:region:account-ID:alias/ExampleAlias .
	//
	// If using a key located in a different Amazon Web Services account than the
	// current Amazon Web Services account, you can specify your KMS key in one of two
	// ways:
	//
	//   - Use the ARN for the KMS key ID. For example,
	//   arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab .
	//
	//   - Use the ARN for the KMS key alias. For example,
	//   arn:aws:kms:region:account-ID:alias/ExampleAlias .
	//
	// If you do not specify an encryption key, your output is encrypted with the
	// default Amazon S3 key (SSE-S3).
	//
	// Note that the role specified in the DataAccessRoleArn request parameter must
	// have permission to use the specified KMS key.
	OutputEncryptionKMSKeyId *string

	// Adds one or more custom tags, each in the form of a key:value pair, to the
	// Medica Scribe job.
	//
	// To learn more about using tags with Amazon Transcribe, refer to [Tagging resources].
	//
	// [Tagging resources]: https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartMedicalScribeJobOutput struct {

	// Provides detailed information about the current Medical Scribe job, including
	// job status and, if applicable, failure reason.
	MedicalScribeJob *types.MedicalScribeJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMedicalScribeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMedicalScribeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMedicalScribeJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMedicalScribeJob"); err != nil {
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
	if err = addOpStartMedicalScribeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMedicalScribeJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMedicalScribeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMedicalScribeJob",
	}
}
