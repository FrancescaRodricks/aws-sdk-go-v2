// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of event integration associations in the account.
func (c *Client) ListEventIntegrationAssociations(ctx context.Context, params *ListEventIntegrationAssociationsInput, optFns ...func(*Options)) (*ListEventIntegrationAssociationsOutput, error) {
	if params == nil {
		params = &ListEventIntegrationAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventIntegrationAssociations", params, optFns, c.addOperationListEventIntegrationAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventIntegrationAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventIntegrationAssociationsInput struct {

	// The name of the event integration.
	//
	// This member is required.
	EventIntegrationName *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEventIntegrationAssociationsOutput struct {

	// The event integration associations.
	EventIntegrationAssociations []types.EventIntegrationAssociation

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventIntegrationAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventIntegrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventIntegrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEventIntegrationAssociations"); err != nil {
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
	if err = addOpListEventIntegrationAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventIntegrationAssociations(options.Region), middleware.Before); err != nil {
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

// ListEventIntegrationAssociationsAPIClient is a client that implements the
// ListEventIntegrationAssociations operation.
type ListEventIntegrationAssociationsAPIClient interface {
	ListEventIntegrationAssociations(context.Context, *ListEventIntegrationAssociationsInput, ...func(*Options)) (*ListEventIntegrationAssociationsOutput, error)
}

var _ ListEventIntegrationAssociationsAPIClient = (*Client)(nil)

// ListEventIntegrationAssociationsPaginatorOptions is the paginator options for
// ListEventIntegrationAssociations
type ListEventIntegrationAssociationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventIntegrationAssociationsPaginator is a paginator for
// ListEventIntegrationAssociations
type ListEventIntegrationAssociationsPaginator struct {
	options   ListEventIntegrationAssociationsPaginatorOptions
	client    ListEventIntegrationAssociationsAPIClient
	params    *ListEventIntegrationAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListEventIntegrationAssociationsPaginator returns a new
// ListEventIntegrationAssociationsPaginator
func NewListEventIntegrationAssociationsPaginator(client ListEventIntegrationAssociationsAPIClient, params *ListEventIntegrationAssociationsInput, optFns ...func(*ListEventIntegrationAssociationsPaginatorOptions)) *ListEventIntegrationAssociationsPaginator {
	if params == nil {
		params = &ListEventIntegrationAssociationsInput{}
	}

	options := ListEventIntegrationAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventIntegrationAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventIntegrationAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventIntegrationAssociations page.
func (p *ListEventIntegrationAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventIntegrationAssociationsOutput, error) {
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

	result, err := p.client.ListEventIntegrationAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventIntegrationAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventIntegrationAssociations",
	}
}
