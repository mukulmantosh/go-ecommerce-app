/*
	Copyright 2022 Google LLC

#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/
package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          string  `gorm:"primaryKey" json:"ID"`
	Name        string  `json:"name"`
	Description string  `gorm:"size:255" json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
}

type UpdateProduct struct {
	Name        string `json:"name"`
	Description string `gorm:"size:255" json:"description"`
}

type ProductParams struct {
	ProductID string `json:"productID"`
}
