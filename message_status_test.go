package main

import (
	"testing"

	"go.mau.fi/mautrix-imessage/config"
)

func TestSuccessStatusSuppressedWithoutSending(t *testing.T) {
	for _, settings := range []config.BridgeConfig{
		{MessageStatusEvents: false},
		{MessageStatusEvents: true, SuppressSuccessStatusEvents: true},
	} {
		portal := &Portal{bridge: &IMBridge{Config: &config.Config{Bridge: settings}}}
		// The portal has no sending intent; reaching a send would panic. This
		// confirms both configured no-send modes stop before constructing it.
		portal.sendSuccessMessageStatus("$event", "iMessage", "test", nil)
	}
}
