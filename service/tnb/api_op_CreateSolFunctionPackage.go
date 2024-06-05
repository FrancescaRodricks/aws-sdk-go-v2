// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a function package.
//
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network. For more information, see [Function packages]in the
// Amazon Web Services Telco Network Builder User Guide.
//
// Creating a function package is the first step for creating a network in AWS
// TNB. This request creates an empty container with an ID. The next step is to
// upload the actual CSAR zip file into that empty container. To upload function
// package content, see [PutSolFunctionPackageContent].
//
// [Function packages]: https://docs.aws.amazon.com/tnb/latest/ug/function-packages.html
// [PutSolFunctionPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolFunctionPackageContent.html
func (c *Client) CreateSolFunctionPackage(ctx context.Context, params *CreateSolFunctionPackageInput, optFns ...func(*Options)) (*CreateSolFunctionPackageOutput, error) {
	if params == nil {
		params = &CreateSolFunctionPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSolFunctionPackage", params, optFns, c.addOperationCreateSolFunctionPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSolFunctionPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolFunctionPackageInput struct {

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSolFunctionPackageOutput struct {

	// Function package ARN.
	//
	// This member is required.
	Arn *string

	// ID of the function package.
	//
	// This member is required.
	Id *string

	// Onboarding state of the function package.
	//
	// This member is required.
	OnboardingState types.OnboardingState

	// Operational state of the function package.
	//
	// This member is required.
	OperationalState types.OperationalState

	// Usage state of the function package.
	//
	// This member is required.
	UsageState types.UsageState

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSolFunctionPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSolFunctionPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSolFunctionPackage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSolFunctionPackage"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolFunctionPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSolFunctionPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSolFunctionPackage",
	}
}
