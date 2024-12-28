package egress

import (
	"testing"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
	"github.com/stretchr/testify/require"
)

func TestGetOutputType(t *testing.T) {
	roomReq := &wirtual.RoomCompositeEgressRequest{
		FileOutputs: []*wirtual.EncodedFileOutput{
			&wirtual.EncodedFileOutput{},
		},
	}

	ot := GetOutputType(roomReq)
	require.Equal(t, OutputTypeFile, ot)

	roomReq = &wirtual.RoomCompositeEgressRequest{
		Output: &wirtual.RoomCompositeEgressRequest_File{
			File: &wirtual.EncodedFileOutput{},
		},
	}

	ot = GetOutputType(roomReq)
	require.Equal(t, OutputTypeFile, ot)

	trackReq := &wirtual.TrackCompositeEgressRequest{
		SegmentOutputs: []*wirtual.SegmentedFileOutput{
			&wirtual.SegmentedFileOutput{},
		},
	}

	ot = GetOutputType(trackReq)
	require.Equal(t, OutputTypeSegments, ot)

	trackReq = &wirtual.TrackCompositeEgressRequest{
		Output: &wirtual.TrackCompositeEgressRequest_Segments{
			Segments: &wirtual.SegmentedFileOutput{},
		},
	}

	ot = GetOutputType(trackReq)
	require.Equal(t, OutputTypeSegments, ot)

	webReq := &wirtual.WebEgressRequest{
		StreamOutputs: []*wirtual.StreamOutput{
			&wirtual.StreamOutput{},
		},
	}

	ot = GetOutputType(webReq)
	require.Equal(t, OutputTypeStream, ot)

	webReq = &wirtual.WebEgressRequest{
		Output: &wirtual.WebEgressRequest_Stream{
			Stream: &wirtual.StreamOutput{},
		},
	}

	ot = GetOutputType(webReq)
	require.Equal(t, OutputTypeStream, ot)

	participantReq := &wirtual.ParticipantEgressRequest{
		ImageOutputs: []*wirtual.ImageOutput{
			&wirtual.ImageOutput{},
		},
	}

	ot = GetOutputType(participantReq)
	require.Equal(t, OutputTypeImages, ot)

	participantReq.SegmentOutputs = []*wirtual.SegmentedFileOutput{
		&wirtual.SegmentedFileOutput{},
	}

	ot = GetOutputType(participantReq)
	require.Equal(t, OutputTypeMultiple, ot)

}
