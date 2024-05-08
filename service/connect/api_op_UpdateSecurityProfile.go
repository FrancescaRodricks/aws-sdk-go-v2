// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a security profile.
func (c *Client) UpdateSecurityProfile(ctx context.Context, params *UpdateSecurityProfileInput, optFns ...func(*Options)) (*UpdateSecurityProfileOutput, error) {
	if params == nil {
		params = &UpdateSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecurityProfile", params, optFns, c.addOperationUpdateSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecurityProfileInput struct {

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The identifier for the security profle.
	//
	// This member is required.
	SecurityProfileId *string

	// The identifier of the hierarchy group that a security profile uses to restrict
	// access to resources in Amazon Connect.
	AllowedAccessControlHierarchyGroupId *string

	// The list of tags that a security profile uses to restrict access to resources
	// in Amazon Connect.
	AllowedAccessControlTags map[string]string

	// A list of the third-party application's metadata.
	Applications []types.Application

	// The description of the security profile.
	Description *string

	// The list of resources that a security profile applies hierarchy restrictions to
	// in Amazon Connect. Following are acceptable ResourceNames: User .
	HierarchyRestrictedResources []string

	// The permissions granted to a security profile. For a list of valid permissions,
	// see [List of security profile permissions].
	//
	// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
	Permissions []string

	// The list of resources that a security profile applies tag restrictions to in
	// Amazon Connect.
	TagRestrictedResources []string

	noSmithyDocumentSerde
}

type UpdateSecurityProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSecurityProfile"); err != nil {
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
	if err = addOpUpdateSecurityProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecurityProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSecurityProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSecurityProfile",
	}
}
