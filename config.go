package pinger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LoadConfigFromFile loads configuration from a config file
func LoadConfigFromFile(filename string) (*PingerConfig, error) {
	config := &PingerConfig{
		EnableDesktop:    true,   // Default to enabled
		DesktopSound:     "Ping", // Default sound
		EnableIMessage:   false,  // Default to disabled
		PhoneNumber:      "",
		EnableTelegram:   false, // Default to disabled
		TelegramBotToken: "",
		TelegramChatID:   "",
	}

	file, err := os.Open(filename)
	if err != nil {
		return config, nil // Return default config if file doesn't exist
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "ENABLE_DESKTOP":
				if enabled, err := strconv.ParseBool(value); err == nil {
					config.EnableDesktop = enabled
				}
			case "DESKTOP_SOUND":
				config.DesktopSound = value
			case "ENABLE_IMESSAGE":
				if enabled, err := strconv.ParseBool(value); err == nil {
					config.EnableIMessage = enabled
				}
			case "PHONE_NUMBER":
				config.PhoneNumber = value
			case "ENABLE_TELEGRAM":
				if enabled, err := strconv.ParseBool(value); err == nil {
					config.EnableTelegram = enabled
				}
			case "TELEGRAM_BOT_TOKEN":
				config.TelegramBotToken = value
			case "TELEGRAM_CHAT_ID":
				config.TelegramChatID = value
			}
		}
	}

	return config, scanner.Err()
}

// SaveConfigToFile saves configuration to a config file
func SaveConfigToFile(config *PingerConfig, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "# go-pinger Configuration\n")
	fmt.Fprintf(file, "# Edit these values and run 'make test' to test your configuration\n\n")
	fmt.Fprintf(file, "# Desktop Notifications\n")
	fmt.Fprintf(file, "ENABLE_DESKTOP=%t\n", config.EnableDesktop)
	fmt.Fprintf(file, "DESKTOP_SOUND=%s\n", config.DesktopSound)
	fmt.Fprintf(file, "\n# iMessage Notifications\n")
	fmt.Fprintf(file, "ENABLE_IMESSAGE=%t\n", config.EnableIMessage)
	fmt.Fprintf(file, "PHONE_NUMBER=%s\n", config.PhoneNumber)
	fmt.Fprintf(file, "\n# Telegram Notifications\n")
	fmt.Fprintf(file, "ENABLE_TELEGRAM=%t\n", config.EnableTelegram)
	fmt.Fprintf(file, "TELEGRAM_BOT_TOKEN=%s\n", config.TelegramBotToken)
	fmt.Fprintf(file, "TELEGRAM_CHAT_ID=%s\n", config.TelegramChatID)
	fmt.Fprintf(file, "\n# Available Desktop Sounds:\n")
	fmt.Fprintf(file, "# Basso, Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi, Submarine, Tink\n")

	return nil
}
