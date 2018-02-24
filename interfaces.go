package fabric

type Authorizer interface {
  SetAuthorization(authorization string)
}

type Apper interface {
  Authorizer
  SetOrganizationID(organizationID string)
  SetAppID(appID string)
}

type Ranger interface {
  Apper
  SetStart(start *int64)
  SetEnd(end *int64)
}
