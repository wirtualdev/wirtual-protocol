// Copyright 2024 Xtressials Corporation, Inc.
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

package ingress

import (
	"testing"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	info := &wirtual.IngressInfo{}

	err := Validate(info)
	require.Error(t, err)

	info.StreamKey = "stream_key"
	err = Validate(info)
	require.Error(t, err)

	info.RoomName = "room_name"
	err = Validate(info)
	require.Error(t, err)

	info.ParticipantIdentity = "participant_identity"
	err = Validate(info)
	require.NoError(t, err)

	// make sure video parameters are validated. Full validation logic tested in the next test
	info.Video = &wirtual.IngressVideoOptions{}
	err = Validate(info)
	require.NoError(t, err)

	info.Video.Source = wirtual.TrackSource_MICROPHONE
	err = Validate(info)
	require.Error(t, err)

	info.Video.Source = wirtual.TrackSource_CAMERA

	// make sure audio parameters are validated. Full validation logic tested in the next test
	info.Audio = &wirtual.IngressAudioOptions{}
	err = Validate(info)
	require.NoError(t, err)

	info.Audio.Source = wirtual.TrackSource_CAMERA
	err = Validate(info)
	require.Error(t, err)

	info.Audio.Source = wirtual.TrackSource_SCREEN_SHARE_AUDIO
	err = Validate(info)
	require.NoError(t, err)
}

func TestValidateBypassTranscoding(t *testing.T) {
	info := &wirtual.IngressInfo{}

	err := ValidateBypassTranscoding(info)
	require.NoError(t, err)

	info.BypassTranscoding = true
	err = ValidateBypassTranscoding(info)
	require.Error(t, err)

	info.InputType = wirtual.IngressInput_WHIP_INPUT
	err = ValidateBypassTranscoding(info)
	require.NoError(t, err)

	info.Video = &wirtual.IngressVideoOptions{}
	err = ValidateBypassTranscoding(info)
	require.NoError(t, err)

	info.Video.EncodingOptions = &wirtual.IngressVideoOptions_Preset{}
	err = ValidateBypassTranscoding(info)
	require.Error(t, err)

	info.Video = nil

	info.Audio = &wirtual.IngressAudioOptions{}
	err = ValidateBypassTranscoding(info)
	require.NoError(t, err)

	info.Audio.EncodingOptions = &wirtual.IngressAudioOptions_Options{
		Options: &wirtual.IngressAudioEncodingOptions{},
	}
	err = ValidateBypassTranscoding(info)
	require.Error(t, err)

}

func TestValidateEnableTranscoding(t *testing.T) {
	info := &wirtual.IngressInfo{}
	T := true
	F := false

	err := ValidateEnableTranscoding(info)
	require.NoError(t, err)

	info.InputType = wirtual.IngressInput_WHIP_INPUT
	err = ValidateEnableTranscoding(info)
	require.NoError(t, err)

	info.Audio = &wirtual.IngressAudioOptions{}
	info.Audio.EncodingOptions = &wirtual.IngressAudioOptions_Options{}
	err = ValidateEnableTranscoding(info)
	require.Error(t, err)

	info.Audio.EncodingOptions = nil

	info.EnableTranscoding = &T
	err = ValidateEnableTranscoding(info)
	require.NoError(t, err)

	info.EnableTranscoding = &F
	err = ValidateEnableTranscoding(info)
	require.NoError(t, err)

	info.Video = &wirtual.IngressVideoOptions{}
	info.Video.EncodingOptions = &wirtual.IngressVideoOptions_Preset{}
	err = ValidateEnableTranscoding(info)
	require.Error(t, err)

	info.Video.EncodingOptions = nil

	info.InputType = wirtual.IngressInput_RTMP_INPUT
	err = ValidateEnableTranscoding(info)
	require.Error(t, err)

	info.EnableTranscoding = &T
	err = ValidateEnableTranscoding(info)
	require.NoError(t, err)
}

func TestValidateVideoOptionsConsistency(t *testing.T) {
	video := &wirtual.IngressVideoOptions{}
	err := ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.Source = wirtual.TrackSource_MICROPHONE
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.Source = wirtual.TrackSource_CAMERA
	video.EncodingOptions = &wirtual.IngressVideoOptions_Preset{
		Preset: wirtual.IngressVideoEncodingPreset(42),
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions = &wirtual.IngressVideoOptions_Preset{
		Preset: wirtual.IngressVideoEncodingPreset_H264_1080P_30FPS_1_LAYER,
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions = &wirtual.IngressVideoOptions_Options{
		Options: &wirtual.IngressVideoEncodingOptions{
			VideoCodec: wirtual.VideoCodec_H264_HIGH,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions = &wirtual.IngressVideoOptions_Options{
		Options: &wirtual.IngressVideoEncodingOptions{
			VideoCodec: wirtual.VideoCodec_DEFAULT_VC,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{
			Width:  640,
			Height: 480,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{
			Width:  641,
			Height: 480,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: wirtual.VideoQuality_HIGH,
		},
		&wirtual.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: wirtual.VideoQuality_LOW,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: wirtual.VideoQuality_HIGH,
		},
		&wirtual.VideoLayer{
			Width:   1280,
			Height:  720,
			Quality: wirtual.VideoQuality_HIGH,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: wirtual.VideoQuality_HIGH,
		},
		&wirtual.VideoLayer{
			Width:   1280,
			Height:  720,
			Quality: wirtual.VideoQuality_LOW,
		},
	}
	err = ValidateVideoOptionsConsistency(video)
	require.Error(t, err)

	video.EncodingOptions.(*wirtual.IngressVideoOptions_Options).Options.Layers = []*wirtual.VideoLayer{
		&wirtual.VideoLayer{
			Width:   640,
			Height:  480,
			Quality: wirtual.VideoQuality_LOW,
		},
		&wirtual.VideoLayer{
			Width:   1280,
			Height:  720,
			Quality: wirtual.VideoQuality_HIGH,
		},
	}

	err = ValidateVideoOptionsConsistency(video)
	require.NoError(t, err)
}

func TestValidateAudioOptionsConsistency(t *testing.T) {
	audio := &wirtual.IngressAudioOptions{}
	err := ValidateAudioOptionsConsistency(audio)
	require.NoError(t, err)

	audio.Source = wirtual.TrackSource_CAMERA
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.Source = wirtual.TrackSource_SCREEN_SHARE_AUDIO
	audio.EncodingOptions = &wirtual.IngressAudioOptions_Preset{
		Preset: wirtual.IngressAudioEncodingPreset(42),
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.EncodingOptions = &wirtual.IngressAudioOptions_Preset{
		Preset: wirtual.IngressAudioEncodingPreset_OPUS_MONO_64KBS,
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.NoError(t, err)

	audio.EncodingOptions = &wirtual.IngressAudioOptions_Options{
		Options: &wirtual.IngressAudioEncodingOptions{
			AudioCodec: wirtual.AudioCodec_AAC,
		},
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.EncodingOptions = &wirtual.IngressAudioOptions_Options{
		Options: &wirtual.IngressAudioEncodingOptions{
			AudioCodec: wirtual.AudioCodec_OPUS,
			Channels:   3,
		},
	}
	err = ValidateAudioOptionsConsistency(audio)
	require.Error(t, err)

	audio.EncodingOptions.(*wirtual.IngressAudioOptions_Options).Options.Channels = 2
	err = ValidateAudioOptionsConsistency(audio)
	require.NoError(t, err)
}
