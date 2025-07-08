package models

import "time"

type Log struct {
    ID        uint      `gorm:"primaryKey"`
    Timestamp time.Time `json:"timestamp"`
    Level     string    `json:"level"`
    Source    string    `json:"source"`
    Message   string    `json:"message"`
    Context   string    `json:"context"` 
}

// LogInput is used just for incoming POST requests
// type LogInput struct {
//     Timestamp time.Time `json:"timestamp" binding:"required"`
//     Level     string    `json:"level" binding:"required"`
//     Source    string    `json:"source" binding:"required"`
//     Message   string    `json:"message" binding:"required"`
//     Context   string    `json:"context"`
// }