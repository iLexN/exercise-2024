package uipath

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

type WebHookPayload struct {
	Type               string    `json:"Type"`
	EventID            string    `json:"EventId"`
	Timestamp          time.Time `json:"Timestamp"`
	Job                Job       `json:"Job"`
	TenantID           int       `json:"TenantId"`
	OrganizationUnitID int       `json:"OrganizationUnitId"`
}

func (w *WebHookPayload) ProcessId() int {
	return w.Job.Release.ID
}

func (w *WebHookPayload) IsSuccessful() bool {
	return w.Job.State == "Successful"
}

type OutEodBalance struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

type JobInterface interface {
	CreateTask() (*asynq.Task, error)
}

type XgroupInput struct {
	ID                string `json:"Id"`
	CreateDate        string `json:"Create date"`
	Account           string `json:"Account"`
	Type              string `json:"Type"`
	Currency          any    `json:"Currency"`
	Amount            string `json:"Amount"`
	TransactionAmount any    `json:"Transaction amount"`
	Comment           string `json:"Comment"`
	TransactionID     any    `json:"Transaction id"`
	ExternalID        any    `json:"External id"`
	Balance           string `json:"Balance"`
}

func (x *XgroupInput) CreateTask() (*asynq.Task, error) {
	payload, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask("job:xgroup", payload), nil
}

type OutputArguments struct {
	OutAvailableBalance any             `json:"outAvailableBalance"`
	OutEodBalance       []OutEodBalance `json:"outEodBalance"`
	InToDateTime        string          `json:"inToDateTime"`
	OutCsvData          json.RawMessage `json:"outCsvData"`
	InFromDateTime      string          `json:"inFromDateTime"`
	OutPeriodType       string          `json:"outPeriodType"`
}

func (o OutputArguments) ToXgroup() ([]XgroupInput, error) {
	var xgroupInputs []XgroupInput

	// Unmarshal the OutCsvData field
	err := json.Unmarshal(o.OutCsvData, &xgroupInputs)
	if err != nil {
		return nil, err
	}

	return xgroupInputs, nil
}

func (o OutputArguments) GetFromDateTime() time.Time {
	datetime, _ := time.Parse("01/02/2006 15:04:05", o.InFromDateTime)
	return datetime
}
func (o OutputArguments) GetToDateTime() time.Time {
	datetime, _ := time.Parse("01/02/2006 15:04:05", o.InToDateTime)
	return datetime
}

type Robot struct {
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
	MachineName string `json:"MachineName"`
}
type Release struct {
	ID         int    `json:"Id"`
	Key        string `json:"Key"`
	ProcessKey string `json:"ProcessKey"`
}
type Job struct {
	ID              int             `json:"Id"`
	Key             string          `json:"Key"`
	State           string          `json:"State"`
	StartTime       time.Time       `json:"StartTime"`
	EndTime         time.Time       `json:"EndTime"`
	Info            string          `json:"Info"`
	OutputArguments OutputArguments `json:"OutputArguments"`
	Robot           Robot           `json:"Robot"`
	Release         Release         `json:"Release"`
}
