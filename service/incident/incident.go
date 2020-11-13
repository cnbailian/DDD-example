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
	"fmt"
	"time"

	"github.com/cnbailian/DDD-example/domain/carrier"
	"github.com/cnbailian/DDD-example/domain/handling"
	"github.com/cnbailian/DDD-example/repository/cargo"
	carrierrepo "github.com/cnbailian/DDD-example/repository/carrier"
	handlingrepo "github.com/cnbailian/DDD-example/repository/handling"
	"github.com/cnbailian/DDD-example/repository/location"
)

func HandlingEvent(payload *HandlingEventPayload) (*handling.Event, error) {
	c, err := cargo.FindByTrackingID(payload.CargoID)
	if err != nil {
		return nil, fmt.Errorf("find cargo: %w", err)
	}
	movement, err := carrierrepo.FindByScheduleID(payload.ScheduleID)
	if err != nil {
		return nil, fmt.Errorf("find carrier movement: %w", err)
	}
	completion, err := time.Parse(time.RFC3339, payload.Time)
	if err != nil {
		return nil, fmt.Errorf("parse time: %w", err)
	}
	event := handling.NewHandlingEvent(c, handling.EventType(payload.EventType), completion, movement)
	if err := handlingrepo.CreateHandlingEvent(event); err != nil {
		return nil, fmt.Errorf("create handling event: %w", err)
	}
	return event, nil
}

func CreateCarrierMovement(payload *CarrierMovementPayload) (*carrier.Movement, error) {
	from, err := location.FindByPortCode(payload.From)
	if err != nil {
		return nil, fmt.Errorf("find from location: %w", err)
	}
	to, err := location.FindByPortCode(payload.To)
	if err != nil {
		return nil, fmt.Errorf("find to location: %w", err)
	}
	movement := carrier.NewCarrierMovement(from, to)
	if err := carrierrepo.CreateCarrierMovement(movement); err != nil {
		return nil, fmt.Errorf("create carrier movement: %w", err)
	}
	return movement, nil
}
