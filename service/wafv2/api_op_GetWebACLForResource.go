// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the WebACL for the specified resource.
//
// This call uses GetWebACL , to verify that your account has permission to access
// the retrieved web ACL. If you get an error that indicates that your account
// isn't authorized to perform wafv2:GetWebACL on the resource, that error won't
// be included in your CloudTrail event history.
//
// For Amazon CloudFront, don't use this call. Instead, call the CloudFront action
// GetDistributionConfig . For information, see [GetDistributionConfig] in the Amazon CloudFront API
// Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for GetWebACLForResource]in the WAF Developer Guide.
//
// [GetDistributionConfig]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistributionConfig.html
// [Permissions for GetWebACLForResource]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-GetWebACLForResource
func (c *Client) GetWebACLForResource(ctx context.Context, params *GetWebACLForResourceInput, optFns ...func(*Options)) (*GetWebACLForResourceOutput, error) {
	if params == nil {
		params = &GetWebACLForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWebACLForResource", params, optFns, c.addOperationGetWebACLForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWebACLForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWebACLForResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource whose web ACL you want to
	// retrieve.
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

type GetWebACLForResourceOutput struct {

	// The web ACL that is associated with the resource. If there is no associated
	// resource, WAF returns a null web ACL.
	WebACL *types.WebACL

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWebACLForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWebACLForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWebACLForResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWebACLForResource"); err != nil {
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
	if err = addOpGetWebACLForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWebACLForResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWebACLForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWebACLForResource",
	}
}
