package fabric

import (
  "fmt"

  "github.com/TommyO/fabric.io/client"
  "github.com/TommyO/fabric.io/client/operations"
  "github.com/TommyO/fabric.io/models"
  runtime "github.com/go-openapi/runtime/client"

)

type Client struct {
  *client.FabricIo
  Token string
}

func Login(username, password string, debug bool) (*Client, error) {
  c := client.NewHTTPClient(nil)
  c.Transport.(*runtime.Runtime).SetDebug(debug)

  DEFAULT_CLIENT_ID := "2c18f8a77609ee6bbac9e53f3768fedc45fb96be0dbcb41defa706dc57d9c931"
  DEFAULT_CLIENT_SECRET := "092ed1cdde336647b13d44178932cba10911577faf0eda894896188a7d900cc9"

  grant := "password"
  scope := "organizations apps issues features account twitter_client_apps beta software answers"

  login := models.OAuthParamsBody{
    GrantType: &grant,
    Scope: &scope,
    Username: &username,
    Password: &password,
    ClientID: &DEFAULT_CLIENT_ID,
    ClientSecret: &DEFAULT_CLIENT_SECRET,
  }

  auth := operations.NewOAuthParams()
  auth.SetBody(&login)

  resp, err := c.Operations.OAuth(auth)
  if err != nil {
    return nil, err
  }

  return &Client{ c, fmt.Sprintf("Bearer %s", resp.Payload.AccessToken) }, nil
}
