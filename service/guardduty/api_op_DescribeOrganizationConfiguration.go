// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the account selected as the delegated administrator
// for GuardDuty.
func (c *Client) DescribeOrganizationConfiguration(ctx context.Context, params *DescribeOrganizationConfigurationInput, optFns ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConfiguration", params, optFns, c.addOperationDescribeOrganizationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConfigurationInput struct {

	// The ID of the detector to retrieve information about the delegated administrator
	// from.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type DescribeOrganizationConfigurationOutput struct {

	// Indicates whether GuardDuty is automatically enabled for accounts added to the
	// organization.
	//
	// This member is required.
	AutoEnable bool

	// Indicates whether the maximum number of allowed member accounts are already
	// associated with the delegated administrator account for your organization.
	//
	// This member is required.
	MemberAccountLimitReached bool

	// Describes which data sources are enabled automatically for member accounts.
	DataSources *types.OrganizationDataSourceConfigurationsResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOrganizationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOrganizationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeOrganizationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConfiguration(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDescribeOrganizationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "DescribeOrganizationConfiguration",
	}
}