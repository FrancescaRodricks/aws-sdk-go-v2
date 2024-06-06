// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the opt-in status of a particular Region.
func (c *Client) GetRegionOptStatus(ctx context.Context, params *GetRegionOptStatusInput, optFns ...func(*Options)) (*GetRegionOptStatusOutput, error) {
	if params == nil {
		params = &GetRegionOptStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRegionOptStatus", params, optFns, c.addOperationGetRegionOptStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRegionOptStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRegionOptStatusInput struct {

	// Specifies the Region-code for a given Region name (for example, af-south-1 ).
	// This function will return the status of whatever Region you pass into this
	// parameter.
	//
	// This member is required.
	RegionName *string

	// Specifies the 12-digit account ID number of the Amazon Web Services account
	// that you want to access or modify with this operation. If you don't specify this
	// parameter, it defaults to the Amazon Web Services account of the identity used
	// to call the operation. To use this parameter, the caller must be an identity in
	// the [organization's management account]or a delegated administrator account. The specified account ID must be a
	// member account in the same organization. The organization must have [all features enabled], and the
	// organization must have [trusted access]enabled for the Account Management service, and
	// optionally a [delegated admin]account assigned.
	//
	// The management account can't specify its own AccountId . It must call the
	// operation in standalone context by not including the AccountId parameter.
	//
	// To call this operation on an account that is not a member of an organization,
	// don't specify this parameter. Instead, call the operation using an identity
	// belonging to the account whose contacts you wish to retrieve or modify.
	//
	// [organization's management account]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account
	// [all features enabled]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
	// [delegated admin]: https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html
	// [trusted access]: https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html
	AccountId *string

	noSmithyDocumentSerde
}

type GetRegionOptStatusOutput struct {

	// The Region code that was passed in.
	RegionName *string

	// One of the potential statuses a Region can undergo (Enabled, Enabling,
	// Disabled, Disabling, Enabled_By_Default).
	RegionOptStatus types.RegionOptStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRegionOptStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRegionOptStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRegionOptStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRegionOptStatus"); err != nil {
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
	if err = addOpGetRegionOptStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRegionOptStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRegionOptStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRegionOptStatus",
	}
}
