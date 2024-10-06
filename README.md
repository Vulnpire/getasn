# goasn
is a simple Go tool that fetches ASN (Autonomous System Number) information from [bgp.he.net](https://bgp.he.net) for a given domain or organization name. It supports input from either stdin or a file and can output results in both standard and verbose modes.

## Features

- **Domain and Organization Support**: Automatically detects if the input is a domain or an organization name and performs the appropriate search.
- **Input Support**: Accepts input via stdin or from a file.
- **Verbose Mode**: Provides more detailed output when using the `-v` flag.

## Installation

`go install github.com/Vulnpire/goasn@latest`

## Usage

### Fetch ASN for a domain via stdin:

`echo "domain.tld" | ./goasn`

### Fetch ASN for an organization via stdin:

`echo "organization_name" | ./goasn`

### Fetch ASN for multiple domains or organizations from a file:

cat urls.txt | ./goasn

## Example

Standard mode:

`$ echo "intigriti.com" | ./goasn`
AS1337

## Chain with other tools:

`echo youtube | ./goasn | asnmap -silent | tlsx -san -cn -silent -resp-only`

Get the IP ranges from Shodan:

`cat orgs.txt | ./goasn | asnmap -silent | sXtract -ir -q "200 OK"`

From a single org (or domain):

`echo tesla | ./goasn | asnmap -silent | sXtract -ir -q "200"`

![image](https://github.com/user-attachments/assets/8b8d27b8-5b7f-4eb8-bc56-56051f57b57d)
