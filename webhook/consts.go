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

package webhook

import "errors"

var (
	ErrNoAuthHeader    = errors.New("authorization header could not be found")
	ErrSecretNotFound  = errors.New("API secret could not be found")
	ErrInvalidChecksum = errors.New("could not verify authenticity of message")
)

const authHeader = "Authorization"

const (
	EventRoomStarted       = "room_started"
	EventRoomFinished      = "room_finished"
	EventParticipantJoined = "participant_joined"
	EventParticipantLeft   = "participant_left"
	EventTrackPublished    = "track_published"
	EventTrackUnpublished  = "track_unpublished"
	EventEgressStarted     = "egress_started"
	EventEgressUpdated     = "egress_updated"
	EventEgressEnded       = "egress_ended"
	EventIngressStarted    = "ingress_started"
	EventIngressEnded      = "ingress_ended"
)
