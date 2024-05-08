// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a list of specified images that are within a repository in a public
// registry. Images are specified with either an imageTag or imageDigest .
//
// You can remove a tag from an image by specifying the image's tag in your
// request. When you remove the last tag from an image, the image is deleted from
// your repository.
//
// You can completely delete an image (and all of its tags) by specifying the
// digest of the image in your request.
func (c *Client) BatchDeleteImage(ctx context.Context, params *BatchDeleteImageInput, optFns ...func(*Options)) (*BatchDeleteImageOutput, error) {
	if params == nil {
		params = &BatchDeleteImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteImage", params, optFns, c.addOperationBatchDeleteImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteImageInput struct {

	// A list of image ID references that correspond to images to delete. The format
	// of the imageIds reference is imageTag=tag or imageDigest=digest .
	//
	// This member is required.
	ImageIds []types.ImageIdentifier

	// The repository in a public registry that contains the image to delete.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID, or registry alias, that's associated with
	// the registry that contains the image to delete. If you do not specify a
	// registry, the default public registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type BatchDeleteImageOutput struct {

	// Any failures associated with the call.
	Failures []types.ImageFailure

	// The image IDs of the deleted images.
	ImageIds []types.ImageIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteImage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchDeleteImage"); err != nil {
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
	if err = addOpBatchDeleteImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchDeleteImage",
	}
}
