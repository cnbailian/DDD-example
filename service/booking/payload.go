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

type CargoBookingPayload struct {
	ShipperID               string `json:"shipper_id"`
	ConsigneeID             string `json:"consignee_id"`
	OriginLocationCode      string `json:"origin_location_code"`
	DestinationLocationCode string `json:"destination_location_code"`
	// Estimated Time of Arrival(ETA)
	ETA string `json:"eta"`
}
