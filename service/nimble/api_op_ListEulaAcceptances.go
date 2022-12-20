// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List EULA acceptances.
func (c *Client) ListEulaAcceptances(ctx context.Context, params *ListEulaAcceptancesInput, optFns ...func(*Options)) (*ListEulaAcceptancesOutput, error) {
	if params == nil {
		params = &ListEulaAcceptancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEulaAcceptances", params, optFns, c.addOperationListEulaAcceptancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEulaAcceptancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEulaAcceptancesInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The list of EULA IDs that have been previously accepted.
	EulaIds []string

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEulaAcceptancesOutput struct {

	// A collection of EULA acceptances.
	EulaAcceptances []types.EulaAcceptance

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEulaAcceptancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEulaAcceptances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEulaAcceptances{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListEulaAcceptancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEulaAcceptances(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListEulaAcceptancesAPIClient is a client that implements the ListEulaAcceptances
// operation.
type ListEulaAcceptancesAPIClient interface {
	ListEulaAcceptances(context.Context, *ListEulaAcceptancesInput, ...func(*Options)) (*ListEulaAcceptancesOutput, error)
}

var _ ListEulaAcceptancesAPIClient = (*Client)(nil)

// ListEulaAcceptancesPaginatorOptions is the paginator options for
// ListEulaAcceptances
type ListEulaAcceptancesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEulaAcceptancesPaginator is a paginator for ListEulaAcceptances
type ListEulaAcceptancesPaginator struct {
	options   ListEulaAcceptancesPaginatorOptions
	client    ListEulaAcceptancesAPIClient
	params    *ListEulaAcceptancesInput
	nextToken *string
	firstPage bool
}

// NewListEulaAcceptancesPaginator returns a new ListEulaAcceptancesPaginator
func NewListEulaAcceptancesPaginator(client ListEulaAcceptancesAPIClient, params *ListEulaAcceptancesInput, optFns ...func(*ListEulaAcceptancesPaginatorOptions)) *ListEulaAcceptancesPaginator {
	if params == nil {
		params = &ListEulaAcceptancesInput{}
	}

	options := ListEulaAcceptancesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEulaAcceptancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEulaAcceptancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEulaAcceptances page.
func (p *ListEulaAcceptancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEulaAcceptancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListEulaAcceptances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEulaAcceptances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "ListEulaAcceptances",
	}
}
