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

package location

import (
	"errors"

	"github.com/cnbailian/DDD-example/domain/location"
)

var database = []*location.Location{
	{
		Port: "010",
		Name: "Beijing",
	},
	{
		Port: "022",
		Name: "Tianjin",
	},
	{
		Port: "0081",
		Name: "Japan",
	},
}

func FindByPortCode(code string) (*location.Location, error) {
	for _, l := range database {
		if l.Port == location.Code(code) {
			return l, nil
		}
	}
	return nil, errors.New("not found")
}

func FindByCityName(name string) ([]*location.Location, error) {
	var res []*location.Location
	for _, l := range database {
		if l.Name == name {
			res = append(res, l)
		}
	}
	return res, nil
}
