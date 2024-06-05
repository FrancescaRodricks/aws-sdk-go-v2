// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more tags to an IAM server certificate. If a tag with the same key
// name already exists, then that tag is overwritten with the new value.
//
// For certificates in a Region supported by Certificate Manager (ACM), we
// recommend that you don't use IAM server certificates. Instead, use ACM to
// provision, manage, and deploy your server certificates. For more information
// about IAM server certificates, [Working with server certificates]in the IAM User Guide.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
//   - Administrative grouping and discovery - Attach tags to resources to aid in
//     organization and search. For example, you could search for all resources with
//     the key name Project and the value MyImportantProject. Or search for all
//     resources with the key name Cost Center and the value 41200.
//
//   - Access control - Include tags in IAM user-based and resource-based
//     policies. You can use tags to restrict access to only a server certificate that
//     has a specified tag attached. For examples of policies that show how to use tags
//     to control access, see [Control access using IAM tags]in the IAM User Guide.
//
//   - Cost allocation - Use tags to help track which individuals and teams are
//     using which Amazon Web Services resources.
//
//   - If any one of the tags is invalid or if you exceed the allowed maximum
//     number of tags, then the entire request fails and the resource is not created.
//     For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
//   - Amazon Web Services always interprets the tag Value as a single string. If
//     you need to store an array, you can store comma-separated values in the string.
//     However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func (c *Client) TagServerCertificate(ctx context.Context, params *TagServerCertificateInput, optFns ...func(*Options)) (*TagServerCertificateOutput, error) {
	if params == nil {
		params = &TagServerCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagServerCertificate", params, optFns, c.addOperationTagServerCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagServerCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagServerCertificateInput struct {

	// The name of the IAM server certificate to which you want to add tags.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	ServerCertificateName *string

	// The list of tags that you want to attach to the IAM server certificate. Each
	// tag consists of a key name and an associated value.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type TagServerCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagServerCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TagServerCertificate"); err != nil {
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
	if err = addOpTagServerCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagServerCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagServerCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TagServerCertificate",
	}
}
