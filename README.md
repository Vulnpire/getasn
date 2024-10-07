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

`$ echo "domain.tld" | goasn`

### Fetch ASN for an organization via stdin:

`$ echo "organization_name" | goasn`

### Fetch ASN for multiple domains or organizations from a file:

`$ cat urls.txt | goasn`

## Example

Standard mode:

```
$ echo "intigriti.com" | goasn

AS16509

```
## Chain with other tools:

Convert domain names to org, get the IP ranges, and subdomains:

```
$ echo tesla.com | dtoconv | goasn | asnmap -silent | tlsx -san -cn -silent -resp-only

gp-zg.ericssonnikolatesla.com
dal11-gpgw1.tesla.com
93.179.67.13
autodiscover.tesla-sv.ru
mail.tesla-sv.ru
autoconfig.tesla-sv.ru
s3.b.smf11.tcs.tesla.com
x3-prod.obs.tesla.com
x3-eng.obs.tesla.com
teslamotors.com
solarcity.com
sg-1.solarcity.com
..SNIP..
```

Get the IP ranges of organizations and extract ports (or exposed services, CVEs, queries, etc.) from Shodan:

```
$ cat orgs.txt | goasn | asnmap -silent | sXtract -ir -q "port:(21 OR 3389 OR 1337 OR 5000 OR 8080)"

205.149.8.234
205.149.15.116
205.149.11.253
205.149.8.242
205.149.8.232
205.149.8.245
205.149.11.31
205.149.11.187
205.149.13.132
205.149.14.61
205.149.15.21
205.149.15.158
205.149.8.26
```

From a single org (or domain):

`$ echo Tesla | goasn | asnmap -silent | sXtract -ir -q "200 OK"`

![image](https://github.com/user-attachments/assets/8b8d27b8-5b7f-4eb8-bc56-56051f57b57d)
