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
  GiphyKey string `json:"giphyKey"`
}

var userPresence map[string]string

func init() {
  userPresence = make(map[string]string)

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
      presence,exists := userPresence[m.Presence.User.ID]

      if !exists && m.Presence.Status == "online"{
        sendMsg(dg, config.AnnounceChannel, greet(getUsername(m.Presence.User, s)) )
      } else {
        if presence != "online" && m.Presence.Status == "online" {
          sendMsg(dg, config.AnnounceChannel, greet(getUsername(m.Presence.User, s)) )
        }
      }

      userPresence[m.Presence.User.ID] = m.Presence.Status
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
  //fmt.Printf("Message %+v\n", m.ContentWithMentionsReplaced())
	if strings.HasPrefix(m.Content, "!") || strings.HasPrefix(m.ContentWithMentionsReplaced(), "@"+s.State.User.Username){
    parts := strings.Fields(m.Content)
    cmd := parts[0]
    ops :=  parts[1:]

    wasSuccessful := runCmd(cmd, ops, s, m)
    if !wasSuccessful {
      glog.Warning("Failed to run command " + cmd)
    }
  }
}
