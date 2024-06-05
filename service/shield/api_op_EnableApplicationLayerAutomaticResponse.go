// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enable the Shield Advanced automatic application layer DDoS mitigation for the
// protected resource.
//
// This feature is available for Amazon CloudFront distributions and Application
// Load Balancers only.
//
// This causes Shield Advanced to create, verify, and apply WAF rules for DDoS
// attacks that it detects for the resource. Shield Advanced applies the rules in a
// Shield rule group inside the web ACL that you've associated with the resource.
// For information about how automatic mitigation works and the requirements for
// using it, see [Shield Advanced automatic application layer DDoS mitigation].
//
// Don't use this action to make changes to automatic mitigation settings when
// it's already enabled for a resource. Instead, use UpdateApplicationLayerAutomaticResponse.
//
// To use this feature, you must associate a web ACL with the protected resource.
// The web ACL must be created using the latest version of WAF (v2). You can
// associate the web ACL through the Shield Advanced console at [https://console.aws.amazon.com/wafv2/shieldv2#/]. For more
// information, see [Getting Started with Shield Advanced]. You can also associate the web ACL to the resource through
// the WAF console or the WAF API, but you must manage Shield Advanced automatic
// mitigation through Shield Advanced. For information about WAF, see [WAF Developer Guide].
//
// [https://console.aws.amazon.com/wafv2/shieldv2#/]: https://console.aws.amazon.com/wafv2/shieldv2#/
// [Getting Started with Shield Advanced]: https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html
// [WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [Shield Advanced automatic application layer DDoS mitigation]: https://docs.aws.amazon.com/waf/latest/developerguide/ddos-advanced-automatic-app-layer-response.html
func (c *Client) EnableApplicationLayerAutomaticResponse(ctx context.Context, params *EnableApplicationLayerAutomaticResponseInput, optFns ...func(*Options)) (*EnableApplicationLayerAutomaticResponseOutput, error) {
	if params == nil {
		params = &EnableApplicationLayerAutomaticResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableApplicationLayerAutomaticResponse", params, optFns, c.addOperationEnableApplicationLayerAutomaticResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableApplicationLayerAutomaticResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableApplicationLayerAutomaticResponseInput struct {

	// Specifies the action setting that Shield Advanced should use in the WAF rules
	// that it creates on behalf of the protected resource in response to DDoS attacks.
	// You specify this as part of the configuration for the automatic application
	// layer DDoS mitigation feature, when you enable or update automatic mitigation.
	// Shield Advanced creates the WAF rules in a Shield Advanced-managed rule group,
	// inside the web ACL that you have associated with the resource.
	//
	// This member is required.
	Action *types.ResponseAction

	// The ARN (Amazon Resource Name) of the protected resource.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type EnableApplicationLayerAutomaticResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableApplicationLayerAutomaticResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableApplicationLayerAutomaticResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableApplicationLayerAutomaticResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableApplicationLayerAutomaticResponse"); err != nil {
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
	if err = addOpEnableApplicationLayerAutomaticResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableApplicationLayerAutomaticResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableApplicationLayerAutomaticResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableApplicationLayerAutomaticResponse",
	}
}
