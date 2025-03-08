package main

import (
    "fmt"
    "log"
    "os"
    "github.com/slack-go/slack"
)

func main() {
    // Get environment variables
    botToken := os.Getenv("SLACK_BOT_TOKEN")
    channelID := os.Getenv("CHANNEL_ID")

    // Validate environment variables
    if botToken == "" {
        log.Fatal("SLACK_BOT_TOKEN is not set")
    }
    if channelID == "" {
        log.Fatal("CHANNEL_ID is not set")
    }

    // Initialize Slack API client
    api := slack.New(botToken)

    // File to upload
    fileArr := []string{"Attention Is All You Need.pdf"}
    channelArr := []string{channelID}

    for i := 0; i < len(fileArr); i++ {
        // Verify file exists
        if _, err := os.Stat(fileArr[i]); os.IsNotExist(err) {
            log.Fatalf("File %s does not exist", fileArr[i])
        }

        // Upload file
        params := slack.FileUploadParameters{
            Channels: channelArr,
            File:     fileArr[i],
        }

        file, err := api.UploadFile(params)
        if err != nil {
            log.Fatalf("Failed to upload file: %v", err)
        }

        fmt.Printf("Successfully uploaded file: %s\nURL: %s\n", file.Name, file.URL)
    }
}

