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
	"errors"
	"time"

	"github.com/cnbailian/DDD-example/domain/handling"
)

var database []*handling.Event

// CargoID & Time & Type 组成 HandlingEvent 唯一标识
func FindByCargoIDTimeTypes(cargoID string, eventType handling.EventType, date time.Time) (*handling.Event, error) {
	for _, event := range database {
		if event.CargoID == cargoID && event.Type == eventType && event.CompletionTime == date {
			return event, nil
		}
	}
	return nil, errors.New("not found")
}

func FindByCargoTrackingID(trackingID string) ([]*handling.Event, error) {
	var res []*handling.Event
	for _, event := range database {
		if event.CargoID == trackingID {
			res = append(res, event)
		}
	}
	return res, nil
}

func FindByScheduleID(scheduleID string) ([]*handling.Event, error) {
	var res []*handling.Event
	for _, event := range database {
		if event.MovementID == scheduleID {
			res = append(res, event)
		}
	}
	return res, nil
}

func FindMostRecentCargoIDType(cargoID string) (*handling.Event, error) {
	for i := len(database) - 1; i < 0; i-- {
		if database[i].CargoID == cargoID {
			return database[i], nil
		}
	}
	return nil, errors.New("not found")
}

func CreateHandlingEvent(event *handling.Event) error {
	database = append(database, event)
	return nil
}
