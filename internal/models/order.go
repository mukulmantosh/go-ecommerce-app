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

type Cart struct {
	gorm.Model
	ID       string    `gorm:"primaryKey" json:"cartId"`
	UserID   string    `json:"user_id"`
	Products []Product `gorm:"many2many:cart_items;"`
}

type Order struct {
	gorm.Model
	ID          string    `gorm:"primaryKey" json:"ID"`
	Total       float32   `gorm:"default:0" json:"total"`
	UserID      string    `json:"user_id"`
	OrderStatus string    `json:"order_status"`
	Products    []Product `gorm:"many2many:order_details;"`
}
