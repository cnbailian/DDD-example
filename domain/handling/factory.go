/*
Copyright 2020 BaiLian.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handling

import (
	"time"

	"github.com/cnbailian/DDD-example/domain/cargo"
	"github.com/cnbailian/DDD-example/domain/carrier"
)

func NewHandlingEvent(c *cargo.Cargo, event EventType, completion time.Time, movement *carrier.Movement) *Event {
	return &Event{
		cargo:          c,
		CargoID:        c.TrackingID,
		Type:           event,
		CompletionTime: completion,
		movement:       movement,
		MovementID:     movement.ScheduleID,
	}
}
