# mail.tm wrapper

A [mail.tm](https://mail.tm) API wrapper written in Golang.

## Installation

If you have [Golang](https://golang.org/) installed and added to the PATH:

```bash
go get github.com/felixstrobel/mailtm
```

## Getting started

### Create a `MailClient`

The default API address is set to `https://api.mail.tm`. If this address should change you can pass the new address as a string to the `NewMailClient` function.
```golang
var client *MailClient = NewMailClient()
```
