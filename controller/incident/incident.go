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

package incident

import (
	"fmt"
	"net/http"

	"github.com/cnbailian/DDD-example/service/incident"
	"github.com/gin-gonic/gin"
)

func HandlingEvent(c *gin.Context) {
	payload := &incident.HandlingEventPayload{}
	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("parse payload: %v", err))
		return
	}
	event, err := incident.HandlingEvent(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("handling event: %v", err))
		return
	}
	c.JSON(http.StatusOK, event)
}

func CreateCarrierMovement(c *gin.Context) {
	payload := &incident.CarrierMovementPayload{}
	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("parse payload: %v", err))
		return
	}
	movement, err := incident.CreateCarrierMovement(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("create carrier movement: %v", err))
		return
	}
	c.JSON(http.StatusOK, movement)
}
