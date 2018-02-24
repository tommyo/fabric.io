package main

import (
  "fmt"
  "encoding/json"
  fabric "github.com/TommyO/fabric.io"
  fabops "github.com/TommyO/fabric.io/client/operations"
)

var client *fabric.Client
var org_id = "57b62e8e01fd3e1880000059"

var ios = "57c09a055547416e5b000054"
var android = "57c09a0c893cfa60f3000034"


func main() {
  var err error
  client, err = fabric.Login("tom@ussieapp.com", "5n8bytdP", true)
  if err != nil {
    panic(err.Error())
  }

// err = ActiveNow(android, "all")
// err = DailyNew(ios, 1480982400, 1512518400)
// err = Active(ios, "daily", 1480982400, 1512518400)
// err = Active(ios, "weekly", 1480982400, 1512518400)
// err = Active(ios, "monthly", 1480982400, 1512518400)
// err = Sessions(ios, 1480982400, 1512518400)
// err = TimeInApp(ios, 1480982400, 1512518400)
  err = SessionsCount(ios, 1480982400, 1512518400)

  if err != nil {
    fmt.Println("ERROR: ", err.Error())
  }
}

func ActiveNow(app string, build string) error {
  opts := fabops.NewActiveNowParams()
  opts.SetAuthorization(client.Token)
  opts.SetOrganizationID(org_id)
  opts.SetAppID(app)
  opts.SetBuild(&build)
//  opts.SetStart(1510012800)
//  opts.SetEnd(1512604800)
  results, err := client.Operations.ActiveNow(opts)
  if err == nil {
    fmt.Printf("%v\n", results.Payload.Cardinality)
    fmt.Printf("%v\n", results.Payload.Timestamp)
    fmt.Printf("%v\n", results.Payload.Deltas.WeekOverWeek.DeltaFraction)
    fmt.Printf("%v\n", results.Payload.Deltas.WeekOverWeek.LastWeekValue)
  }
  return err
}

func DailyNew(app string, start float64, end float64) error {
  opts := fabops.NewDailyNewParams()
  opts.SetAuthorization(client.Token)
  opts.SetOrganizationID(org_id)
  opts.SetAppID(app)
  opts.SetStart(&start)
  opts.SetEnd(&end)
  results, err := client.Operations.DailyNew(opts)
  if err == nil {
    j, _ := json.Marshal(results.Payload)
    fmt.Printf("%s\n", j)
  }
  return err
}

func Active(app string, period string, start float64, end float64) error {
  opts := fabops.NewActiveParams()
  opts.SetAuthorization(client.Token)
  opts.SetOrganizationID(org_id)
  opts.SetAppID(app)
  opts.SetPeriod(period)
  opts.SetStart(&start)
  opts.SetEnd(&end)
  results, err := client.Operations.Active(opts)
  if err == nil {
    j, _ := json.Marshal(results.Payload)
    fmt.Printf("%s\n", j)
  }
  return err
}
/*
func Sessions(app string, start float64, end float64) error {
  opts := fabops.NewSessionsParams()
  opts.SetAuthorization(client.Token)
  opts.SetOrganizationID(org_id)
  opts.SetAppID(app)
  opts.SetStart(&start)
  opts.SetEnd(&end)
  results, err := client.Operations.Sessions(opts)
  if err == nil {
    j, _ := json.Marshal(results.Payload)
    fmt.Printf("%s\n", j)
  }
  return err
}
*/

func TimeInApp(app string, start float64, end float64) error {
  opts := fabops.NewTimeInAppParams()
  opts.SetAuthorization(client.Token)
  opts.SetOrganizationID(org_id)
  opts.SetAppID(app)
  opts.SetStart(&start)
  opts.SetEnd(&end)
  results, err := client.Operations.TimeInApp(opts)
  if err == nil {
    j, _ := json.Marshal(results.Payload)
    fmt.Printf("%s\n", j)
  }
  return err
}

func SessionsCount(app string, start float64, end float64) error {
  opts := fabops.NewSessionsCountParams()
  opts.SetAuthorization(client.Token)
  opts.SetOrganizationID(org_id)
  opts.SetAppID(app)
  opts.SetStart(&start)
  opts.SetEnd(&end)
  results, err := client.Operations.SessionsCount(opts)
  if err == nil {
    j, _ := json.Marshal(results.Payload)
    fmt.Printf("%s\n", j)
  }
  return err
}
