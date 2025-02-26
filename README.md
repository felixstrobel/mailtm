# mailtm

The `mailtm` library allows communication with the [Mail.tm API](https://api.mail.tm), providing functionality for
creating disposable email accounts and managing messages. Below are detailed instructions on how to install and use the
wrapper
effectively.

### Documentation

---

#### 1. **Installing the Package**

Run the following command to install the `mailtm` module:

```bash
go get github.com/felixstrobel/mailtm
```

---

#### 2. **Creating a New `Client`**

Start by creating a `Client` object. This object will act as the main entry point for interacting with the API.

```go
import "github.com/felixstrobel/mailtm"

func main() {
   client, err := mailtm.New()
   if err != nil {
    panic(err)
   }

   // Use the `client` object to perform further operations
}
```

#### Using `WithBaseURL` and `WithHTTPClient` Options

When creating a `MailClient`, you can customize its behavior using the `WithBaseURL` and `WithHTTPClient` options:

1. **WithBaseURL**
    - Use this option to specify a custom base URL for the Mail.tm API.
      ```go
      client, err := mailtm.New(
          mailtm.WithBaseURL("https://custom.mail.api"), // Specify a custom API URL
      )
      ```

2. **WithHTTPClient**
    - Use this option to specify a custom HTTP client for the `MailClient`.
      ```go
      client, err := mailtm.New(
          mailtm.WithHTTPClient(&http.Client{}), // Use a custom HTTP client
      )
      ```

These options allow you to tailor the `MailClient` for specific requirements, such as using a proxy-enabled HTTP client
or pointing to another instance of the Mail.tm API.
