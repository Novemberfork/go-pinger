# go-pinger

A simple Go library for sending notifications across multiple channels (desktop, iMessage, and more coming soon).

## Features

- 🖥️ Desktop notifications (macOS)
- 📱 iMessage notifications (macOS)
- 🔧 Simple configuration
- 🧪 Built-in testing
- 📦 Easy to integrate

## Quick Start

1. **Initialize configuration:**
   ```bash
   make init
   ```

2. **Edit configuration:**
   Edit `pinger.conf` to set your preferences:
   ```ini
   ENABLE_DESKTOP=true
   ENABLE_IMESSAGE=true
   PHONE_NUMBER=+1234567890
   ```

3. **Test your configuration:**
   ```bash
   make test
   ```

## Usage as a Library

```go
package main

import (
    "github.com/NovemberFork/go-pinger"
)

func main() {
    // Create configuration
    config := pinger.PingerConfig{
        EnableDesktop:  true,
        EnableIMessage: true,
        PhoneNumber:    "+1234567890",
    }
    
    // Create pinger instance
    p := pinger.NewPinger(config)
    
    // Send notifications
    p.Ping("Alert", "Something important happened!")
    p.PingSimple("Simple message")
}
```

## Configuration

| Setting | Description | Default |
|---------|-------------|---------|
| `ENABLE_DESKTOP` | Enable desktop notifications | `true` |
| `ENABLE_IMESSAGE` | Enable iMessage notifications | `false` |
| `PHONE_NUMBER` | Phone number for iMessage (optional) | `""` |

## Commands

- `make init` - Create a new configuration file
- `make test` - Test the current configuration
- `make build` - Build the binary
- `make clean` - Clean build artifacts

## Future Features

- Email notifications
- SMS notifications (via Twilio, etc.)
- Telegram notifications
- Slack notifications
- Discord notifications
