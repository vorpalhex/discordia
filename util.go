package main

import (
  "fmt"
  "math/rand"

  "github.com/bwmarrin/discordgo"
  "github.com/golang/glog"
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

func getUsername(u *discordgo.User, s *discordgo.Session) string {
  if len(u.Username) > 0 {
    return u.Username
  }
  user,err := s.User(u.ID)
  if err != nil {
    glog.Error("Failed to resolve username", err)
  }
  return user.Username
}

func sendMsg(s *discordgo.Session, cID string, content string) {
  _, err := s.ChannelMessageSend(cID, content)
  if err != nil {
    glog.Error("Failed to send message", err)
  }
  return
}

func greet(who string) string {
  greetings := config.Greetings

  greeting := greetings[ rand.Intn(len(greetings)) ]

  return fmt.Sprintf(greeting, who)
}
