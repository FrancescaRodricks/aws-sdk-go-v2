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

// You might need to reboot your DB cluster, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB cluster
// parameter group associated with the DB cluster, reboot the DB cluster for the
// changes to take effect.
//
// Rebooting a DB cluster restarts the database engine service. Rebooting a DB
// cluster results in a momentary outage, during which the DB cluster status is set
// to rebooting.
//
// Use this operation only for a non-Aurora Multi-AZ DB cluster.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func (c *Client) RebootDBCluster(ctx context.Context, params *RebootDBClusterInput, optFns ...func(*Options)) (*RebootDBClusterOutput, error) {
	if params == nil {
		params = &RebootDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootDBCluster", params, optFns, c.addOperationRebootDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootDBClusterInput struct {

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing DBCluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	noSmithyDocumentSerde
}

type RebootDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster.
	//
	// For an Amazon Aurora DB cluster, this data type is used as a response element
	// in the operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , PromoteReadReplicaDBCluster ,
	// RestoreDBClusterFromS3 , RestoreDBClusterFromSnapshot ,
	// RestoreDBClusterToPointInTime , StartDBCluster , and StopDBCluster .
	//
	// For a Multi-AZ DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , RebootDBCluster ,
	// RestoreDBClusterFromSnapshot , and RestoreDBClusterToPointInTime .
	//
	// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
	// User Guide.
	//
	// For more information on Multi-AZ DB clusters, see [Multi-AZ deployments with two readable standby DB instances] in the Amazon RDS User
	// Guide.
	//
	// [Multi-AZ deployments with two readable standby DB instances]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
	// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRebootDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebootDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RebootDBCluster"); err != nil {
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
	if err = addOpRebootDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebootDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebootDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RebootDBCluster",
	}
}
