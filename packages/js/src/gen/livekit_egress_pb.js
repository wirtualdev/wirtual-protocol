// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-es v1.7.2 with parameter "target=js+dts"
// @generated from file livekit_egress.proto (package livekit, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { AudioCodec, ImageCodec, VideoCodec } from "./livekit_models_pb.js";

/**
 * @generated from enum livekit.EncodedFileType
 */
export const EncodedFileType = proto3.makeEnum(
  "livekit.EncodedFileType",
  [
    {no: 0, name: "DEFAULT_FILETYPE"},
    {no: 1, name: "MP4"},
    {no: 2, name: "OGG"},
  ],
);

/**
 * @generated from enum livekit.SegmentedFileProtocol
 */
export const SegmentedFileProtocol = proto3.makeEnum(
  "livekit.SegmentedFileProtocol",
  [
    {no: 0, name: "DEFAULT_SEGMENTED_FILE_PROTOCOL"},
    {no: 1, name: "HLS_PROTOCOL"},
  ],
);

/**
 * @generated from enum livekit.SegmentedFileSuffix
 */
export const SegmentedFileSuffix = proto3.makeEnum(
  "livekit.SegmentedFileSuffix",
  [
    {no: 0, name: "INDEX"},
    {no: 1, name: "TIMESTAMP"},
  ],
);

/**
 * @generated from enum livekit.ImageFileSuffix
 */
export const ImageFileSuffix = proto3.makeEnum(
  "livekit.ImageFileSuffix",
  [
    {no: 0, name: "IMAGE_SUFFIX_INDEX"},
    {no: 1, name: "IMAGE_SUFFIX_TIMESTAMP"},
  ],
);

/**
 * @generated from enum livekit.StreamProtocol
 */
export const StreamProtocol = proto3.makeEnum(
  "livekit.StreamProtocol",
  [
    {no: 0, name: "DEFAULT_PROTOCOL"},
    {no: 1, name: "RTMP"},
  ],
);

/**
 * @generated from enum livekit.EncodingOptionsPreset
 */
export const EncodingOptionsPreset = proto3.makeEnum(
  "livekit.EncodingOptionsPreset",
  [
    {no: 0, name: "H264_720P_30"},
    {no: 1, name: "H264_720P_60"},
    {no: 2, name: "H264_1080P_30"},
    {no: 3, name: "H264_1080P_60"},
    {no: 4, name: "PORTRAIT_H264_720P_30"},
    {no: 5, name: "PORTRAIT_H264_720P_60"},
    {no: 6, name: "PORTRAIT_H264_1080P_30"},
    {no: 7, name: "PORTRAIT_H264_1080P_60"},
  ],
);

/**
 * @generated from enum livekit.EgressStatus
 */
export const EgressStatus = proto3.makeEnum(
  "livekit.EgressStatus",
  [
    {no: 0, name: "EGRESS_STARTING"},
    {no: 1, name: "EGRESS_ACTIVE"},
    {no: 2, name: "EGRESS_ENDING"},
    {no: 3, name: "EGRESS_COMPLETE"},
    {no: 4, name: "EGRESS_FAILED"},
    {no: 5, name: "EGRESS_ABORTED"},
    {no: 6, name: "EGRESS_LIMIT_REACHED"},
  ],
);

/**
 * composite using a web browser
 *
 * @generated from message livekit.RoomCompositeEgressRequest
 */
export const RoomCompositeEgressRequest = proto3.makeMessageType(
  "livekit.RoomCompositeEgressRequest",
  () => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "layout", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "audio_only", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "video_only", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "custom_base_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "file", kind: "message", T: EncodedFileOutput, oneof: "output" },
    { no: 7, name: "stream", kind: "message", T: StreamOutput, oneof: "output" },
    { no: 10, name: "segments", kind: "message", T: SegmentedFileOutput, oneof: "output" },
    { no: 8, name: "preset", kind: "enum", T: proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
    { no: 9, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
    { no: 11, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
    { no: 12, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
    { no: 13, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
    { no: 14, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true },
  ],
);

/**
 * record any website
 *
 * @generated from message livekit.WebEgressRequest
 */
export const WebEgressRequest = proto3.makeMessageType(
  "livekit.WebEgressRequest",
  () => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "audio_only", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "video_only", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 12, name: "await_start_signal", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "file", kind: "message", T: EncodedFileOutput, oneof: "output" },
    { no: 5, name: "stream", kind: "message", T: StreamOutput, oneof: "output" },
    { no: 6, name: "segments", kind: "message", T: SegmentedFileOutput, oneof: "output" },
    { no: 7, name: "preset", kind: "enum", T: proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
    { no: 8, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
    { no: 9, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
    { no: 10, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
    { no: 11, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
    { no: 13, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true },
  ],
);

/**
 * record audio and video from a single participant
 *
 * @generated from message livekit.ParticipantEgressRequest
 */
export const ParticipantEgressRequest = proto3.makeMessageType(
  "livekit.ParticipantEgressRequest",
  () => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "identity", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "screen_share", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "preset", kind: "enum", T: proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
    { no: 5, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
    { no: 6, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
    { no: 7, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
    { no: 8, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
    { no: 9, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true },
  ],
);

/**
 * containerize up to one audio and one video track
 *
 * @generated from message livekit.TrackCompositeEgressRequest
 */
export const TrackCompositeEgressRequest = proto3.makeMessageType(
  "livekit.TrackCompositeEgressRequest",
  () => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "audio_track_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "video_track_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "file", kind: "message", T: EncodedFileOutput, oneof: "output" },
    { no: 5, name: "stream", kind: "message", T: StreamOutput, oneof: "output" },
    { no: 8, name: "segments", kind: "message", T: SegmentedFileOutput, oneof: "output" },
    { no: 6, name: "preset", kind: "enum", T: proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
    { no: 7, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
    { no: 11, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
    { no: 12, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
    { no: 13, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
    { no: 14, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true },
  ],
);

/**
 * record tracks individually, without transcoding
 *
 * @generated from message livekit.TrackEgressRequest
 */
export const TrackEgressRequest = proto3.makeMessageType(
  "livekit.TrackEgressRequest",
  () => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "track_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "file", kind: "message", T: DirectFileOutput, oneof: "output" },
    { no: 4, name: "websocket_url", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "output" },
  ],
);

/**
 * @generated from message livekit.EncodedFileOutput
 */
export const EncodedFileOutput = proto3.makeMessageType(
  "livekit.EncodedFileOutput",
  () => [
    { no: 1, name: "file_type", kind: "enum", T: proto3.getEnumType(EncodedFileType) },
    { no: 2, name: "filepath", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "disable_manifest", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
    { no: 4, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
    { no: 5, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
    { no: 7, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" },
  ],
);

/**
 * Used to generate HLS segments or other kind of segmented output
 *
 * @generated from message livekit.SegmentedFileOutput
 */
export const SegmentedFileOutput = proto3.makeMessageType(
  "livekit.SegmentedFileOutput",
  () => [
    { no: 1, name: "protocol", kind: "enum", T: proto3.getEnumType(SegmentedFileProtocol) },
    { no: 2, name: "filename_prefix", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "playlist_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "live_playlist_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "segment_duration", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 10, name: "filename_suffix", kind: "enum", T: proto3.getEnumType(SegmentedFileSuffix) },
    { no: 8, name: "disable_manifest", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
    { no: 6, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
    { no: 7, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
    { no: 9, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" },
  ],
);

/**
 * @generated from message livekit.DirectFileOutput
 */
export const DirectFileOutput = proto3.makeMessageType(
  "livekit.DirectFileOutput",
  () => [
    { no: 1, name: "filepath", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "disable_manifest", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
    { no: 3, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
    { no: 4, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
    { no: 6, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" },
  ],
);

/**
 * @generated from message livekit.ImageOutput
 */
export const ImageOutput = proto3.makeMessageType(
  "livekit.ImageOutput",
  () => [
    { no: 1, name: "capture_interval", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "width", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "height", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "filename_prefix", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "filename_suffix", kind: "enum", T: proto3.getEnumType(ImageFileSuffix) },
    { no: 6, name: "image_codec", kind: "enum", T: proto3.getEnumType(ImageCodec) },
    { no: 7, name: "disable_manifest", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 8, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
    { no: 9, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
    { no: 10, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
    { no: 11, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" },
  ],
);

/**
 * @generated from message livekit.S3Upload
 */
export const S3Upload = proto3.makeMessageType(
  "livekit.S3Upload",
  () => [
    { no: 1, name: "access_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "secret", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "bucket", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "force_path_style", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "metadata", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 8, name: "tagging", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "content_disposition", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.GCPUpload
 */
export const GCPUpload = proto3.makeMessageType(
  "livekit.GCPUpload",
  () => [
    { no: 1, name: "credentials", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "bucket", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.AzureBlobUpload
 */
export const AzureBlobUpload = proto3.makeMessageType(
  "livekit.AzureBlobUpload",
  () => [
    { no: 1, name: "account_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "account_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "container_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.AliOSSUpload
 */
export const AliOSSUpload = proto3.makeMessageType(
  "livekit.AliOSSUpload",
  () => [
    { no: 1, name: "access_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "secret", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "bucket", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.StreamOutput
 */
export const StreamOutput = proto3.makeMessageType(
  "livekit.StreamOutput",
  () => [
    { no: 1, name: "protocol", kind: "enum", T: proto3.getEnumType(StreamProtocol) },
    { no: 2, name: "urls", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message livekit.EncodingOptions
 */
export const EncodingOptions = proto3.makeMessageType(
  "livekit.EncodingOptions",
  () => [
    { no: 1, name: "width", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "height", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "depth", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "framerate", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "audio_codec", kind: "enum", T: proto3.getEnumType(AudioCodec) },
    { no: 6, name: "audio_bitrate", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "audio_quality", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "audio_frequency", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "video_codec", kind: "enum", T: proto3.getEnumType(VideoCodec) },
    { no: 9, name: "video_bitrate", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 12, name: "video_quality", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "key_frame_interval", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * @generated from message livekit.UpdateLayoutRequest
 */
export const UpdateLayoutRequest = proto3.makeMessageType(
  "livekit.UpdateLayoutRequest",
  () => [
    { no: 1, name: "egress_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "layout", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.UpdateStreamRequest
 */
export const UpdateStreamRequest = proto3.makeMessageType(
  "livekit.UpdateStreamRequest",
  () => [
    { no: 1, name: "egress_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "add_output_urls", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "remove_output_urls", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message livekit.ListEgressRequest
 */
export const ListEgressRequest = proto3.makeMessageType(
  "livekit.ListEgressRequest",
  () => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "egress_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "active", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message livekit.ListEgressResponse
 */
export const ListEgressResponse = proto3.makeMessageType(
  "livekit.ListEgressResponse",
  () => [
    { no: 1, name: "items", kind: "message", T: EgressInfo, repeated: true },
  ],
);

/**
 * @generated from message livekit.StopEgressRequest
 */
export const StopEgressRequest = proto3.makeMessageType(
  "livekit.StopEgressRequest",
  () => [
    { no: 1, name: "egress_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.EgressInfo
 */
export const EgressInfo = proto3.makeMessageType(
  "livekit.EgressInfo",
  () => [
    { no: 1, name: "egress_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(EgressStatus) },
    { no: 10, name: "started_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 11, name: "ended_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 18, name: "updated_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 9, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "room_composite", kind: "message", T: RoomCompositeEgressRequest, oneof: "request" },
    { no: 14, name: "web", kind: "message", T: WebEgressRequest, oneof: "request" },
    { no: 19, name: "participant", kind: "message", T: ParticipantEgressRequest, oneof: "request" },
    { no: 5, name: "track_composite", kind: "message", T: TrackCompositeEgressRequest, oneof: "request" },
    { no: 6, name: "track", kind: "message", T: TrackEgressRequest, oneof: "request" },
    { no: 7, name: "stream", kind: "message", T: StreamInfoList, oneof: "result" },
    { no: 8, name: "file", kind: "message", T: FileInfo, oneof: "result" },
    { no: 12, name: "segments", kind: "message", T: SegmentsInfo, oneof: "result" },
    { no: 15, name: "stream_results", kind: "message", T: StreamInfo, repeated: true },
    { no: 16, name: "file_results", kind: "message", T: FileInfo, repeated: true },
    { no: 17, name: "segment_results", kind: "message", T: SegmentsInfo, repeated: true },
    { no: 20, name: "image_results", kind: "message", T: ImagesInfo, repeated: true },
  ],
);

/**
 * @generated from message livekit.StreamInfoList
 * @deprecated
 */
export const StreamInfoList = proto3.makeMessageType(
  "livekit.StreamInfoList",
  () => [
    { no: 1, name: "info", kind: "message", T: StreamInfo, repeated: true },
  ],
);

/**
 * @generated from message livekit.StreamInfo
 */
export const StreamInfo = proto3.makeMessageType(
  "livekit.StreamInfo",
  () => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "started_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "ended_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "duration", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "status", kind: "enum", T: proto3.getEnumType(StreamInfo_Status) },
    { no: 6, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from enum livekit.StreamInfo.Status
 */
export const StreamInfo_Status = proto3.makeEnum(
  "livekit.StreamInfo.Status",
  [
    {no: 0, name: "ACTIVE"},
    {no: 1, name: "FINISHED"},
    {no: 2, name: "FAILED"},
  ],
);

/**
 * @generated from message livekit.FileInfo
 */
export const FileInfo = proto3.makeMessageType(
  "livekit.FileInfo",
  () => [
    { no: 1, name: "filename", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "started_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "ended_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "duration", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "size", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message livekit.SegmentsInfo
 */
export const SegmentsInfo = proto3.makeMessageType(
  "livekit.SegmentsInfo",
  () => [
    { no: 1, name: "playlist_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "live_playlist_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "duration", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "size", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "playlist_location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "live_playlist_location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "segment_count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "started_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 7, name: "ended_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message livekit.ImagesInfo
 */
export const ImagesInfo = proto3.makeMessageType(
  "livekit.ImagesInfo",
  () => [
    { no: 1, name: "image_count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "started_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "ended_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message livekit.AutoParticipantEgress
 */
export const AutoParticipantEgress = proto3.makeMessageType(
  "livekit.AutoParticipantEgress",
  () => [
    { no: 1, name: "preset", kind: "enum", T: proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
    { no: 2, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
    { no: 3, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
    { no: 4, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
  ],
);

/**
 * @generated from message livekit.AutoTrackEgress
 */
export const AutoTrackEgress = proto3.makeMessageType(
  "livekit.AutoTrackEgress",
  () => [
    { no: 1, name: "filepath", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "disable_manifest", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
    { no: 3, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
    { no: 4, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
  ],
);
