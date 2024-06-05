// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the notification rules for an Amazon Web Services account.
func (c *Client) ListNotificationRules(ctx context.Context, params *ListNotificationRulesInput, optFns ...func(*Options)) (*ListNotificationRulesOutput, error) {
	if params == nil {
		params = &ListNotificationRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotificationRules", params, optFns, c.addOperationListNotificationRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotificationRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotificationRulesInput struct {

	// The filters to use to return information by service or resource type. For valid
	// values, see ListNotificationRulesFilter.
	//
	// A filter with the same name can appear more than once when used with OR
	// statements. Filters with different names should be applied with AND statements.
	Filters []types.ListNotificationRulesFilter

	// A non-negative integer used to limit the number of returned results. The
	// maximum number of results that can be returned is 100.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNotificationRulesOutput struct {

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// The list of notification rules for the Amazon Web Services account, by Amazon
	// Resource Name (ARN) and ID.
	NotificationRules []types.NotificationRuleSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotificationRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNotificationRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNotificationRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNotificationRules"); err != nil {
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
	if err = addOpListNotificationRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotificationRules(options.Region), middleware.Before); err != nil {
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

// ListNotificationRulesAPIClient is a client that implements the
// ListNotificationRules operation.
type ListNotificationRulesAPIClient interface {
	ListNotificationRules(context.Context, *ListNotificationRulesInput, ...func(*Options)) (*ListNotificationRulesOutput, error)
}

var _ ListNotificationRulesAPIClient = (*Client)(nil)

// ListNotificationRulesPaginatorOptions is the paginator options for
// ListNotificationRules
type ListNotificationRulesPaginatorOptions struct {
	// A non-negative integer used to limit the number of returned results. The
	// maximum number of results that can be returned is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotificationRulesPaginator is a paginator for ListNotificationRules
type ListNotificationRulesPaginator struct {
	options   ListNotificationRulesPaginatorOptions
	client    ListNotificationRulesAPIClient
	params    *ListNotificationRulesInput
	nextToken *string
	firstPage bool
}

// NewListNotificationRulesPaginator returns a new ListNotificationRulesPaginator
func NewListNotificationRulesPaginator(client ListNotificationRulesAPIClient, params *ListNotificationRulesInput, optFns ...func(*ListNotificationRulesPaginatorOptions)) *ListNotificationRulesPaginator {
	if params == nil {
		params = &ListNotificationRulesInput{}
	}

	options := ListNotificationRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotificationRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotificationRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotificationRules page.
func (p *ListNotificationRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotificationRulesOutput, error) {
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

	result, err := p.client.ListNotificationRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNotificationRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNotificationRules",
	}
}
