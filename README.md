# LlamaStore Go SDK 0.0.3

A Go SDK for LlamaStore.

- API version: 0.1.7
- SDK version: 0.0.3

The llama store API! Get details on all your favorite llamas. ## To use this API - You will need to register a user, once done you can request an API token. - You can then use your API token to get details about the llamas. ## User registration To register a user, send a POST request to `/user` with the following body: `json {      email :  <your email> ,      password :  <your password>  } ` This API has a maximum of 1000 current users. Once this is exceeded, older users will be deleted. If your user is deleted, you will need to register again. ## Get an API token To get an API token, send a POST request to `/token` with the following body: `json {      email :  <your email> ,      password :  <your password>  } ` This will return a token that you can use to authenticate with the API: `json {    access_token :  <your new token> ,    token_type :  bearer  } ` ## Use the API token To use the API token, add it to the `Authorization` header of your request: `Authorization: Bearer <your token>`

## Table of Contents

- [Authentication](#authentication)
- [Services](#services)

## Authentication

### Access Token

The llama-store API uses a access token as a form of authentication.

The access token can be set when initializing the SDK like this:

```go
// Constructor initialization
```

Or at a later stage:

```go
// Setter initialization
```

## Services

### LlamaPictureService

#### GetLlamaPictureByLlamaId

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
)

sdk := llamastore.NewLlamaStore()


params := llamapicture.GetLlamaPictureByLlamaIdRequestParams{}
params.SetLlamaId(int64(123))

response, err := sdk.LlamaPicture.GetLlamaPictureByLlamaId("llamaId")
```

#### CreateLlamaPicture

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/llamapicture"
)

sdk := llamastore.NewLlamaStore()


params := llamapicture.CreateLlamaPictureRequestParams{}
params.SetLlamaId(int64(123))

llamaId, err := sdk.LlamaPicture.CreateLlamaPicture("llamaId", request)

jsonData, err := json.MarshalIndent(llamaId)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

#### UpdateLlamaPicture

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/llamapicture"
)

sdk := llamastore.NewLlamaStore()


params := llamapicture.UpdateLlamaPictureRequestParams{}
params.SetLlamaId(int64(123))

llamaId, err := sdk.LlamaPicture.UpdateLlamaPicture("llamaId", request)

jsonData, err := json.MarshalIndent(llamaId)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

#### DeleteLlamaPicture

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
)

sdk := llamastore.NewLlamaStore()


params := llamapicture.DeleteLlamaPictureRequestParams{}
params.SetLlamaId(int64(123))

response, err := sdk.LlamaPicture.DeleteLlamaPicture("llamaId")
```

### LlamaService

#### GetLlamas

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
)

sdk := llamastore.NewLlamaStore()

response, err := sdk.Llama.GetLlamas()
```

#### CreateLlama

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama"
)

sdk := llamastore.NewLlamaStore()

llamaColor := "brown"

request := llama.LlamaCreate{}
request.SetName("Name")
request.SetAge(int64(123))
request.SetColor(llamaColor)
request.SetRating(int64(123))

llama, err := sdk.Llama.CreateLlama(request)

jsonData, err := json.MarshalIndent(llama)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

#### GetLlamaById

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama"
)

sdk := llamastore.NewLlamaStore()


params := llama.GetLlamaByIdRequestParams{}
params.SetLlamaId(int64(123))

llama, err := sdk.Llama.GetLlamaById("llamaId")

jsonData, err := json.MarshalIndent(llama)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

#### UpdateLlama

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama"
)

sdk := llamastore.NewLlamaStore()

llamaColor := "brown"

request := llama.LlamaCreate{}
request.SetName("Name")
request.SetAge(int64(123))
request.SetColor(llamaColor)
request.SetRating(int64(123))


params := llama.UpdateLlamaRequestParams{}
params.SetLlamaId(int64(123))

llama, err := sdk.Llama.UpdateLlama("llamaId", request)

jsonData, err := json.MarshalIndent(llama)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

#### DeleteLlama

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
)

sdk := llamastore.NewLlamaStore()


params := llama.DeleteLlamaRequestParams{}
params.SetLlamaId(int64(123))

response, err := sdk.Llama.DeleteLlama("llamaId")
```

### TokenService

#### CreateApiToken

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/token"
)

sdk := llamastore.NewLlamaStore()


request := token.ApiTokenRequest{}
request.SetEmail("Email")
request.SetPassword("Password")

apiToken, err := sdk.Token.CreateApiToken(request)

jsonData, err := json.MarshalIndent(apiToken)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

### UserService

#### GetUserByEmail

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/user"
)

sdk := llamastore.NewLlamaStore()


params := user.GetUserByEmailRequestParams{}
params.SetEmail("string")

user, err := sdk.User.GetUserByEmail("email")

jsonData, err := json.MarshalIndent(user)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```

#### RegisterUser

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/liblaber/llama-store-sdk-go/pkg/llama-store"
  "github.com/liblaber/llama-store-sdk-go/pkg/user"
)

sdk := llamastore.NewLlamaStore()


request := user.UserRegistration{}
request.SetEmail("Email")
request.SetPassword("Password")

user, err := sdk.User.RegisterUser(request)

jsonData, err := json.MarshalIndent(user)
if err != nil {
  fmt.Printf("Error marshalling JSON: %v\n", err)
  return
}

fmt.Printf("%s\n", jsonData)
```
