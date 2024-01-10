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

type Category struct {
	gorm.Model
	ID          string    `gorm:"primaryKey;uniqueIndex" json:"ID"`
	Name        string    `gorm:"uniqueIndex;size:255" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID;references:ID" json:"product"`
}
