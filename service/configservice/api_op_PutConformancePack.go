// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a conformance pack. A conformance pack is a collection of
// Config rules that can be easily deployed in an account and a region and across
// an organization. For information on how many conformance packs you can have per
// account, see [Service Limits]in the Config Developer Guide.
//
// This API creates a service-linked role AWSServiceRoleForConfigConforms in your
// account. The service-linked role is created only when the role does not exist in
// your account.
//
// You must specify only one of the follow parameters: TemplateS3Uri , TemplateBody
// or TemplateSSMDocumentDetails .
//
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
func (c *Client) PutConformancePack(ctx context.Context, params *PutConformancePackInput, optFns ...func(*Options)) (*PutConformancePackOutput, error) {
	if params == nil {
		params = &PutConformancePackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConformancePack", params, optFns, c.addOperationPutConformancePackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConformancePackInput struct {

	// The unique name of the conformance pack you want to deploy.
	//
	// This member is required.
	ConformancePackName *string

	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []types.ConformancePackInputParameter

	// The name of the Amazon S3 bucket where Config stores conformance pack templates.
	//
	// This field is optional.
	DeliveryS3Bucket *string

	// The prefix for the Amazon S3 bucket.
	//
	// This field is optional.
	DeliveryS3KeyPrefix *string

	// A string containing the full conformance pack template body. The structure
	// containing the template body has a minimum length of 1 byte and a maximum length
	// of 51,200 bytes.
	//
	// You can use a YAML template with two resource types: Config rule (
	// AWS::Config::ConfigRule ) and remediation action (
	// AWS::Config::RemediationConfiguration ).
	TemplateBody *string

	// The location of the file containing the template body ( s3://bucketname/prefix
	// ). The uri must point to a conformance pack template (max size: 300 KB) that is
	// located in an Amazon S3 bucket in the same Region as the conformance pack.
	//
	// You must have access to read Amazon S3 bucket. In addition, in order to ensure
	// a successful deployment, the template object must not be in an [archived storage class]if this
	// parameter is passed.
	//
	// [archived storage class]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html
	TemplateS3Uri *string

	// An object of type TemplateSSMDocumentDetails , which contains the name or the
	// Amazon Resource Name (ARN) of the Amazon Web Services Systems Manager document
	// (SSM document) and the version of the SSM document that is used to create a
	// conformance pack.
	TemplateSSMDocumentDetails *types.TemplateSSMDocumentDetails

	noSmithyDocumentSerde
}

type PutConformancePackOutput struct {

	// ARN of the conformance pack.
	ConformancePackArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutConformancePackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutConformancePack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConformancePack{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutConformancePack"); err != nil {
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
	if err = addOpPutConformancePackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConformancePack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConformancePack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutConformancePack",
	}
}
