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
package abstract

import "github.com/labstack/echo/v4"

type Category interface {
	AddCategory(ctx echo.Context) error
	GetCategoryById(ctx echo.Context) error
	UpdateCategory(ctx echo.Context) error
	DeleteCategory(ctx echo.Context) error
}
