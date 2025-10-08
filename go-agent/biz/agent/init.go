package agent

import (
	"context"
	"log"
)

func InitAll(ctx context.Context) {
	log.Printf("init all agent")
	InitPlanRunner(ctx)
	InitRecommendRunner(ctx)
	log.Printf("init all agent done")
}