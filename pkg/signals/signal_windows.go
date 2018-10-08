/*
 * Licensed Materials - Property of tenxcloud.com
 * (C) Copyright 2018 Dreamxos. All Rights Reserved.
 * 2018-10-08  @author gaozh
 */

package signals

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
