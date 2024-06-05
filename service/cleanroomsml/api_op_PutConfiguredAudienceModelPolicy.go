// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create or update the resource policy for a configured audience model.
func (c *Client) PutConfiguredAudienceModelPolicy(ctx context.Context, params *PutConfiguredAudienceModelPolicyInput, optFns ...func(*Options)) (*PutConfiguredAudienceModelPolicyOutput, error) {
	if params == nil {
		params = &PutConfiguredAudienceModelPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfiguredAudienceModelPolicy", params, optFns, c.addOperationPutConfiguredAudienceModelPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfiguredAudienceModelPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConfiguredAudienceModelPolicyInput struct {

	// The Amazon Resource Name (ARN) of the configured audience model that the
	// resource policy will govern.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The IAM resource policy.
	//
	// This member is required.
	ConfiguredAudienceModelPolicy *string

	// Use this to prevent unexpected concurrent modification of the policy.
	PolicyExistenceCondition types.PolicyExistenceCondition

	// A cryptographic hash of the contents of the policy used to prevent unexpected
	// concurrent modification of the policy.
	PreviousPolicyHash *string

	noSmithyDocumentSerde
}

type PutConfiguredAudienceModelPolicyOutput struct {

	// The IAM resource policy.
	//
	// This member is required.
	ConfiguredAudienceModelPolicy *string

	// A cryptographic hash of the contents of the policy used to prevent unexpected
	// concurrent modification of the policy.
	//
	// This member is required.
	PolicyHash *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutConfiguredAudienceModelPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutConfiguredAudienceModelPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfiguredAudienceModelPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutConfiguredAudienceModelPolicy"); err != nil {
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
	if err = addOpPutConfiguredAudienceModelPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfiguredAudienceModelPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfiguredAudienceModelPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutConfiguredAudienceModelPolicy",
	}
}
