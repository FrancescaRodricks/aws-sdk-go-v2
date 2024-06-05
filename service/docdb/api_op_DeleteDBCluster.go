// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a previously provisioned cluster. When you delete a cluster, all
// automated backups for that cluster are deleted and can't be recovered. Manual DB
// cluster snapshots of the specified cluster are not deleted.
func (c *Client) DeleteDBCluster(ctx context.Context, params *DeleteDBClusterInput, optFns ...func(*Options)) (*DeleteDBClusterOutput, error) {
	if params == nil {
		params = &DeleteDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBCluster", params, optFns, c.addOperationDeleteDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DeleteDBCluster.
type DeleteDBClusterInput struct {

	// The cluster identifier for the cluster to be deleted. This parameter isn't case
	// sensitive.
	//
	// Constraints:
	//
	//   - Must match an existing DBClusterIdentifier .
	//
	// This member is required.
	DBClusterIdentifier *string

	//  The cluster snapshot identifier of the new cluster snapshot created when
	// SkipFinalSnapshot is set to false .
	//
	// Specifying this parameter and also setting the SkipFinalShapshot parameter to
	// true results in an error.
	//
	// Constraints:
	//
	//   - Must be from 1 to 255 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	FinalDBSnapshotIdentifier *string

	//  Determines whether a final cluster snapshot is created before the cluster is
	// deleted. If true is specified, no cluster snapshot is created. If false is
	// specified, a cluster snapshot is created before the DB cluster is deleted.
	//
	// If SkipFinalSnapshot is false , you must specify a FinalDBSnapshotIdentifier
	// parameter.
	//
	// Default: false
	SkipFinalSnapshot *bool

	noSmithyDocumentSerde
}

type DeleteDBClusterOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDBCluster"); err != nil {
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
	if err = addOpDeleteDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDBCluster",
	}
}
