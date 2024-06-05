// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List profile notifications.
func (c *Client) ListProfileNotifications(ctx context.Context, params *ListProfileNotificationsInput, optFns ...func(*Options)) (*ListProfileNotificationsOutput, error) {
	if params == nil {
		params = &ListProfileNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfileNotifications", params, optFns, c.addOperationListProfileNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfileNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfileNotificationsInput struct {

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	noSmithyDocumentSerde
}

type ListProfileNotificationsOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// Notification summaries.
	NotificationSummaries []types.ProfileNotificationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfileNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfileNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfileNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProfileNotifications"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfileNotifications(options.Region), middleware.Before); err != nil {
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

// ListProfileNotificationsAPIClient is a client that implements the
// ListProfileNotifications operation.
type ListProfileNotificationsAPIClient interface {
	ListProfileNotifications(context.Context, *ListProfileNotificationsInput, ...func(*Options)) (*ListProfileNotificationsOutput, error)
}

var _ ListProfileNotificationsAPIClient = (*Client)(nil)

// ListProfileNotificationsPaginatorOptions is the paginator options for
// ListProfileNotifications
type ListProfileNotificationsPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfileNotificationsPaginator is a paginator for ListProfileNotifications
type ListProfileNotificationsPaginator struct {
	options   ListProfileNotificationsPaginatorOptions
	client    ListProfileNotificationsAPIClient
	params    *ListProfileNotificationsInput
	nextToken *string
	firstPage bool
}

// NewListProfileNotificationsPaginator returns a new
// ListProfileNotificationsPaginator
func NewListProfileNotificationsPaginator(client ListProfileNotificationsAPIClient, params *ListProfileNotificationsInput, optFns ...func(*ListProfileNotificationsPaginatorOptions)) *ListProfileNotificationsPaginator {
	if params == nil {
		params = &ListProfileNotificationsInput{}
	}

	options := ListProfileNotificationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfileNotificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfileNotificationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfileNotifications page.
func (p *ListProfileNotificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfileNotificationsOutput, error) {
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

	result, err := p.client.ListProfileNotifications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProfileNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProfileNotifications",
	}
}
