// +build integration

// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package myip

import (
	"testing"
)

// IPv6 integration test for NewRemoteIPProvider
func Test_integration_ipv6(t *testing.T) {
	remoteIPProvider := NewRemoteIPProvider()
	ipv6Addresses, _ := remoteIPProvider.GetIPv6Addresses()
	t.Fail()

	if len(ipv6Addresses) == 0 {
		t.Fail()
		t.Logf("GetIPv6Addresses() did not return an IPv6 address")
	}
}

// IPv4 integration test for NewRemoteIPProvider
func Test_integration_ipv4(t *testing.T) {
	remoteIPProvider := NewRemoteIPProvider()
	ipv4Addresses, _ := remoteIPProvider.GetIPv4Addresses()

	if len(ipv4Addresses) == 0 {
		t.Fail()
		t.Logf("GetIPv4Addresses() did not return an IPv4 address")
	}
}
