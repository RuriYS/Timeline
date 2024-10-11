# Timeline

Timeline is a Discord bot built with [DiscordGO](https://github.com/bwmarrin/discordgo) that logs all activities within your Discord server and saves it to a database. This project is currently a work in progress (WIP) and aims to provide robust logging capabilities to help server administrators keep track of events, messages, and user activities.

## Features

- Nothing yet

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (1.23 or higher)
- [DiscordGO](https://github.com/bwmarrin/discordgo)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/RuriYoshinova/Timeline.git
   ```

2. **Navigate to the Directory**:
   ```bash
   cd Timeline
   ```

3. **Install Dependencies**:
   Ensure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

4. **Configure Your Bot Token and Prefix**:
   Create a `.env` file in the root directory and add your bot token:
   ```
   DISCORD_TOKEN=your-bot-token
   PREFIX=$
   ```

5. **Run the Bot**:
   ```bash
   go run main.go
   ```

## Usage

- Run the bot with `go run main.go`

## Contributing

Contributions are welcome! If you'd like to contribute to Timeline, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

## License

This project is licensed under the WTFPL License - see the [LICENSE](LICENSE) file for details.

---

*Last updated on Fri, Oct 4, 2024.*
