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

// Lists member users in a group.
func (c *Client) ListGroupMemberships(ctx context.Context, params *ListGroupMembershipsInput, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) {
	if params == nil {
		params = &ListGroupMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupMemberships", params, optFns, c.addOperationListGroupMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupMembershipsInput struct {

	// The ID for the Amazon Web Services account that the group is in. Currently, you
	// use the ID for the Amazon Web Services account that contains your Amazon
	// QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The name of the group that you want to see a membership list of.
	//
	// This member is required.
	GroupName *string

	// The namespace of the group that you want a list of users from.
	//
	// This member is required.
	Namespace *string

	// The maximum number of results to return from this request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupMembershipsOutput struct {

	// The list of the members of the group.
	GroupMemberList []types.GroupMember

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

func (c *Client) addOperationListGroupMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGroupMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroupMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroupMemberships"); err != nil {
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
	if err = addOpListGroupMembershipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupMemberships(options.Region), middleware.Before); err != nil {
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

// ListGroupMembershipsAPIClient is a client that implements the
// ListGroupMemberships operation.
type ListGroupMembershipsAPIClient interface {
	ListGroupMemberships(context.Context, *ListGroupMembershipsInput, ...func(*Options)) (*ListGroupMembershipsOutput, error)
}

var _ ListGroupMembershipsAPIClient = (*Client)(nil)

// ListGroupMembershipsPaginatorOptions is the paginator options for
// ListGroupMemberships
type ListGroupMembershipsPaginatorOptions struct {
	// The maximum number of results to return from this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupMembershipsPaginator is a paginator for ListGroupMemberships
type ListGroupMembershipsPaginator struct {
	options   ListGroupMembershipsPaginatorOptions
	client    ListGroupMembershipsAPIClient
	params    *ListGroupMembershipsInput
	nextToken *string
	firstPage bool
}

// NewListGroupMembershipsPaginator returns a new ListGroupMembershipsPaginator
func NewListGroupMembershipsPaginator(client ListGroupMembershipsAPIClient, params *ListGroupMembershipsInput, optFns ...func(*ListGroupMembershipsPaginatorOptions)) *ListGroupMembershipsPaginator {
	if params == nil {
		params = &ListGroupMembershipsInput{}
	}

	options := ListGroupMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupMemberships page.
func (p *ListGroupMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) {
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

	result, err := p.client.ListGroupMemberships(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroupMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroupMemberships",
	}
}
