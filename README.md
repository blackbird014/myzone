# MyZone Blockchain

MyZone is a custom blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk), designed to provide a scalable and modular blockchain application framework.

## Features

- Custom `mybank` module for token transfers
- CometBFT consensus engine
- CLI tools for interacting with the blockchain

## Prerequisites

- [Go](https://golang.org/doc/install) version 1.21 or higher
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installation

### Build from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/myzone.git
cd myzone

# Install dependencies
go mod tidy

# Build the project
make build
```

The build process will create the `myzoned` binary in the project directory.

### Install Binary

To install the `myzoned` binary to your GOPATH/bin:

```bash
make install
```

## Usage

### Initialize a New Chain

```bash
# Initialize a new chain
myzoned init <moniker> --chain-id myzone-1

# Create a key for the genesis account
myzoned keys add <key_name>

# Add the account to genesis with tokens
myzoned add-genesis-account <key_name> 10000000token

# Generate the genesis transaction
myzoned gentx <key_name> 1000000token --chain-id myzone-1

# Collect genesis transactions
myzoned collect-gentxs
```

### Start the Chain

```bash
# Start the chain
myzoned start
```

### CLI Commands

```bash
# Get status of the node
myzoned status

# Send tokens
myzoned tx bank send <from_key> <to_address> <amount> --chain-id myzone-1

# Query account balance
myzoned query bank balances <address>
```

## Project Structure

- `app/`: Application setup and configuration
- `cmd/`: Command-line interface
- `x/`: Custom modules
  - `mybank/`: Custom banking module

## Contributing

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin feature/my-new-feature`
5. Submit a pull request

## License

This project is licensed under the [Apache 2.0 License](LICENSE).

## Acknowledgments

- [Cosmos SDK](https://github.com/cosmos/cosmos-sdk)
- [CometBFT](https://github.com/cometbft/cometbft) 