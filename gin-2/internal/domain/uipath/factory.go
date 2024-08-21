package uipath

import (
	"github.com/hibiken/asynq"
	"payment-portal/internal/domain/gateway"
)

type JobGroupData struct {
	GatewayId uint
	JobData   []JobInterface
}

func (j *JobGroupData) Enqueue(client *asynq.Client) error {
	for _, job := range j.JobData {
		task, err := job.CreateTask()
		if err != nil {
			return err
		}

		// Enqueue the task using the asynq package
		if _, err := client.Enqueue(task); err != nil {
			return err
		}
	}

	return nil
}

func GetGatewayJob(a *OutputArguments, g *gateway.Gateway) (*JobGroupData, error) {
	// Function body
	xgroupJobs, err := a.ToXgroup()
	if err != nil {
		return nil, err
	}

	// Convert the XgroupJob instances to JobInterface instances
	jobData := make([]JobInterface, len(xgroupJobs))
	for i, job := range xgroupJobs {
		jobData[i] = &job
	}

	return &JobGroupData{
		GatewayId: g.ID,
		JobData:   jobData,
	}, nil
}
