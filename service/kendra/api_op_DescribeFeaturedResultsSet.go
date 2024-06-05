// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a set of featured results. Features results are placed
// above all other results for certain queries. If there's an exact match of a
// query, then one or more specific documents are featured in the search results.
func (c *Client) DescribeFeaturedResultsSet(ctx context.Context, params *DescribeFeaturedResultsSetInput, optFns ...func(*Options)) (*DescribeFeaturedResultsSetOutput, error) {
	if params == nil {
		params = &DescribeFeaturedResultsSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFeaturedResultsSet", params, optFns, c.addOperationDescribeFeaturedResultsSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFeaturedResultsSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFeaturedResultsSetInput struct {

	// The identifier of the set of featured results that you want to get information
	// on.
	//
	// This member is required.
	FeaturedResultsSetId *string

	// The identifier of the index used for featuring results.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeFeaturedResultsSetOutput struct {

	// The Unix timestamp when the set of the featured results was created.
	CreationTimestamp *int64

	// The description for the set of featured results.
	Description *string

	// The list of document IDs that don't exist but you have specified as featured
	// documents. Amazon Kendra cannot feature these documents if they don't exist in
	// the index. You can check the status of a document and its ID or check for
	// documents with status errors using the [BatchGetDocumentStatus]API.
	//
	// [BatchGetDocumentStatus]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchGetDocumentStatus.html
	FeaturedDocumentsMissing []types.FeaturedDocumentMissing

	// The list of document IDs for the documents you want to feature with their
	// metadata information. For more information on the list of featured documents,
	// see [FeaturedResultsSet].
	//
	// [FeaturedResultsSet]: https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html
	FeaturedDocumentsWithMetadata []types.FeaturedDocumentWithMetadata

	// The identifier of the set of featured results.
	FeaturedResultsSetId *string

	// The name for the set of featured results.
	FeaturedResultsSetName *string

	// The timestamp when the set of featured results was last updated.
	LastUpdatedTimestamp *int64

	// The list of queries for featuring results. For more information on the list of
	// queries, see [FeaturedResultsSet].
	//
	// [FeaturedResultsSet]: https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html
	QueryTexts []string

	// The current status of the set of featured results. When the value is ACTIVE ,
	// featured results are ready for use. You can still configure your settings before
	// setting the status to ACTIVE . You can set the status to ACTIVE or INACTIVE
	// using the [UpdateFeaturedResultsSet]API. The queries you specify for featured results must be unique per
	// featured results set for each index, whether the status is ACTIVE or INACTIVE .
	//
	// [UpdateFeaturedResultsSet]: https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateFeaturedResultsSet.html
	Status types.FeaturedResultsSetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFeaturedResultsSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFeaturedResultsSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFeaturedResultsSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFeaturedResultsSet"); err != nil {
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
	if err = addOpDescribeFeaturedResultsSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFeaturedResultsSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFeaturedResultsSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFeaturedResultsSet",
	}
}
