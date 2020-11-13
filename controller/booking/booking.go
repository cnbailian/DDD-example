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
	"net/http"

	"github.com/cnbailian/DDD-example/service/booking"
	"github.com/gin-gonic/gin"
)

func CargoBooking(c *gin.Context) {
	payload := &booking.CargoBookingPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("parse data: %v", err))
		return
	}
	cargo, err := booking.CargoBooking(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("booking cargo: %v", err))
		return
	}
	c.JSON(http.StatusOK, cargo)
}
