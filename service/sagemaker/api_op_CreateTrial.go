// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an SageMaker trial. A trial is a set of steps called trial components
// that produce a machine learning model. A trial is part of a single SageMaker
// experiment.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to a trial and then use the [Search] API to search for the tags.
//
// To get a list of all your trials, call the [ListTrials] API. To view a trial's properties,
// call the [DescribeTrial]API. To create a trial component, call the [CreateTrialComponent] API.
//
// [DescribeTrial]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrial.html
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [ListTrials]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListTrials.html
// [CreateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrialComponent.html
func (c *Client) CreateTrial(ctx context.Context, params *CreateTrialInput, optFns ...func(*Options)) (*CreateTrialOutput, error) {
	if params == nil {
		params = &CreateTrialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrial", params, optFns, c.addOperationCreateTrialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrialInput struct {

	// The name of the experiment to associate the trial with.
	//
	// This member is required.
	ExperimentName *string

	// The name of the trial. The name must be unique in your Amazon Web Services
	// account and is not case-sensitive.
	//
	// This member is required.
	TrialName *string

	// The name of the trial as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialName is displayed.
	DisplayName *string

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// A list of tags to associate with the trial. You can use [Search] API to search on the
	// tags.
	//
	// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTrialOutput struct {

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrial{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrial{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrial"); err != nil {
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
	if err = addOpCreateTrialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrial(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrial(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrial",
	}
}
