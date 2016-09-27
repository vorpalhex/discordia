# Discordia

A [Discord](http://discordapp.com) bot written in Golang

## Registering

Visit `https://discordapp.com/oauth2/authorize?client_id={MY_CLIENT_ID}&scope=bot&permissions=261120`

## Requirements to Build

[DiscordGo](https://github.com/bwmarrin/discordgo) via `go get https://github.com/bwmarrin/discordgo`

## Running

`discordia -t {my_bot_token}`

## Commands

+ !teams [team1] [team2]...
  Ask Discordia to assign everyone currently online to teams, randomly
+ !gif [optional search terms]
  Pull a (hopefully relevant) gif from Giphy
