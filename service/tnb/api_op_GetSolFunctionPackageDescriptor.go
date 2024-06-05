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

// Gets a function package descriptor in a function package.
//
// A function package descriptor is a .yaml file in a function package that uses
// the TOSCA standard to describe how the network function in the function package
// should run on your network.
//
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func (c *Client) GetSolFunctionPackageDescriptor(ctx context.Context, params *GetSolFunctionPackageDescriptorInput, optFns ...func(*Options)) (*GetSolFunctionPackageDescriptorOutput, error) {
	if params == nil {
		params = &GetSolFunctionPackageDescriptorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSolFunctionPackageDescriptor", params, optFns, c.addOperationGetSolFunctionPackageDescriptorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSolFunctionPackageDescriptorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolFunctionPackageDescriptorInput struct {

	// Indicates which content types, expressed as MIME types, the client is able to
	// understand.
	//
	// This member is required.
	Accept types.DescriptorContentType

	// ID of the function package.
	//
	// This member is required.
	VnfPkgId *string

	noSmithyDocumentSerde
}

type GetSolFunctionPackageDescriptorOutput struct {

	// Indicates the media type of the resource.
	ContentType types.DescriptorContentType

	// Contents of the function package descriptor.
	Vnfd []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSolFunctionPackageDescriptorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSolFunctionPackageDescriptor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSolFunctionPackageDescriptor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSolFunctionPackageDescriptor"); err != nil {
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
	if err = addOpGetSolFunctionPackageDescriptorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolFunctionPackageDescriptor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSolFunctionPackageDescriptor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSolFunctionPackageDescriptor",
	}
}
