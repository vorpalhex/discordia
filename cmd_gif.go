package main

import (
  "strings"
  "math/rand"

  "github.com/bwmarrin/discordgo"
  "github.com/paddycarey/gophy"
  "github.com/golang/glog"
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
      glog.Error("Failed to get trending gif", err)
      return false
    }
    gif = gifs[rand.Intn(len(gifs))].URL
  }else{
    gifs, _, err := gophyClient.SearchGifs(strings.Join(search, " "), "", 50, 0)
    if err != nil {
      glog.Error("Failed to get gif via search", err)
      return false
    }
    gif = gifs[rand.Intn(len(gifs))].URL
  }

  sendMsg(s, m.ChannelID, gif)
  return true
}

func init(){
  gophyOpts = &gophy.ClientOptions{}
  if len(config.GiphyKey) > 0 {
    gophyOpts.ApiKey = config.GiphyKey
  }
  gophyClient = gophy.NewClient(gophyOpts)

  registerCmd("!gif", cmd_gif)
}
