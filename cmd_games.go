package main

import (
  "math/rand"

  "github.com/bwmarrin/discordgo"
)

func cmd_teams(teams []string, s *discordgo.Session, m *discordgo.MessageCreate) bool {
  members := getOnlineUsers(s)

  if len(teams) < 2 {
    sendMsg(s, m.ChannelID, "You need to give me at least two teams...")
  }
  output := "Team's are:\n"
  for _,v := range members {
    teamChoice := rand.Intn(len(teams))
    team := teams[teamChoice]
    output += v + ":" + team + "\n"
  }
  sendMsg(s, m.ChannelID, output)
  return true
}

func init(){
  registerCmd("!teams", cmd_teams)
}
