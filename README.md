# Ethereum Vanity Wallet Generator

A multi-threaded Ethereum wallet generator written in Go. This tool allows users to search for Ethereum addresses with custom prefixes and/or suffixes, providing real-time statistics on the generation process.

## Features

- **Concurrent Processing**: Utilizes Go's goroutines for efficient, multi-threaded wallet generation.
- **Customizable Address Criteria**: Specify desired starting and/or ending strings for the Ethereum address.
- **Live Statistics**: Displays real-time information on generation speed, difficulty, and probability.
- **Performance Optimized**: Designed for optimal performance on multi-core systems.

## Requirements

- Go 1.16 or higher
- Ethereum Crypto Package: `github.com/ethereum/go-ethereum/crypto`

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sjwhole/eth-vanity-wallet-generator.git
   cd eth-vanity-wallet-generator
   ```

2. Install required dependencies:
   ```bash
   go get github.com/ethereum/go-ethereum/crypto
   ```

## Usage

### Building the Program

Compile the program using the following command:

```bash
go build main.go
```

### Running the Generator

Execute the compiled program:

```bash
./main
```

You will be prompted to enter the following information:

1. **Starting String**: The hexadecimal string you want the Ethereum address to start with (excluding the '0x' prefix). Leave blank if not needed.
2. **Ending String**: The hexadecimal string you want the Ethereum address to end with. Leave blank if not needed.
3. **Number of Goroutines**: The number of concurrent threads to use for wallet generation. It's recommended to use a number up to your available CPU cores.

### Example Input

```
Enter the starting string (without '0x'): 777
Enter the ending string: 777
Enter the number of goroutines to use (Available CPUs: 8): 4
```

### Output

The program will display real-time statistics:

```
--- Statistics ---
Difficulty: 16777216
Generated: 50000 addresses
50% Probability: 11629079 addresses
Speed: 50000 addr/s
Current Probability: 0.298023%
------------------
```

- **Difficulty**: The average number of addresses that need to be generated to find a match.
- **Generated**: Total number of addresses generated so far.
- **50% Probability**: Number of addresses needed for a 50% chance of finding a match.
- **Speed**: Number of addresses being generated per second.
- **Current Probability**: The probability of having found a match based on the number of generated addresses.

### When a Match is Found

Upon finding a matching address, the program will display:

```
Found matching address!
Address: 0x777...777
Private Key: e3b0c44298fc1c149afbf4c8996fb924...
```
