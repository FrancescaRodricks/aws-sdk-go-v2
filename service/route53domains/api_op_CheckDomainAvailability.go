// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation checks the availability of one domain name. Note that if the
// availability status of a domain is pending, you must submit another request to
// determine the availability of the domain name.
func (c *Client) CheckDomainAvailability(ctx context.Context, params *CheckDomainAvailabilityInput, optFns ...func(*Options)) (*CheckDomainAvailabilityOutput, error) {
	if params == nil {
		params = &CheckDomainAvailabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckDomainAvailability", params, optFns, c.addOperationCheckDomainAvailabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckDomainAvailabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CheckDomainAvailability request contains the following elements.
type CheckDomainAvailabilityInput struct {

	// The name of the domain that you want to get availability for. The top-level
	// domain (TLD), such as .com, must be a TLD that Route 53 supports. For a list of
	// supported TLDs, see [Domains that You Can Register with Amazon Route 53]in the Amazon Route 53 Developer Guide.
	//
	// The domain name can contain only the following characters:
	//
	//   - Letters a through z. Domain names are not case sensitive.
	//
	//   - Numbers 0 through 9.
	//
	//   - Hyphen (-). You can't specify a hyphen at the beginning or end of a label.
	//
	//   - Period (.) to separate the labels in the name, such as the . in example.com .
	//
	// Internationalized domain names are not supported for some top-level domains. To
	// determine whether the TLD that you want to use supports internationalized domain
	// names, see [Domains that You Can Register with Amazon Route 53]. For more information, see [Formatting Internationalized Domain Names].
	//
	// [Formatting Internationalized Domain Names]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html#domain-name-format-idns
	// [Domains that You Can Register with Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html
	//
	// This member is required.
	DomainName *string

	// Reserved for future use.
	IdnLangCode *string

	noSmithyDocumentSerde
}

// The CheckDomainAvailability response includes the following elements.
type CheckDomainAvailabilityOutput struct {

	// Whether the domain name is available for registering.
	//
	// You can register only domains designated as AVAILABLE .
	//
	// Valid values:
	//
	// AVAILABLE The domain name is available.
	//
	// AVAILABLE_RESERVED The domain name is reserved under specific conditions.
	//
	// AVAILABLE_PREORDER The domain name is available and can be preordered.
	//
	// DONT_KNOW The TLD registry didn't reply with a definitive answer about whether
	// the domain name is available. Route 53 can return this response for a variety of
	// reasons, for example, the registry is performing maintenance. Try again later.
	//
	// INVALID_NAME_FOR_TLD The TLD isn't valid. For example, it can contain
	// characters that aren't allowed.
	//
	// PENDING The TLD registry didn't return a response in the expected amount of
	// time. When the response is delayed, it usually takes just a few extra seconds.
	// You can resubmit the request immediately.
	//
	// RESERVED The domain name has been reserved for another person or organization.
	//
	// UNAVAILABLE The domain name is not available.
	//
	// UNAVAILABLE_PREMIUM The domain name is not available.
	//
	// UNAVAILABLE_RESTRICTED The domain name is forbidden.
	Availability types.DomainAvailability

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCheckDomainAvailabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCheckDomainAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCheckDomainAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CheckDomainAvailability"); err != nil {
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
	if err = addOpCheckDomainAvailabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckDomainAvailability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckDomainAvailability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CheckDomainAvailability",
	}
}
