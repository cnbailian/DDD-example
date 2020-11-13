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

package carrier

import (
	"errors"

	"github.com/cnbailian/DDD-example/domain/carrier"
	"github.com/cnbailian/DDD-example/domain/location"
)

var database []*carrier.Movement

func FindByScheduleID(scheduleID string) (*carrier.Movement, error) {
	for _, movement := range database {
		if movement.ScheduleID == scheduleID {
			return movement, nil
		}
	}
	return nil, errors.New("not found")
}

func FindByFromTo(from, to *location.Location) (*carrier.Movement, error) {
	for _, movement := range database {
		if movement.From == from && movement.To == to {
			return movement, nil
		}
	}
	return nil, errors.New("not found")
}

func CreateCarrierMovement(movement *carrier.Movement) error {
	database = append(database, movement)
	return nil
}
