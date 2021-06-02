package hostcommands

import (
	"context"

	"github.com/openshift/assisted-service/models"
)

type emptyCmd struct {
}

func NewEmptyCmd() *emptyCmd {
	return &emptyCmd{}
}

func (c *emptyCmd) GetSteps(ctx context.Context, host *models.Host) ([]*models.Step, error) {
	step := &models.Step{
		StepType: "",
		Command:  "",
		Args:     nil,
	}

	return []*models.Step{step}, nil
}
