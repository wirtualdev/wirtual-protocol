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

package guid

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
)

func TestMarshalUnmarshal(t *testing.T) {
	id0 := wirtual.TrackID(New(TrackPrefix))
	b0 := Marshal(id0)
	id1 := Unmarshal[wirtual.TrackID](b0)
	b1 := Marshal(id1)
	require.EqualValues(t, id0, id1)
	require.EqualValues(t, b0, b1)
}

func TestMarshalAppend(t *testing.T) {
	id0 := wirtual.RoomID(New(RoomPrefix))
	id1 := wirtual.RoomID(New(RoomPrefix))
	var b []byte
	b = MarshalAppend(b, id0)
	b = MarshalAppend(b, id1)
	require.Equal(t, id0, Unmarshal[wirtual.RoomID](wirtual.GuidBlock(b[0:])))
	require.Equal(t, id1, Unmarshal[wirtual.RoomID](wirtual.GuidBlock(b[9:])))
}

func BenchmarkNew(b *testing.B) {
	b.Run("new", func(b *testing.B) {
		var guid string
		for i := 0; i < b.N; i++ {
			guid = New(TrackPrefix)
		}
		_ = guid
	})
}
