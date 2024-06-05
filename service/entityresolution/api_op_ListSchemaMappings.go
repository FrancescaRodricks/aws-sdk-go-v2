// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the SchemaMappings that have been created for an Amazon
// Web Services account.
func (c *Client) ListSchemaMappings(ctx context.Context, params *ListSchemaMappingsInput, optFns ...func(*Options)) (*ListSchemaMappingsOutput, error) {
	if params == nil {
		params = &ListSchemaMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchemaMappings", params, optFns, c.addOperationListSchemaMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchemaMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemaMappingsInput struct {

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSchemaMappingsOutput struct {

	// The pagination token from the previous API call.
	NextToken *string

	// A list of SchemaMappingSummary objects, each of which contain the fields
	// SchemaName , SchemaArn , CreatedAt , UpdatedAt .
	SchemaList []types.SchemaMappingSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchemaMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSchemaMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSchemaMappings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSchemaMappings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemaMappings(options.Region), middleware.Before); err != nil {
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

// ListSchemaMappingsAPIClient is a client that implements the ListSchemaMappings
// operation.
type ListSchemaMappingsAPIClient interface {
	ListSchemaMappings(context.Context, *ListSchemaMappingsInput, ...func(*Options)) (*ListSchemaMappingsOutput, error)
}

var _ ListSchemaMappingsAPIClient = (*Client)(nil)

// ListSchemaMappingsPaginatorOptions is the paginator options for
// ListSchemaMappings
type ListSchemaMappingsPaginatorOptions struct {
	// The maximum number of objects returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchemaMappingsPaginator is a paginator for ListSchemaMappings
type ListSchemaMappingsPaginator struct {
	options   ListSchemaMappingsPaginatorOptions
	client    ListSchemaMappingsAPIClient
	params    *ListSchemaMappingsInput
	nextToken *string
	firstPage bool
}

// NewListSchemaMappingsPaginator returns a new ListSchemaMappingsPaginator
func NewListSchemaMappingsPaginator(client ListSchemaMappingsAPIClient, params *ListSchemaMappingsInput, optFns ...func(*ListSchemaMappingsPaginatorOptions)) *ListSchemaMappingsPaginator {
	if params == nil {
		params = &ListSchemaMappingsInput{}
	}

	options := ListSchemaMappingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchemaMappingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchemaMappingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchemaMappings page.
func (p *ListSchemaMappingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchemaMappingsOutput, error) {
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

	result, err := p.client.ListSchemaMappings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSchemaMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSchemaMappings",
	}
}
