# Go Cloudflare IP Scanner

This program scans Cloudflare CIDR ranges to find valid IP addresses based on your ISP's network rules and connectivity.

## Build

Build a smaller stripped binary:

```bash
go build -ldflags="-s -w" -o main .
```

The `-s -w` linker flags remove symbol tables and debug information, reducing the binary size.

## Run

Run the program with optional positional arguments:

```bash
./main [bots] [timeout] [method]
```

Example:

```bash
./main 500 1 1
```

Arguments:

- `bots`: Number of concurrent scan workers.
- `timeout`: Connection timeout in seconds.
- `method`: Scan method. Use `1` for ICMP ping or `2` for TCP ping on port `443`.

When no arguments are provided, the program uses these defaults:

```text
bots:    300
timeout: 1 second
method:  2
```

Run with the defaults:

```bash
./main
```

## How It Works

The program loads Cloudflare CIDR ranges, expands them into individual IP addresses, and distributes those addresses among the configured number of workers.

With method `1`, each worker sends an ICMP echo request and waits for a reply until the configured timeout expires:

```bash
./main 500 1 1
```

With method `2`, each worker attempts a TCP connection to port `443` and waits until the configured timeout expires:

```bash
./main 500 1 2
```

Responsive IP addresses are appended to `hit.txt`.
