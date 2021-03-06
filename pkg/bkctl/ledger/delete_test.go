// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteArgError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := testLedgerCommands(deleteCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the ledger id is not specified or the ledger id is specified more than one",
		nameErr.Error())

	args = []string{"delete", "a"}
	_, execErr, _, _ := testLedgerCommands(deleteCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid ledger id a", execErr.Error())

	args = []string{"delete", "--", "-1"}
	_, execErr, _, _ = testLedgerCommands(deleteCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid ledger id -1", execErr.Error())
}
