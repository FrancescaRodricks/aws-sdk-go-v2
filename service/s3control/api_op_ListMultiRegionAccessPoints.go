// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation is not supported by directory buckets.
//
// Returns a list of the Multi-Region Access Points currently associated with the
// specified Amazon Web Services account. Each call can return up to 100
// Multi-Region Access Points, the maximum number of Multi-Region Access Points
// that can be associated with a single account.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to ListMultiRegionAccessPoint :
//
// [CreateMultiRegionAccessPoint]
//
// [DeleteMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [GetMultiRegionAccessPoint]
//
// [DeleteMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [CreateMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func (c *Client) ListMultiRegionAccessPoints(ctx context.Context, params *ListMultiRegionAccessPointsInput, optFns ...func(*Options)) (*ListMultiRegionAccessPointsOutput, error) {
	if params == nil {
		params = &ListMultiRegionAccessPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMultiRegionAccessPoints", params, optFns, c.addOperationListMultiRegionAccessPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMultiRegionAccessPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMultiRegionAccessPointsInput struct {

	// The Amazon Web Services account ID for the owner of the Multi-Region Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// Not currently used. Do not use this parameter.
	MaxResults int32

	// Not currently used. Do not use this parameter.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListMultiRegionAccessPointsInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type ListMultiRegionAccessPointsOutput struct {

	// The list of Multi-Region Access Points associated with the user.
	AccessPoints []types.MultiRegionAccessPointReport

	// If the specified bucket has more Multi-Region Access Points than can be
	// returned in one call to this action, this field contains a continuation token.
	// You can use this token tin subsequent calls to this action to retrieve
	// additional Multi-Region Access Points.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMultiRegionAccessPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListMultiRegionAccessPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListMultiRegionAccessPoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMultiRegionAccessPoints"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListMultiRegionAccessPointsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListMultiRegionAccessPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMultiRegionAccessPoints(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListMultiRegionAccessPointsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListMultiRegionAccessPointsMiddleware struct {
}

func (*endpointPrefix_opListMultiRegionAccessPointsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListMultiRegionAccessPointsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*ListMultiRegionAccessPointsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListMultiRegionAccessPointsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListMultiRegionAccessPointsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListMultiRegionAccessPointsAPIClient is a client that implements the
// ListMultiRegionAccessPoints operation.
type ListMultiRegionAccessPointsAPIClient interface {
	ListMultiRegionAccessPoints(context.Context, *ListMultiRegionAccessPointsInput, ...func(*Options)) (*ListMultiRegionAccessPointsOutput, error)
}

var _ ListMultiRegionAccessPointsAPIClient = (*Client)(nil)

// ListMultiRegionAccessPointsPaginatorOptions is the paginator options for
// ListMultiRegionAccessPoints
type ListMultiRegionAccessPointsPaginatorOptions struct {
	// Not currently used. Do not use this parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMultiRegionAccessPointsPaginator is a paginator for
// ListMultiRegionAccessPoints
type ListMultiRegionAccessPointsPaginator struct {
	options   ListMultiRegionAccessPointsPaginatorOptions
	client    ListMultiRegionAccessPointsAPIClient
	params    *ListMultiRegionAccessPointsInput
	nextToken *string
	firstPage bool
}

// NewListMultiRegionAccessPointsPaginator returns a new
// ListMultiRegionAccessPointsPaginator
func NewListMultiRegionAccessPointsPaginator(client ListMultiRegionAccessPointsAPIClient, params *ListMultiRegionAccessPointsInput, optFns ...func(*ListMultiRegionAccessPointsPaginatorOptions)) *ListMultiRegionAccessPointsPaginator {
	if params == nil {
		params = &ListMultiRegionAccessPointsInput{}
	}

	options := ListMultiRegionAccessPointsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMultiRegionAccessPointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMultiRegionAccessPointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMultiRegionAccessPoints page.
func (p *ListMultiRegionAccessPointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMultiRegionAccessPointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListMultiRegionAccessPoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMultiRegionAccessPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMultiRegionAccessPoints",
	}
}

func copyListMultiRegionAccessPointsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListMultiRegionAccessPointsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListMultiRegionAccessPointsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *ListMultiRegionAccessPointsInput) copy() interface{} {
	v := *in
	return &v
}
func backFillListMultiRegionAccessPointsAccountID(input interface{}, v string) error {
	in := input.(*ListMultiRegionAccessPointsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListMultiRegionAccessPointsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListMultiRegionAccessPointsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
