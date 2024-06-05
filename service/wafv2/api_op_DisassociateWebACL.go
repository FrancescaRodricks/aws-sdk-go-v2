// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the specified regional application resource from any existing web
// ACL association. A resource can have at most one web ACL association. A regional
// application can be an Application Load Balancer (ALB), an Amazon API Gateway
// REST API, an AppSync GraphQL API, an Amazon Cognito user pool, an App Runner
// service, or an Amazon Web Services Verified Access instance.
//
// For Amazon CloudFront, don't use this call. Instead, use your CloudFront
// distribution configuration. To disassociate a web ACL, provide an empty web ACL
// ID in the CloudFront call UpdateDistribution . For information, see [UpdateDistribution] in the
// Amazon CloudFront API Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for DisassociateWebACL]in the WAF Developer Guide.
//
// [Permissions for DisassociateWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-DisassociateWebACL
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
func (c *Client) DisassociateWebACL(ctx context.Context, params *DisassociateWebACLInput, optFns ...func(*Options)) (*DisassociateWebACLOutput, error) {
	if params == nil {
		params = &DisassociateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateWebACL", params, optFns, c.addOperationDisassociateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateWebACLInput struct {

	// The Amazon Resource Name (ARN) of the resource to disassociate from the web
	// ACL.
	//
	// The ARN must be in one of the following formats:
	//
	//   - For an Application Load Balancer:
	//   arn:partition:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//   - For an Amazon API Gateway REST API:
	//   arn:partition:apigateway:region::/restapis/api-id/stages/stage-name
	//
	//   - For an AppSync GraphQL API:
	//   arn:partition:appsync:region:account-id:apis/GraphQLApiId
	//
	//   - For an Amazon Cognito user pool:
	//   arn:partition:cognito-idp:region:account-id:userpool/user-pool-id
	//
	//   - For an App Runner service:
	//   arn:partition:apprunner:region:account-id:service/apprunner-service-name/apprunner-service-id
	//
	//   - For an Amazon Web Services Verified Access instance:
	//   arn:partition:ec2:region:account-id:verified-access-instance/instance-id
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type DisassociateWebACLOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateWebACL"); err != nil {
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
	if err = addOpDisassociateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateWebACL",
	}
}
