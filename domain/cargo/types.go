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

package cargo

import (
	"time"

	"github.com/cnbailian/DDD-example/domain/customer"
	"github.com/cnbailian/DDD-example/domain/location"
)

// Entity
// 通过引用表示关联 Entity 还是 Value Object
type Cargo struct {
	// 标识符
	TrackingID string `json:"tracking_id"`
	// Customer 是关联，但 Role 是 Value Object，所以 Role 内关联 Customer
	// 逻辑说得通，但还看不出这样的好处
	Shipper   Shipper   `json:"shipper"`
	Consignee Consignee `json:"consignee"`
	// Delivery Specification
	Spec DeliverySpecification `json:"spec"`
	// Delivery History
	//History *DeliveryHistory `json:"history"`
}

// Role value object
type Shipper *customer.Customer
type Consignee *customer.Customer

type DeliverySpecification struct {
	Origin      *location.Location `json:"origin"`
	Destination *location.Location `json:"destination"`
	// Estimated Time of Arrival(ETA)
	ETA time.Time `json:"eta"`
}

// 重构：Cargo Aggregate 的另一种设计
// 更新 HandlingEvent 时需要更新 DeliveryHistory，而更新 DeliveryHistory 会在事务中牵扯 Cargo Aggregate，所以需要剥离。
// 我的想法：如果在业务逻辑中 Cargo 的更新并不频繁，大部分的更新都是添加 HandlingEvent 的话，在一起也未尝不可。
//type DeliveryHistory struct {
//	// 双向关联
//	// 双向关联存在于业务逻辑中，如果在代码中没有体现出，则会导致代码 entity 职责划分不明确
//	cargo          *Cargo
//	HandlingEvents []*HandlingEvent `json:"handling_events"`
//}

// Prototype
type Prototype struct {
	Shipper       Shipper               `json:"shipper"`
	Consignee     Consignee             `json:"consignee"`
	Specification DeliverySpecification `json:"specification"`
}
