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

package controller

import (
	"net/http"

	"github.com/cnbailian/DDD-example/controller/booking"
	"github.com/cnbailian/DDD-example/controller/incident"
	"github.com/cnbailian/DDD-example/controller/tracking"
	"github.com/cnbailian/DDD-example/pkg/apiserver"
	"github.com/gin-gonic/gin"
)

func init() {
	// register routes
	apiserver.RegisterRoutes(AddToGinRoute)
	apiserver.RegisterRoutes(BookingAddToGinRoute)
	apiserver.RegisterRoutes(IncidentAddToGinRoute)
	apiserver.RegisterRoutes(TrackingAddToGinRoute)
}

func AddToGinRoute(r *gin.Engine) {
	r.GET("index", func(c *gin.Context) {
		c.JSON(http.StatusOK, "index")
	})
}

func BookingAddToGinRoute(r *gin.Engine) {
	g := r.Group("booking")
	g.POST("cargo", booking.CargoBooking)
}

func IncidentAddToGinRoute(r *gin.Engine) {
	g := r.Group("incident")
	g.POST("event", incident.HandlingEvent)
	g.POST("carrier_movement", incident.CreateCarrierMovement)
}

func TrackingAddToGinRoute(r *gin.Engine) {
	g := r.Group("tracking")
	g.GET("cargo/:id", tracking.GetCargo)
	g.GET("cargo/:id/events", tracking.GetCargoHandlingEvents)
}
