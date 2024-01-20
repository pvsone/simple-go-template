package activities

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
)

// Basic activity definition
func Echo1(ctx context.Context, val string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Echo1 activity started", "val", val)

	time.Sleep(1 * time.Second)

	result := val
	return result, nil
}

// Struct-based activity definitions
type EchoActivities struct {
	Number int
}

func (a *EchoActivities) Echo2(ctx context.Context, val string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Echo2 activity started", "val", val)

	time.Sleep(1 * time.Second)

	result := val
	return result, nil
}

func (a *EchoActivities) Echo3(ctx context.Context, val string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Echo3 activity started", "val", val)

	time.Sleep(1 * time.Second)

	result := val
	return result, nil
}

// Structs for input and output
type EchoInput struct {
	Val string
}

type EchoOutput struct {
	Result string
}

func (a *EchoActivities) Echo4(ctx context.Context, in EchoInput) (*EchoOutput, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Echo4 activity started", "val", in.Val)

	time.Sleep(1 * time.Second)

	var result string
	for i := 0; i < a.Number; i++ {
		result += fmt.Sprintf("%s ", in.Val)
	}
	out := &EchoOutput{
		Result: result,
	}
	return out, nil
}
