// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a previously associated principal ARN from a specified portfolio.
//
// The PrincipalType and PrincipalARN must match the
// AssociatePrincipalWithPortfolio call request details. For example, to
// disassociate an association created with a PrincipalARN of PrincipalType IAM
// you must use the PrincipalType IAM when calling
// DisassociatePrincipalFromPortfolio .
//
// For portfolios that have been shared with principal name sharing enabled: after
// disassociating a principal, share recipient accounts will no longer be able to
// provision products in this portfolio using a role matching the name of the
// associated principal.
//
// For more information, review [associate-principal-with-portfolio] in the Amazon Web Services CLI Command Reference.
//
// If you disassociate a principal from a portfolio, with PrincipalType as IAM ,
// the same principal will still have access to the portfolio if it matches one of
// the associated principals of type IAM_PATTERN . To fully remove access for a
// principal, verify all the associated Principals of type IAM_PATTERN , and then
// ensure you disassociate any IAM_PATTERN principals that match the principal
// whose access you are removing.
//
// [associate-principal-with-portfolio]: https://docs.aws.amazon.com/cli/latest/reference/servicecatalog/associate-principal-with-portfolio.html#options
func (c *Client) DisassociatePrincipalFromPortfolio(ctx context.Context, params *DisassociatePrincipalFromPortfolioInput, optFns ...func(*Options)) (*DisassociatePrincipalFromPortfolioOutput, error) {
	if params == nil {
		params = &DisassociatePrincipalFromPortfolioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociatePrincipalFromPortfolio", params, optFns, c.addOperationDisassociatePrincipalFromPortfolioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociatePrincipalFromPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociatePrincipalFromPortfolioInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The ARN of the principal (user, role, or group). This field allows an ARN with
	// no accountID with or without wildcard characters if PrincipalType is IAM_PATTERN
	// .
	//
	// This member is required.
	PrincipalARN *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// The supported value is IAM if you use a fully defined ARN, or IAM_PATTERN if
	// you specify an IAM ARN with no AccountId, with or without wildcard characters.
	PrincipalType types.PrincipalType

	noSmithyDocumentSerde
}

type DisassociatePrincipalFromPortfolioOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociatePrincipalFromPortfolioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociatePrincipalFromPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociatePrincipalFromPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociatePrincipalFromPortfolio"); err != nil {
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
	if err = addOpDisassociatePrincipalFromPortfolioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociatePrincipalFromPortfolio(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociatePrincipalFromPortfolio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociatePrincipalFromPortfolio",
	}
}
