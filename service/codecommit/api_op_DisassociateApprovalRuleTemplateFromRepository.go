// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the association between a template and a repository so that approval
// rules based on the template are not automatically created when pull requests are
// created in the specified repository. This does not delete any approval rules
// previously created for pull requests through the template association.
func (c *Client) DisassociateApprovalRuleTemplateFromRepository(ctx context.Context, params *DisassociateApprovalRuleTemplateFromRepositoryInput, optFns ...func(*Options)) (*DisassociateApprovalRuleTemplateFromRepositoryOutput, error) {
	if params == nil {
		params = &DisassociateApprovalRuleTemplateFromRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateApprovalRuleTemplateFromRepository", params, optFns, c.addOperationDisassociateApprovalRuleTemplateFromRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateApprovalRuleTemplateFromRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateApprovalRuleTemplateFromRepositoryInput struct {

	// The name of the approval rule template to disassociate from a specified
	// repository.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The name of the repository you want to disassociate from the template.
	//
	// This member is required.
	RepositoryName *string

	noSmithyDocumentSerde
}

type DisassociateApprovalRuleTemplateFromRepositoryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateApprovalRuleTemplateFromRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateApprovalRuleTemplateFromRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateApprovalRuleTemplateFromRepository{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateApprovalRuleTemplateFromRepository"); err != nil {
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
	if err = addOpDisassociateApprovalRuleTemplateFromRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateApprovalRuleTemplateFromRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateApprovalRuleTemplateFromRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateApprovalRuleTemplateFromRepository",
	}
}
