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

// Returns a paginated list of DataIntegration associations in the account.
//
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func (c *Client) ListDataIntegrationAssociations(ctx context.Context, params *ListDataIntegrationAssociationsInput, optFns ...func(*Options)) (*ListDataIntegrationAssociationsOutput, error) {
	if params == nil {
		params = &ListDataIntegrationAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataIntegrationAssociations", params, optFns, c.addOperationListDataIntegrationAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataIntegrationAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataIntegrationAssociationsInput struct {

	// A unique identifier for the DataIntegration.
	//
	// This member is required.
	DataIntegrationIdentifier *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataIntegrationAssociationsOutput struct {

	// The Amazon Resource Name (ARN) and unique ID of the DataIntegration association.
	DataIntegrationAssociations []types.DataIntegrationAssociationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataIntegrationAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataIntegrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataIntegrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataIntegrationAssociations"); err != nil {
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
	if err = addOpListDataIntegrationAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataIntegrationAssociations(options.Region), middleware.Before); err != nil {
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

// ListDataIntegrationAssociationsAPIClient is a client that implements the
// ListDataIntegrationAssociations operation.
type ListDataIntegrationAssociationsAPIClient interface {
	ListDataIntegrationAssociations(context.Context, *ListDataIntegrationAssociationsInput, ...func(*Options)) (*ListDataIntegrationAssociationsOutput, error)
}

var _ ListDataIntegrationAssociationsAPIClient = (*Client)(nil)

// ListDataIntegrationAssociationsPaginatorOptions is the paginator options for
// ListDataIntegrationAssociations
type ListDataIntegrationAssociationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataIntegrationAssociationsPaginator is a paginator for
// ListDataIntegrationAssociations
type ListDataIntegrationAssociationsPaginator struct {
	options   ListDataIntegrationAssociationsPaginatorOptions
	client    ListDataIntegrationAssociationsAPIClient
	params    *ListDataIntegrationAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListDataIntegrationAssociationsPaginator returns a new
// ListDataIntegrationAssociationsPaginator
func NewListDataIntegrationAssociationsPaginator(client ListDataIntegrationAssociationsAPIClient, params *ListDataIntegrationAssociationsInput, optFns ...func(*ListDataIntegrationAssociationsPaginatorOptions)) *ListDataIntegrationAssociationsPaginator {
	if params == nil {
		params = &ListDataIntegrationAssociationsInput{}
	}

	options := ListDataIntegrationAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataIntegrationAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataIntegrationAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataIntegrationAssociations page.
func (p *ListDataIntegrationAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataIntegrationAssociationsOutput, error) {
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

	result, err := p.client.ListDataIntegrationAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataIntegrationAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataIntegrationAssociations",
	}
}
