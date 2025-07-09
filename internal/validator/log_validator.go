package validator

import (
    "encoding/json"
    "errors"
    "github.com/Salah-ZEddine/incident-dashboard-common/models"
    "log"
)

func ValidateLog(raw []byte) (*models.Log, error) {
    var logEntry models.Log
    if err := json.Unmarshal(raw, &logEntry); err != nil {
        return nil, errors.New("invalid JSON structure")
    }

    if err := logEntry.Validate(); err != nil {
        log.Printf("invalid log: %v", err)
        return nil, err
    }

    return &logEntry, nil
}
