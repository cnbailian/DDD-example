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

type Event struct {
	// 单向关联
	cargo *cargo.Cargo
	// 重构：Cargo Aggregate 的另一种设计
	// 在 DeliveryHistory 中不使用 HandlingEvents 集合，而是通过查询代替，所以 Repository 需要存储 Cargo 标识符
	CargoID        string    `json:"cargo_id"`
	Type           EventType `json:"type"`
	CompletionTime time.Time `json:"completion_time"`
	movement       *carrier.Movement
	MovementID     string `json:"movement_id"`
}

type EventType string

const (
	SentEventType EventType = "Sent"
	// Arrival of Goods(AOG)
	AOGEventType      EventType = "AOG"
	DeliveryEventType EventType = "Delivery"
	// 海关
	CustomsDeclarationEventType EventType = "Customs Declaration"
	CustomsClearanceEventType   EventType = "Customs Clearance"
	CustomsInspectionEventType  EventType = "Customs Inspection"
	CustomsReleaseEventType     EventType = "Customs Release"
)
