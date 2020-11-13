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

package tracking

import (
	"fmt"
	"net/http"

	"github.com/cnbailian/DDD-example/repository/cargo"
	"github.com/cnbailian/DDD-example/repository/handling"
	"github.com/gin-gonic/gin"
)

func GetCargo(c *gin.Context) {
	id := c.Param("id")
	cargoEntity, err := cargo.FindByTrackingID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("find cargo: %v", err))
		return
	}
	c.JSON(http.StatusOK, cargoEntity)
}

func GetCargoHandlingEvents(c *gin.Context) {
	id := c.Param("id")
	events, err := handling.FindByCargoTrackingID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("find events: %v", err))
		return
	}
	c.JSON(http.StatusOK, events)
}
