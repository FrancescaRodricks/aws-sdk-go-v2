// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Associates a target account, organizational unit, or the root with a specified
//
// configuration. The target can be associated with a configuration policy or
// self-managed behavior. Only the Security Hub delegated administrator can invoke
// this operation from the home Region.
func (c *Client) StartConfigurationPolicyAssociation(ctx context.Context, params *StartConfigurationPolicyAssociationInput, optFns ...func(*Options)) (*StartConfigurationPolicyAssociationOutput, error) {
	if params == nil {
		params = &StartConfigurationPolicyAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartConfigurationPolicyAssociation", params, optFns, c.addOperationStartConfigurationPolicyAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartConfigurationPolicyAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartConfigurationPolicyAssociationInput struct {

	//  The Amazon Resource Name (ARN) of a configuration policy, the universally
	// unique identifier (UUID) of a configuration policy, or a value of
	// SELF_MANAGED_SECURITY_HUB for a self-managed configuration.
	//
	// This member is required.
	ConfigurationPolicyIdentifier *string

	//  The identifier of the target account, organizational unit, or the root to
	// associate with the specified configuration.
	//
	// This member is required.
	Target types.Target

	noSmithyDocumentSerde
}

type StartConfigurationPolicyAssociationOutput struct {

	//  The current status of the association between the specified target and the
	// configuration.
	AssociationStatus types.ConfigurationPolicyAssociationStatus

	//  An explanation for a FAILED value for AssociationStatus .
	AssociationStatusMessage *string

	//  Indicates whether the association between the specified target and the
	// configuration was directly applied by the Security Hub delegated administrator
	// or inherited from a parent.
	AssociationType types.AssociationType

	//  The UUID of the configuration policy.
	ConfigurationPolicyId *string

	//  The identifier of the target account, organizational unit, or the organization
	// root with which the configuration is associated.
	TargetId *string

	//  Indicates whether the target is an Amazon Web Services account, organizational
	// unit, or the organization root.
	TargetType types.TargetType

	//  The date and time, in UTC and ISO 8601 format, that the configuration policy
	// association was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartConfigurationPolicyAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartConfigurationPolicyAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartConfigurationPolicyAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartConfigurationPolicyAssociation"); err != nil {
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
	if err = addOpStartConfigurationPolicyAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartConfigurationPolicyAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartConfigurationPolicyAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartConfigurationPolicyAssociation",
	}
}
