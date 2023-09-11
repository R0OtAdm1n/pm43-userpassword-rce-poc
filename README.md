# Honeywell PM43 Userpassword Command Injection POC

This is a proof-of-concept (POC) code to test the command injection vulnerability via the `userpassword` parameter in PM43 printer devices. It allows you to check if a specific version of PM43 (P10.11.013310 or earlier) is vulnerable to command injection.

## Prerequisites

To use this POC, you need to meet the following requirements:

- Go programming language version 1.17 or higher installed on your computer.
- Access to a PM43 printer device running the vulnerable version.

## Usage

Follow the instructions below to use this POC:

1. Clone or download the POC code to your local computer.
2. Open a terminal and navigate to the project directory.

### Single URL Testing

If you want to test a single URL, use the following command:

```bash
go run main.go -url <target URL>
```

Replace `<target URL>` with the URL of the PM43 printer device you want to test. For example:

```bash
go run main.go -url http://printer.example.com
```

### Multiple URLs Testing

If you have a file containing multiple URLs that need to be tested, use the following command:

```bash
go run main.go -file <input file>
```

Replace `<input file>` with the path to the file containing the URLs. Each URL in the file should be on a separate line.

## Output

The POC will send the command injection payload to the specified URL(s) and check if the response contains the expected result. If a vulnerability is detected, it will display a success message along with the payload used. If no vulnerability is detected, it will display a failure message.

Note: The POC assumes a specific format for the command injection payload. Depending on the specific vulnerability you are testing, you may need to modify the payload in the code.

## Disclaimer

This POC is intended for educational and testing purposes only. Use it responsibly and obtain necessary authorizations. The author is not responsible for any misuse or damage caused by the use of this POC.

## License

This POC is released under the [MIT License](LICENSE).

## Credits

This POC is based on the original code from the OpenAI GPT-3.5 model.

**Use this POC responsibly and respect the security and privacy of others.**