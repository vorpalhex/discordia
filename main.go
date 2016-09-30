package main

import (
  "fmt"
  "flag"
  "strings"
  "os"
  "encoding/json"

  "github.com/bwmarrin/discordgo"
  "github.com/golang/glog"
)

// Variables used for command line parameters
var (
	token    string
)

var config struct {
  AnnounceEnabled bool `json:"announceEnabled"`
  AnnounceChannel string `json:"announceChannel"`
  ClientToken string `json:"clientToken"`
  Greetings []string `json:"greetings"`
}

func init() {
	flag.StringVar(&token, "t", "", "Account Token")
	flag.Parse()

  configFile, err := os.Open("config.json")
  if err != nil {
    glog.Error("Failed to load config", err)
    return
  }

  jsonParser := json.NewDecoder(configFile)
  if err = jsonParser.Decode(&config); err != nil {
    glog.Error("Failed to parse config", err)
    return
  }

  if len(token) < 1 {
    token = config.ClientToken
  }

}

func main() {
  dg, err := discordgo.New("Bot " + token);
  if err != nil {
    glog.Error("Failed to log into Discord", err)
    return
  }

  dg.AddHandler(getMsg)
  if config.AnnounceEnabled {
    dg.AddHandler(func(s *discordgo.Session, m *discordgo.PresenceUpdate) {
      //fmt.Printf("Status Update %+v \n", s.State.Guilds[0].Presences[m.Presence.User.ID])
      presences := s.State.Guilds[0].Presences;
      for _,v := range presences {
        if v.User.ID == m.Presence.User.ID {
          fmt.Printf("%+v %+v \n", v, m)
          if v.Status != "online" && m.Presence.Status == "online" {
            sendMsg(dg, config.AnnounceChannel, greet(getUsername(m.Presence.User, s)) )
          }
        }
      }
    })
  }
  err = dg.Open()
  if err != nil {
    glog.Error("Failed to connect to Discord", err)
    return
  }
  fmt.Println("Discordia is listening...")
  <-make(chan struct{})
  return
}

func getMsg(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!") {
    parts := strings.Fields(m.Content)
    cmd := parts[0]
    ops :=  parts[1:]

    wasSuccessful := runCmd(cmd, ops, s, m)
    if !wasSuccessful {
      glog.Warning("Failed to run command " + cmd)
    }
  }
}
