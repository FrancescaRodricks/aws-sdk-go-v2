// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of CreateCaseOption types along with the corresponding supported
// hours and language availability. You can specify the language categoryCode ,
// issueType and serviceCode used to retrieve the CreateCaseOptions.
//
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func (c *Client) DescribeCreateCaseOptions(ctx context.Context, params *DescribeCreateCaseOptionsInput, optFns ...func(*Options)) (*DescribeCreateCaseOptionsOutput, error) {
	if params == nil {
		params = &DescribeCreateCaseOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCreateCaseOptions", params, optFns, c.addOperationDescribeCreateCaseOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCreateCaseOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCreateCaseOptionsInput struct {

	// The category of problem for the support case. You also use the DescribeServices operation to
	// get the category code for a service. Each Amazon Web Services service defines
	// its own set of category codes.
	//
	// This member is required.
	CategoryCode *string

	// The type of issue for the case. You can specify customer-service or technical .
	// If you don't specify a value, the default is technical .
	//
	// This member is required.
	IssueType *string

	// The language in which Amazon Web Services Support handles the case. Amazon Web
	// Services Support currently supports Chinese (“zh”), English ("en"), Japanese
	// ("ja") and Korean (“ko”). You must specify the ISO 639-1 code for the language
	// parameter if you want support in that language.
	//
	// This member is required.
	Language *string

	// The code for the Amazon Web Services service. You can use the DescribeServices operation to get
	// the possible serviceCode values.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type DescribeCreateCaseOptionsOutput struct {

	//  A JSON-formatted array that contains the available communication type options,
	// along with the available support timeframes for the given inputs.
	CommunicationTypes []types.CommunicationTypeOptions

	// Language availability can be any of the following:
	//
	//   - available
	//
	//   - best_effort
	//
	//   - unavailable
	LanguageAvailability *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCreateCaseOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCreateCaseOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCreateCaseOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCreateCaseOptions"); err != nil {
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
	if err = addOpDescribeCreateCaseOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCreateCaseOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCreateCaseOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCreateCaseOptions",
	}
}
