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

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wirtualdev/wirtual-protocol/auth"
	"github.com/wirtualdev/wirtual-protocol/wirtual"
)

// Receive reads and verifies incoming webhook is signed with key/secret pair
// closes body after reading
func Receive(r *http.Request, provider auth.KeyProvider) ([]byte, error) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	authToken := r.Header.Get(authHeader)
	if authToken == "" {
		return nil, ErrNoAuthHeader
	}

	v, err := auth.ParseAPIToken(authToken)
	if err != nil {
		return nil, err
	}

	secret := provider.GetSecret(v.APIKey())
	if secret == "" {
		return nil, ErrSecretNotFound
	}

	claims, err := v.Verify(secret)
	if err != nil {
		return nil, err
	}

	// verify checksum
	sha := sha256.Sum256(data)
	hash := base64.StdEncoding.EncodeToString(sha[:])

	if claims.Sha256 != hash {
		return nil, ErrInvalidChecksum
	}

	return data, nil
}

// ReceiveWebhookEvent reads and verifies incoming webhook, and returns a parsed WebhookEvent
func ReceiveWebhookEvent(r *http.Request, provider auth.KeyProvider) (*wirtual.WebhookEvent, error) {
	data, err := Receive(r, provider)
	if err != nil {
		return nil, err
	}
	unmarshalOpts := protojson.UnmarshalOptions{
		DiscardUnknown: true,
		AllowPartial:   true,
	}
	event := wirtual.WebhookEvent{}
	if err = unmarshalOpts.Unmarshal(data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
