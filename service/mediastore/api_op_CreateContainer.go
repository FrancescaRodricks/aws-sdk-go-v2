// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a storage container to hold objects. A container is similar to a bucket
// in the Amazon S3 service.
func (c *Client) CreateContainer(ctx context.Context, params *CreateContainerInput, optFns ...func(*Options)) (*CreateContainerOutput, error) {
	if params == nil {
		params = &CreateContainerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContainer", params, optFns, c.addOperationCreateContainerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContainerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerInput struct {

	// The name for the container. The name must be from 1 to 255 characters.
	// Container names must be unique to your AWS account within a specific region. As
	// an example, you could create a container named movies in every region, as long
	// as you don’t have an existing container with that name.
	//
	// This member is required.
	ContainerName *string

	// An array of key:value pairs that you define. These values can be anything that
	// you want. Typically, the tag key represents a category (such as "environment")
	// and the tag value represents a specific value within that category (such as
	// "test," "development," or "production"). You can add up to 50 tags to each
	// container. For more information about tagging, including naming and usage
	// conventions, see [Tagging Resources in MediaStore].
	//
	// [Tagging Resources in MediaStore]: https://docs.aws.amazon.com/mediastore/latest/ug/tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateContainerOutput struct {

	// ContainerARN: The Amazon Resource Name (ARN) of the newly created container.
	// The ARN has the following format: arn:aws:::container/. For example:
	// arn:aws:mediastore:us-west-2:111122223333:container/movies
	//
	// ContainerName: The container name as specified in the request.
	//
	// CreationTime: Unix time stamp.
	//
	// Status: The status of container creation or deletion. The status is one of the
	// following: CREATING , ACTIVE , or DELETING . While the service is creating the
	// container, the status is CREATING . When an endpoint is available, the status
	// changes to ACTIVE .
	//
	// The return value does not include the container's endpoint. To make downstream
	// requests, you must obtain this value by using DescribeContaineror ListContainers.
	//
	// This member is required.
	Container *types.Container

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContainer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContainer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContainer"); err != nil {
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
	if err = addOpCreateContainerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContainer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContainer",
	}
}
