/*
 * Copyright 1999-2018 Alibaba Group.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/alibaba/Dragonfly/dfget/config"
	"github.com/alibaba/Dragonfly/dfget/util"
	"github.com/go-check/check"
	"github.com/sirupsen/logrus"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type CoreTestSuite struct{}

func init() {
	check.Suite(&CoreTestSuite{})
}

func (suite *CoreTestSuite) TestPrepare(c *check.C) {
	tmpDir, _ := ioutil.TempDir("/tmp", "dfget-TestCore-")
	defer os.RemoveAll(tmpDir)

	buf := &bytes.Buffer{}

	ctx := config.NewContext()
	ctx.WorkHome = tmpDir
	ctx.MetaPath = path.Join(tmpDir, "meta", "host.meta")
	ctx.SystemDataDir = path.Join(tmpDir, "data")
	ctx.Output = path.Join(tmpDir, "test.output")
	ctx.ClientLogger = logrus.StandardLogger()
	ctx.ClientLogger.Out = buf
	util.Printer.Out = buf

	err := prepare(ctx)
	fmt.Printf("%s\nerror:%v", buf.String(), err)
}
