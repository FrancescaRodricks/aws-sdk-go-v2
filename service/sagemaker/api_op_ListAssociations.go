// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the associations in your account and their properties.
func (c *Client) ListAssociations(ctx context.Context, params *ListAssociationsInput, optFns ...func(*Options)) (*ListAssociationsOutput, error) {
	if params == nil {
		params = &ListAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociations", params, optFns, c.addOperationListAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociationsInput struct {

	// A filter that returns only associations of the specified type.
	AssociationType types.AssociationEdgeType

	// A filter that returns only associations created on or after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only associations created on or before the specified time.
	CreatedBefore *time.Time

	// A filter that returns only associations with the specified destination Amazon
	// Resource Name (ARN).
	DestinationArn *string

	// A filter that returns only associations with the specified destination type.
	DestinationType *string

	// The maximum number of associations to return in the response. The default value
	// is 10.
	MaxResults *int32

	// If the previous call to ListAssociations didn't return the full set of
	// associations, the call returns a token for getting the next set of associations.
	NextToken *string

	// The property used to sort results. The default value is CreationTime .
	SortBy types.SortAssociationsBy

	// The sort order. The default value is Descending .
	SortOrder types.SortOrder

	// A filter that returns only associations with the specified source ARN.
	SourceArn *string

	// A filter that returns only associations with the specified source type.
	SourceType *string

	noSmithyDocumentSerde
}

type ListAssociationsOutput struct {

	// A list of associations and their properties.
	AssociationSummaries []types.AssociationSummary

	// A token for getting the next set of associations, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociations(options.Region), middleware.Before); err != nil {
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

// ListAssociationsAPIClient is a client that implements the ListAssociations
// operation.
type ListAssociationsAPIClient interface {
	ListAssociations(context.Context, *ListAssociationsInput, ...func(*Options)) (*ListAssociationsOutput, error)
}

var _ ListAssociationsAPIClient = (*Client)(nil)

// ListAssociationsPaginatorOptions is the paginator options for ListAssociations
type ListAssociationsPaginatorOptions struct {
	// The maximum number of associations to return in the response. The default value
	// is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssociationsPaginator is a paginator for ListAssociations
type ListAssociationsPaginator struct {
	options   ListAssociationsPaginatorOptions
	client    ListAssociationsAPIClient
	params    *ListAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListAssociationsPaginator returns a new ListAssociationsPaginator
func NewListAssociationsPaginator(client ListAssociationsAPIClient, params *ListAssociationsInput, optFns ...func(*ListAssociationsPaginatorOptions)) *ListAssociationsPaginator {
	if params == nil {
		params = &ListAssociationsInput{}
	}

	options := ListAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssociations page.
func (p *ListAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssociationsOutput, error) {
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

	result, err := p.client.ListAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssociations",
	}
}
