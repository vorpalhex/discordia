package main

import (
  "fmt"
  
  "github.com/bwmarrin/discordgo"
)

func getOnlineUsers(s *discordgo.Session) []string{
  m := s.State.Guilds[0].Presences
  users := make([]string, 0)

  for _, v := range m {
    user,_ := s.User(v.User.ID)
    if v.Status == "online" && user.Username != "discordia" {
      users = append(users, user.Username)
    }
  }
  return users
}

func sendMsg(s *discordgo.Session, cID string, content string) {
  _, err := s.ChannelMessageSend(cID, content)
  if err != nil {
    fmt.Println("Failed to send message", err)
  }
  return
}
