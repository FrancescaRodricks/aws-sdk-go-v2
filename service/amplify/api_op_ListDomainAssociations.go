// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the domain associations for an Amplify app.
func (c *Client) ListDomainAssociations(ctx context.Context, params *ListDomainAssociationsInput, optFns ...func(*Options)) (*ListDomainAssociationsOutput, error) {
	if params == nil {
		params = &ListDomainAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainAssociations", params, optFns, c.addOperationListDomainAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list domain associations request.
type ListDomainAssociationsInput struct {

	//  The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	//  The maximum number of records to list in a single response.
	MaxResults int32

	//  A pagination token. Set to null to start listing apps from the start. If
	// non-null, a pagination token is returned in a result. Pass its value in here to
	// list more projects.
	NextToken *string

	noSmithyDocumentSerde
}

// The result structure for the list domain association request.
type ListDomainAssociationsOutput struct {

	//  A list of domain associations.
	//
	// This member is required.
	DomainAssociations []types.DomainAssociation

	//  A pagination token. If non-null, a pagination token is returned in a result.
	// Pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDomainAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDomainAssociations"); err != nil {
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
	if err = addOpListDomainAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainAssociations(options.Region), middleware.Before); err != nil {
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

// ListDomainAssociationsAPIClient is a client that implements the
// ListDomainAssociations operation.
type ListDomainAssociationsAPIClient interface {
	ListDomainAssociations(context.Context, *ListDomainAssociationsInput, ...func(*Options)) (*ListDomainAssociationsOutput, error)
}

var _ ListDomainAssociationsAPIClient = (*Client)(nil)

// ListDomainAssociationsPaginatorOptions is the paginator options for
// ListDomainAssociations
type ListDomainAssociationsPaginatorOptions struct {
	//  The maximum number of records to list in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDomainAssociationsPaginator is a paginator for ListDomainAssociations
type ListDomainAssociationsPaginator struct {
	options   ListDomainAssociationsPaginatorOptions
	client    ListDomainAssociationsAPIClient
	params    *ListDomainAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListDomainAssociationsPaginator returns a new ListDomainAssociationsPaginator
func NewListDomainAssociationsPaginator(client ListDomainAssociationsAPIClient, params *ListDomainAssociationsInput, optFns ...func(*ListDomainAssociationsPaginatorOptions)) *ListDomainAssociationsPaginator {
	if params == nil {
		params = &ListDomainAssociationsInput{}
	}

	options := ListDomainAssociationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDomainAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDomainAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDomainAssociations page.
func (p *ListDomainAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDomainAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDomainAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDomainAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDomainAssociations",
	}
}
