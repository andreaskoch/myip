// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package myip

import (
	"fmt"
	"net"
	"os"
	"testing"
)

// Instantiating a new LocalIPProvider for determining the local IPv6 address(es)
func ExampleNewLocalIPProvider() {

	// create a new local IP provider instance
	localIPProvider, ipProviderError := NewLocalIPProvider()
	if ipProviderError != nil {
		fmt.Fprintf(os.Stderr, "Failed to create a new local IP provider: %s", ipProviderError.Error())
		os.Exit(1)
	}

	// get the local IPv6 addresses
	localIPv6Addresses, localIPError := localIPProvider.GetIPv6Addresses()
	if localIPError != nil {
		fmt.Fprintf(os.Stderr, "Failed to determine the local IPv6 addresses: %s", localIPError.Error())
		os.Exit(1)
	}

	// print the the local IPv6 addresses
	fmt.Fprintf(os.Stdout, "%s", localIPv6Addresses)
}

func Test_GetIPs_NoInterfacesSupplied_NoIPsAreReturned(t *testing.T) {
	// arrange
	interfaces := []net.Interface{}
	addressProvider := interfaceAddressProvider{interfaces}

	// act
	ips, _ := addressProvider.GetIPs()

	// assert
	if len(ips) > 0 {
		t.Fail()
		t.Logf("GetIPs() should return an empty list of IPs but returned %v instead.", ips)
	}
}
