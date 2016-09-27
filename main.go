package main

import (
  "fmt"
  "flag"
  "strings"

  "github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	token    string
)

func init() {
	flag.StringVar(&token, "t", "", "Account Token")
	flag.Parse()
}

func main() {
  dg, err := discordgo.New("", "", "Bot " + token);
  if err != nil {
    fmt.Println("Failed to log into Discord", err)
    return
  }

  dg.AddHandler(getMsg)
  err = dg.Open()
  if err != nil {
    fmt.Println("Failed to connect to Discord", err)
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
      fmt.Println("Failed to run command " + cmd)
    }
  }
}
