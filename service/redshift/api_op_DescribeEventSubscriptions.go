// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists descriptions of all the Amazon Redshift event notification subscriptions
// for a customer account. If you specify a subscription name, lists the
// description for that subscription.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all event notification subscriptions that match any combination
// of the specified keys and values. For example, if you have owner and environment
// for tag keys, and admin and test for tag values, all subscriptions that have
// any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, subscriptions are
// returned regardless of whether they have tag keys or values associated with
// them.
func (c *Client) DescribeEventSubscriptions(ctx context.Context, params *DescribeEventSubscriptionsInput, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) {
	if params == nil {
		params = &DescribeEventSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventSubscriptions", params, optFns, c.addOperationDescribeEventSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventSubscriptionsInput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeEventSubscriptions request
	// exceed the value specified in MaxRecords , Amazon Web Services returns a value
	// in the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The name of the Amazon Redshift event notification subscription to be described.
	SubscriptionName *string

	// A tag key or keys for which you want to return all matching event notification
	// subscriptions that are associated with the specified key or keys. For example,
	// suppose that you have subscriptions that are tagged with keys called owner and
	// environment . If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the subscriptions that have either or both of
	// these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching event
	// notification subscriptions that are associated with the specified tag value or
	// values. For example, suppose that you have subscriptions that are tagged with
	// values called admin and test . If you specify both of these tag values in the
	// request, Amazon Redshift returns a response with the subscriptions that have
	// either or both of these tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

type DescribeEventSubscriptionsOutput struct {

	// A list of event subscriptions.
	EventSubscriptionsList []types.EventSubscription

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEventSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEventSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEventSubscriptions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventSubscriptions(options.Region), middleware.Before); err != nil {
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

// DescribeEventSubscriptionsAPIClient is a client that implements the
// DescribeEventSubscriptions operation.
type DescribeEventSubscriptionsAPIClient interface {
	DescribeEventSubscriptions(context.Context, *DescribeEventSubscriptionsInput, ...func(*Options)) (*DescribeEventSubscriptionsOutput, error)
}

var _ DescribeEventSubscriptionsAPIClient = (*Client)(nil)

// DescribeEventSubscriptionsPaginatorOptions is the paginator options for
// DescribeEventSubscriptions
type DescribeEventSubscriptionsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventSubscriptionsPaginator is a paginator for
// DescribeEventSubscriptions
type DescribeEventSubscriptionsPaginator struct {
	options   DescribeEventSubscriptionsPaginatorOptions
	client    DescribeEventSubscriptionsAPIClient
	params    *DescribeEventSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventSubscriptionsPaginator returns a new
// DescribeEventSubscriptionsPaginator
func NewDescribeEventSubscriptionsPaginator(client DescribeEventSubscriptionsAPIClient, params *DescribeEventSubscriptionsInput, optFns ...func(*DescribeEventSubscriptionsPaginatorOptions)) *DescribeEventSubscriptionsPaginator {
	if params == nil {
		params = &DescribeEventSubscriptionsInput{}
	}

	options := DescribeEventSubscriptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEventSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEventSubscriptions page.
func (p *DescribeEventSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeEventSubscriptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEventSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEventSubscriptions",
	}
}
