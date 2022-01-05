# mailtm

The `mailtm` module wraps the [Mail.tm API](https://api.mail.tm) and provides full functionality. 

This is the seconds version. It solves the login problem from version 1 and adds a login with existing credentials.

If you like it, please consider giving it a :star:

###### Install

`go get github.com/felixstrobel/mailtm`

###### Documentation

Firstly, create an `MailClient` object. This is the object that communicates with the API and contains all important
information.
```go
import "github.com/felixstrobel/mailtm"

func main() {
    client, err := mailtm.NewMailClient()
}
```

After that you can either create an account:
```go
client.GetDomains()
client.CreateAccount()
```
You have to fetch the available domains before creating an account (see above). The created account can be accessed through `client.Account`. The account's address and password are randomized.

If you already have an account and know the address and password, you can skip this step and go on with fetching the JWT token. 
```go
// if you don't have valid credentials:
client.GetAuthToken()

// if you have valid credentials:
client.GetAuthTokenCredentials(yourAddress, yourPassword) 
```
---
You are authenticated now! You can begin...

...fetching your messages:
```go
messages, err := client.GetMessages()
```
...get details of a message:
```go
message, err := client.GetMessageByID(messages[0].ID)
```
...delete one:
```go
client.DeleteMessageByID(message.ID)
```
...or update the seen status to true:
```go
client.SeenMessageByID(message.ID)
```
---
If you decide to delete an account, you can use the `DeleteAccountByID` function to do so:
```go
client.DeleteAccountByID(c.Account.ID)
```