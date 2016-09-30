package main

import (
  "github.com/bwmarrin/discordgo"
)

type cmdHandl func([]string, *discordgo.Session, *discordgo.MessageCreate) bool

var (
  cmds map[string]cmdHandl
)

func init(){
  cmds = make(map[string]cmdHandl)
}


func registerCmd(prefix string, handl cmdHandl) {
  cmds[prefix] = handl
}

func runCmd(cmd string, ops []string, s *discordgo.Session, m *discordgo.MessageCreate) bool {
  hndl, exists := cmds[cmd]
  if exists {
    go hndl(ops, s, m)
    return true
  } else {
    return false
  }
}
