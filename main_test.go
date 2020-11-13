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

package main

import (
	"log"

	"github.com/cnbailian/DDD-example/domain/cargo"
	"github.com/cnbailian/DDD-example/domain/handling"
	"github.com/cnbailian/DDD-example/service/booking"
	"github.com/cnbailian/DDD-example/service/incident"
)

// 完整的业务逻辑流程
func CargoShipping() *cargo.Cargo {
	// 预约 Cargo
	cargo, err := booking.CargoBooking(&booking.CargoBookingPayload{
		// Customer 拥有默认存储
		ShipperID:   "1",
		ConsigneeID: "2",
		// Location 拥有默认存储
		OriginLocationCode:      "010",
		DestinationLocationCode: "0081",
		ETA:                     "2020-11-11T11:11:11Z",
	})
	if err != nil {
		log.Fatal(err)
	}
	// 创建 Carrier Movements
	beijingToTianjin, err := incident.CreateCarrierMovement(&incident.CarrierMovementPayload{
		From: "010",
		To:   "022",
	})
	if err != nil {
		log.Fatal(err)
	}
	TianjinToJapan, err := incident.CreateCarrierMovement(&incident.CarrierMovementPayload{
		From: "022",
		To:   "0081",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Handling Events
	// 从北京发往天津
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.SentEventType),
		Time:       "2020-11-12T11:11:11Z",
		ScheduleID: beijingToTianjin.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 抵达天津
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.AOGEventType),
		Time:       "2020-11-13T11:11:11Z",
		ScheduleID: beijingToTianjin.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 从天津发往日本
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.SentEventType),
		Time:       "2020-11-14T11:11:11Z",
		ScheduleID: TianjinToJapan.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 海关报关
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.CustomsDeclarationEventType),
		Time:       "2020-11-15T11:11:11Z",
		ScheduleID: TianjinToJapan.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 海关清关
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.CustomsClearanceEventType),
		Time:       "2020-11-16T11:11:11Z",
		ScheduleID: TianjinToJapan.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 海关放行
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.CustomsReleaseEventType),
		Time:       "2020-11-17T11:11:11Z",
		ScheduleID: TianjinToJapan.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 抵达日本
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.AOGEventType),
		Time:       "2020-11-18T11:11:11Z",
		ScheduleID: TianjinToJapan.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	// 交付货物
	if _, err := incident.HandlingEvent(&incident.HandlingEventPayload{
		CargoID:    cargo.TrackingID,
		EventType:  string(handling.DeliveryEventType),
		Time:       "2020-11-18T15:11:11Z",
		ScheduleID: TianjinToJapan.ScheduleID,
	}); err != nil {
		log.Fatal(err)
	}
	return cargo
}

func ExampleCargo() {
	CargoShipping()
	// Output:
}
