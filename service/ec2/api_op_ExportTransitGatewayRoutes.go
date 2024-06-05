// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports routes from the specified transit gateway route table to the specified
// S3 bucket. By default, all routes are exported. Alternatively, you can filter by
// CIDR range.
//
// The routes are saved to the specified bucket in a JSON file. For more
// information, see [Export Route Tables to Amazon S3]in Transit Gateways.
//
// [Export Route Tables to Amazon S3]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-route-tables.html#tgw-export-route-tables
func (c *Client) ExportTransitGatewayRoutes(ctx context.Context, params *ExportTransitGatewayRoutesInput, optFns ...func(*Options)) (*ExportTransitGatewayRoutesOutput, error) {
	if params == nil {
		params = &ExportTransitGatewayRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportTransitGatewayRoutes", params, optFns, c.addOperationExportTransitGatewayRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportTransitGatewayRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportTransitGatewayRoutesInput struct {

	// The name of the S3 bucket.
	//
	// This member is required.
	S3Bucket *string

	// The ID of the route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//   - attachment.transit-gateway-attachment-id - The id of the transit gateway
	//   attachment.
	//
	//   - attachment.resource-id - The resource id of the transit gateway attachment.
	//
	//   - route-search.exact-match - The exact match of the specified filter.
	//
	//   - route-search.longest-prefix-match - The longest prefix that matches the
	//   route.
	//
	//   - route-search.subnet-of-match - The routes with a subnet that match the
	//   specified CIDR filter.
	//
	//   - route-search.supernet-of-match - The routes with a CIDR that encompass the
	//   CIDR filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31 routes in your
	//   route table and you specify supernet-of-match as 10.0.1.0/30, then the result
	//   returns 10.0.1.0/29.
	//
	//   - state - The state of the route ( active | blackhole ).
	//
	//   - transit-gateway-route-destination-cidr-block - The CIDR range.
	//
	//   - type - The type of route ( propagated | static ).
	Filters []types.Filter

	noSmithyDocumentSerde
}

type ExportTransitGatewayRoutesOutput struct {

	// The URL of the exported file in Amazon S3. For example,
	// s3://bucket_name/VPCTransitGateway/TransitGatewayRouteTables/file_name.
	S3Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportTransitGatewayRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpExportTransitGatewayRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpExportTransitGatewayRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportTransitGatewayRoutes"); err != nil {
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
	if err = addOpExportTransitGatewayRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportTransitGatewayRoutes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportTransitGatewayRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportTransitGatewayRoutes",
	}
}
