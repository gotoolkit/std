// Copyright 2019 llitfkitfk@gmail.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package signal

import (
	"os"
	gosignal "os/signal"
	"sync/atomic"
	"syscall"
)

func Trap(cleanup func(), logger interface {
	Info(args ...interface{})
}) {
	c := make(chan os.Signal, 1)
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGPIPE}
	gosignal.Notify(c, signals...)
	go func() {
		interruptCount := uint32(0)
		for sig := range c {
			if sig == syscall.SIGPIPE {
				continue
			}
			go func(sig os.Signal) {
				switch sig {
				case os.Interrupt, syscall.SIGTERM:
					if atomic.LoadUint32(&interruptCount) < 3 {
						if atomic.AddUint32(&interruptCount, 1) == 1 {
							cleanup()
							os.Exit(0)
						} else {
							return
						}
					} else {
						logger.Info("Forcing shutdown without cleanup; 3 interrupts received")
					}
				case syscall.SIGQUIT:
					logger.Info("Forcing shutdown without cleanup on SIGQUIT")
				}
				os.Exit(128 + int(sig.(syscall.Signal)))
			}(sig)
		}
	}()
}
