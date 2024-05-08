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

// Lists the mailbox permissions associated with a user, group, or resource
// mailbox.
func (c *Client) ListMailboxPermissions(ctx context.Context, params *ListMailboxPermissionsInput, optFns ...func(*Options)) (*ListMailboxPermissionsOutput, error) {
	if params == nil {
		params = &ListMailboxPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMailboxPermissions", params, optFns, c.addOperationListMailboxPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMailboxPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMailboxPermissionsInput struct {

	// The identifier of the user, or resource for which to list mailbox permissions.
	//
	// The entity ID can accept UserId or ResourceId, Username or Resourcename, or
	// email.
	//
	//   - Entity ID: 12345678-1234-1234-1234-123456789012, or
	//   r-0123456789a0123456789b0123456789
	//
	//   - Email address: entity@domain.tld
	//
	//   - Entity name: entity
	//
	// This member is required.
	EntityId *string

	// The identifier of the organization under which the user, group, or resource
	// exists.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMailboxPermissionsOutput struct {

	// The token to use to retrieve the next page of results. The value is "null" when
	// there are no more results to return.
	NextToken *string

	// One page of the user, group, or resource mailbox permissions.
	Permissions []types.Permission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMailboxPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMailboxPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMailboxPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMailboxPermissions"); err != nil {
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
	if err = addOpListMailboxPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMailboxPermissions(options.Region), middleware.Before); err != nil {
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

// ListMailboxPermissionsAPIClient is a client that implements the
// ListMailboxPermissions operation.
type ListMailboxPermissionsAPIClient interface {
	ListMailboxPermissions(context.Context, *ListMailboxPermissionsInput, ...func(*Options)) (*ListMailboxPermissionsOutput, error)
}

var _ ListMailboxPermissionsAPIClient = (*Client)(nil)

// ListMailboxPermissionsPaginatorOptions is the paginator options for
// ListMailboxPermissions
type ListMailboxPermissionsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMailboxPermissionsPaginator is a paginator for ListMailboxPermissions
type ListMailboxPermissionsPaginator struct {
	options   ListMailboxPermissionsPaginatorOptions
	client    ListMailboxPermissionsAPIClient
	params    *ListMailboxPermissionsInput
	nextToken *string
	firstPage bool
}

// NewListMailboxPermissionsPaginator returns a new ListMailboxPermissionsPaginator
func NewListMailboxPermissionsPaginator(client ListMailboxPermissionsAPIClient, params *ListMailboxPermissionsInput, optFns ...func(*ListMailboxPermissionsPaginatorOptions)) *ListMailboxPermissionsPaginator {
	if params == nil {
		params = &ListMailboxPermissionsInput{}
	}

	options := ListMailboxPermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMailboxPermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMailboxPermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMailboxPermissions page.
func (p *ListMailboxPermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMailboxPermissionsOutput, error) {
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

	result, err := p.client.ListMailboxPermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMailboxPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMailboxPermissions",
	}
}
