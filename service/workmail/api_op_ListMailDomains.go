// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the mail domains in a given WorkMail organization.
func (c *Client) ListMailDomains(ctx context.Context, params *ListMailDomainsInput, optFns ...func(*Options)) (*ListMailDomainsOutput, error) {
	if params == nil {
		params = &ListMailDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMailDomains", params, optFns, c.addOperationListMailDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMailDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMailDomainsInput struct {

	// The WorkMail organization for which to list domains.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results. The first call does not
	// require a token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMailDomainsOutput struct {

	// The list of mail domain summaries, specifying domains that exist in the
	// specified WorkMail organization, along with the information about whether the
	// domain is or isn't the default.
	MailDomains []types.MailDomainSummary

	// The token to use to retrieve the next page of results. The value becomes null
	// when there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMailDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMailDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMailDomains{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMailDomains"); err != nil {
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
	if err = addOpListMailDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMailDomains(options.Region), middleware.Before); err != nil {
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

// ListMailDomainsAPIClient is a client that implements the ListMailDomains
// operation.
type ListMailDomainsAPIClient interface {
	ListMailDomains(context.Context, *ListMailDomainsInput, ...func(*Options)) (*ListMailDomainsOutput, error)
}

var _ ListMailDomainsAPIClient = (*Client)(nil)

// ListMailDomainsPaginatorOptions is the paginator options for ListMailDomains
type ListMailDomainsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMailDomainsPaginator is a paginator for ListMailDomains
type ListMailDomainsPaginator struct {
	options   ListMailDomainsPaginatorOptions
	client    ListMailDomainsAPIClient
	params    *ListMailDomainsInput
	nextToken *string
	firstPage bool
}

// NewListMailDomainsPaginator returns a new ListMailDomainsPaginator
func NewListMailDomainsPaginator(client ListMailDomainsAPIClient, params *ListMailDomainsInput, optFns ...func(*ListMailDomainsPaginatorOptions)) *ListMailDomainsPaginator {
	if params == nil {
		params = &ListMailDomainsInput{}
	}

	options := ListMailDomainsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMailDomainsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMailDomainsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMailDomains page.
func (p *ListMailDomainsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMailDomainsOutput, error) {
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

	result, err := p.client.ListMailDomains(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMailDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMailDomains",
	}
}
