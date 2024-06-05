// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about device pools.
func (c *Client) ListDevicePools(ctx context.Context, params *ListDevicePoolsInput, optFns ...func(*Options)) (*ListDevicePoolsOutput, error) {
	if params == nil {
		params = &ListDevicePoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevicePools", params, optFns, c.addOperationListDevicePoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevicePoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the result of a list device pools request.
type ListDevicePoolsInput struct {

	// The project ARN.
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The device pools' type.
	//
	// Allowed values include:
	//
	//   - CURATED: A device pool that is created and managed by AWS Device Farm.
	//
	//   - PRIVATE: A device pool that is created and managed by the device pool
	//   developer.
	Type types.DevicePoolType

	noSmithyDocumentSerde
}

// Represents the result of a list device pools request.
type ListDevicePoolsOutput struct {

	// Information about the device pools.
	DevicePools []types.DevicePool

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevicePoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDevicePools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDevicePools{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDevicePools"); err != nil {
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
	if err = addOpListDevicePoolsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevicePools(options.Region), middleware.Before); err != nil {
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

// ListDevicePoolsAPIClient is a client that implements the ListDevicePools
// operation.
type ListDevicePoolsAPIClient interface {
	ListDevicePools(context.Context, *ListDevicePoolsInput, ...func(*Options)) (*ListDevicePoolsOutput, error)
}

var _ ListDevicePoolsAPIClient = (*Client)(nil)

// ListDevicePoolsPaginatorOptions is the paginator options for ListDevicePools
type ListDevicePoolsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevicePoolsPaginator is a paginator for ListDevicePools
type ListDevicePoolsPaginator struct {
	options   ListDevicePoolsPaginatorOptions
	client    ListDevicePoolsAPIClient
	params    *ListDevicePoolsInput
	nextToken *string
	firstPage bool
}

// NewListDevicePoolsPaginator returns a new ListDevicePoolsPaginator
func NewListDevicePoolsPaginator(client ListDevicePoolsAPIClient, params *ListDevicePoolsInput, optFns ...func(*ListDevicePoolsPaginatorOptions)) *ListDevicePoolsPaginator {
	if params == nil {
		params = &ListDevicePoolsInput{}
	}

	options := ListDevicePoolsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevicePoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevicePoolsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDevicePools page.
func (p *ListDevicePoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevicePoolsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListDevicePools(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDevicePools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDevicePools",
	}
}
