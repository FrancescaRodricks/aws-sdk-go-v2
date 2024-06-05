// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about an array of TableRestoreStatus objects.
func (c *Client) ListTableRestoreStatus(ctx context.Context, params *ListTableRestoreStatusInput, optFns ...func(*Options)) (*ListTableRestoreStatusOutput, error) {
	if params == nil {
		params = &ListTableRestoreStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTableRestoreStatus", params, optFns, c.addOperationListTableRestoreStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTableRestoreStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTableRestoreStatusInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to display the next page of results.
	MaxResults *int32

	// The namespace from which to list all of the statuses of RestoreTableFromSnapshot
	// operations .
	NamespaceName *string

	// If your initial ListTableRestoreStatus operation returns a nextToken, you can
	// include the returned nextToken in following ListTableRestoreStatus operations.
	// This will return results on the next page.
	NextToken *string

	// The workgroup from which to list all of the statuses of RestoreTableFromSnapshot
	// operations.
	WorkgroupName *string

	noSmithyDocumentSerde
}

type ListTableRestoreStatusOutput struct {

	// If your initial ListTableRestoreStatus operation returns a nextToken , you can
	// include the returned nextToken in following ListTableRestoreStatus operations.
	// This will returns results on the next page.
	NextToken *string

	// The array of returned TableRestoreStatus objects.
	TableRestoreStatuses []types.TableRestoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTableRestoreStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTableRestoreStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTableRestoreStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTableRestoreStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTableRestoreStatus(options.Region), middleware.Before); err != nil {
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

// ListTableRestoreStatusAPIClient is a client that implements the
// ListTableRestoreStatus operation.
type ListTableRestoreStatusAPIClient interface {
	ListTableRestoreStatus(context.Context, *ListTableRestoreStatusInput, ...func(*Options)) (*ListTableRestoreStatusOutput, error)
}

var _ ListTableRestoreStatusAPIClient = (*Client)(nil)

// ListTableRestoreStatusPaginatorOptions is the paginator options for
// ListTableRestoreStatus
type ListTableRestoreStatusPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to display the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTableRestoreStatusPaginator is a paginator for ListTableRestoreStatus
type ListTableRestoreStatusPaginator struct {
	options   ListTableRestoreStatusPaginatorOptions
	client    ListTableRestoreStatusAPIClient
	params    *ListTableRestoreStatusInput
	nextToken *string
	firstPage bool
}

// NewListTableRestoreStatusPaginator returns a new ListTableRestoreStatusPaginator
func NewListTableRestoreStatusPaginator(client ListTableRestoreStatusAPIClient, params *ListTableRestoreStatusInput, optFns ...func(*ListTableRestoreStatusPaginatorOptions)) *ListTableRestoreStatusPaginator {
	if params == nil {
		params = &ListTableRestoreStatusInput{}
	}

	options := ListTableRestoreStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTableRestoreStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTableRestoreStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTableRestoreStatus page.
func (p *ListTableRestoreStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTableRestoreStatusOutput, error) {
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

	result, err := p.client.ListTableRestoreStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTableRestoreStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTableRestoreStatus",
	}
}
