// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an Amazon RDS DB instance that was stopped using the Amazon Web Services
// console, the stop-db-instance CLI command, or the StopDBInstance operation.
//
// For more information, see [Starting an Amazon RDS DB instance That Was Previously Stopped] in the Amazon RDS User Guide.
//
// This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.
// For Aurora DB clusters, use StartDBCluster instead.
//
// [Starting an Amazon RDS DB instance That Was Previously Stopped]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StartInstance.html
func (c *Client) StartDBInstance(ctx context.Context, params *StartDBInstanceInput, optFns ...func(*Options)) (*StartDBInstanceOutput, error) {
	if params == nil {
		params = &StartDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDBInstance", params, optFns, c.addOperationStartDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDBInstanceInput struct {

	// The user-supplied instance identifier.
	//
	// This member is required.
	DBInstanceIdentifier *string

	noSmithyDocumentSerde
}

type StartDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the operations CreateDBInstance
	// , CreateDBInstanceReadReplica , DeleteDBInstance , DescribeDBInstances ,
	// ModifyDBInstance , PromoteReadReplica , RebootDBInstance ,
	// RestoreDBInstanceFromDBSnapshot , RestoreDBInstanceFromS3 ,
	// RestoreDBInstanceToPointInTime , StartDBInstance , and StopDBInstance .
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartDBInstance"); err != nil {
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
	if err = addOpStartDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDBInstance",
	}
}
