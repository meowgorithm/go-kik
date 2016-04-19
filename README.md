# Go Kik

Bindings for the Kik API. No external dependencies.

## Installation

    go get github.com/meowgorithm/go-kik

## Overview

Kik operates over a send-receive system utilizing REST calls for sending and a
webhook for receiving. Note that SSL is required on the webhook end.

For local developement you could use something like [Localtunnel][lt] to expose
your development machine over SSL.

lt: [http://localtunnel.me/]

## Quick Example

	package main

    import (
        "kik"
        "log"
        "net/http"
    )

    var (
        client *kik.Client
    )

    func main() {

        // Client for making API requests
        client = &kik.Client{
            Username: "username",
            ApiKey:   "api-key",
            Callback: handleMessages,
            Verbose:  true,
        }

        // Check Kik config
        if _, err := client.GetConfig(), err != nil {
            log.Printf("Error reading config: %s", err)
        }

        // Set Kik config
        config := kik.Config{
            Callback: &kikWebhookUrl,
            Features: kik.Features{
                ReceiveReadReceipts: false,
                ReceiveIsTyping: false,
                ManuallySendReadReceipts: false,
                ReceiveDeliveryReceipts: false,
            },
        }
        if err := client.SetConfig(config); err != nil {
            log.Printf("Error setting config: %s\n", err)
        }

        // Incoming webhook handler for Kik
        http.HandleFunc("/", client.Webhook)
        http.ListenAndServe(":8000", nil)
    }

    // Handle incoming messages
    func handleMessages(p kik.Payload, err error) {

        // Custom keyboard
        k := []kik.Keyboard{
            kik.Keyboard{
                Hidden: false,
                Type:   kik.Suggested,
                Responses: []kik.KeyboardResponse{
                    kik.KeyboardResponse{
                        Type: kik.Text,
                        Body: "What a Button",
                    },
                },
            },
        }

        // Interate over incoming messages
        var m []kik.Message
        for _, in := range p.Messages {
            out := Message{
                ChatId: in.ChatId,
                To: in.From,
                Body: "Hello world!",
                Keyboards: k
            }
            m = append(m, out)
        }

        client.SendMessages(m)
    }

## Author

Christian Rocha <christian@rocha.is>
