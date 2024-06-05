// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Sets the resource policy on a repository that specifies permissions to access
//
// it.
//
// When you call PutRepositoryPermissionsPolicy , the resource policy on the
// repository is ignored when evaluting permissions. This ensures that the owner of
// a repository cannot lock themselves out of the repository, which would prevent
// them from being able to update the resource policy.
func (c *Client) PutRepositoryPermissionsPolicy(ctx context.Context, params *PutRepositoryPermissionsPolicyInput, optFns ...func(*Options)) (*PutRepositoryPermissionsPolicyOutput, error) {
	if params == nil {
		params = &PutRepositoryPermissionsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRepositoryPermissionsPolicy", params, optFns, c.addOperationPutRepositoryPermissionsPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRepositoryPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRepositoryPermissionsPolicyInput struct {

	//  The name of the domain containing the repository to set the resource policy
	// on.
	//
	// This member is required.
	Domain *string

	//  A valid displayable JSON Aspen policy string to be set as the access control
	// resource policy on the provided repository.
	//
	// This member is required.
	PolicyDocument *string

	//  The name of the repository to set the resource policy on.
	//
	// This member is required.
	Repository *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	//  Sets the revision of the resource policy that specifies permissions to access
	// the repository. This revision is used for optimistic locking, which prevents
	// others from overwriting your changes to the repository's resource policy.
	PolicyRevision *string

	noSmithyDocumentSerde
}

type PutRepositoryPermissionsPolicyOutput struct {

	//  The resource policy that was set after processing the request.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRepositoryPermissionsPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRepositoryPermissionsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRepositoryPermissionsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRepositoryPermissionsPolicy"); err != nil {
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
	if err = addOpPutRepositoryPermissionsPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRepositoryPermissionsPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRepositoryPermissionsPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRepositoryPermissionsPolicy",
	}
}
