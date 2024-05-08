// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a mobile device access rule for the specified WorkMail organization.
//
// Deleting already deleted and non-existing rules does not produce an error. In
// those cases, the service sends back an HTTP 200 response with an empty HTTP
// body.
func (c *Client) DeleteMobileDeviceAccessRule(ctx context.Context, params *DeleteMobileDeviceAccessRuleInput, optFns ...func(*Options)) (*DeleteMobileDeviceAccessRuleOutput, error) {
	if params == nil {
		params = &DeleteMobileDeviceAccessRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMobileDeviceAccessRule", params, optFns, c.addOperationDeleteMobileDeviceAccessRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMobileDeviceAccessRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMobileDeviceAccessRuleInput struct {

	// The identifier of the rule to be deleted.
	//
	// This member is required.
	MobileDeviceAccessRuleId *string

	// The WorkMail organization under which the rule will be deleted.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type DeleteMobileDeviceAccessRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMobileDeviceAccessRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMobileDeviceAccessRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMobileDeviceAccessRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteMobileDeviceAccessRule"); err != nil {
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
	if err = addOpDeleteMobileDeviceAccessRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMobileDeviceAccessRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMobileDeviceAccessRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteMobileDeviceAccessRule",
	}
}
