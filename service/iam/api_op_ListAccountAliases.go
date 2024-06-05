// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the account alias associated with the Amazon Web Services account (Note:
// you can have only one). For information about using an Amazon Web Services
// account alias, see [Creating, deleting, and listing an Amazon Web Services account alias]in the Amazon Web Services Sign-In User Guide.
//
// [Creating, deleting, and listing an Amazon Web Services account alias]: https://docs.aws.amazon.com/signin/latest/userguide/CreateAccountAlias.html
func (c *Client) ListAccountAliases(ctx context.Context, params *ListAccountAliasesInput, optFns ...func(*Options)) (*ListAccountAliasesOutput, error) {
	if params == nil {
		params = &ListAccountAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountAliases", params, optFns, c.addOperationListAccountAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountAliasesInput struct {

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	MaxItems *int32

	noSmithyDocumentSerde
}

// Contains the response to a successful ListAccountAliases request.
type ListAccountAliasesOutput struct {

	// A list of aliases associated with the account. Amazon Web Services supports
	// only one alias per account.
	//
	// This member is required.
	AccountAliases []string

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListAccountAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListAccountAliases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccountAliases"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountAliases(options.Region), middleware.Before); err != nil {
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

// ListAccountAliasesAPIClient is a client that implements the ListAccountAliases
// operation.
type ListAccountAliasesAPIClient interface {
	ListAccountAliases(context.Context, *ListAccountAliasesInput, ...func(*Options)) (*ListAccountAliasesOutput, error)
}

var _ ListAccountAliasesAPIClient = (*Client)(nil)

// ListAccountAliasesPaginatorOptions is the paginator options for
// ListAccountAliases
type ListAccountAliasesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountAliasesPaginator is a paginator for ListAccountAliases
type ListAccountAliasesPaginator struct {
	options   ListAccountAliasesPaginatorOptions
	client    ListAccountAliasesAPIClient
	params    *ListAccountAliasesInput
	nextToken *string
	firstPage bool
}

// NewListAccountAliasesPaginator returns a new ListAccountAliasesPaginator
func NewListAccountAliasesPaginator(client ListAccountAliasesAPIClient, params *ListAccountAliasesInput, optFns ...func(*ListAccountAliasesPaginatorOptions)) *ListAccountAliasesPaginator {
	if params == nil {
		params = &ListAccountAliasesInput{}
	}

	options := ListAccountAliasesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountAliases page.
func (p *ListAccountAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountAliasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListAccountAliases(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAccountAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountAliases",
	}
}
