// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about an access control configuration that you created for
// your documents in an index. This includes user and group access information for
// your documents. This is useful for user context filtering, where search results
// are filtered based on the user or their group access to documents.
func (c *Client) DescribeAccessControlConfiguration(ctx context.Context, params *DescribeAccessControlConfigurationInput, optFns ...func(*Options)) (*DescribeAccessControlConfigurationOutput, error) {
	if params == nil {
		params = &DescribeAccessControlConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccessControlConfiguration", params, optFns, c.addOperationDescribeAccessControlConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccessControlConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccessControlConfigurationInput struct {

	// The identifier of the access control configuration you want to get information
	// on.
	//
	// This member is required.
	Id *string

	// The identifier of the index for an access control configuration.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeAccessControlConfigurationOutput struct {

	// The name for the access control configuration.
	//
	// This member is required.
	Name *string

	// Information on principals (users and/or groups) and which documents they should
	// have access to. This is useful for user context filtering, where search results
	// are filtered based on the user or their group access to documents.
	AccessControlList []types.Principal

	// The description for the access control configuration.
	Description *string

	// The error message containing details if there are issues processing the access
	// control configuration.
	ErrorMessage *string

	// The list of [principal] lists that define the hierarchy for which documents users should
	// have access to.
	//
	// [principal]: https://docs.aws.amazon.com/kendra/latest/dg/API_Principal.html
	HierarchicalAccessControlList []types.HierarchicalPrincipal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccessControlConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccessControlConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccessControlConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAccessControlConfiguration"); err != nil {
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
	if err = addOpDescribeAccessControlConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccessControlConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccessControlConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAccessControlConfiguration",
	}
}
