package main

import (
  "strings"
  "math/rand"

  "github.com/bwmarrin/discordgo"
  "github.com/paddycarey/gophy"
)

var (
  gophyOpts *gophy.ClientOptions
  gophyClient *gophy.Client
)

func cmd_gif(search []string, s *discordgo.Session, m *discordgo.MessageCreate) bool {
  var gif string

  if len(search) < 1 {
    gifs, err := gophyClient.TrendingGifs("", 20)
    if err != nil {
      return false
    }
    gif = gifs[rand.Intn(len(gifs))].URL
  }else{
    gifs, _, err := gophyClient.SearchGifs(strings.Join(search, " "), "", 1, 0)
    if err != nil {
      return false
    }
    gif = gifs[0].URL
  }

  sendMsg(s, m.ChannelID, gif)
  return true
}

func init(){
  gophyOpts = &gophy.ClientOptions{}
  gophyClient = gophy.NewClient(gophyOpts)

  registerCmd("!gif", cmd_gif)
}
