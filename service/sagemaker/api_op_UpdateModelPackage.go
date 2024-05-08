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

// Updates a versioned model.
func (c *Client) UpdateModelPackage(ctx context.Context, params *UpdateModelPackageInput, optFns ...func(*Options)) (*UpdateModelPackageOutput, error) {
	if params == nil {
		params = &UpdateModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelPackage", params, optFns, c.addOperationUpdateModelPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelPackageInput struct {

	// The Amazon Resource Name (ARN) of the model package.
	//
	// This member is required.
	ModelPackageArn *string

	// An array of additional Inference Specification objects to be added to the
	// existing array additional Inference Specification. Total number of additional
	// Inference Specifications can not exceed 15. Each additional Inference
	// Specification specifies artifacts based on this model package that can be used
	// on inference endpoints. Generally used with SageMaker Neo to store the compiled
	// artifacts.
	AdditionalInferenceSpecificationsToAdd []types.AdditionalInferenceSpecificationDefinition

	// A description for the approval status of the model.
	ApprovalDescription *string

	// The metadata properties associated with the model package versions.
	CustomerMetadataProperties map[string]string

	// The metadata properties associated with the model package versions to remove.
	CustomerMetadataPropertiesToRemove []string

	// Specifies details about inference jobs that you can run with models based on
	// this model package, including the following information:
	//
	//   - The Amazon ECR paths of containers that contain the inference code and
	//   model artifacts.
	//
	//   - The instance types that the model package supports for transform jobs and
	//   real-time endpoints used for inference.
	//
	//   - The input and output content formats that the model package supports for
	//   inference.
	InferenceSpecification *types.InferenceSpecification

	// The approval status of the model.
	ModelApprovalStatus types.ModelApprovalStatus

	// The URI of the source for the model package.
	SourceUri *string

	noSmithyDocumentSerde
}

type UpdateModelPackageOutput struct {

	// The Amazon Resource Name (ARN) of the model.
	//
	// This member is required.
	ModelPackageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateModelPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateModelPackage"); err != nil {
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
	if err = addOpUpdateModelPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModelPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModelPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateModelPackage",
	}
}
