# MyClient - Go Client for MyServer API

MyClient is a Go client designed to interact with the MyServer API, enabling CRUD (Create, Read, Update, Delete) operations on MyInfo entities. It provides an easy-to-use interface to manage MyInfo data through HTTP requests.

## Table of Contents

- [MyClient - Go Client for MyServer API](#myclient---go-client-for-myserver-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Methods](#api-methods)
    - [CreateMyInfo](#createmyinfo)
    - [GetMyInfo](#getmyinfo)
    - [UpdateMyInfo](#updatemyinfo)
    - [DeleteMyInfo](#deletemyinfo)
  - [Examples](#examples)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

To use MyClient in your Go project, you need to install it using `go get`:

```bash
go get -u github.com/srikanthbhandary-teach/my-client
```

## Usage

1. Import the MyClient package into your Go code:

   ```go
   import "github.com/srikanthbhandary-teach/my-client"
   ```

2. Create a new instance of the client by calling `NewClient`:

   ```go
   apiKey := "your-api-key"
   client := myclient.NewClient(apiKey)
   ```

   Replace `"your-api-key"` with your actual API key.

3. Use the client to interact with the MyServer API. Example:

   ```go
   err := client.CreateMyInfo("1", "Alice", 30)
   if err != nil {
       fmt.Printf("Error creating MyInfo: %v\n", err)
   }
   ```

## API Methods

### CreateMyInfo

```go
func (c *Client) CreateMyInfo(id, name string, age int) error
```

Sends a POST request to create a new MyInfo entity.

- `id`: The ID of the MyInfo entity.
- `name`: The name of the MyInfo entity.
- `age`: The age of the MyInfo entity.

### GetMyInfo

```go
func (c *Client) GetMyInfo(id string) ([]byte, error)
```

Sends a GET request to retrieve a MyInfo entity by its ID.

- `id`: The ID of the MyInfo entity.

### UpdateMyInfo

```go
func (c *Client) UpdateMyInfo(id, name string, age int) error
```

Sends a PUT request to update an existing MyInfo entity.

- `id`: The ID of the MyInfo entity to update.
- `name`: The updated name for the MyInfo entity.
- `age`: The updated age for the MyInfo entity.

### DeleteMyInfo

```go
func (c *Client) DeleteMyInfo(id string) error
```

Sends a DELETE request to delete a MyInfo entity by its ID.

- `id`: The ID of the MyInfo entity to delete.

## Examples

Here's an example of how to use MyClient to interact with the MyServer API:

```go
package main

import (
    "fmt"
    "github.com/srikanthbhandary-teach/my-client"
)

func main() {
    apiKey := "your-api-key"
    client := myclient.NewClient(apiKey)

    // Create a new MyInfo entity
    err := client.CreateMyInfo("1", "Alice", 30)
    if err != nil {
        fmt.Printf("Error creating MyInfo: %v\n", err)
        return
    }

    // Retrieve a MyInfo entity
    data, err := client.GetMyInfo("1")
    if err != nil {
        fmt.Printf("Error retrieving MyInfo: %v\n", err)
        return
    }
    fmt.Printf("Retrieved MyInfo: %s\n", string(data))
}
```

Replace `"your-api-key"` with your real API key.

## Contributing

Contributions are welcome! Please feel free to open issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize this README to suit your project's specific details and requirements.