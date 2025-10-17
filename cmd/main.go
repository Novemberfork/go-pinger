package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NovemberFork/go-pinger"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-pinger <command>")
		fmt.Println("Commands:")
		fmt.Println("  test    - Test the current configuration")
		fmt.Println("  init    - Initialize a new configuration file")
		fmt.Println("  help    - Show this help message")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "test":
		testConfiguration()
	case "init":
		initConfiguration()
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
		os.Exit(1)
	}
}

func testConfiguration() {
	fmt.Println("🧪 Testing go-pinger configuration...")

	// Load configuration
	config, err := pinger.LoadConfigFromFile("pinger.conf")
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	// Create pinger instance
	p := pinger.NewPinger(*config)

	// Test the configuration
	err = p.TestConnection()
	if err != nil {
		log.Fatalf("❌ Test failed: %v", err)
	}
}

func initConfiguration() {
	fmt.Println("🔧 Initializing go-pinger configuration...")

	// Create default configuration
	config := &pinger.PingerConfig{
		EnableDesktop:  true,
		DesktopSound:   "Ping",
		EnableIMessage: false,
		PhoneNumber:    "1112223333",
	}

	// Save to file
	err := pinger.SaveConfigToFile(config, "pinger.conf")
	if err != nil {
		log.Fatalf("❌ Failed to save configuration: %v", err)
	}

	fmt.Println("✅ Configuration file created: pinger.conf")
	fmt.Println("📝 Edit pinger.conf to configure your notification preferences")
	fmt.Println("🧪 Run 'make test' to test your configuration")
}

func showHelp() {
	fmt.Println("go-pinger - Simple notification library for Go")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  test    - Test the current configuration")
	fmt.Println("  init    - Initialize a new configuration file")
	fmt.Println("  help    - Show this help message")
	fmt.Println("")
	fmt.Println("Configuration:")
	fmt.Println("  Edit pinger.conf to configure notification preferences")
	fmt.Println("  ENABLE_DESKTOP=true/false  - Enable/disable desktop notifications")
	fmt.Println("  DESKTOP_SOUND=SoundName    - Desktop notification sound")
	fmt.Println("  ENABLE_IMESSAGE=true/false - Enable/disable iMessage notifications")
	fmt.Println("  PHONE_NUMBER=+1234567890   - Phone number for iMessage (optional)")
	fmt.Println("")
	fmt.Println("Available desktop sounds: Basso, Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi, Submarine, Tink")
}
