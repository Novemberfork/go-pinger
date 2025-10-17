# go-pinger

[![Version](https://img.shields.io/badge/version-0.0.2-blue.svg)](https://github.com/NovemberFork/go-pinger/releases)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

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
   ENABLE_IMESSAGE=true
   DESKTOP_SOUND=Ping
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
        DesktopSound:   "Ping",
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

| Setting           | Description                          | Default |
| ----------------- | ------------------------------------ | ------- |
| `ENABLE_DESKTOP`  | Enable desktop notifications         | `true`  |
| `DESKTOP_SOUND`   | Sound for desktop notifications      | `Ping`  |
| `ENABLE_IMESSAGE` | Enable iMessage notifications        | `false` |
| `PHONE_NUMBER`    | Phone number for iMessage (optional) | `""`    |

## Installation

```bash
go get github.com/NovemberFork/go-pinger@v0.0.2
```

## Commands

- `make init` - Create a new configuration file
- `make test` - Test the current configuration
- `make build` - Build the binary
- `make version` - Show version information
- `make tag` - Create git tag for current version
- `make release` - Create and push release tag
- `make clean` - Clean build artifacts

## Future Features

- Email notifications
- SMS notifications (via Twilio, etc.)
- Telegram notifications
- Slack notifications
- Discord notifications
