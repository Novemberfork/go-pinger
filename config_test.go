package pinger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test.conf")

	// Test loading from non-existent file (should return defaults)
	config, err := LoadConfigFromFile(configFile)
	if err != nil {
		t.Errorf("LoadConfigFromFile returned error: %v", err)
	}
	if config.EnableDesktop != true {
		t.Error("Default EnableDesktop should be true")
	}
	if config.DesktopSound != "Ping" {
		t.Error("Default DesktopSound should be 'Ping'")
	}
	if config.EnableIMessage != false {
		t.Error("Default EnableIMessage should be false")
	}
	if config.EnableTelegram != false {
		t.Error("Default EnableTelegram should be false")
	}

	// Create a test config file
	configContent := `# Test configuration
ENABLE_DESKTOP=false
DESKTOP_SOUND=Basso
ENABLE_IMESSAGE=true
PHONE_NUMBER=+1234567890
ENABLE_TELEGRAM=true
TELEGRAM_BOT_TOKEN=test_token_123
TELEGRAM_CHAT_ID=987654321
`
	err = os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load the config file
	config, err = LoadConfigFromFile(configFile)
	if err != nil {
		t.Errorf("LoadConfigFromFile returned error: %v", err)
	}

	// Verify loaded values
	if config.EnableDesktop != false {
		t.Error("EnableDesktop should be false")
	}
	if config.DesktopSound != "Basso" {
		t.Errorf("DesktopSound should be 'Basso', got '%s'", config.DesktopSound)
	}
	if config.EnableIMessage != true {
		t.Error("EnableIMessage should be true")
	}
	if config.PhoneNumber != "+1234567890" {
		t.Errorf("PhoneNumber should be '+1234567890', got '%s'", config.PhoneNumber)
	}
	if config.EnableTelegram != true {
		t.Error("EnableTelegram should be true")
	}
	if config.TelegramBotToken != "test_token_123" {
		t.Errorf("TelegramBotToken should be 'test_token_123', got '%s'", config.TelegramBotToken)
	}
	if config.TelegramChatID != "987654321" {
		t.Errorf("TelegramChatID should be '987654321', got '%s'", config.TelegramChatID)
	}
}

func TestSaveConfigToFile(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test.conf")

	config := &PingerConfig{
		EnableDesktop:    true,
		DesktopSound:     "Ping",
		EnableIMessage:   true,
		PhoneNumber:      "+9876543210",
		EnableTelegram:   true,
		TelegramBotToken: "save_test_token",
		TelegramChatID:   "111222333",
	}

	// Save the config
	err := SaveConfigToFile(config, configFile)
	if err != nil {
		t.Fatalf("SaveConfigToFile returned error: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Fatal("Config file was not created")
	}

	// Load it back and verify
	loadedConfig, err := LoadConfigFromFile(configFile)
	if err != nil {
		t.Fatalf("LoadConfigFromFile returned error: %v", err)
	}

	if loadedConfig.EnableDesktop != config.EnableDesktop {
		t.Error("EnableDesktop mismatch")
	}
	if loadedConfig.DesktopSound != config.DesktopSound {
		t.Error("DesktopSound mismatch")
	}
	if loadedConfig.EnableIMessage != config.EnableIMessage {
		t.Error("EnableIMessage mismatch")
	}
	if loadedConfig.PhoneNumber != config.PhoneNumber {
		t.Error("PhoneNumber mismatch")
	}
	if loadedConfig.EnableTelegram != config.EnableTelegram {
		t.Error("EnableTelegram mismatch")
	}
	if loadedConfig.TelegramBotToken != config.TelegramBotToken {
		t.Errorf("TelegramBotToken mismatch: expected '%s', got '%s'", config.TelegramBotToken, loadedConfig.TelegramBotToken)
	}
	if loadedConfig.TelegramChatID != config.TelegramChatID {
		t.Errorf("TelegramChatID mismatch: expected '%s', got '%s'", config.TelegramChatID, loadedConfig.TelegramChatID)
	}
}

func TestLoadConfigWithComments(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test.conf")

	// Create config with comments and empty lines
	configContent := `# Comment line
# Another comment
ENABLE_TELEGRAM=true

# More comments
TELEGRAM_BOT_TOKEN=token_with_comments
TELEGRAM_CHAT_ID=123
`
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	config, err := LoadConfigFromFile(configFile)
	if err != nil {
		t.Errorf("LoadConfigFromFile returned error: %v", err)
	}

	if !config.EnableTelegram {
		t.Error("EnableTelegram should be true")
	}
	if config.TelegramBotToken != "token_with_comments" {
		t.Errorf("TelegramBotToken should be 'token_with_comments', got '%s'", config.TelegramBotToken)
	}
	if config.TelegramChatID != "123" {
		t.Errorf("TelegramChatID should be '123', got '%s'", config.TelegramChatID)
	}
}
