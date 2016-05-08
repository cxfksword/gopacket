// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// +build linux

package afpacket

import (
	"fmt"
	"testing"
)

func TestGetOSVersion(t *testing.T) {
	fmt.Println(getOSVersion())
}
