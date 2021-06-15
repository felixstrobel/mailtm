# mail.tm wrapper

A [mail.tm](https://mail.tm) API wrapper written in Golang.

## Installation

If you have [Golang](https://golang.org/) installed and added to the PATH:

```bash
go get github.com/felixstrobel/mailtm
```

## Getting started

### Create a client

```golang
var client *MailClient = NewMailClient()
```

The default API address is set to `https://api.mail.tm`. If this address should change you can set the new one liek this:
```golang
client.URL =  "NEW_API_ADDRESS"
```

### Registering an email

Before registering your first email you have to fetch the latest domains.
```golang
var domains []Domain = client.GetAvailableDomains()
```

Then you are able to register a new email. You can do this by using the `Register` function that takes a random username, one domain of the list you got in the step before and a password of your choice.
```golang
client.Register("USERNAME", domains[0], "PASSWORD")
```

### Logging in

After registering an email you have to log in to ave the full funcionality of the API. You don't have to do anything more than just calling the `Login` function.
```golang
client.Login()
```
This step fetched a Bearer-Token in the background which is used for authorization reasons. It's value is saved in:
```golang
client.BearerToken
```

### Get all messages

To get all messages in your inbox from a page use:
```golang
var messages []Message = client.GetMessages(PAGE_NUMBER)
```

### Get one message

Every message has its own id. You can access it via:
```golang
var messageId string = message.MessageId
```

To read just one message add this id to the `GetMessage` function:
```golang
var message Message = client.GetMessage(MESSAGE_ID)
```

### Marking a message as seen

To mark a message as seen you can simply do:
```golang
client.MarkMessageAsSeen(MESSAGE_ID)
```

### Getting the raw message

Getting the raw message in plaintext can be done like this:
```golang
var rawMessage string = client.GetMessageSource(MESSAGE_ID)
```

### Deleting a message

Deleting a message can be archieved through caling the `DeleteMessage` function and passing in the message id.
```golang
client.DeleteMessage(MESSAGE_ID)
```

### Deleting an email

To delete an email you can simply call the `Delete` function and the email will be deleted.
```golang
client.Delete()
```



