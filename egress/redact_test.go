package egress

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
	"github.com/stretchr/testify/require"
)

var (
	file = &wirtual.EncodedFileOutput{
		Output: &wirtual.EncodedFileOutput_S3{
			S3: &wirtual.S3Upload{
				AccessKey: "ACCESS_KEY",
				Secret:    "LONG_SECRET_STRING",
			},
		},
	}

	image = &wirtual.ImageOutput{
		Output: &wirtual.ImageOutput_AliOSS{
			AliOSS: &wirtual.AliOSSUpload{
				AccessKey: "ACCESS_KEY",
				Secret:    "LONG_SECRET_STRING",
			},
		},
	}

	segments = &wirtual.SegmentedFileOutput{
		Output: &wirtual.SegmentedFileOutput_Gcp{
			Gcp: &wirtual.GCPUpload{
				Credentials: "CREDENTIALS",
			},
		},
	}

	directFile = &wirtual.DirectFileOutput{
		Output: &wirtual.DirectFileOutput_Azure{
			Azure: &wirtual.AzureBlobUpload{
				AccountName: "ACCOUNT_NAME",
				AccountKey:  "ACCOUNT_KEY",
			},
		},
	}
)

func TestRedactUpload(t *testing.T) {
	cl := proto.Clone(file)
	RedactUpload(cl.(UploadRequest))

	require.Equal(t, "{access_key}", cl.(*wirtual.EncodedFileOutput).Output.(*wirtual.EncodedFileOutput_S3).S3.AccessKey)
	require.Equal(t, "{secret}", cl.(*wirtual.EncodedFileOutput).Output.(*wirtual.EncodedFileOutput_S3).S3.Secret)

	cl = proto.Clone(image)
	RedactUpload(cl.(UploadRequest))

	require.Equal(t, "{access_key}", cl.(*wirtual.ImageOutput).Output.(*wirtual.ImageOutput_AliOSS).AliOSS.AccessKey)
	require.Equal(t, "{secret}", cl.(*wirtual.ImageOutput).Output.(*wirtual.ImageOutput_AliOSS).AliOSS.Secret)

	cl = proto.Clone(segments)
	RedactUpload(cl.(UploadRequest))

	require.Equal(t, "{credentials}", cl.(*wirtual.SegmentedFileOutput).Output.(*wirtual.SegmentedFileOutput_Gcp).Gcp.Credentials)

	cl = proto.Clone(directFile)
	RedactUpload(cl.(UploadRequest))

	require.Equal(t, "{account_name}", cl.(*wirtual.DirectFileOutput).Output.(*wirtual.DirectFileOutput_Azure).Azure.AccountName)
	require.Equal(t, "{account_key}", cl.(*wirtual.DirectFileOutput).Output.(*wirtual.DirectFileOutput_Azure).Azure.AccountKey)
}

func TestRedactStreamOutput(t *testing.T) {
	so := &wirtual.StreamOutput{
		Urls: []string{
			"rtmps://foo.bar.com/app/secret_stream_key",
		},
	}

	RedactStreamKeys(so)
	require.Equal(t, "rtmps://foo.bar.com/app/{sec...key}", so.Urls[0])
}

func TestRedactEncodedOutputs(t *testing.T) {
	trackComposite := &wirtual.TrackCompositeEgressRequest{
		FileOutputs: []*wirtual.EncodedFileOutput{
			file,
		},
	}

	cl := proto.Clone(trackComposite)
	RedactEncodedOutputs(cl.(EncodedOutput))

	require.Equal(t, "{access_key}", cl.(*wirtual.TrackCompositeEgressRequest).FileOutputs[0].Output.(*wirtual.EncodedFileOutput_S3).S3.AccessKey)
	require.Equal(t, "{secret}", cl.(*wirtual.TrackCompositeEgressRequest).FileOutputs[0].Output.(*wirtual.EncodedFileOutput_S3).S3.Secret)

	roomComposite := &wirtual.RoomCompositeEgressRequest{
		ImageOutputs: []*wirtual.ImageOutput{
			image,
		},
	}

	cl = proto.Clone(roomComposite)
	RedactEncodedOutputs(cl.(EncodedOutput))

	require.Equal(t, "{access_key}", cl.(*wirtual.RoomCompositeEgressRequest).ImageOutputs[0].Output.(*wirtual.ImageOutput_AliOSS).AliOSS.AccessKey)
	require.Equal(t, "{secret}", cl.(*wirtual.RoomCompositeEgressRequest).ImageOutputs[0].Output.(*wirtual.ImageOutput_AliOSS).AliOSS.Secret)

	participant := &wirtual.ParticipantEgressRequest{
		SegmentOutputs: []*wirtual.SegmentedFileOutput{
			segments,
		},
	}

	cl = proto.Clone(participant)
	RedactEncodedOutputs(cl.(EncodedOutput))

	require.Equal(t, "{credentials}", cl.(*wirtual.ParticipantEgressRequest).SegmentOutputs[0].Output.(*wirtual.SegmentedFileOutput_Gcp).Gcp.Credentials)
}

func TestRedactDirectOutput(t *testing.T) {
	track := &wirtual.TrackEgressRequest{
		Output: &wirtual.TrackEgressRequest_File{
			File: &wirtual.DirectFileOutput{
				Output: &wirtual.DirectFileOutput_S3{
					S3: &wirtual.S3Upload{
						AccessKey: "ACCESS_KEY",
						Secret:    "SECRET",
					},
				},
			},
		},
	}

	RedactDirectOutputs(track)
	require.Equal(t, "{access_key}", track.Output.(*wirtual.TrackEgressRequest_File).File.Output.(*wirtual.DirectFileOutput_S3).S3.AccessKey)
	require.Equal(t, "{secret}", track.Output.(*wirtual.TrackEgressRequest_File).File.Output.(*wirtual.DirectFileOutput_S3).S3.Secret)
}
