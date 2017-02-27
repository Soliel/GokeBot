package main

import (
  "github.com/bwmarrin/discordgo"
  "strconv"
  "strings"
  "bytes"
  "fmt"
)

const (
  PREFIX = "!goke "
)

var (
  BotID   string
  karaMap map[string][]string
)

func main () {
  karaMap = make(map[string][]string)
  karaMap["276616615582367747"] = []string{}
  
  fmt.Println(karaMap)
  
  dg, err := discordgo.New("Bot " + "MjEwNTU3NjQ5ODE5OTI2NTMx.C2HnBQ.wLuZ-SVkUHMiz2JwfqMUq96rckc")
  if err != nil {
    fmt.Println("Failed to authenticate with discord ", err)
    return
  }
  
  u, err := dg.User("@me")
  if err != nil {
    fmt.Println("Failed to obtain current user ", err)
    return
  }
  
  dg.AddHandler(onMessageRecieved)
  
  err = dg.Open()
  if err != nil {
    fmt.Println("Error starting discord listener ", err)
    return
  }
  
  fmt.Println("Bot is now running as user: ", u.Username)
  
  <- make(chan struct{})
}

func onMessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == BotID {
    return
  }
  
  if len(m.Content) < len(PREFIX) {
    return
  }
  
  if m.Content[:len(PREFIX)] != PREFIX {
    return
  }
  
  content := m.Content[len(PREFIX):]
  if len(content) < 1 {
    return
  }
  
  content = strings.ToLower(content)
  args :=  strings.Fields(content)
  
  fmt.Println("Looking up commands.")
  
  switch args[0] {
    case "queue":
      go karaokeQueue(m.ChannelID, s)
  }
}

func karaokeQueue(cID string, s *discordgo.Session) {
  fmt.Println("Started command at channel: " + cID)
  fmt.Println(karaMap[cID])
  fmt.Println(len(karaMap[cID]))
  var QueueList bytes.Buffer
  
  queueArray := karaMap[cID]
  
  if len(queueArray) >= 2 {
  
    QueueList.WriteString("**Current: " + queueArray[0] + "**")
    QueueList.WriteString("\n\nUp Next: " + queueArray[1])
    QueueList.WriteString("\n\n__Queue__")
  
    for index, name := range karaMap[cID] {
      QueueList.WriteString("\n"+ strconv.FormatInt(int64(index), 10) + ": " + name)
    }
  } else if len(karaMap[cID]) == 1 {
     QueueList.WriteString("**Current: " + queueArray[0] + "**")
  } else {
    return
  }
  
  embed := &discordgo.MessageEmbed{
    Color: 0xFFB6C1,
    Fields: []*discordgo.MessageEmbedField{
      {"Karaoke List", QueueList.String(), true},
    },
  }
  
  _, err := s.ChannelMessageSendEmbed(cID, embed)
  if err != nil {
    fmt.Println(err)
  }
}


