// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds tags to a resource group with the specified ARN. Existing tags on a
// resource group are not changed if they are not specified in the request
// parameters.
//
// Do not store personally identifiable information (PII) or other confidential or
// sensitive information in tags. We use tags to provide you with billing and
// administration services. Tags are not intended to be used for private or
// sensitive data.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
//   - resource-groups:Tag
func (c *Client) Tag(ctx context.Context, params *TagInput, optFns ...func(*Options)) (*TagOutput, error) {
	if params == nil {
		params = &TagInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Tag", params, optFns, c.addOperationTagMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagInput struct {

	// The ARN of the resource group to which to add tags.
	//
	// This member is required.
	Arn *string

	// The tags to add to the specified resource group. A tag is a string-to-string
	// map of key-value pairs.
	//
	// This member is required.
	Tags map[string]string

	noSmithyDocumentSerde
}

type TagOutput struct {

	// The ARN of the tagged resource.
	Arn *string

	// The tags that have been added to the specified resource group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTag{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTag{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Tag"); err != nil {
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
	if err = addOpTagValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTag(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTag(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Tag",
	}
}
