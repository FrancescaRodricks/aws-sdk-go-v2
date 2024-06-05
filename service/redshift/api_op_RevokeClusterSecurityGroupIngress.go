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

// Revokes an ingress rule in an Amazon Redshift security group for a previously
// authorized IP range or Amazon EC2 security group. To add an ingress rule, see AuthorizeClusterSecurityGroupIngress.
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups]in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Cluster Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
func (c *Client) RevokeClusterSecurityGroupIngress(ctx context.Context, params *RevokeClusterSecurityGroupIngressInput, optFns ...func(*Options)) (*RevokeClusterSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &RevokeClusterSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeClusterSecurityGroupIngress", params, optFns, c.addOperationRevokeClusterSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeClusterSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeClusterSecurityGroupIngressInput struct {

	// The name of the security Group from which to revoke the ingress rule.
	//
	// This member is required.
	ClusterSecurityGroupName *string

	// The IP range for which to revoke access. This range must be a valid Classless
	// Inter-Domain Routing (CIDR) block of IP addresses. If CIDRIP is specified,
	// EC2SecurityGroupName and EC2SecurityGroupOwnerId cannot be provided.
	CIDRIP *string

	// The name of the EC2 Security Group whose access is to be revoked. If
	// EC2SecurityGroupName is specified, EC2SecurityGroupOwnerId must also be
	// provided and CIDRIP cannot be provided.
	EC2SecurityGroupName *string

	// The Amazon Web Services account number of the owner of the security group
	// specified in the EC2SecurityGroupName parameter. The Amazon Web Services access
	// key ID is not an acceptable value. If EC2SecurityGroupOwnerId is specified,
	// EC2SecurityGroupName must also be provided. and CIDRIP cannot be provided.
	//
	// Example: 111122223333
	EC2SecurityGroupOwnerId *string

	noSmithyDocumentSerde
}

type RevokeClusterSecurityGroupIngressOutput struct {

	// Describes a security group.
	ClusterSecurityGroup *types.ClusterSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeClusterSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRevokeClusterSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeClusterSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeClusterSecurityGroupIngress"); err != nil {
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
	if err = addOpRevokeClusterSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeClusterSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeClusterSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeClusterSecurityGroupIngress",
	}
}
