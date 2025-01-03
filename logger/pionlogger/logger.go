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

package pionlogger

import (
	"strings"

	"github.com/pion/logging"

	"github.com/wirtualdev/wirtual-protocol/logger"
)

var (
	pionIgnoredPrefixes = map[string]*prefixSet{
		"ice": {
			"pingAllCandidates called with no candidate pairs",
			"failed to send packet: io: read/write on closed pipe",
			"Ignoring remote candidate with tcpType active",
			"discard message from",
			"Failed to discover mDNS candidate",
			"Failed to read from candidate tcp",
			"remote mDNS candidate added, but mDNS is disabled",
		},
		"pc": {
			"Failed to accept RTCP stream is already closed",
			"Failed to accept RTP stream is already closed",
			"Incoming unhandled RTCP ssrc",
		},
		"tcp_mux": {
			"Error reading first packet from",
			"error closing connection",
		},
		"turn": {
			"error when handling datagram",
			"Failed to send ChannelData from allocation",
			"Failed to handle datagram",
		},
	}
)

type prefixSet []string

func (s *prefixSet) Match(msg string) bool {
	if s == nil {
		return false
	}

	for _, prefix := range *s {
		if strings.HasPrefix(msg, prefix) {
			return true
		}
	}
	return false
}

func NewLoggerFactory(l logger.Logger) logging.LoggerFactory {
	if zl, ok := l.(logger.ZapLogger); ok {
		return &zapLoggerFactory{zl}
	}
	return &loggerFactory{l}
}

// zapLoggerFactory implements logging.LoggerFactory interface for zap loggers
type zapLoggerFactory struct {
	logger logger.ZapLogger
}

func (f *zapLoggerFactory) NewLogger(scope string) logging.LeveledLogger {
	return &zapLogAdapter{
		logger:          f.logger,
		level:           f.logger.ComponentLeveler().ComponentLevel(formatComponent(scope)),
		scope:           scope,
		ignoredPrefixes: pionIgnoredPrefixes[scope],
	}
}

// loggerFactory implements logging.LoggerFactory interface for generic loggers
type loggerFactory struct {
	logger logger.Logger
}

func (f *loggerFactory) NewLogger(scope string) logging.LeveledLogger {
	return &logAdapter{
		logger:          f.logger.WithComponent(formatComponent(scope)).WithCallDepth(1),
		ignoredPrefixes: pionIgnoredPrefixes[scope],
	}
}

func formatComponent(scope string) string {
	return "pion." + scope
}
