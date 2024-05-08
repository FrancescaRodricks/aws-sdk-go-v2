// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified policy template. You can update only the description and
// the some elements of the [policyBody].
//
// Changes you make to the policy template content are immediately (within the
// constraints of eventual consistency) reflected in authorization decisions that
// involve all template-linked policies instantiated from this template.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [policyBody]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyTemplate.html#amazonverifiedpermissions-UpdatePolicyTemplate-request-policyBody
func (c *Client) UpdatePolicyTemplate(ctx context.Context, params *UpdatePolicyTemplateInput, optFns ...func(*Options)) (*UpdatePolicyTemplateOutput, error) {
	if params == nil {
		params = &UpdatePolicyTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePolicyTemplate", params, optFns, c.addOperationUpdatePolicyTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePolicyTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePolicyTemplateInput struct {

	// Specifies the ID of the policy store that contains the policy template that you
	// want to update.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies the ID of the policy template that you want to update.
	//
	// This member is required.
	PolicyTemplateId *string

	// Specifies new statement content written in Cedar policy language to replace the
	// current body of the policy template.
	//
	// You can change only the following elements of the policy body:
	//
	//   - The action referenced by the policy template.
	//
	//   - Any conditional clauses, such as when or unless clauses.
	//
	// You can't change the following elements:
	//
	//   - The effect ( permit or forbid ) of the policy template.
	//
	//   - The principal referenced by the policy template.
	//
	//   - The resource referenced by the policy template.
	//
	// This member is required.
	Statement *string

	// Specifies a new description to apply to the policy template.
	Description *string

	noSmithyDocumentSerde
}

type UpdatePolicyTemplateOutput struct {

	// The date and time that the policy template was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time that the policy template was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy store that contains the updated policy template.
	//
	// This member is required.
	PolicyStoreId *string

	// The ID of the updated policy template.
	//
	// This member is required.
	PolicyTemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePolicyTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdatePolicyTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdatePolicyTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePolicyTemplate"); err != nil {
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
	if err = addOpUpdatePolicyTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePolicyTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePolicyTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePolicyTemplate",
	}
}
