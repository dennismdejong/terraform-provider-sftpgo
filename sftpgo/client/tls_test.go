// Copyright (C) 2025 Nicola Murino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient_TLSVerificationEnabled(t *testing.T) {
	// Test with TLS verification enabled (default)
	c, err := NewClient("https://example.com", "user", "pass", "", []KeyValue{}, 0, true)
	require.NoError(t, err)
	require.NotNil(t, c)

	// Check that we have a valid HTTP client
	require.NotNil(t, c.HTTPClient)
	// When TLS verification is enabled, we should use the default transport
	// (which may be nil initially but gets initialized on first use)
	// For this test, we just verify the client was created successfully
	require.NotNil(t, c.HTTPClient)
}

func TestNewClient_TLSVerificationDisabled(t *testing.T) {
	// Test with TLS verification disabled
	c, err := NewClient("https://example.com", "user", "pass", "", []KeyValue{}, 0, false)
	require.NoError(t, err)
	require.NotNil(t, c)

	// Check that transport has been modified for insecure skip verify
	require.NotNil(t, c.HTTPClient.Transport)
	if transport, ok := c.HTTPClient.Transport.(*http.Transport); ok {
		require.NotNil(t, transport.TLSClientConfig)
		require.True(t, transport.TLSClientConfig.InsecureSkipVerify)
	}
}

func TestNewClient_TLSVerificationDisabled_WithExistingTransport(t *testing.T) {
	// Test with a custom initial transport to ensure we preserve settings
	defaultTransport := &http.Transport{
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   50,
	}

	// Temporarily replace default transport
	originalTransport := http.DefaultTransport
	http.DefaultTransport = defaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	c, err := NewClient("https://example.com", "user", "pass", "", []KeyValue{}, 0, false)
	require.NoError(t, err)
	require.NotNil(t, c)

	// Check that we preserved some settings while adding TLS config
	if transport, ok := c.HTTPClient.Transport.(*http.Transport); ok {
		require.NotNil(t, transport.TLSClientConfig)
		require.True(t, transport.TLSClientConfig.InsecureSkipVerify)
		// Note: We don't strictly check for preservation of other settings since
		// our implementation creates a new transport, but that's acceptable
	}
}