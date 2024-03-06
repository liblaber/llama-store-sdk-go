# LlamaStore Go SDK 0.0.3

A Go SDK for LlamaStore.

- API version: 0.1.7
- SDK version: 0.0.3

The llama store API! Get details on all your favorite llamas.

## To use this API

- You will need to register a user, once done you can request an API token.
- You can then use your API token to get details about the llamas.

## User registration

To register a user, send a POST request to `/user` with the following body:

```json
{
  "email": "<your email>",
  "password": "<your password>"
}
```

This API has a maximum of 1000 current users. Once this is exceeded, older users will be deleted. If your user is deleted, you will need to register again.

## Get an API token

To get an API token, send a POST request to `/token` with the following body:

```json
{
  "email": "<your email>",
  "password": "<your password>"
}
```

This will return a token that you can use to authenticate with the API:

```json
{
  "access_token": "<your new token>",
  "token_type": "bearer"
}
```

## Use the API token

To use the API token, add it to the `Authorization` header of your request:

```
Authorization: Bearer <your token>
```

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

#### CreateLlamaPicture

#### UpdateLlamaPicture

#### DeleteLlamaPicture

### LlamaService

#### GetLlamas

#### CreateLlama

#### GetLlamaById

#### UpdateLlama

#### DeleteLlama

### TokenService

#### CreateApiToken

### UserService

#### GetUserByEmail

#### RegisterUser
