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

package incident

import (
	"testing"
	"time"

	"github.com/cnbailian/DDD-example/domain/handling"
	"github.com/cnbailian/DDD-example/service/booking"
)

func TestHandlingEvent(t *testing.T) {
	cargo, err := booking.CargoBooking(&booking.CargoBookingPayload{
		ShipperID:               "1",
		ConsigneeID:             "2",
		OriginLocationCode:      "010",
		DestinationLocationCode: "0081",
		ETA:                     "2020-11-03T10:00:00Z",
	})
	if err != nil {
		t.Error(err)
	}

	movement, err := CreateCarrierMovement(&CarrierMovementPayload{
		From: "010",
		To:   "022",
	})
	if err != nil {
		t.Error(err)
	}

	payload := &HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.SentEventType),
		Time:       time.Now().Format(time.RFC3339),
		ScheduleID: movement.ScheduleID,
	}
	if _, err := HandlingEvent(payload); err != nil {
		t.Error(err)
	}
}
