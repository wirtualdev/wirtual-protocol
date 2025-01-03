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

package mono

import "time"

var (
	epoch     = time.Now()
	epochNano = epoch.UnixNano()
)

func FromTime(t time.Time) time.Time {
	return epoch.Add(t.Sub(epoch))
}

func Now() time.Time {
	return epoch.Add(time.Since(epoch))
}

func UnixNano() int64 {
	return epochNano + int64(time.Since(epoch))
}

func UnixMicro() int64 {
	return UnixNano() / 1000
}
