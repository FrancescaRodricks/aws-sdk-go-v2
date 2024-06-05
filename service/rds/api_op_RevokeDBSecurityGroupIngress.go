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

// Revokes ingress from a DBSecurityGroup for previously authorized IP ranges or
// EC2 or VPC security groups. Required parameters for this API are one of CIDRIP,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId).
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func (c *Client) RevokeDBSecurityGroupIngress(ctx context.Context, params *RevokeDBSecurityGroupIngressInput, optFns ...func(*Options)) (*RevokeDBSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &RevokeDBSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeDBSecurityGroupIngress", params, optFns, c.addOperationRevokeDBSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeDBSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeDBSecurityGroupIngressInput struct {

	// The name of the DB security group to revoke ingress from.
	//
	// This member is required.
	DBSecurityGroupName *string

	// The IP range to revoke access from. Must be a valid CIDR range. If CIDRIP is
	// specified, EC2SecurityGroupName , EC2SecurityGroupId and EC2SecurityGroupOwnerId
	// can't be provided.
	CIDRIP *string

	// The id of the EC2 security group to revoke access from. For VPC DB security
	// groups, EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId
	// and either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupId *string

	// The name of the EC2 security group to revoke access from. For VPC DB security
	// groups, EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId
	// and either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupName *string

	// The Amazon Web Services account number of the owner of the EC2 security group
	// specified in the EC2SecurityGroupName parameter. The Amazon Web Services access
	// key ID isn't an acceptable value. For VPC DB security groups, EC2SecurityGroupId
	// must be provided. Otherwise, EC2SecurityGroupOwnerId and either
	// EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupOwnerId *string

	noSmithyDocumentSerde
}

type RevokeDBSecurityGroupIngressOutput struct {

	// Contains the details for an Amazon RDS DB security group.
	//
	// This data type is used as a response element in the DescribeDBSecurityGroups
	// action.
	DBSecurityGroup *types.DBSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeDBSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRevokeDBSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeDBSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeDBSecurityGroupIngress"); err != nil {
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
	if err = addOpRevokeDBSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeDBSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeDBSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeDBSecurityGroupIngress",
	}
}
