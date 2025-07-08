package validator

import (
    "encoding/json"
    "errors"
    "log-processor/internal/model"
)

func ValidateLog(raw []byte) (*models.Log, error) {
    var logEntry models.Log
    if err := json.Unmarshal(raw, &logEntry); err != nil {
        return nil, errors.New("invalid JSON structure")
    }

    if logEntry.Level == "" || logEntry.Source == "" || logEntry.Message == "" || logEntry.Timestamp.IsZero() {
        return nil, errors.New("missing required fields")
    }

    return &logEntry, nil
}
