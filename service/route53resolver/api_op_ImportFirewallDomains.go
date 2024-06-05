// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports domain names from a file into a domain list, for use in a DNS firewall
// rule group.
//
// Each domain specification in your domain list must satisfy the following
// requirements:
//
//   - It can optionally start with * (asterisk).
//
//   - With the exception of the optional starting asterisk, it must only contain
//     the following characters: A-Z , a-z , 0-9 , - (hyphen).
//
//   - It must be from 1-255 characters in length.
func (c *Client) ImportFirewallDomains(ctx context.Context, params *ImportFirewallDomainsInput, optFns ...func(*Options)) (*ImportFirewallDomainsOutput, error) {
	if params == nil {
		params = &ImportFirewallDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportFirewallDomains", params, optFns, c.addOperationImportFirewallDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportFirewallDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportFirewallDomainsInput struct {

	// The fully qualified URL or URI of the file stored in Amazon Simple Storage
	// Service (Amazon S3) that contains the list of domains to import.
	//
	// The file must be in an S3 bucket that's in the same Region as your DNS
	// Firewall. The file must be a text file and must contain a single domain per
	// line.
	//
	// This member is required.
	DomainFileUrl *string

	// The ID of the domain list that you want to modify with the import operation.
	//
	// This member is required.
	FirewallDomainListId *string

	// What you want DNS Firewall to do with the domains that are listed in the file.
	// This must be set to REPLACE , which updates the domain list to exactly match the
	// list in the file.
	//
	// This member is required.
	Operation types.FirewallDomainImportOperation

	noSmithyDocumentSerde
}

type ImportFirewallDomainsOutput struct {

	// The Id of the firewall domain list that DNS Firewall just updated.
	Id *string

	// The name of the domain list.
	Name *string

	// Status of the import request.
	Status types.FirewallDomainListStatus

	// Additional information about the status of the list, if available.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportFirewallDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportFirewallDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportFirewallDomains{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportFirewallDomains"); err != nil {
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
	if err = addOpImportFirewallDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportFirewallDomains(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportFirewallDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportFirewallDomains",
	}
}
