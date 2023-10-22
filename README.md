# mailtm

The `mailtm` module wraps the [Mail.tm API](https://api.mail.tm) and provides full functionality.

### Install

`go get github.com/felixstrobel/mailtm`

### Documentation

Firstly, create an `MailClient` object. This is the object that communicates with the API.
```go
import "github.com/felixstrobel/mailtm"

func main() {
    client, err := mailtm.New()
}
```

After that you can directly create an account:
```go
// With random password
client.NewAccount()

// With custom password
client.NewAccountWithPassword("password")
```

If you already have an account and know the address and password, you can simply sign back in: 
```go
client.RetrieveAccount("your@email.com", "your_password") 
```
---

##### Fetching your messages of the first page (around 30 messages):
```go
client.GetMessages(&account, 1)
```
##### Get detailed message:
```go
client.GetMessageByID(&account, "the_id_of_the_message")
```
#### Delete a message:
```go
client.DeleteMessageByID(&account, "the_id_of_the_message")
```
#### Mark a message as seen:
```go
client.SeenMessageByID(&account, "the_id_of_the_message")
```
---
If you decide to delete an account, you can use the `DeleteAccount` function to do so:
```go
client.DeleteAccount(&account)
```