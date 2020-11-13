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

package booking

import (
	"fmt"
	"time"

	"github.com/cnbailian/DDD-example/domain/cargo"
	cargorepo "github.com/cnbailian/DDD-example/repository/cargo"
	"github.com/cnbailian/DDD-example/repository/customer"
	"github.com/cnbailian/DDD-example/repository/location"
)

func CargoBooking(payload *CargoBookingPayload) (*cargo.Cargo, error) {
	shipper, err := customer.FindByCustomerID(payload.ShipperID)
	if err != nil {
		return nil, fmt.Errorf("find shipper: %w", err)
	}
	consignee, err := customer.FindByCustomerID(payload.ConsigneeID)
	if err != nil {
		return nil, fmt.Errorf("find consignee: %w", err)
	}
	origin, err := location.FindByPortCode(payload.OriginLocationCode)
	if err != nil {
		return nil, fmt.Errorf("find origin: %w", err)
	}
	destination, err := location.FindByPortCode(payload.DestinationLocationCode)
	if err != nil {
		return nil, fmt.Errorf("find destination: %w", err)
	}
	eta, err := time.Parse(time.RFC3339, payload.ETA)
	if err != nil {
		return nil, fmt.Errorf("parse eta time: %w", err)
	}
	c := cargo.NewCargo(&cargo.Prototype{
		Shipper:   shipper,
		Consignee: consignee,
		Specification: cargo.DeliverySpecification{
			Origin:      origin,
			Destination: destination,
			ETA:         eta,
		},
	})
	if err := cargorepo.CreateCargo(c); err != nil {
		return nil, fmt.Errorf("create cargo: %w", err)
	}
	return c, nil
}
