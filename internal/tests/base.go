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
package tests

import (
	"github.com/mukulmantosh/go-ecommerce-app/internal/database"
	"log/slog"
	"os"
	"path/filepath"
)

func Setup() (database.DBClient, error) {
	slog.Info("Initializing Setup..")
	testDb, err := database.NewTestDBClient()
	if err != nil {
		panic(err)
	}
	err = testDb.RunMigration()
	if err != nil {
		return nil, err
	}
	return testDb, err
}

func Teardown(testDB database.DBClient) {
	slog.Info("Removing TestDB...")
	testDB.CloseConnection()
	currentDirectory, _ := os.Getwd()
	filePath := filepath.Join(currentDirectory, "test.db")
	err := os.Remove(filePath)
	if err != nil {
		return
	}
}
