// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the availability options configured for a domain. By default, shows the
// configuration with any pending changes. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see [Configuring Availability Options]
// in the Amazon CloudSearch Developer Guide.
//
// [Configuring Availability Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-availability-options.html
func (c *Client) DescribeAvailabilityOptions(ctx context.Context, params *DescribeAvailabilityOptionsInput, optFns ...func(*Options)) (*DescribeAvailabilityOptionsOutput, error) {
	if params == nil {
		params = &DescribeAvailabilityOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAvailabilityOptions", params, optFns, c.addOperationDescribeAvailabilityOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAvailabilityOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeAvailabilityOptions operation. Specifies the name of the
// domain you want to describe. To show the active configuration and exclude any
// pending changes, set the Deployed option to true .
type DescribeAvailabilityOptionsInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration ( true ) or include any pending
	// changes ( false ). Defaults to false .
	Deployed *bool

	noSmithyDocumentSerde
}

// The result of a DescribeAvailabilityOptions request. Indicates whether or not
// the Multi-AZ option is enabled for the domain specified in the request.
type DescribeAvailabilityOptionsOutput struct {

	// The availability options configured for the domain. Indicates whether Multi-AZ
	// is enabled for the domain.
	AvailabilityOptions *types.AvailabilityOptionsStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAvailabilityOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAvailabilityOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAvailabilityOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAvailabilityOptions"); err != nil {
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
	if err = addOpDescribeAvailabilityOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAvailabilityOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAvailabilityOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAvailabilityOptions",
	}
}
