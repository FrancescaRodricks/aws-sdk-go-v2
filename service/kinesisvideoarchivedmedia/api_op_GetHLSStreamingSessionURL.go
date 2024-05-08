// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves an HTTP Live Streaming (HLS) URL for the stream. You can then open
// the URL in a browser or media player to view the stream contents.
//
// Both the StreamName and the StreamARN parameters are optional, but you must
// specify either the StreamName or the StreamARN when invoking this API operation.
//
// An Amazon Kinesis video stream has the following requirements for providing
// data through HLS:
//
//   - For streaming video, the media must contain H.264 or H.265 encoded video
//     and, optionally, AAC encoded audio. Specifically, the codec ID of track 1 should
//     be V_MPEG/ISO/AVC (for H.264) or V_MPEG/ISO/HEVC (for H.265). Optionally, the
//     codec ID of track 2 should be A_AAC . For audio only streaming, the codec ID
//     of track 1 should be A_AAC .
//
//   - Data retention must be greater than 0.
//
//   - The video track of each fragment must contain codec private data in the
//     Advanced Video Coding (AVC) for H.264 format or HEVC for H.265 format ([MPEG-4 specification ISO/IEC 14496-15] ). For
//     information about adapting stream data to a given format, see [NAL Adaptation Flags].
//
//   - The audio track (if present) of each fragment must contain codec private
//     data in the AAC format ([AAC specification ISO/IEC 13818-7] ).
//
// Kinesis Video Streams HLS sessions contain fragments in the fragmented MPEG-4
// form (also called fMP4 or CMAF) or the MPEG-2 form (also called TS chunks, which
// the HLS specification also supports). For more information about HLS fragment
// types, see the [HLS specification].
//
// The following procedure shows how to use HLS with Kinesis Video Streams:
//
//   - Get an endpoint using [GetDataEndpoint], specifying GET_HLS_STREAMING_SESSION_URL for the
//     APIName parameter.
//
//   - Retrieve the HLS URL using GetHLSStreamingSessionURL . Kinesis Video Streams
//     creates an HLS streaming session to be used for accessing content in a stream
//     using the HLS protocol. GetHLSStreamingSessionURL returns an authenticated URL
//     (that includes an encrypted session token) for the session's HLS master playlist
//     (the root resource needed for streaming with HLS).
//
// Don't share or store this token where an unauthorized entity could access it.
//
//	The token provides access to the content of the stream. Safeguard the token with
//	the same measures that you would use with your Amazon Web Services credentials.
//
// The media that is made available through the playlist consists only of the
//
//	requested stream, time range, and format. No other media data (such as frames
//	outside the requested window or alternate bitrates) is made available.
//
//	- Provide the URL (containing the encrypted session token) for the HLS master
//	playlist to a media player that supports the HLS protocol. Kinesis Video Streams
//	makes the HLS media playlist, initialization fragment, and media fragments
//	available through the master playlist URL. The initialization fragment contains
//	the codec private data for the stream, and other data needed to set up the video
//	or audio decoder and renderer. The media fragments contain H.264-encoded video
//	frames or AAC-encoded audio samples.
//
//	- The media player receives the authenticated URL and requests stream
//	metadata and media data normally. When the media player requests data, it calls
//	the following actions:
//
//	- GetHLSMasterPlaylist: Retrieves an HLS master playlist, which contains a
//	URL for the GetHLSMediaPlaylist action for each track, and additional metadata
//	for the media player, including estimated bitrate and resolution.
//
//	- GetHLSMediaPlaylist: Retrieves an HLS media playlist, which contains a URL
//	to access the MP4 initialization fragment with the GetMP4InitFragment action,
//	and URLs to access the MP4 media fragments with the GetMP4MediaFragment
//	actions. The HLS media playlist also contains metadata about the stream that the
//	player needs to play it, such as whether the PlaybackMode is LIVE or ON_DEMAND
//	. The HLS media playlist is typically static for sessions with a PlaybackType
//	of ON_DEMAND . The HLS media playlist is continually updated with new
//	fragments for sessions with a PlaybackType of LIVE . There is a distinct HLS
//	media playlist for the video track and the audio track (if applicable) that
//	contains MP4 media URLs for the specific track.
//
//	- GetMP4InitFragment: Retrieves the MP4 initialization fragment. The media
//	player typically loads the initialization fragment before loading any media
//	fragments. This fragment contains the " fytp " and " moov " MP4 atoms, and the
//	child atoms that are needed to initialize the media player decoder.
//
// The initialization fragment does not correspond to a fragment in a Kinesis
//
//	video stream. It contains only the codec private data for the stream and
//	respective track, which the media player needs to decode the media frames.
//
//	- GetMP4MediaFragment: Retrieves MP4 media fragments. These fragments contain
//	the " moof " and " mdat " MP4 atoms and their child atoms, containing the
//	encoded fragment's media frames and their timestamps.
//
// For the HLS streaming session, in-track codec private data (CPD) changes are
//
//	supported. After the first media fragment is made available in a streaming
//	session, fragments can contain CPD changes for each track. Therefore, the
//	fragments in a session can have a different resolution, bit rate, or other
//	information in the CPD without interrupting playback. However, any change made
//	in the track number or track codec format can return an error when those
//	different media fragments are loaded. For example, streaming will fail if the
//	fragments in the stream change from having only video to having both audio and
//	video, or if an AAC audio track is changed to an ALAW audio track. For each
//	streaming session, only 500 CPD changes are allowed.
//
// Data retrieved with this action is billable. For information, see [Pricing].
//
//   - GetTSFragment: Retrieves MPEG TS fragments containing both initialization
//     and media data for all tracks in the stream.
//
// If the ContainerFormat is MPEG_TS , this API is used instead of
//
//	GetMP4InitFragment and GetMP4MediaFragment to retrieve stream media.
//
// Data retrieved with this action is billable. For more information, see [Kinesis Video Streams pricing].
//
// A streaming session URL must not be shared between players. The service might
// throttle a session if multiple media players are sharing it. For connection
// limits, see [Kinesis Video Streams Limits].
//
// You can monitor the amount of data that the media player consumes by monitoring
// the GetMP4MediaFragment.OutgoingBytes Amazon CloudWatch metric. For information
// about using CloudWatch to monitor Kinesis Video Streams, see [Monitoring Kinesis Video Streams]. For pricing
// information, see [Amazon Kinesis Video Streams Pricing]and [Amazon Web Services Pricing]. Charges for both HLS sessions and outgoing Amazon Web
// Services data apply.
//
// For more information about HLS, see [HTTP Live Streaming] on the [Apple Developer site].
//
// If an error is thrown after invoking a Kinesis Video Streams archived media
// API, in addition to the HTTP status code and the response body, it includes the
// following pieces of information:
//
//   - x-amz-ErrorType HTTP header – contains a more specific error type in
//     addition to what the HTTP status code provides.
//
//   - x-amz-RequestId HTTP header – if you want to report an issue to Amazon Web
//     Services, the support team can better diagnose the problem if given the Request
//     Id.
//
// Both the HTTP status code and the ErrorType header can be utilized to make
// programmatic decisions about whether errors are retry-able and under what
// conditions, as well as provide information on what actions the client programmer
// might need to take in order to successfully try again.
//
// For more information, see the Errors section at the bottom of this topic, as
// well as [Common Errors].
//
// [Amazon Web Services Pricing]: https://aws.amazon.com/pricing/
// [Common Errors]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html
// [Kinesis Video Streams Limits]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/limits.html
// [MPEG-4 specification ISO/IEC 14496-15]: https://www.iso.org/standard/55980.html
// [HLS specification]: https://tools.ietf.org/html/draft-pantos-http-live-streaming-23
// [AAC specification ISO/IEC 13818-7]: https://www.iso.org/standard/43345.html
// [HTTP Live Streaming]: https://developer.apple.com/streaming/
// [NAL Adaptation Flags]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/producer-reference-nal.html
// [Apple Developer site]: https://developer.apple.com
// [Kinesis Video Streams pricing]: https://aws.amazon.com/kinesis/video-streams/pricing/
// [Pricing]: https://aws.amazon.com/kinesis/video-streams/pricing/
// [Amazon Kinesis Video Streams Pricing]: https://aws.amazon.com/kinesis/video-streams/pricing/
// [Monitoring Kinesis Video Streams]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/monitoring.html
// [GetDataEndpoint]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_GetDataEndpoint.html
func (c *Client) GetHLSStreamingSessionURL(ctx context.Context, params *GetHLSStreamingSessionURLInput, optFns ...func(*Options)) (*GetHLSStreamingSessionURLOutput, error) {
	if params == nil {
		params = &GetHLSStreamingSessionURLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHLSStreamingSessionURL", params, optFns, c.addOperationGetHLSStreamingSessionURLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHLSStreamingSessionURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHLSStreamingSessionURLInput struct {

	// Specifies which format should be used for packaging the media. Specifying the
	// FRAGMENTED_MP4 container format packages the media into MP4 fragments (fMP4 or
	// CMAF). This is the recommended packaging because there is minimal packaging
	// overhead. The other container format option is MPEG_TS . HLS has supported MPEG
	// TS chunks since it was released and is sometimes the only supported packaging on
	// older HLS players. MPEG TS typically has a 5-25 percent packaging overhead. This
	// means MPEG TS typically requires 5-25 percent more bandwidth and cost than fMP4.
	//
	// The default is FRAGMENTED_MP4 .
	ContainerFormat types.ContainerFormat

	// Specifies when flags marking discontinuities between fragments are added to the
	// media playlists.
	//
	// Media players typically build a timeline of media content to play, based on the
	// timestamps of each fragment. This means that if there is any overlap or gap
	// between fragments (as is typical if HLSFragmentSelectoris set to SERVER_TIMESTAMP ), the media
	// player timeline will also have small gaps between fragments in some places, and
	// will overwrite frames in other places. Gaps in the media player timeline can
	// cause playback to stall and overlaps can cause playback to be jittery. When
	// there are discontinuity flags between fragments, the media player is expected to
	// reset the timeline, resulting in the next fragment being played immediately
	// after the previous fragment.
	//
	// The following modes are supported:
	//
	//   - ALWAYS : a discontinuity marker is placed between every fragment in the HLS
	//   media playlist. It is recommended to use a value of ALWAYS if the fragment
	//   timestamps are not accurate.
	//
	//   - NEVER : no discontinuity markers are placed anywhere. It is recommended to
	//   use a value of NEVER to ensure the media player timeline most accurately maps
	//   to the producer timestamps.
	//
	//   - ON_DISCONTINUITY : a discontinuity marker is placed between fragments that
	//   have a gap or overlap of more than 50 milliseconds. For most playback scenarios,
	//   it is recommended to use a value of ON_DISCONTINUITY so that the media player
	//   timeline is only reset when there is a significant issue with the media timeline
	//   (e.g. a missing fragment).
	//
	// The default is ALWAYS when HLSFragmentSelector is set to SERVER_TIMESTAMP , and NEVER when it is
	// set to PRODUCER_TIMESTAMP .
	DiscontinuityMode types.HLSDiscontinuityMode

	// Specifies when the fragment start timestamps should be included in the HLS
	// media playlist. Typically, media players report the playhead position as a time
	// relative to the start of the first fragment in the playback session. However,
	// when the start timestamps are included in the HLS media playlist, some media
	// players might report the current playhead as an absolute time based on the
	// fragment timestamps. This can be useful for creating a playback experience that
	// shows viewers the wall-clock time of the media.
	//
	// The default is NEVER . When HLSFragmentSelector is SERVER_TIMESTAMP , the timestamps will be the
	// server start timestamps. Similarly, when HLSFragmentSelectoris PRODUCER_TIMESTAMP , the timestamps
	// will be the producer start timestamps.
	DisplayFragmentTimestamp types.HLSDisplayFragmentTimestamp

	// The time in seconds until the requested session expires. This value can be
	// between 300 (5 minutes) and 43200 (12 hours).
	//
	// When a session expires, no new calls to GetHLSMasterPlaylist ,
	// GetHLSMediaPlaylist , GetMP4InitFragment , GetMP4MediaFragment , or
	// GetTSFragment can be made for that session.
	//
	// The default is 300 (5 minutes).
	Expires *int32

	// The time range of the requested fragment and the source of the timestamps.
	//
	// This parameter is required if PlaybackMode is ON_DEMAND or LIVE_REPLAY . This
	// parameter is optional if PlaybackMode is LIVE . If PlaybackMode is LIVE , the
	// FragmentSelectorType can be set, but the TimestampRange should not be set. If
	// PlaybackMode is ON_DEMAND or LIVE_REPLAY , both FragmentSelectorType and
	// TimestampRange must be set.
	HLSFragmentSelector *types.HLSFragmentSelector

	// The maximum number of fragments that are returned in the HLS media playlists.
	//
	// When the PlaybackMode is LIVE , the most recent fragments are returned up to
	// this value. When the PlaybackMode is ON_DEMAND , the oldest fragments are
	// returned, up to this maximum number.
	//
	// When there are a higher number of fragments available in a live HLS media
	// playlist, video players often buffer content before starting playback.
	// Increasing the buffer size increases the playback latency, but it decreases the
	// likelihood that rebuffering will occur during playback. We recommend that a live
	// HLS media playlist have a minimum of 3 fragments and a maximum of 10 fragments.
	//
	// The default is 5 fragments if PlaybackMode is LIVE or LIVE_REPLAY , and 1,000 if
	// PlaybackMode is ON_DEMAND .
	//
	// The maximum value of 5,000 fragments corresponds to more than 80 minutes of
	// video on streams with 1-second fragments, and more than 13 hours of video on
	// streams with 10-second fragments.
	MaxMediaPlaylistFragmentResults *int64

	// Whether to retrieve live, live replay, or archived, on-demand data.
	//
	// Features of the three types of sessions include the following:
	//
	//   - LIVE : For sessions of this type, the HLS media playlist is continually
	//   updated with the latest fragments as they become available. We recommend that
	//   the media player retrieve a new playlist on a one-second interval. When this
	//   type of session is played in a media player, the user interface typically
	//   displays a "live" notification, with no scrubber control for choosing the
	//   position in the playback window to display.
	//
	// In LIVE mode, the newest available fragments are included in an HLS media
	//   playlist, even if there is a gap between fragments (that is, if a fragment is
	//   missing). A gap like this might cause a media player to halt or cause a jump in
	//   playback. In this mode, fragments are not added to the HLS media playlist if
	//   they are older than the newest fragment in the playlist. If the missing fragment
	//   becomes available after a subsequent fragment is added to the playlist, the
	//   older fragment is not added, and the gap is not filled.
	//
	//   - LIVE_REPLAY : For sessions of this type, the HLS media playlist is updated
	//   similarly to how it is updated for LIVE mode except that it starts by
	//   including fragments from a given start time. Instead of fragments being added as
	//   they are ingested, fragments are added as the duration of the next fragment
	//   elapses. For example, if the fragments in the session are two seconds long, then
	//   a new fragment is added to the media playlist every two seconds. This mode is
	//   useful to be able to start playback from when an event is detected and continue
	//   live streaming media that has not yet been ingested as of the time of the
	//   session creation. This mode is also useful to stream previously archived media
	//   without being limited by the 1,000 fragment limit in the ON_DEMAND mode.
	//
	//   - ON_DEMAND : For sessions of this type, the HLS media playlist contains all
	//   the fragments for the session, up to the number that is specified in
	//   MaxMediaPlaylistFragmentResults . The playlist must be retrieved only once for
	//   each session. When this type of session is played in a media player, the user
	//   interface typically displays a scrubber control for choosing the position in the
	//   playback window to display.
	//
	// In all playback modes, if FragmentSelectorType is PRODUCER_TIMESTAMP , and if
	// there are multiple fragments with the same start timestamp, the fragment that
	// has the largest fragment number (that is, the newest fragment) is included in
	// the HLS media playlist. The other fragments are not included. Fragments that
	// have different timestamps but have overlapping durations are still included in
	// the HLS media playlist. This can lead to unexpected behavior in the media
	// player.
	//
	// The default is LIVE .
	PlaybackMode types.HLSPlaybackMode

	// The Amazon Resource Name (ARN) of the stream for which to retrieve the HLS
	// master playlist URL.
	//
	// You must specify either the StreamName or the StreamARN .
	StreamARN *string

	// The name of the stream for which to retrieve the HLS master playlist URL.
	//
	// You must specify either the StreamName or the StreamARN .
	StreamName *string

	noSmithyDocumentSerde
}

type GetHLSStreamingSessionURLOutput struct {

	// The URL (containing the session token) that a media player can use to retrieve
	// the HLS master playlist.
	HLSStreamingSessionURL *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHLSStreamingSessionURLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetHLSStreamingSessionURL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetHLSStreamingSessionURL{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetHLSStreamingSessionURL"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHLSStreamingSessionURL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHLSStreamingSessionURL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetHLSStreamingSessionURL",
	}
}
