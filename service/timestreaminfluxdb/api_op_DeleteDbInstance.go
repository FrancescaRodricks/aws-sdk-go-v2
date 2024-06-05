// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreaminfluxdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Timestream for InfluxDB DB instance.
func (c *Client) DeleteDbInstance(ctx context.Context, params *DeleteDbInstanceInput, optFns ...func(*Options)) (*DeleteDbInstanceOutput, error) {
	if params == nil {
		params = &DeleteDbInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDbInstance", params, optFns, c.addOperationDeleteDbInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDbInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDbInstanceInput struct {

	// The id of the DB instance.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type DeleteDbInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the DB instance.
	//
	// This member is required.
	Arn *string

	// A service-generated unique identifier.
	//
	// This member is required.
	Id *string

	// The customer-supplied name that uniquely identifies the DB instance when
	// interacting with the Amazon Timestream for InfluxDB API and CLI commands.
	//
	// This member is required.
	Name *string

	// A list of VPC subnet IDs associated with the DB instance.
	//
	// This member is required.
	VpcSubnetIds []string

	// The amount of storage allocated for your DB storage type (in gibibytes).
	AllocatedStorage *int32

	// The Availability Zone in which the DB instance resides.
	AvailabilityZone *string

	// The Timestream for InfluxDB instance type that InfluxDB runs on.
	DbInstanceType types.DbInstanceType

	// The id of the DB parameter group assigned to your DB instance.
	DbParameterGroupIdentifier *string

	// The Timestream for InfluxDB DB storage type that InfluxDB stores data on.
	DbStorageType types.DbStorageType

	// Specifies whether the Timestream for InfluxDB is deployed as Single-AZ or with
	// a MultiAZ Standby for High availability.
	DeploymentType types.DeploymentType

	// The endpoint used to connect to InfluxDB. The default InfluxDB port is 8086.
	Endpoint *string

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the
	// initial InfluxDB authorization parameters. The secret value is a JSON formatted
	// key-value pair holding InfluxDB authorization values: organization, bucket,
	// username, and password.
	InfluxAuthParametersSecretArn *string

	// Configuration for sending InfluxDB engine logs to send to specified S3 bucket.
	LogDeliveryConfiguration *types.LogDeliveryConfiguration

	// Indicates if the DB instance has a public IP to facilitate access.
	PubliclyAccessible *bool

	// The Availability Zone in which the standby instance is located when deploying
	// with a MultiAZ standby instance.
	SecondaryAvailabilityZone *string

	// The status of the DB instance.
	Status types.Status

	// A list of VPC security group IDs associated with the DB instance.
	VpcSecurityGroupIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDbInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteDbInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteDbInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDbInstance"); err != nil {
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
	if err = addOpDeleteDbInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDbInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDbInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDbInstance",
	}
}
