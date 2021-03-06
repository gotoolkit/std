// Copyright 2019 llitfkitfk@gmail.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package signal

import (
	"context"
	"os"
	gosignal "os/signal"
	"syscall"
)

func SigTermCancelContext(ctx context.Context) context.Context {
	term := make(chan os.Signal)
	gosignal.Notify(term, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	c, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case <-term:
			cancel()
		case <-c.Done():
		}
	}()

	return c
}
