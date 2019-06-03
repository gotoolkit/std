// Copyright 2019 llitfkitfk@gmail.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import (
	"os"
	"runtime"
)

func Pid() int {
	return os.Getpid()
}

func GoVersion() string {
	return runtime.Version()
}
