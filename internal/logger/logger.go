/*
 Copyright 2024 Oussama Jaaouani <oussama@jaaouani.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package logger

import (
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var (
	globalLogger zerolog.Logger
	once         sync.Once
)

func initialize() {
	buildInfo, _ := debug.ReadBuildInfo()

	globalLogger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}).Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}

func Get() zerolog.Logger {
	once.Do(initialize)
	return globalLogger
}

func SetGlobalLevel(level zerolog.Level) {
	once.Do(initialize)
	globalLogger = globalLogger.Level(level)
}
