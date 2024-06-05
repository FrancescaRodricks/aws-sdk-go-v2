// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an orchestration of data collection rules. The Amazon Web Services IoT
// FleetWise Edge Agent software running in vehicles uses campaigns to decide how
// to collect and transfer data to the cloud. You create campaigns in the cloud.
// After you or your team approve campaigns, Amazon Web Services IoT FleetWise
// automatically deploys them to vehicles.
//
// For more information, see [Collect and transfer data with campaigns] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// [Collect and transfer data with campaigns]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/campaigns.html
func (c *Client) CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) {
	if params == nil {
		params = &CreateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCampaign", params, optFns, c.addOperationCreateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCampaignInput struct {

	//  The data collection scheme associated with the campaign. You can specify a
	// scheme that collects data based on time or an event.
	//
	// This member is required.
	CollectionScheme types.CollectionScheme

	//  The name of the campaign to create.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the signal catalog to associate with the
	// campaign.
	//
	// This member is required.
	SignalCatalogArn *string

	//  The ARN of the vehicle or fleet to deploy a campaign to.
	//
	// This member is required.
	TargetArn *string

	//  (Optional) Whether to compress signals before transmitting data to Amazon Web
	// Services IoT FleetWise. If you don't want to compress the signals, use OFF . If
	// it's not specified, SNAPPY is used.
	//
	// Default: SNAPPY
	Compression types.Compression

	// The destination where the campaign sends data. You can choose to send data to
	// be stored in Amazon S3 or Amazon Timestream.
	//
	// Amazon S3 optimizes the cost of data storage and provides additional mechanisms
	// to use vehicle data, such as data lakes, centralized data storage, data
	// processing pipelines, and analytics. Amazon Web Services IoT FleetWise supports
	// at-least-once file delivery to S3. Your vehicle data is stored on multiple
	// Amazon Web Services IoT FleetWise servers for redundancy and high availability.
	//
	// You can use Amazon Timestream to access and analyze time series data, and
	// Timestream to query vehicle data so that you can identify trends and patterns.
	DataDestinationConfigs []types.DataDestinationConfig

	//  (Optional) A list of vehicle attributes to associate with a campaign.
	//
	// Enrich the data with specified vehicle attributes. For example, add make and
	// model to the campaign, and Amazon Web Services IoT FleetWise will associate the
	// data with those attributes as dimensions in Amazon Timestream. You can then
	// query the data against make and model .
	//
	// Default: An empty array
	DataExtraDimensions []string

	// An optional description of the campaign to help identify its purpose.
	Description *string

	//  (Optional) Option for a vehicle to send diagnostic trouble codes to Amazon Web
	// Services IoT FleetWise. If you want to send diagnostic trouble codes, use
	// SEND_ACTIVE_DTCS . If it's not specified, OFF is used.
	//
	// Default: OFF
	DiagnosticsMode types.DiagnosticsMode

	//  (Optional) The time the campaign expires, in seconds since epoch (January 1,
	// 1970 at midnight UTC time). Vehicle data isn't collected after the campaign
	// expires.
	//
	// Default: 253402214400 (December 31, 9999, 00:00:00 UTC)
	ExpiryTime *time.Time

	//  (Optional) How long (in milliseconds) to collect raw data after a triggering
	// event initiates the collection. If it's not specified, 0 is used.
	//
	// Default: 0
	PostTriggerCollectionDuration *int64

	// (Optional) A number indicating the priority of one campaign over another
	// campaign for a certain vehicle or fleet. A campaign with the lowest value is
	// deployed to vehicles before any other campaigns. If it's not specified, 0 is
	// used.
	//
	// Default: 0
	Priority *int32

	// (Optional) A list of information about signals to collect.
	SignalsToCollect []types.SignalInformation

	// (Optional) Whether to store collected data after a vehicle lost a connection
	// with the cloud. After a connection is re-established, the data is automatically
	// forwarded to Amazon Web Services IoT FleetWise. If you want to store collected
	// data when a vehicle loses connection with the cloud, use TO_DISK . If it's not
	// specified, OFF is used.
	//
	// Default: OFF
	SpoolingMode types.SpoolingMode

	// (Optional) The time, in milliseconds, to deliver a campaign after it was
	// approved. If it's not specified, 0 is used.
	//
	// Default: 0
	StartTime *time.Time

	// Metadata that can be used to manage the campaign.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCampaignOutput struct {

	//  The ARN of the created campaign.
	Arn *string

	// The name of the created campaign.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCampaign"); err != nil {
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
	if err = addOpCreateCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCampaign",
	}
}
