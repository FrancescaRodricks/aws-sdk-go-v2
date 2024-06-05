// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stores a pronunciation lexicon in an Amazon Web Services Region. If a lexicon
// with the same name already exists in the region, it is overwritten by the new
// lexicon. Lexicon operations have eventual consistency, therefore, it might take
// some time before the lexicon is available to the SynthesizeSpeech operation.
//
// For more information, see [Managing Lexicons].
//
// [Managing Lexicons]: https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html
func (c *Client) PutLexicon(ctx context.Context, params *PutLexiconInput, optFns ...func(*Options)) (*PutLexiconOutput, error) {
	if params == nil {
		params = &PutLexiconInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLexicon", params, optFns, c.addOperationPutLexiconMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLexiconOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLexiconInput struct {

	// Content of the PLS lexicon as string data.
	//
	// This member is required.
	Content *string

	// Name of the lexicon. The name must follow the regular express format
	// [0-9A-Za-z]{1,20}. That is, the name is a case-sensitive alphanumeric string up
	// to 20 characters long.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type PutLexiconOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLexiconMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLexicon{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLexicon{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutLexicon"); err != nil {
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
	if err = addOpPutLexiconValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLexicon(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLexicon(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutLexicon",
	}
}
