// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package lsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apache/skywalking-banyandb/pkg/index/testcases"
	"github.com/apache/skywalking-banyandb/pkg/logger"
	"github.com/apache/skywalking-banyandb/pkg/test"
)

func TestStore_MatchTerm(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	testcases.SetUp(tester, s)
	testcases.RunServiceName(t, s)
}

func TestStore_Iterator(t *testing.T) {
	tester := assert.New(t)
	path, fn := setUp(require.New(t))
	s, err := NewStore(StoreOpts{
		Path:   path,
		Logger: logger.GetLogger("test"),
	})
	defer func() {
		tester.NoError(s.Close())
		fn()
	}()
	tester.NoError(err)
	data := testcases.SetUpDuration(tester, s)
	testcases.RunDuration(t, data, s)
}

func setUp(t *require.Assertions) (tempDir string, deferFunc func()) {
	t.NoError(logger.Init(logger.Logging{
		Env:   "dev",
		Level: "warn",
	}))
	tempDir, deferFunc = test.Space(t)
	return tempDir, deferFunc
}
