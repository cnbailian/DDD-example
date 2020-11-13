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

package customer

import (
	"errors"

	"github.com/cnbailian/DDD-example/domain/customer"
	"github.com/cnbailian/DDD-example/repository/cargo"
)

var database = []*customer.Customer{
	{
		ID:   "1",
		Name: "Shibata Jun",
	},
	{
		ID:   "2",
		Name: "Maksim Mrvica",
	},
	{
		ID:   "3",
		Name: "Halie Loren",
	},
	{
		ID:   "4",
		Name: "Gummy",
	},
	{
		ID:   "5",
		Name: "Stefanie Sun",
	},
	{
		ID:   "6",
		Name: "Karen Mok",
	},
}

func FindByCustomerID(customerID string) (*customer.Customer, error) {
	for _, c := range database {
		if c.ID == customerID {
			return c, nil
		}
	}
	return nil, errors.New("not found")
}

func FindByName(name string) (*customer.Customer, error) {
	for _, c := range database {
		if c.Name == name {
			return c, nil
		}
	}
	return nil, errors.New("not found")
}

func FindByCargoTrackingID(trackingID string) ([]*customer.Customer, error) {
	cargoEntity, err := cargo.FindByTrackingID(trackingID)
	if err != nil {
		return nil, errors.New("not found cargo")
	}
	var res []*customer.Customer
	for _, c := range database {
		if c == cargoEntity.Shipper || c == cargoEntity.Consignee {
			res = append(res, c)
		}
	}
	return res, nil
}
