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

// Deletes a tenant database from your DB instance. This command only applies to
// RDS for Oracle container database (CDB) instances.
//
// You can't delete a tenant database when it is the only tenant in the DB
// instance.
func (c *Client) DeleteTenantDatabase(ctx context.Context, params *DeleteTenantDatabaseInput, optFns ...func(*Options)) (*DeleteTenantDatabaseOutput, error) {
	if params == nil {
		params = &DeleteTenantDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTenantDatabase", params, optFns, c.addOperationDeleteTenantDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTenantDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTenantDatabaseInput struct {

	// The user-supplied identifier for the DB instance that contains the tenant
	// database that you want to delete.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The user-supplied name of the tenant database that you want to remove from your
	// DB instance. Amazon RDS deletes the tenant database with this name. This
	// parameter isn’t case-sensitive.
	//
	// This member is required.
	TenantDBName *string

	// The DBSnapshotIdentifier of the new DBSnapshot created when the
	// SkipFinalSnapshot parameter is disabled.
	//
	// If you enable this parameter and also enable SkipFinalShapshot , the command
	// results in an error.
	FinalDBSnapshotIdentifier *string

	// Specifies whether to skip the creation of a final DB snapshot before removing
	// the tenant database from your DB instance. If you enable this parameter, RDS
	// doesn't create a DB snapshot. If you don't enable this parameter, RDS creates a
	// DB snapshot before it deletes the tenant database. By default, RDS doesn't skip
	// the final snapshot. If you don't enable this parameter, you must specify the
	// FinalDBSnapshotIdentifier parameter.
	SkipFinalSnapshot *bool

	noSmithyDocumentSerde
}

type DeleteTenantDatabaseOutput struct {

	// A tenant database in the DB instance. This data type is an element in the
	// response to the DescribeTenantDatabases action.
	TenantDatabase *types.TenantDatabase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTenantDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteTenantDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteTenantDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteTenantDatabase"); err != nil {
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
	if err = addOpDeleteTenantDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTenantDatabase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTenantDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteTenantDatabase",
	}
}
