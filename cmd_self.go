package main

import (
  "time"
  "fmt"

  "github.com/bwmarrin/discordgo"
  "github.com/golang/glog"
)

var startTime time.Time

func init(){
  startTime = time.Now()
}

func cmd_self(subcmds []string, s *discordgo.Session, m *discordgo.MessageCreate) bool {
  switch subcmds[0] {
    case "status":
      endTime := time.Now().Sub(startTime)
      msg := fmt.Sprintf("I've been running for %v during which time I've encountered %+v lines of errors.", endTime, glog.Stats.Error.Lines())
      sendMsg(s, m.ChannelID, msg)

    default:
      sendMsg(s, m.ChannelID, "I'm not sure what you're asking...")
  }

  return true
}

func init(){
  registerCmd("!discordia", cmd_self)
}
