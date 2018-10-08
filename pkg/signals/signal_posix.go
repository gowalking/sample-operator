/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamx. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

// +build !windows
package signals

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
