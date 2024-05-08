// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified policy from the specified certificate.
//
// Note: This action is deprecated and works as expected for backward
// compatibility, but we won't add enhancements. Use DetachPolicyinstead.
//
// Requires permission to access the [DetachPrincipalPolicy] action.
//
// Deprecated: This operation has been deprecated.
//
// [DetachPrincipalPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) DetachPrincipalPolicy(ctx context.Context, params *DetachPrincipalPolicyInput, optFns ...func(*Options)) (*DetachPrincipalPolicyOutput, error) {
	if params == nil {
		params = &DetachPrincipalPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachPrincipalPolicy", params, optFns, c.addOperationDetachPrincipalPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachPrincipalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DetachPrincipalPolicy operation.
type DetachPrincipalPolicyInput struct {

	// The name of the policy to detach.
	//
	// This member is required.
	PolicyName *string

	// The principal.
	//
	// Valid principals are CertificateArn
	// (arn:aws:iot:region:accountId:cert/certificateId), thingGroupArn
	// (arn:aws:iot:region:accountId:thinggroup/groupName) and CognitoId (region:id).
	//
	// This member is required.
	Principal *string

	noSmithyDocumentSerde
}

type DetachPrincipalPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachPrincipalPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDetachPrincipalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachPrincipalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetachPrincipalPolicy"); err != nil {
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
	if err = addOpDetachPrincipalPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachPrincipalPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachPrincipalPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetachPrincipalPolicy",
	}
}
