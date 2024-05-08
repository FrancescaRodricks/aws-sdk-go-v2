// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the descriptions of all the current mount targets, or a specific mount
// target, for a file system. When requesting all of the current mount targets, the
// order of mount targets returned in the response is unspecified.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeMountTargets action, on either the file system ID that
// you specify in FileSystemId , or on the file system of the mount target that you
// specify in MountTargetId .
func (c *Client) DescribeMountTargets(ctx context.Context, params *DescribeMountTargetsInput, optFns ...func(*Options)) (*DescribeMountTargetsOutput, error) {
	if params == nil {
		params = &DescribeMountTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMountTargets", params, optFns, c.addOperationDescribeMountTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMountTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMountTargetsInput struct {

	// (Optional) The ID of the access point whose mount targets that you want to
	// list. It must be included in your request if a FileSystemId or MountTargetId is
	// not included in your request. Accepts either an access point ID or ARN as input.
	AccessPointId *string

	// (Optional) ID of the file system whose mount targets you want to list (String).
	// It must be included in your request if an AccessPointId or MountTargetId is not
	// included. Accepts either a file system ID or ARN as input.
	FileSystemId *string

	// (Optional) Opaque pagination token returned from a previous DescribeMountTargets
	// operation (String). If present, it specifies to continue the list from where the
	// previous returning call left off.
	Marker *string

	// (Optional) Maximum number of mount targets to return in the response.
	// Currently, this number is automatically set to 10, and other values are ignored.
	// The response is paginated at 100 per page if you have more than 100 mount
	// targets.
	MaxItems *int32

	// (Optional) ID of the mount target that you want to have described (String). It
	// must be included in your request if FileSystemId is not included. Accepts
	// either a mount target ID or ARN as input.
	MountTargetId *string

	noSmithyDocumentSerde
}

type DescribeMountTargetsOutput struct {

	// If the request included the Marker , the response returns that value in this
	// field.
	Marker *string

	// Returns the file system's mount targets as an array of MountTargetDescription
	// objects.
	MountTargets []types.MountTargetDescription

	// If a value is present, there are more mount targets to return. In a subsequent
	// request, you can provide Marker in your request with this value to retrieve the
	// next set of mount targets.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMountTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMountTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMountTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMountTargets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMountTargets(options.Region), middleware.Before); err != nil {
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

// DescribeMountTargetsAPIClient is a client that implements the
// DescribeMountTargets operation.
type DescribeMountTargetsAPIClient interface {
	DescribeMountTargets(context.Context, *DescribeMountTargetsInput, ...func(*Options)) (*DescribeMountTargetsOutput, error)
}

var _ DescribeMountTargetsAPIClient = (*Client)(nil)

// DescribeMountTargetsPaginatorOptions is the paginator options for
// DescribeMountTargets
type DescribeMountTargetsPaginatorOptions struct {
	// (Optional) Maximum number of mount targets to return in the response.
	// Currently, this number is automatically set to 10, and other values are ignored.
	// The response is paginated at 100 per page if you have more than 100 mount
	// targets.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMountTargetsPaginator is a paginator for DescribeMountTargets
type DescribeMountTargetsPaginator struct {
	options   DescribeMountTargetsPaginatorOptions
	client    DescribeMountTargetsAPIClient
	params    *DescribeMountTargetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMountTargetsPaginator returns a new DescribeMountTargetsPaginator
func NewDescribeMountTargetsPaginator(client DescribeMountTargetsAPIClient, params *DescribeMountTargetsInput, optFns ...func(*DescribeMountTargetsPaginatorOptions)) *DescribeMountTargetsPaginator {
	if params == nil {
		params = &DescribeMountTargetsInput{}
	}

	options := DescribeMountTargetsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMountTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMountTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMountTargets page.
func (p *DescribeMountTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMountTargetsOutput, error) {
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

	result, err := p.client.DescribeMountTargets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeMountTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMountTargets",
	}
}
