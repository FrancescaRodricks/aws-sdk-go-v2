// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deprecated. Use DescribeLoa instead.
//
// Gets the LOA-CFA for a connection.
//
// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA) is a
// document that your APN partner or service provider uses when establishing your
// cross connect to Amazon Web Services at the colocation facility. For more
// information, see [Requesting Cross Connects at Direct Connect Locations]in the Direct Connect User Guide.
//
// Deprecated: This operation has been deprecated.
//
// [Requesting Cross Connects at Direct Connect Locations]: https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html
func (c *Client) DescribeConnectionLoa(ctx context.Context, params *DescribeConnectionLoaInput, optFns ...func(*Options)) (*DescribeConnectionLoaOutput, error) {
	if params == nil {
		params = &DescribeConnectionLoaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectionLoa", params, optFns, c.addOperationDescribeConnectionLoaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectionLoaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectionLoaInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// The standard media type for the LOA-CFA document. The only supported value is
	// application/pdf.
	LoaContentType types.LoaContentType

	// The name of the APN partner or service provider who establishes connectivity on
	// your behalf. If you specify this parameter, the LOA-CFA lists the provider name
	// alongside your company name as the requester of the cross connect.
	ProviderName *string

	noSmithyDocumentSerde
}

type DescribeConnectionLoaOutput struct {

	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *types.Loa

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectionLoaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConnectionLoa{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConnectionLoa{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConnectionLoa"); err != nil {
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
	if err = addOpDescribeConnectionLoaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectionLoa(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnectionLoa(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConnectionLoa",
	}
}
