package main

import (
    "fmt"
    "time"
    "strings"
)

const (
    ColorReset  = "\033[0m"
    ColorBlue   = "\033[34m" // Email
    ColorGreen  = "\033[32m" // SMS
    ColorCyan   = "\033[36m" // Push
    ColorRed    = "\033[31m" // Erreur
)

type Notificator interface {
    Send(message string) error
}

type EmailNotificator struct {
    Address string
}

func (e EmailNotificator) Send(message string) error {
    fmt.Printf("%sEmail envoyé à %s: %s%s\n", ColorBlue, e.Address, message, ColorReset)
    return nil
}

type SMSNotificator struct {
    Phone string
}

func (s SMSNotificator) Send(message string) error {
    if !strings.HasPrefix(s.Phone, "06") {
        return fmt.Errorf("%sNuméro invalide: %s%s", ColorRed, s.Phone, ColorReset)
    }
    fmt.Printf("%sSMS envoyé à %s: %s%s\n", ColorGreen, s.Phone, message, ColorReset)
    return nil
}

type PushNotificator struct {
    DeviceID string
}

func (p PushNotificator) Send(message string) error {
    fmt.Printf("%sPush envoyé à %s: %s%s\n", ColorCyan, p.DeviceID, message, ColorReset)
    return nil
}

type NotificationLog struct {
    Message   string
    Timestamp time.Time
}

type Storer struct {
    Logs []NotificationLog
}

func (s *Storer) Archive(message string) {
    s.Logs = append(s.Logs, NotificationLog{
        Message:   fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), message),
        Timestamp: time.Now(),
    })
}

func (s Storer) PrintHistory() {
    fmt.Println("\nHistorique des notifications archivées :")
    for _, log := range s.Logs {
        fmt.Println(log.Message)
    }
}

func main() {
    notificators := []Notificator{
        EmailNotificator{Address: "user@example.com"},
        SMSNotificator{Phone: "0612345678"},
        SMSNotificator{Phone: "0123456789"}, // Erreur attendue
        PushNotificator{DeviceID: "device123"},
    }

    storer := &Storer{}

    for _, n := range notificators {
        message := "Hello, notification !"
        err := n.Send(message)
        if err != nil {
            fmt.Printf("%sErreur d'envoi : %v%s \n", ColorRed, err, ColorReset)
        } else {
            storer.Archive(message)
        }
    }

    storer.PrintHistory()
}