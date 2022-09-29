/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	C "gopkg.in/check.v1"
)

var originStdout = os.Stdout
var redirectStdoutFile = "testdata/stdout"
var redirectStdout, _ = os.OpenFile(redirectStdoutFile, os.O_CREATE|os.O_RDWR, 0644)

func Test(t *testing.T) { C.TestingT(t) }

type testWrapper struct{}

func init() {
	DebugFile = "" // disable outside debug file
	testWrapper := &testWrapper{}
	C.Suite(testWrapper)
}

func redirectOutput() {
	os.Stdout = redirectStdout
}
func restoreOutput() {
	os.Stdout = originStdout
}
func resetOutput() {
	_ = redirectStdout.Truncate(0)
	_, _ = redirectStdout.Seek(0, io.SeekStart)
}
func readOutput() string {
	fileContent, err := ioutil.ReadFile(redirectStdoutFile)
	if err != nil {
		std.Println("read stdout file failed:", err)
	}
	return string(fileContent)
}
func checkOutput(c *C.C, regfmt string, preferResult bool) {
	output := readOutput()
	result, _ := regexp.MatchString(regfmt, output)
	if result != preferResult {
		c.Errorf("match output failed: `%s`, `%#v`", regfmt, output)
	}
}

func (*testWrapper) BenchmarkSyslog(c *C.C) {
	b := newBackendSyslog("benchSyslog")
	for i := 0; i < c.N; i++ {
		_ = b.log(LevelInfo, "test")
	}
}

func (*testWrapper) TestGeneral(c *C.C) {
	defer func() {
		if err := recover(); err != nil {
			std.Println("catch panic:", err)
		}
	}()

	redirectOutput()
	defer restoreOutput()
	defer resetOutput()

	logger := NewLogger("logger_test")
	logger.SetLogLevel(LevelDebug)

	resetOutput()
	logger.Debug("test debug")
	checkOutput(c, `^<debug> logger_test.go:\d+: test debug\n$`, true)

	resetOutput()
	logger.Info("test info")
	checkOutput(c, `^<info> logger_test.go:\d+: test info\n$`, true)

	resetOutput()
	logger.Info("test info multi-lines\n\nthe thread line and following two empty lines\n\n")
	checkOutput(c, `^<info> logger_test.go:\d+: test info multi-lines\n\nthe thread line and following two empty lines\n\n\n$`, true)

	resetOutput()
	logger.Warning("test warning:", fmt.Errorf("error message"), "append string")
	checkOutput(c, `^<warning> logger_test.go:\d+: test warning: error message append string\n$`, true)

	resetOutput()
	logger.Warning("test warning:", fmt.Errorf("error message"))
	checkOutput(c, `^<warning> logger_test.go:\d+: test warning: error message\n$`, true)

	resetOutput()
	logger.Warningf("test warningf: %v", fmt.Errorf("error message"))
	checkOutput(c, `^<warning> logger_test.go:\d+: test warningf: error message\n$`, true)

	resetOutput()
	logger.Error("test error:", fmt.Errorf("error message"))
	checkOutput(c, `^<error> logger_test.go:\d+: test error: error message\n(  ->  \w+\.\w+:\d+\n)+$`, true)

	resetOutput()
	logger.Errorf("test errorf: %v", fmt.Errorf("error message"))
	checkOutput(c, `^<error> logger_test.go:\d+: test errorf: error message\n(  ->  \w+\.\w+:\d+\n)+$`, true)

	testPanicFunc := func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Info("got panic")
			}
		}()
		logger.Panic("test panic")
	}
	resetOutput()
	testPanicFunc()
	checkOutput(c, `^<error> logger_test.go:\d+: test panic\n(  ->  \w+\.\w+:\d+\n)+<info> logger_test.go:\d+: got panic\n$`, true)
}

// TODO: remove
func (*testWrapper) TestFuncTracing(c *C.C) {
	defer func() {
		if err := recover(); err != nil {
			std.Println("catch error:", err)
		}
	}()

	logger := NewLogger("logger_test")

	logger.BeginTracing()
	defer logger.EndTracing()
	defer func() {
		logger.EndTracing()
	}()
	logger.EndTracing()

	subFunc := func() {
		logger.BeginTracing()
		logger.EndTracing()
	}
	go subFunc()

	panic("test panic")
}

func (*testWrapper) TestIsNeedLog(c *C.C) {
	logger := &Logger{}
	logger.SetLogLevel(LevelInfo)
	c.Check(logger.isNeedLog(LevelDebug), C.Equals, false)
	c.Check(logger.isNeedLog(LevelInfo), C.Equals, true)
	c.Check(logger.isNeedLog(LevelWarning), C.Equals, true)
	c.Check(logger.isNeedLog(LevelError), C.Equals, true)
	c.Check(logger.isNeedLog(LevelPanic), C.Equals, true)
	c.Check(logger.isNeedLog(LevelFatal), C.Equals, true)
	logger.SetLogLevel(LevelDebug)
	c.Check(logger.isNeedLog(LevelDebug), C.Equals, true)
	c.Check(logger.isNeedLog(LevelInfo), C.Equals, true)
	c.Check(logger.isNeedLog(LevelWarning), C.Equals, true)
	c.Check(logger.isNeedLog(LevelError), C.Equals, true)
	c.Check(logger.isNeedLog(LevelPanic), C.Equals, true)
	c.Check(logger.isNeedLog(LevelFatal), C.Equals, true)
}

func (*testWrapper) TestIsNeedTraceMore(c *C.C) {
	logger := &Logger{}
	logger.SetLogLevel(LevelInfo)
	c.Check(logger.isNeedTraceMore(LevelDebug), C.Equals, false)
	c.Check(logger.isNeedTraceMore(LevelInfo), C.Equals, false)
	c.Check(logger.isNeedTraceMore(LevelWarning), C.Equals, false)
	c.Check(logger.isNeedTraceMore(LevelError), C.Equals, true)
	c.Check(logger.isNeedTraceMore(LevelPanic), C.Equals, true)
	c.Check(logger.isNeedTraceMore(LevelFatal), C.Equals, true)
}

func (*testWrapper) TestAddRemoveBackend(c *C.C) {
	logger := &Logger{}

	var backendNull Backend
	logger.AddBackend(backendNull)
	c.Check(len(logger.backends), C.Equals, 0)
	var backendConsoleNull *backendConsole
	logger.AddBackend(backendConsoleNull)
	c.Check(len(logger.backends), C.Equals, 0)
	logger.ResetBackends()

	logger.AddBackendConsole()
	c.Check(len(logger.backends), C.Equals, 1)
	logger.AddBackendConsole()
	c.Check(len(logger.backends), C.Equals, 2)
	logger.RemoveBackendConsole()
	c.Check(len(logger.backends), C.Equals, 0)
}

func (*testWrapper) TestDebugFile(c *C.C) {
	oldDebugFile := DebugFile
	DebugFile = "testdata/dde_debug"
	defer func() { DebugFile = oldDebugFile }()

	os.Clearenv()
	defer os.Clearenv()

	_ = os.Remove(DebugFile)
	c.Check(getDefaultLogLevel("test_debug_file"), C.Equals, LevelInfo)

	_, _ =os.Create(DebugFile)
	c.Check(getDefaultLogLevel("test_debug_file"), C.Equals, LevelDebug)
}

func (*testWrapper) TestDebugEnv(c *C.C) {
	os.Clearenv()
	defer os.Clearenv()

	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelInfo)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG", "")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelDebug)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG", "1")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelDebug)
}

func (*testWrapper) TestDebugLevelEnv(c *C.C) {
	os.Clearenv()
	defer os.Clearenv()

	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelInfo)

	_ = os.Setenv("DDE_DEBUG_LEVEL", "debug")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelDebug)

	_ = os.Setenv("DDE_DEBUG_LEVEL", "warning")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelWarning)
}

func (*testWrapper) TestDebugMatchEnv(c *C.C) {
	os.Clearenv()
	defer os.Clearenv()

	_ = os.Setenv("DDE_DEBUG_MATCH", "test1")
	c.Check(getDefaultLogLevel("test1"), C.Equals, LevelDebug)
	c.Check(getDefaultLogLevel("test2"), C.Equals, LevelDisable)

	_ = os.Setenv("DDE_DEBUG_MATCH", "test1|test2")
	c.Check(getDefaultLogLevel("test1"), C.Equals, LevelDebug)
	c.Check(getDefaultLogLevel("test2"), C.Equals, LevelDebug)

	_ = os.Setenv("DDE_DEBUG_MATCH", "not match")
	c.Check(getDefaultLogLevel("test1"), C.Equals, LevelDisable)
	c.Check(getDefaultLogLevel("test2"), C.Equals, LevelDisable)
}

func (*testWrapper) TestDebugMixEnv(c *C.C) {
	os.Clearenv()
	defer os.Clearenv()

	_ = os.Setenv("DDE_DEBUG", "1")
	_ = os.Setenv("DDE_DEBUG_LEVEL", "warning")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelWarning)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG_LEVEL", "error")
	_ = os.Setenv("DDE_DEBUG_MATCH", "test_env")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelError)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG_LEVEL", "error")
	_ = os.Setenv("DDE_DEBUG_MATCH", "not match")
	c.Check(getDefaultLogLevel("test_env"), C.Equals, LevelDisable)
}

func (*testWrapper) TestDebugConsoleEnv(c *C.C) {
	os.Clearenv()
	defer os.Clearenv()

	_ = os.Setenv("DDE_DEBUG_CONSOLE", "1")
	console := newBackendConsole("test-console")
	c.Check(console.syslogMode, C.Equals, true)

	redirectOutput()
	defer restoreOutput()
	resetOutput()
	_ = console.log(LevelInfo, "this line shows as syslog format in console")
	checkOutput(c, `\w+ \d+ \d{2}:\d{2}:\d{2} .* test-console\[\d+\]: <info> this line shows as syslog format in console\n$`, true)
}

func (*testWrapper) TestFmtSprint(c *C.C) {
	c.Check(fmtSprint(""), C.Equals, "")
	c.Check(fmtSprint("a", "b", "c"), C.Equals, "a b c")
	c.Check(fmtSprint("a\n", "b\n", "c\n"), C.Equals, "a\n b\n c\n")
}
