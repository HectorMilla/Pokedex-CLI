# Pokédex CLI

A command-line interface Pokédex application built in Go that interacts with the [PokéAPI](https://pokeapi.co/). This application allows you to explore Pokémon locations, catch Pokémon, and manage your personal Pokédex.

## Features

- Explore Pokémon locations with pagination (map/mapb commands)
- Catch Pokémon with a chance-based catching system
- View detailed information about caught Pokémon
- Cache system for improved performance
- Interactive command-line interface

## Installation

```bash
# Clone the repository
git clone https://github.com/HectorMilla/Pokedex-CLI.git

# Navigate to the project directory
cd Pokedex-CLI

# Build the project
go build
```

## Usage

Start the application:
```bash
./Pokedex-CLI
```

### Available Commands

- `help` - Show all available commands
- `exit` - Exit the Pokédex
- `map` - List Pokémon location areas
- `mapb` - Go back to previous location areas
- `catch <pokemon_name>` - Try to catch a Pokémon
- `inspect <pokemon_name>` - View details of a caught Pokémon
- `pokedex` - View all Pokémon in your Pokédex

### Examples

```bash
# List Pokémon locations
> map

# Try to catch a Pokémon
> catch pikachu

# View your caught Pokémon
> pokedex

# Inspect a specific Pokémon
> inspect charizard
```

## Technical Details

- Built in Go
- Uses the PokéAPI for Pokémon data
- Implements caching for API responses
- Concurrent operations for improved performance
- Error handling for API requests and user input

## Upcoming Features

- Command History
  - Support for "up" arrow to cycle through previous commands
- Enhanced Battle System
  - Simulate battles between Pokémon
  - Random encounters with wild Pokémon
- Improved Pokémon Management
  - Keep Pokémon in a "party" system
  - Pokémon leveling and experience system
  - Evolution system based on time
  - Persistence of Pokédex data between sessions
- Enhanced Gameplay Mechanics
  - Different types of Poké Balls with varying catch rates
  - More intuitive area navigation with directional commands
- Technical Improvements
  - Additional unit test coverage
  - Code refactoring for better organization and testability
