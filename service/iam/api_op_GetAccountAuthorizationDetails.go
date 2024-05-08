// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about all IAM users, groups, roles, and policies in your
// Amazon Web Services account, including their relationships to one another. Use
// this operation to obtain a snapshot of the configuration of IAM permissions
// (users, groups, roles, and policies) in your account.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality.
//
// You can optionally filter the results using the Filter parameter. You can
// paginate the results using the MaxItems and Marker parameters.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
func (c *Client) GetAccountAuthorizationDetails(ctx context.Context, params *GetAccountAuthorizationDetailsInput, optFns ...func(*Options)) (*GetAccountAuthorizationDetailsOutput, error) {
	if params == nil {
		params = &GetAccountAuthorizationDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccountAuthorizationDetails", params, optFns, c.addOperationGetAccountAuthorizationDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccountAuthorizationDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccountAuthorizationDetailsInput struct {

	// A list of entity types used to filter the results. Only the entities that match
	// the types you specify are included in the output. Use the value
	// LocalManagedPolicy to include customer managed policies.
	//
	// The format for this parameter is a comma-separated (if more than one) list of
	// strings. Each string value in the list must be one of the valid values listed
	// below.
	Filter []types.EntityType

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

// Contains the response to a successful GetAccountAuthorizationDetails request.
type GetAccountAuthorizationDetailsOutput struct {

	// A list containing information about IAM groups.
	GroupDetailList []types.GroupDetail

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

	// A list containing information about managed policies.
	Policies []types.ManagedPolicyDetail

	// A list containing information about IAM roles.
	RoleDetailList []types.RoleDetail

	// A list containing information about IAM users.
	UserDetailList []types.UserDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccountAuthorizationDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetAccountAuthorizationDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetAccountAuthorizationDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccountAuthorizationDetails"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccountAuthorizationDetails(options.Region), middleware.Before); err != nil {
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

// GetAccountAuthorizationDetailsAPIClient is a client that implements the
// GetAccountAuthorizationDetails operation.
type GetAccountAuthorizationDetailsAPIClient interface {
	GetAccountAuthorizationDetails(context.Context, *GetAccountAuthorizationDetailsInput, ...func(*Options)) (*GetAccountAuthorizationDetailsOutput, error)
}

var _ GetAccountAuthorizationDetailsAPIClient = (*Client)(nil)

// GetAccountAuthorizationDetailsPaginatorOptions is the paginator options for
// GetAccountAuthorizationDetails
type GetAccountAuthorizationDetailsPaginatorOptions struct {
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

// GetAccountAuthorizationDetailsPaginator is a paginator for
// GetAccountAuthorizationDetails
type GetAccountAuthorizationDetailsPaginator struct {
	options   GetAccountAuthorizationDetailsPaginatorOptions
	client    GetAccountAuthorizationDetailsAPIClient
	params    *GetAccountAuthorizationDetailsInput
	nextToken *string
	firstPage bool
}

// NewGetAccountAuthorizationDetailsPaginator returns a new
// GetAccountAuthorizationDetailsPaginator
func NewGetAccountAuthorizationDetailsPaginator(client GetAccountAuthorizationDetailsAPIClient, params *GetAccountAuthorizationDetailsInput, optFns ...func(*GetAccountAuthorizationDetailsPaginatorOptions)) *GetAccountAuthorizationDetailsPaginator {
	if params == nil {
		params = &GetAccountAuthorizationDetailsInput{}
	}

	options := GetAccountAuthorizationDetailsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAccountAuthorizationDetailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAccountAuthorizationDetailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAccountAuthorizationDetails page.
func (p *GetAccountAuthorizationDetailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAccountAuthorizationDetailsOutput, error) {
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

	result, err := p.client.GetAccountAuthorizationDetails(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetAccountAuthorizationDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccountAuthorizationDetails",
	}
}
