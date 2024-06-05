// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon QuickSight groups that an Amazon QuickSight user is a member
// of.
func (c *Client) ListUserGroups(ctx context.Context, params *ListUserGroupsInput, optFns ...func(*Options)) (*ListUserGroupsOutput, error) {
	if params == nil {
		params = &ListUserGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserGroups", params, optFns, c.addOperationListUserGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserGroupsInput struct {

	// The Amazon Web Services account ID that the user is in. Currently, you use the
	// ID for the Amazon Web Services account that contains your Amazon QuickSight
	// account.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace. Currently, you should set this to default .
	//
	// This member is required.
	Namespace *string

	// The Amazon QuickSight user name that you want to list group memberships for.
	//
	// This member is required.
	UserName *string

	// The maximum number of results to return from this request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUserGroupsOutput struct {

	// The list of groups the user is a member of.
	GroupList []types.Group

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUserGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUserGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUserGroups"); err != nil {
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
	if err = addOpListUserGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserGroups(options.Region), middleware.Before); err != nil {
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

// ListUserGroupsAPIClient is a client that implements the ListUserGroups
// operation.
type ListUserGroupsAPIClient interface {
	ListUserGroups(context.Context, *ListUserGroupsInput, ...func(*Options)) (*ListUserGroupsOutput, error)
}

var _ ListUserGroupsAPIClient = (*Client)(nil)

// ListUserGroupsPaginatorOptions is the paginator options for ListUserGroups
type ListUserGroupsPaginatorOptions struct {
	// The maximum number of results to return from this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserGroupsPaginator is a paginator for ListUserGroups
type ListUserGroupsPaginator struct {
	options   ListUserGroupsPaginatorOptions
	client    ListUserGroupsAPIClient
	params    *ListUserGroupsInput
	nextToken *string
	firstPage bool
}

// NewListUserGroupsPaginator returns a new ListUserGroupsPaginator
func NewListUserGroupsPaginator(client ListUserGroupsAPIClient, params *ListUserGroupsInput, optFns ...func(*ListUserGroupsPaginatorOptions)) *ListUserGroupsPaginator {
	if params == nil {
		params = &ListUserGroupsInput{}
	}

	options := ListUserGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUserGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUserGroups page.
func (p *ListUserGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserGroupsOutput, error) {
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

	result, err := p.client.ListUserGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUserGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUserGroups",
	}
}
