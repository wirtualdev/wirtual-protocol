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

package sip

import (
	"time"

	"github.com/wirtualdev/wirtual-protocol/auth"
	"github.com/wirtualdev/wirtual-protocol/wirtual"
)

type SIPTokenParams struct {
	APIKey                string
	APISecret             string
	RoomName              string
	ParticipantIdentity   string
	ParticipantName       string
	ParticipantMetadata   string
	ParticipantAttributes map[string]string
	RoomPreset            string
	RoomConfig            *wirtual.RoomConfiguration
}

func BuildSIPToken(params SIPTokenParams) (string, error) {
	t := true
	at := auth.NewAccessToken(params.APIKey, params.APISecret).
		SetVideoGrant(&auth.VideoGrant{
			RoomJoin:             true,
			Room:                 params.RoomName,
			CanSubscribe:         &t,
			CanPublish:           &t,
			CanPublishData:       &t,
			CanUpdateOwnMetadata: &t,
		}).
		SetIdentity(params.ParticipantIdentity).
		SetName(params.ParticipantName).
		SetMetadata(params.ParticipantMetadata).
		SetAttributes(params.ParticipantAttributes).
		SetRoomPreset(params.RoomPreset).
		SetRoomConfig(params.RoomConfig).
		SetKind(wirtual.ParticipantInfo_SIP).
		SetValidFor(24 * time.Hour)

	return at.ToJWT()
}
