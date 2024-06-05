// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists configured audience model associations within a collaboration.
func (c *Client) ListCollaborationConfiguredAudienceModelAssociations(ctx context.Context, params *ListCollaborationConfiguredAudienceModelAssociationsInput, optFns ...func(*Options)) (*ListCollaborationConfiguredAudienceModelAssociationsOutput, error) {
	if params == nil {
		params = &ListCollaborationConfiguredAudienceModelAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollaborationConfiguredAudienceModelAssociations", params, optFns, c.addOperationListCollaborationConfiguredAudienceModelAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollaborationConfiguredAudienceModelAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollaborationConfiguredAudienceModelAssociationsInput struct {

	// A unique identifier for the collaboration that the configured audience model
	// association belongs to. Accepts a collaboration ID.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollaborationConfiguredAudienceModelAssociationsOutput struct {

	// The metadata of the configured audience model association within a
	// collaboration.
	//
	// This member is required.
	CollaborationConfiguredAudienceModelAssociationSummaries []types.CollaborationConfiguredAudienceModelAssociationSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollaborationConfiguredAudienceModelAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCollaborationConfiguredAudienceModelAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCollaborationConfiguredAudienceModelAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCollaborationConfiguredAudienceModelAssociations"); err != nil {
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
	if err = addOpListCollaborationConfiguredAudienceModelAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollaborationConfiguredAudienceModelAssociations(options.Region), middleware.Before); err != nil {
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

// ListCollaborationConfiguredAudienceModelAssociationsAPIClient is a client that
// implements the ListCollaborationConfiguredAudienceModelAssociations operation.
type ListCollaborationConfiguredAudienceModelAssociationsAPIClient interface {
	ListCollaborationConfiguredAudienceModelAssociations(context.Context, *ListCollaborationConfiguredAudienceModelAssociationsInput, ...func(*Options)) (*ListCollaborationConfiguredAudienceModelAssociationsOutput, error)
}

var _ ListCollaborationConfiguredAudienceModelAssociationsAPIClient = (*Client)(nil)

// ListCollaborationConfiguredAudienceModelAssociationsPaginatorOptions is the
// paginator options for ListCollaborationConfiguredAudienceModelAssociations
type ListCollaborationConfiguredAudienceModelAssociationsPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollaborationConfiguredAudienceModelAssociationsPaginator is a paginator
// for ListCollaborationConfiguredAudienceModelAssociations
type ListCollaborationConfiguredAudienceModelAssociationsPaginator struct {
	options   ListCollaborationConfiguredAudienceModelAssociationsPaginatorOptions
	client    ListCollaborationConfiguredAudienceModelAssociationsAPIClient
	params    *ListCollaborationConfiguredAudienceModelAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListCollaborationConfiguredAudienceModelAssociationsPaginator returns a new
// ListCollaborationConfiguredAudienceModelAssociationsPaginator
func NewListCollaborationConfiguredAudienceModelAssociationsPaginator(client ListCollaborationConfiguredAudienceModelAssociationsAPIClient, params *ListCollaborationConfiguredAudienceModelAssociationsInput, optFns ...func(*ListCollaborationConfiguredAudienceModelAssociationsPaginatorOptions)) *ListCollaborationConfiguredAudienceModelAssociationsPaginator {
	if params == nil {
		params = &ListCollaborationConfiguredAudienceModelAssociationsInput{}
	}

	options := ListCollaborationConfiguredAudienceModelAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollaborationConfiguredAudienceModelAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollaborationConfiguredAudienceModelAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next
// ListCollaborationConfiguredAudienceModelAssociations page.
func (p *ListCollaborationConfiguredAudienceModelAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollaborationConfiguredAudienceModelAssociationsOutput, error) {
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

	result, err := p.client.ListCollaborationConfiguredAudienceModelAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCollaborationConfiguredAudienceModelAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCollaborationConfiguredAudienceModelAssociations",
	}
}
