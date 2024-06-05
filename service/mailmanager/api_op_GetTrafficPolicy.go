// Code generated by smithy-go-codegen DO NOT EDIT.

package mailmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mailmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Fetch attributes of a traffic policy resource.
func (c *Client) GetTrafficPolicy(ctx context.Context, params *GetTrafficPolicyInput, optFns ...func(*Options)) (*GetTrafficPolicyOutput, error) {
	if params == nil {
		params = &GetTrafficPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTrafficPolicy", params, optFns, c.addOperationGetTrafficPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTrafficPolicyInput struct {

	// The identifier of the traffic policy resource.
	//
	// This member is required.
	TrafficPolicyId *string

	noSmithyDocumentSerde
}

type GetTrafficPolicyOutput struct {

	// The identifier of the traffic policy resource.
	//
	// This member is required.
	TrafficPolicyId *string

	// A user-friendly name for the traffic policy resource.
	//
	// This member is required.
	TrafficPolicyName *string

	// The timestamp of when the traffic policy was created.
	CreatedTimestamp *time.Time

	// The default action of the traffic policy.
	DefaultAction types.AcceptAction

	// The timestamp of when the traffic policy was last updated.
	LastUpdatedTimestamp *time.Time

	// The maximum message size in bytes of email which is allowed in by this traffic
	// policy—anything larger will be blocked.
	MaxMessageSizeBytes *int32

	// The list of conditions which are in the traffic policy resource.
	PolicyStatements []types.PolicyStatement

	// The Amazon Resource Name (ARN) of the traffic policy resource.
	TrafficPolicyArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTrafficPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTrafficPolicy"); err != nil {
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
	if err = addOpGetTrafficPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrafficPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTrafficPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTrafficPolicy",
	}
}
