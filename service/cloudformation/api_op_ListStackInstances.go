// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns summary information about stack instances that are associated with the
// specified stack set. You can filter for stack instances that are associated with
// a specific Amazon Web Services account name or Region, or that have a specific
// status.
func (c *Client) ListStackInstances(ctx context.Context, params *ListStackInstancesInput, optFns ...func(*Options)) (*ListStackInstancesOutput, error) {
	if params == nil {
		params = &ListStackInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackInstances", params, optFns, c.addOperationListStackInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackInstancesInput struct {

	// The name or unique ID of the stack set that you want to list stack instances
	// for.
	//
	// This member is required.
	StackSetName *string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account.
	//
	// By default, SELF is specified. Use SELF for stack sets with self-managed
	// permissions.
	//
	//   - If you are signed in to the management account, specify SELF .
	//
	//   - If you are signed in to a delegated administrator account, specify
	//   DELEGATED_ADMIN .
	//
	// Your Amazon Web Services account must be registered as a delegated
	//   administrator in the management account. For more information, see [Register a delegated administrator]in the
	//   CloudFormation User Guide.
	//
	// [Register a delegated administrator]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html
	CallAs types.CallAs

	// The filter to apply to stack instances
	Filters []types.StackInstanceFilter

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int32

	// If the previous request didn't return all the remaining results, the response's
	// NextToken parameter value is set to a token. To retrieve the next set of
	// results, call ListStackInstances again and assign that token to the request
	// object's NextToken parameter. If there are no remaining results, the previous
	// response object's NextToken parameter is set to null .
	NextToken *string

	// The name of the Amazon Web Services account that you want to list stack
	// instances for.
	StackInstanceAccount *string

	// The name of the Region where you want to list stack instances.
	StackInstanceRegion *string

	noSmithyDocumentSerde
}

type ListStackInstancesOutput struct {

	// If the request doesn't return all the remaining results, NextToken is set to a
	// token. To retrieve the next set of results, call ListStackInstances again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null .
	NextToken *string

	// A list of StackInstanceSummary structures that contain information about the
	// specified stack instances.
	Summaries []types.StackInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStackInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStackInstances"); err != nil {
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
	if err = addOpListStackInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStackInstances(options.Region), middleware.Before); err != nil {
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

// ListStackInstancesAPIClient is a client that implements the ListStackInstances
// operation.
type ListStackInstancesAPIClient interface {
	ListStackInstances(context.Context, *ListStackInstancesInput, ...func(*Options)) (*ListStackInstancesOutput, error)
}

var _ ListStackInstancesAPIClient = (*Client)(nil)

// ListStackInstancesPaginatorOptions is the paginator options for
// ListStackInstances
type ListStackInstancesPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStackInstancesPaginator is a paginator for ListStackInstances
type ListStackInstancesPaginator struct {
	options   ListStackInstancesPaginatorOptions
	client    ListStackInstancesAPIClient
	params    *ListStackInstancesInput
	nextToken *string
	firstPage bool
}

// NewListStackInstancesPaginator returns a new ListStackInstancesPaginator
func NewListStackInstancesPaginator(client ListStackInstancesAPIClient, params *ListStackInstancesInput, optFns ...func(*ListStackInstancesPaginatorOptions)) *ListStackInstancesPaginator {
	if params == nil {
		params = &ListStackInstancesInput{}
	}

	options := ListStackInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStackInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStackInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStackInstances page.
func (p *ListStackInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStackInstancesOutput, error) {
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

	result, err := p.client.ListStackInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStackInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStackInstances",
	}
}
