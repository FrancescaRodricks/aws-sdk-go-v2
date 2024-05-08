// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB shard group for Aurora Limitless Database. You must enable
// Aurora Limitless Database to create a DB shard group.
//
// Valid for: Aurora DB clusters only
func (c *Client) CreateDBShardGroup(ctx context.Context, params *CreateDBShardGroupInput, optFns ...func(*Options)) (*CreateDBShardGroupOutput, error) {
	if params == nil {
		params = &CreateDBShardGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBShardGroup", params, optFns, c.addOperationCreateDBShardGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBShardGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBShardGroupInput struct {

	// The name of the primary DB cluster for the DB shard group.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the DB shard group.
	//
	// This member is required.
	DBShardGroupIdentifier *string

	// The maximum capacity of the DB shard group in Aurora capacity units (ACUs).
	//
	// This member is required.
	MaxACU *float64

	// Specifies whether to create standby instances for the DB shard group. Valid
	// values are the following:
	//
	//   - 0 - Creates a single, primary DB instance for each physical shard. This is
	//   the default value, and the only one supported for the preview.
	//
	//   - 1 - Creates a primary DB instance and a standby instance in a different
	//   Availability Zone (AZ) for each physical shard.
	//
	//   - 2 - Creates a primary DB instance and two standby instances in different
	//   AZs for each physical shard.
	ComputeRedundancy *int32

	// Specifies whether the DB shard group is publicly accessible.
	//
	// When the DB shard group is publicly accessible, its Domain Name System (DNS)
	// endpoint resolves to the private IP address from within the DB shard group's
	// virtual private cloud (VPC). It resolves to the public IP address from outside
	// of the DB shard group's VPC. Access to the DB shard group is ultimately
	// controlled by the security group it uses. That public access is not permitted if
	// the security group assigned to the DB shard group doesn't permit it.
	//
	// When the DB shard group isn't publicly accessible, it is an internal DB shard
	// group with a DNS name that resolves to a private IP address.
	//
	// Default: The default behavior varies depending on whether DBSubnetGroupName is
	// specified.
	//
	// If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified,
	// the following applies:
	//
	//   - If the default VPC in the target Region doesn’t have an internet gateway
	//   attached to it, the DB shard group is private.
	//
	//   - If the default VPC in the target Region has an internet gateway attached to
	//   it, the DB shard group is public.
	//
	// If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the
	// following applies:
	//
	//   - If the subnets are part of a VPC that doesn’t have an internet gateway
	//   attached to it, the DB shard group is private.
	//
	//   - If the subnets are part of a VPC that has an internet gateway attached to
	//   it, the DB shard group is public.
	PubliclyAccessible *bool

	noSmithyDocumentSerde
}

type CreateDBShardGroupOutput struct {

	// Specifies whether to create standby instances for the DB shard group. Valid
	// values are the following:
	//
	//   - 0 - Creates a single, primary DB instance for each physical shard. This is
	//   the default value, and the only one supported for the preview.
	//
	//   - 1 - Creates a primary DB instance and a standby instance in a different
	//   Availability Zone (AZ) for each physical shard.
	//
	//   - 2 - Creates a primary DB instance and two standby instances in different
	//   AZs for each physical shard.
	ComputeRedundancy *int32

	// The name of the primary DB cluster for the DB shard group.
	DBClusterIdentifier *string

	// The name of the DB shard group.
	DBShardGroupIdentifier *string

	// The Amazon Web Services Region-unique, immutable identifier for the DB shard
	// group.
	DBShardGroupResourceId *string

	// The connection endpoint for the DB shard group.
	Endpoint *string

	// The maximum capacity of the DB shard group in Aurora capacity units (ACUs).
	MaxACU *float64

	// Indicates whether the DB shard group is publicly accessible.
	//
	// When the DB shard group is publicly accessible, its Domain Name System (DNS)
	// endpoint resolves to the private IP address from within the DB shard group's
	// virtual private cloud (VPC). It resolves to the public IP address from outside
	// of the DB shard group's VPC. Access to the DB shard group is ultimately
	// controlled by the security group it uses. That public access isn't permitted if
	// the security group assigned to the DB shard group doesn't permit it.
	//
	// When the DB shard group isn't publicly accessible, it is an internal DB shard
	// group with a DNS name that resolves to a private IP address.
	//
	// For more information, see CreateDBShardGroup.
	//
	// This setting is only for Aurora Limitless Database.
	PubliclyAccessible *bool

	// The status of the DB shard group.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBShardGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBShardGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBShardGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDBShardGroup"); err != nil {
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
	if err = addOpCreateDBShardGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBShardGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBShardGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDBShardGroup",
	}
}
