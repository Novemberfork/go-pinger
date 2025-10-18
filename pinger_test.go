package pinger

import (
	"testing"
)

func TestNewPinger(t *testing.T) {
	config := PingerConfig{
		EnableDesktop:    true,
		DesktopSound:     "Ping",
		EnableIMessage:   false,
		PhoneNumber:      "",
		EnableTelegram:   false,
		TelegramBotToken: "",
		TelegramChatID:   "",
	}

	pinger := NewPinger(config)
	if pinger == nil {
		t.Fatal("NewPinger returned nil")
	}

	if pinger.config.DesktopSound != "Ping" {
		t.Errorf("Expected DesktopSound to be 'Ping', got '%s'", pinger.config.DesktopSound)
	}
}

func TestValidateSound(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Ping", "Ping"},
		{"ping", "Ping"},
		{"PING", "Ping"},
		{"Basso", "Basso"},
		{"Invalid", "Ping"},
		{"", "Ping"},
	}

	for _, test := range tests {
		result := validateSound(test.input)
		if result != test.expected {
			t.Errorf("validateSound(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestPingerConfig(t *testing.T) {
	config := PingerConfig{
		EnableDesktop:    true,
		DesktopSound:     "Ping",
		EnableIMessage:   true,
		PhoneNumber:      "+1234567890",
		EnableTelegram:   true,
		TelegramBotToken: "test_token",
		TelegramChatID:   "123456",
	}

	if !config.EnableDesktop {
		t.Error("EnableDesktop should be true")
	}

	if !config.EnableIMessage {
		t.Error("EnableIMessage should be true")
	}

	if !config.EnableTelegram {
		t.Error("EnableTelegram should be true")
	}

	if config.TelegramBotToken != "test_token" {
		t.Errorf("Expected TelegramBotToken to be 'test_token', got '%s'", config.TelegramBotToken)
	}

	if config.TelegramChatID != "123456" {
		t.Errorf("Expected TelegramChatID to be '123456', got '%s'", config.TelegramChatID)
	}
}

func TestPingWithTelegramDisabled(t *testing.T) {
	config := PingerConfig{
		EnableDesktop:    false,
		EnableIMessage:   false,
		EnableTelegram:   false,
		TelegramBotToken: "",
		TelegramChatID:   "",
	}

	pinger := NewPinger(config)
	err := pinger.Ping("Test", "This should not send anything")
	if err != nil {
		t.Errorf("Ping returned error: %v", err)
	}
}

func TestPingSimple(t *testing.T) {
	config := PingerConfig{
		EnableDesktop:  false,
		EnableIMessage: false,
		EnableTelegram: false,
	}

	pinger := NewPinger(config)
	err := pinger.PingSimple("Test message")
	if err != nil {
		t.Errorf("PingSimple returned error: %v", err)
	}
}
