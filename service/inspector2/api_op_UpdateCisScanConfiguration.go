// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a CIS scan configuration.
func (c *Client) UpdateCisScanConfiguration(ctx context.Context, params *UpdateCisScanConfigurationInput, optFns ...func(*Options)) (*UpdateCisScanConfigurationOutput, error) {
	if params == nil {
		params = &UpdateCisScanConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCisScanConfiguration", params, optFns, c.addOperationUpdateCisScanConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCisScanConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCisScanConfigurationInput struct {

	// The CIS scan configuration ARN.
	//
	// This member is required.
	ScanConfigurationArn *string

	// The scan name for the CIS scan configuration.
	ScanName *string

	// The schedule for the CIS scan configuration.
	Schedule types.Schedule

	//  The security level for the CIS scan configuration. Security level refers to
	// the Benchmark levels that CIS assigns to a profile.
	SecurityLevel types.CisSecurityLevel

	// The targets for the CIS scan configuration.
	Targets *types.UpdateCisTargets

	noSmithyDocumentSerde
}

type UpdateCisScanConfigurationOutput struct {

	// The CIS scan configuration ARN.
	//
	// This member is required.
	ScanConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCisScanConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCisScanConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCisScanConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCisScanConfiguration"); err != nil {
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
	if err = addOpUpdateCisScanConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCisScanConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCisScanConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCisScanConfiguration",
	}
}
