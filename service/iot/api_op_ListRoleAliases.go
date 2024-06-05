// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the role aliases registered in your account.
//
// Requires permission to access the [ListRoleAliases] action.
//
// [ListRoleAliases]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListRoleAliases(ctx context.Context, params *ListRoleAliasesInput, optFns ...func(*Options)) (*ListRoleAliasesOutput, error) {
	if params == nil {
		params = &ListRoleAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoleAliases", params, optFns, c.addOperationListRoleAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoleAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoleAliasesInput struct {

	// Return the list of role aliases in ascending alphabetical order.
	AscendingOrder bool

	// A marker used to get the next set of results.
	Marker *string

	// The maximum number of results to return at one time.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListRoleAliasesOutput struct {

	// A marker used to get the next set of results.
	NextMarker *string

	// The role aliases.
	RoleAliases []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoleAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRoleAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoleAliases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRoleAliases"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoleAliases(options.Region), middleware.Before); err != nil {
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

// ListRoleAliasesAPIClient is a client that implements the ListRoleAliases
// operation.
type ListRoleAliasesAPIClient interface {
	ListRoleAliases(context.Context, *ListRoleAliasesInput, ...func(*Options)) (*ListRoleAliasesOutput, error)
}

var _ ListRoleAliasesAPIClient = (*Client)(nil)

// ListRoleAliasesPaginatorOptions is the paginator options for ListRoleAliases
type ListRoleAliasesPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoleAliasesPaginator is a paginator for ListRoleAliases
type ListRoleAliasesPaginator struct {
	options   ListRoleAliasesPaginatorOptions
	client    ListRoleAliasesAPIClient
	params    *ListRoleAliasesInput
	nextToken *string
	firstPage bool
}

// NewListRoleAliasesPaginator returns a new ListRoleAliasesPaginator
func NewListRoleAliasesPaginator(client ListRoleAliasesAPIClient, params *ListRoleAliasesInput, optFns ...func(*ListRoleAliasesPaginatorOptions)) *ListRoleAliasesPaginator {
	if params == nil {
		params = &ListRoleAliasesInput{}
	}

	options := ListRoleAliasesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoleAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoleAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoleAliases page.
func (p *ListRoleAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoleAliasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListRoleAliases(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRoleAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRoleAliases",
	}
}
