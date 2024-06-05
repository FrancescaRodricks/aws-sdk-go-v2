// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Retrieves a sortable, filterable list of existing Glue machine learning
//
// transforms in this Amazon Web Services account, or the resources with the
// specified tag. This operation takes the optional Tags field, which you can use
// as a filter of the responses so that tagged resources can be retrieved as a
// group. If you choose to use tag filtering, only resources with the tags are
// retrieved.
func (c *Client) ListMLTransforms(ctx context.Context, params *ListMLTransformsInput, optFns ...func(*Options)) (*ListMLTransformsOutput, error) {
	if params == nil {
		params = &ListMLTransformsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMLTransforms", params, optFns, c.addOperationListMLTransformsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMLTransformsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMLTransformsInput struct {

	// A TransformFilterCriteria used to filter the machine learning transforms.
	Filter *types.TransformFilterCriteria

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// A TransformSortCriteria used to sort the machine learning transforms.
	Sort *types.TransformSortCriteria

	// Specifies to return only these tagged resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ListMLTransformsOutput struct {

	// The identifiers of all the machine learning transforms in the account, or the
	// machine learning transforms with the specified tags.
	//
	// This member is required.
	TransformIds []string

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMLTransformsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMLTransforms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMLTransforms{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMLTransforms"); err != nil {
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
	if err = addOpListMLTransformsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMLTransforms(options.Region), middleware.Before); err != nil {
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

// ListMLTransformsAPIClient is a client that implements the ListMLTransforms
// operation.
type ListMLTransformsAPIClient interface {
	ListMLTransforms(context.Context, *ListMLTransformsInput, ...func(*Options)) (*ListMLTransformsOutput, error)
}

var _ ListMLTransformsAPIClient = (*Client)(nil)

// ListMLTransformsPaginatorOptions is the paginator options for ListMLTransforms
type ListMLTransformsPaginatorOptions struct {
	// The maximum size of a list to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMLTransformsPaginator is a paginator for ListMLTransforms
type ListMLTransformsPaginator struct {
	options   ListMLTransformsPaginatorOptions
	client    ListMLTransformsAPIClient
	params    *ListMLTransformsInput
	nextToken *string
	firstPage bool
}

// NewListMLTransformsPaginator returns a new ListMLTransformsPaginator
func NewListMLTransformsPaginator(client ListMLTransformsAPIClient, params *ListMLTransformsInput, optFns ...func(*ListMLTransformsPaginatorOptions)) *ListMLTransformsPaginator {
	if params == nil {
		params = &ListMLTransformsInput{}
	}

	options := ListMLTransformsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMLTransformsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMLTransformsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMLTransforms page.
func (p *ListMLTransformsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMLTransformsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListMLTransforms(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListMLTransforms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMLTransforms",
	}
}
