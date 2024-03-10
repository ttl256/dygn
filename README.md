# dygn

A CLI tool to produce an IPv6 address with the host portion being a modified EUI-64.

## Installation

### macOS

```
brew tap ttl256/dygn
brew install dygn
```

## Usage

```
Usage of dygn:
  -mac value
        MAC address
        Any format will do: xx:xx:xx:xx:xx:xx, xx-xx-xx-xx-xx-xx, xxxx.xxxx.xxxx, xxxxxxxxxxxx
  -prefix value
        IPv6 prefix of prefix-length 64
```

## Example

```
$ dygn --prefix 2001:db8:dead:beef::/64 --mac fc:99:47:75:ce:e0
2001:db8:dead:beef:fe99:47ff:fe75:cee0
```