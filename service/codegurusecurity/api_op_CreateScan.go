// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurusecurity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use to create a scan using code uploaded to an S3 bucket.
func (c *Client) CreateScan(ctx context.Context, params *CreateScanInput, optFns ...func(*Options)) (*CreateScanOutput, error) {
	if params == nil {
		params = &CreateScanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateScan", params, optFns, c.addOperationCreateScanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScanInput struct {

	// The identifier for an input resource used to create a scan.
	//
	// This member is required.
	ResourceId types.ResourceId

	// The unique name that CodeGuru Security uses to track revisions across multiple
	// scans of the same resource. Only allowed for a STANDARD scan type. If not
	// specified, it will be auto generated.
	//
	// This member is required.
	ScanName *string

	// The type of analysis you want CodeGuru Security to perform in the scan, either
	// Security or All . The Secuirty type only generates findings related to
	// security. The All type generates both security findings and quality findings.
	// Defaults to Security type if missing.
	AnalysisType types.AnalysisType

	// The idempotency token for the request. Amazon CodeGuru Security uses this value
	// to prevent the accidental creation of duplicate scans if there are failures and
	// retries.
	ClientToken *string

	// The type of scan, either Standard or Express . Defaults to Standard type if
	// missing. Express scans run on limited resources and use a limited set of
	// detectors to analyze your code in near-real time. Standard scans have standard
	// resource limits and use the full set of detectors to analyze your code.
	ScanType types.ScanType

	// An array of key-value pairs used to tag a scan. A tag is a custom attribute
	// label with two parts:
	//   - A tag key. For example, CostCenter , Environment , or Secret . Tag keys are
	//   case sensitive.
	//   - An optional tag value field. For example, 111122223333 , Production , or a
	//   team name. Omitting the tag value is the same as using an empty string. Tag
	//   values are case sensitive.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateScanOutput struct {

	// The identifier for the resource object that contains resources that were
	// scanned.
	//
	// This member is required.
	ResourceId types.ResourceId

	// UUID that identifies the individual scan run.
	//
	// This member is required.
	RunId *string

	// The name of the scan.
	//
	// This member is required.
	ScanName *string

	// The current state of the scan. Returns either InProgress , Successful , or
	// Failed .
	//
	// This member is required.
	ScanState types.ScanState

	// The ARN for the scan name.
	ScanNameArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateScanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateScan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateScan{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateScanMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateScanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScan(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

type idempotencyToken_initializeOpCreateScan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateScan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateScan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateScanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateScanInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateScanMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateScan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateScan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-security",
		OperationName: "CreateScan",
	}
}
