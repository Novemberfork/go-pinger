# go-pinger

[![Version](https://img.shields.io/badge/version-0.0.3-blue.svg)](https://github.com/NovemberFork/go-pinger/releases)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A simple Go library for sending notifications across multiple channels (desktop, iMessage, Telegram, and more).

## Features

- 🖥️ Desktop notifications (macOS)
- 📱 iMessage notifications (macOS)
- ✈️ Telegram notifications (cross-platform)
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
   
   ENABLE_TELEGRAM=true
   TELEGRAM_BOT_TOKEN=your_bot_token_here
   TELEGRAM_CHAT_ID=your_chat_id_here
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
        EnableDesktop:    true,
        DesktopSound:     "Ping",
        EnableIMessage:   true,
        PhoneNumber:      "+1234567890",
        EnableTelegram:   true,
        TelegramBotToken: "your_bot_token",
        TelegramChatID:   "your_chat_id",
    }

    // Create pinger instance
    p := pinger.NewPinger(config)

    // Send notifications
    p.Ping("Alert", "Something important happened!")
    p.PingSimple("Simple message")
}
```

## Configuration

| Setting              | Description                             | Default |
| -------------------- | --------------------------------------- | ------- |
| `ENABLE_DESKTOP`     | Enable desktop notifications            | `true`  |
| `DESKTOP_SOUND`      | Sound for desktop notifications         | `Ping`  |
| `ENABLE_IMESSAGE`    | Enable iMessage notifications           | `false` |
| `PHONE_NUMBER`       | Phone number for iMessage (optional)    | `""`    |
| `ENABLE_TELEGRAM`    | Enable Telegram notifications           | `false` |
| `TELEGRAM_BOT_TOKEN` | Bot token from @BotFather               | `""`    |
| `TELEGRAM_CHAT_ID`   | Chat ID to send messages to             | `""`    |

## Telegram Setup

To use Telegram notifications:

1. **Create a Telegram Bot:**
   - Open Telegram and search for `@BotFather`
   - Send `/newbot` and follow the instructions
   - Save the bot token provided

2. **Get your Chat ID:**
   - Start a chat with your new bot
   - Send any message to the bot
   - Visit `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`
   - Look for the `"chat":{"id":...}` field in the response
   - Use this ID as your `TELEGRAM_CHAT_ID`

3. **For channel notifications:**
   - Add your bot to the channel as an administrator
   - Use the channel ID (usually starts with `-100`) as `TELEGRAM_CHAT_ID`

## Installation

```bash
go get github.com/NovemberFork/go-pinger@latest
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
- Slack notifications
- Discord notifications
- Webhooks support
