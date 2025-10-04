package agent

import "context"

func InitAll(ctx context.Context) {
	InitPlanRunner(ctx)
	InitRecommendRunner(ctx)
}