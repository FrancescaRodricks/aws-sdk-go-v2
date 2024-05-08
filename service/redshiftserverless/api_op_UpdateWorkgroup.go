// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a workgroup with the specified configuration settings. You can't update
// multiple parameters in one request. For example, you can update baseCapacity or
// port in a single request, but you can't update both in the same request.
func (c *Client) UpdateWorkgroup(ctx context.Context, params *UpdateWorkgroupInput, optFns ...func(*Options)) (*UpdateWorkgroupOutput, error) {
	if params == nil {
		params = &UpdateWorkgroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkgroup", params, optFns, c.addOperationUpdateWorkgroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkgroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkgroupInput struct {

	// The name of the workgroup to update. You can't update the name of a workgroup
	// once it is created.
	//
	// This member is required.
	WorkgroupName *string

	// The new base data warehouse capacity in Redshift Processing Units (RPUs).
	BaseCapacity *int32

	// An array of parameters to set for advanced control over a database. The options
	// are auto_mv , datestyle , enable_case_sensitive_identifier ,
	// enable_user_activity_logging , query_group , search_path , require_ssl ,
	// use_fips_ssl , and query monitoring metrics that let you define performance
	// boundaries. For more information about query monitoring rules and available
	// metrics, see [Query monitoring metrics for Amazon Redshift Serverless].
	//
	// [Query monitoring metrics for Amazon Redshift Serverless]: https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless
	ConfigParameters []types.ConfigParameter

	// The value that specifies whether to turn on enhanced virtual private cloud
	// (VPC) routing, which forces Amazon Redshift Serverless to route traffic through
	// your VPC.
	EnhancedVpcRouting *bool

	// The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve
	// queries. The max capacity is specified in RPUs.
	MaxCapacity *int32

	// The custom port to use when connecting to a workgroup. Valid port ranges are
	// 5431-5455 and 8191-8215. The default is 5439.
	Port *int32

	// A value that specifies whether the workgroup can be accessible from a public
	// network.
	PubliclyAccessible *bool

	// An array of security group IDs to associate with the workgroup.
	SecurityGroupIds []string

	// An array of VPC subnet IDs to associate with the workgroup.
	SubnetIds []string

	noSmithyDocumentSerde
}

type UpdateWorkgroupOutput struct {

	// The updated workgroup object.
	//
	// This member is required.
	Workgroup *types.Workgroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkgroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkgroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkgroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateWorkgroup"); err != nil {
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
	if err = addOpUpdateWorkgroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkgroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkgroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateWorkgroup",
	}
}
