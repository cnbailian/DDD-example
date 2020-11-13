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
	"github.com/cnbailian/DDD-example/pkg/util"
)

func NewCargo(payload *Prototype) *Cargo {
	cargo := &Cargo{
		TrackingID: util.GetUUID(),
		Shipper:    payload.Shipper,
		Consignee:  payload.Consignee,
		Spec:       payload.Specification,
	}
	// 重构：Cargo Aggregate 的另一种设计
	//cargo.History = newDeliveryHistory(cargo)
	return cargo
}

func (c *Cargo) CopyCargo() *Cargo {
	cargo := &Cargo{
		TrackingID: util.GetUUID(),
		Shipper:    c.Shipper,
		Consignee:  c.Consignee,
		Spec:       c.Spec,
	}
	// 重构：Cargo Aggregate 的另一种设计
	//cargo.History = newDeliveryHistory(cargo)
	return cargo
}

// 重构：Cargo Aggregate 的另一种设计
// 这样 Cargo 就可以不存储 DeliveryHistory，而是在需要时在 HandlingEvents 集合中查询，临时构建
//func newDeliveryHistory(c *Cargo) *DeliveryHistory {
//	return &DeliveryHistory{
//		cargo:          c,
//		HandlingEvents: nil,
//	}
//}
