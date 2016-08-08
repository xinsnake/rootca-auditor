# RootCA Auditor

I am trying to build an auditor for system root certificates. This project
is just started, so bugs and problems are expected.

PRs are welcomed.

## Work in progress

Working platform
* OS X El Capitan ([from Apple](https://support.apple.com/en-au/HT205204))

## Installation and Usage

As this project is just few days old, currently there's no binary available.
Please clone the source code and build it on your machine yourself.

*Usage*
* compile `go build -o main *.go`
* run `./main`

*Sample Output*
```
REMOTE NOT EXIST: Developer ID Certification Authority, 187AA9A8C296210C
REMOTE NOT EXIST: OISTE WISeKey Global Root GB CA, 76B1205274F0858746B3F8231AF6C2C0
```

## License

MIT License

