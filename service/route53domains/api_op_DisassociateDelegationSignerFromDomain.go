// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a delegation signer (DS) record in the registry zone for this domain
// name.
func (c *Client) DisassociateDelegationSignerFromDomain(ctx context.Context, params *DisassociateDelegationSignerFromDomainInput, optFns ...func(*Options)) (*DisassociateDelegationSignerFromDomainOutput, error) {
	if params == nil {
		params = &DisassociateDelegationSignerFromDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDelegationSignerFromDomain", params, optFns, c.addOperationDisassociateDelegationSignerFromDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDelegationSignerFromDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDelegationSignerFromDomainInput struct {

	// Name of the domain.
	//
	// This member is required.
	DomainName *string

	// An internal identification number assigned to each DS record after it’s
	// created. You can retrieve it as part of DNSSEC information returned by [GetDomainDetail].
	//
	// [GetDomainDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetDomainDetail.html
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DisassociateDelegationSignerFromDomainOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use [GetOperationDetail].
	//
	// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateDelegationSignerFromDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDelegationSignerFromDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDelegationSignerFromDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateDelegationSignerFromDomain"); err != nil {
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
	if err = addOpDisassociateDelegationSignerFromDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDelegationSignerFromDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateDelegationSignerFromDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateDelegationSignerFromDomain",
	}
}
