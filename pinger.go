package pinger

import (
	"fmt"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// PingerConfig holds the configuration for the pinger
type PingerConfig struct {
	// Desktop notifications
	EnableDesktop bool
	DesktopSound  string

	// iMessage notifications
	EnableIMessage bool
	PhoneNumber    string

	// Telegram notifications
	EnableTelegram   bool
	TelegramBotToken string
	TelegramChatID   string
}

// Pinger provides a simple interface for sending notifications
type Pinger struct {
	config PingerConfig
}

// Valid sounds for macOS notifications
var validSounds = []string{
	"Basso", "Blow", "Bottle", "Frog", "Funk", "Glass",
	"Hero", "Morse", "Ping", "Pop", "Purr", "Sosumi",
	"Submarine", "Tink",
}

// validateSound checks if a sound is valid and returns default if not
func validateSound(sound string) string {
	if sound == "" {
		return "Ping" // Default sound
	}

	for _, validSound := range validSounds {
		if strings.EqualFold(sound, validSound) {
			return validSound
		}
	}

	return "Ping" // Fallback to default
}

// NewPinger creates a new pinger instance with the given configuration
func NewPinger(config PingerConfig) *Pinger {
	// Validate and set default sound for desktop notifications
	config.DesktopSound = validateSound(config.DesktopSound)

	return &Pinger{
		config: config,
	}
}

// Ping sends a notification with the given title and message
func (p *Pinger) Ping(title, message string) error {
	// Send desktop notification if enabled
	if p.config.EnableDesktop {
		p.sendDesktopNotification(title, message)
	}

	// Send iMessage if enabled
	if p.config.EnableIMessage && p.config.PhoneNumber != "" {
		p.sendIMessage(message)
	}

	// Send Telegram message if enabled
	if p.config.EnableTelegram && p.config.TelegramBotToken != "" && p.config.TelegramChatID != "" {
		p.sendTelegramMessage(title, message)
	}

	return nil
}

// PingSimple sends a simple message (uses "Notification" as title)
func (p *Pinger) PingSimple(message string) error {
	return p.Ping("Notification", message)
}

// TestConnection tests all enabled notification methods
func (p *Pinger) TestConnection() error {
	fmt.Println("🧪 Testing notification system...")

	// Test desktop notification
	if p.config.EnableDesktop {
		fmt.Printf("🖥️  Testing desktop notification (sound: %s)...\n", p.config.DesktopSound)
		p.sendDesktopNotification("go-pinger Test", "Testing desktop notifications...")
	}

	// Test iMessage if enabled
	if p.config.EnableIMessage && p.config.PhoneNumber != "" {
		fmt.Printf("📱 Testing iMessage to %s (uses your default ringtone)...\n", p.config.PhoneNumber)
		p.sendIMessage("🧪 go-pinger Test - All systems working!")
	}

	// Test Telegram if enabled
	if p.config.EnableTelegram && p.config.TelegramBotToken != "" && p.config.TelegramChatID != "" {
		fmt.Printf("✈️  Testing Telegram to chat %s...\n", p.config.TelegramChatID)
		p.sendTelegramMessage("go-pinger Test", "🧪 Testing Telegram notifications - All systems working!")
	}

	fmt.Println("✅ Notification test complete!")
	return nil
}

// sendDesktopNotification sends a macOS desktop notification
func (p *Pinger) sendDesktopNotification(title, message string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s" sound name "%s" subtitle "go-pinger"`, message, title, p.config.DesktopSound))
	cmd.Run() // Ignore errors, just try to send
}

// sendIMessage attempts to send an iMessage
func (p *Pinger) sendIMessage(message string) {
	// Simplified AppleScript approach
	script := fmt.Sprintf(`
		tell application "Messages"
			set targetService to 1st service whose service type = iMessage
			set targetBuddy to buddy "%s" of targetService
			send "%s" to targetBuddy
		end tell
	`, p.config.PhoneNumber, message)

	cmd := exec.Command("osascript", "-e", script)
	err := cmd.Run()
	if err != nil {
		// Silently fail - we don't want to spam errors if iMessage fails
		fmt.Printf("📱 (iMessage failed, but desktop notification sent)\n")
	} else {
		fmt.Printf("📱 iMessage sent to %s\n", p.config.PhoneNumber)
	}
}

// sendTelegramMessage sends a message to Telegram
func (p *Pinger) sendTelegramMessage(title, message string) {
	bot, err := tgbotapi.NewBotAPI(p.config.TelegramBotToken)
	if err != nil {
		fmt.Printf("✈️  (Telegram failed: %v)\n", err)
		return
	}

	// Combine title and message for Telegram
	fullMessage := fmt.Sprintf("*%s*\n\n%s", title, message)

	msg := tgbotapi.NewMessage(0, fullMessage)
	msg.ParseMode = "Markdown"

	// Parse chat ID
	var chatID int64
	fmt.Sscanf(p.config.TelegramChatID, "%d", &chatID)
	msg.ChatID = chatID

	_, err = bot.Send(msg)
	if err != nil {
		fmt.Printf("✈️  (Telegram failed: %v)\n", err)
	} else {
		fmt.Printf("✈️  Telegram message sent to chat %s\n", p.config.TelegramChatID)
	}
}
