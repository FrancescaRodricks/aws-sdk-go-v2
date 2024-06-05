// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the list of Identity and Access Management (IAM) roles that can be
// used by the cluster to access other Amazon Web Services services.
//
// The maximum number of IAM roles that you can associate is subject to a quota.
// For more information, go to [Quotas and limits]in the Amazon Redshift Cluster Management Guide.
//
// [Quotas and limits]: https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html
func (c *Client) ModifyClusterIamRoles(ctx context.Context, params *ModifyClusterIamRolesInput, optFns ...func(*Options)) (*ModifyClusterIamRolesOutput, error) {
	if params == nil {
		params = &ModifyClusterIamRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterIamRoles", params, optFns, c.addOperationModifyClusterIamRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterIamRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterIamRolesInput struct {

	// The unique identifier of the cluster for which you want to associate or
	// disassociate IAM roles.
	//
	// This member is required.
	ClusterIdentifier *string

	// Zero or more IAM roles to associate with the cluster. The roles must be in
	// their Amazon Resource Name (ARN) format.
	AddIamRoles []string

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the
	// cluster when the cluster was last modified.
	DefaultIamRoleArn *string

	// Zero or more IAM roles in ARN format to disassociate from the cluster.
	RemoveIamRoles []string

	noSmithyDocumentSerde
}

type ModifyClusterIamRolesOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClusterIamRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterIamRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterIamRoles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyClusterIamRoles"); err != nil {
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
	if err = addOpModifyClusterIamRolesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterIamRoles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClusterIamRoles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyClusterIamRoles",
	}
}
