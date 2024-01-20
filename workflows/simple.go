package workflows

import (
	"simple-go/activities"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type SimpleInput struct {
	Val string
}

type SimpleOutput struct {
	Result string
}

func Simple(ctx workflow.Context, input SimpleInput) (*SimpleOutput, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Simple workflow started", "val", input.Val)

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	var result1 string
	err := workflow.ExecuteActivity(ctx, activities.Echo1, input.Val).Get(ctx, &result1)
	if err != nil {
		return nil, err
	}

	logger.Info("Sleeping for 1 second...")
	workflow.Sleep(ctx, 1*time.Second)

	var echoActivity *activities.EchoActivities

	var result2 string
	err = workflow.ExecuteActivity(ctx, echoActivity.Echo2, result1).Get(ctx, &result2)
	if err != nil {
		return nil, err
	}

	var result3 string
	err = workflow.ExecuteActivity(ctx, echoActivity.Echo3, result2).Get(ctx, &result3)
	if err != nil {
		return nil, err
	}

	echoInput := activities.EchoInput{Val: result3}
	var echoOutput activities.EchoOutput

	echoFuture := workflow.ExecuteActivity(ctx, echoActivity.Echo4, echoInput)
	err = echoFuture.Get(ctx, &echoOutput)
	if err != nil {
		return nil, err
	}

	output := &SimpleOutput{
		Result: echoOutput.Result,
	}
	return output, nil
}
