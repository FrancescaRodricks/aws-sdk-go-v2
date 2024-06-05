// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a template for approval rules that can then be associated with one or
// more repositories in your Amazon Web Services account. When you associate a
// template with a repository, CodeCommit creates an approval rule that matches the
// conditions of the template for all pull requests that meet the conditions of the
// template. For more information, see AssociateApprovalRuleTemplateWithRepository.
func (c *Client) CreateApprovalRuleTemplate(ctx context.Context, params *CreateApprovalRuleTemplateInput, optFns ...func(*Options)) (*CreateApprovalRuleTemplateOutput, error) {
	if params == nil {
		params = &CreateApprovalRuleTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApprovalRuleTemplate", params, optFns, c.addOperationCreateApprovalRuleTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApprovalRuleTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApprovalRuleTemplateInput struct {

	// The content of the approval rule that is created on pull requests in associated
	// repositories. If you specify one or more destination references (branches),
	// approval rules are created in an associated repository only if their destination
	// references (branches) match those specified in the template.
	//
	// When you create the content of the approval rule template, you can specify
	// approvers in an approval pool in one of two ways:
	//
	//   - CodeCommitApprovers: This option only requires an Amazon Web Services
	//   account and a resource. It can be used for both IAM users and federated access
	//   users whose name matches the provided resource name. This is a very powerful
	//   option that offers a great deal of flexibility. For example, if you specify the
	//   Amazon Web Services account 123456789012 and Mary_Major, all of the following
	//   are counted as approvals coming from that user:
	//
	//   - An IAM user in the account (arn:aws:iam::123456789012:user/Mary_Major)
	//
	//   - A federated user identified in IAM as Mary_Major
	//   (arn:aws:sts::123456789012:federated-user/Mary_Major)
	//
	// This option does not recognize an active session of someone assuming the role
	//   of CodeCommitReview with a role session name of Mary_Major
	//   (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major) unless you
	//   include a wildcard (*Mary_Major).
	//
	//   - Fully qualified ARN: This option allows you to specify the fully qualified
	//   Amazon Resource Name (ARN) of the IAM user or role.
	//
	// For more information about IAM ARNs, wildcards, and formats, see [IAM Identifiers] in the IAM
	// User Guide.
	//
	// [IAM Identifiers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html
	//
	// This member is required.
	ApprovalRuleTemplateContent *string

	// The name of the approval rule template. Provide descriptive names, because this
	// name is applied to the approval rules created automatically in associated
	// repositories.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The description of the approval rule template. Consider providing a description
	// that explains what this template does and when it might be appropriate to
	// associate it with repositories.
	ApprovalRuleTemplateDescription *string

	noSmithyDocumentSerde
}

type CreateApprovalRuleTemplateOutput struct {

	// The content and structure of the created approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplate *types.ApprovalRuleTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApprovalRuleTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApprovalRuleTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApprovalRuleTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApprovalRuleTemplate"); err != nil {
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
	if err = addOpCreateApprovalRuleTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApprovalRuleTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApprovalRuleTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApprovalRuleTemplate",
	}
}
