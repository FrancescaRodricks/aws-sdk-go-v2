// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists user associations for an identity provider.
func (c *Client) ListUserAssociations(ctx context.Context, params *ListUserAssociationsInput, optFns ...func(*Options)) (*ListUserAssociationsOutput, error) {
	if params == nil {
		params = &ListUserAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserAssociations", params, optFns, c.addOperationListUserAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserAssociationsInput struct {

	// An object that specifies details for the identity provider.
	//
	// This member is required.
	IdentityProvider types.IdentityProvider

	// The ID of the EC2 instance, which provides user-based subscriptions.
	//
	// This member is required.
	InstanceId *string

	// An array of structures that you can use to filter the results to those that
	// match one or more sets of key-value pairs that you specify.
	Filters []types.Filter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUserAssociationsOutput struct {

	// Metadata that describes the list user association operation.
	InstanceUserSummaries []types.InstanceUserSummary

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUserAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUserAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUserAssociations"); err != nil {
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
	if err = addOpListUserAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserAssociations(options.Region), middleware.Before); err != nil {
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

// ListUserAssociationsAPIClient is a client that implements the
// ListUserAssociations operation.
type ListUserAssociationsAPIClient interface {
	ListUserAssociations(context.Context, *ListUserAssociationsInput, ...func(*Options)) (*ListUserAssociationsOutput, error)
}

var _ ListUserAssociationsAPIClient = (*Client)(nil)

// ListUserAssociationsPaginatorOptions is the paginator options for
// ListUserAssociations
type ListUserAssociationsPaginatorOptions struct {
	// Maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserAssociationsPaginator is a paginator for ListUserAssociations
type ListUserAssociationsPaginator struct {
	options   ListUserAssociationsPaginatorOptions
	client    ListUserAssociationsAPIClient
	params    *ListUserAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListUserAssociationsPaginator returns a new ListUserAssociationsPaginator
func NewListUserAssociationsPaginator(client ListUserAssociationsAPIClient, params *ListUserAssociationsInput, optFns ...func(*ListUserAssociationsPaginatorOptions)) *ListUserAssociationsPaginator {
	if params == nil {
		params = &ListUserAssociationsInput{}
	}

	options := ListUserAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUserAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUserAssociations page.
func (p *ListUserAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserAssociationsOutput, error) {
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

	result, err := p.client.ListUserAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUserAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUserAssociations",
	}
}
