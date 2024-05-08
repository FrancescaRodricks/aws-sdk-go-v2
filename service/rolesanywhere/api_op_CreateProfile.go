// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a profile, a list of the roles that Roles Anywhere service is trusted
// to assume. You use profiles to intersect permissions with IAM managed policies.
//
// Required permissions: rolesanywhere:CreateProfile .
func (c *Client) CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) {
	if params == nil {
		params = &CreateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfile", params, optFns, c.addOperationCreateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileInput struct {

	// The name of the profile.
	//
	// This member is required.
	Name *string

	// A list of IAM roles that this profile can assume in a temporary credential
	// request.
	//
	// This member is required.
	RoleArns []string

	//  Used to determine how long sessions vended using this profile are valid for.
	// See the Expiration section of the [CreateSession API documentation] page for more details. In requests, if this
	// value is not provided, the default value will be 3600.
	//
	// [CreateSession API documentation]: https://docs.aws.amazon.com/rolesanywhere/latest/userguide/authentication-create-session.html#credentials-object
	DurationSeconds *int32

	// Specifies whether the profile is enabled.
	Enabled *bool

	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns []string

	// Specifies whether instance properties are required in temporary credential
	// requests with this profile.
	RequireInstanceProperties *bool

	// A session policy that applies to the trust boundary of the vended session
	// credentials.
	SessionPolicy *string

	// The tags to attach to the profile.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProfileOutput struct {

	// The state of the profile after a read or write operation.
	Profile *types.ProfileDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProfile"); err != nil {
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
	if err = addOpCreateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProfile",
	}
}
