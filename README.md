# gopassgen (Go password generator)

A command-line application to generate random passwords and keys, with options to customize the password characteristics and copy the generated password to the clipboard.

## Features

- Generate random passwords of specified length
- Exclude uppercase letters, lowercase letters, digits, or special characters
- Include specific characters in the generated password
- Copy the generated password to the clipboard

## Installation

To install the application, first clone the repository and navigate to the project directory:

```sh
git clone https://github.com/mohrezfadaei/passgen-cli-cli.git
cd passgen-cli-cli
```

Build the application:

```sh
make build
```

## Usage

To generate a password, run the application with the desired flags:

```sh
./passgen-cli [flags]
```

### Available Flags

- `--length, -l [number]`: The length of the password (default: 16)
- `--exclude-uppercase`: Exclude uppercase letters (default: false)
- `--exclude-lowercase`: Exclude lowercase letters (default: false)
- `--exclude-digits`: Exclude digits (default: false)
- `--exclude-special-characters`: Exclude special characters (default: false)
- `--include-chars [chars]`: Characters to include in the password
- `--copy`: Copy the generated password to the clipboard (default: false)

### Examples

Generate a default password of length 16:

```sh
./passgen-cli
```

Generate a password of length 20 without uppercase letters:

```sh
./passgen-cli --length 20 --exclude-uppercase
```

Generate a password that includes specific characters:

```sh
./passgen-cli --include-chars "&$Df"
```

Generate a password and copy it to the clipboard:

```sh
./passgen-cli --length 20 --copy
```

## Development

### Running Tests

To run the tests for the application:

```sh
make test
```

### Formatting Code

To format the code:

```sh
make fmt
```

### Linting Code

To lint the code:

```sh
make lint
```

### Cleaning Up

To clean up the build artifacts:

```sh
make clean
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Acknowledgments

This application uses the following libraries:

- [Cobra](https://github.com/spf13/cobra) for command-line interface
- [Clipboard](https://github.com/atotto/clipboard) for clipboard functionality
