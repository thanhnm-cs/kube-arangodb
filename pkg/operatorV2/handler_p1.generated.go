package operator

import (
	"context"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
)

type HandleP1Func[P1 any] func(ctx context.Context, p1 P1) (bool, error)

type HandleP1ConditionFunc[P1 any] func(ctx context.Context, p1 P1) (*Condition, bool, error)

type HandleP1ConditionExtract[P1 any] func(ctx context.Context, p1 P1) *api.ConditionList

func HandleP1[P1 any](ctx context.Context, p1 P1, handler ...HandleP1Func[P1]) (bool, error) {
	isChanged := false
	for _, h := range handler {
		changed, err := h(ctx, p1)
		if changed {
			isChanged = true
		}

		if err != nil {
			return isChanged, err
		}
	}

	return isChanged, nil
}

func HandleP1WithStop[P1 any](ctx context.Context, p1 P1, handler ...HandleP1Func[P1]) (bool, error) {
	changed, err := HandleP1[P1](ctx, p1, handler...)
	if IsStop(err) {
		return changed, nil
	}

	return changed, err
}

func HandleP1WithCondition[P1 any](ctx context.Context, conditions *api.ConditionList, condition api.ConditionType, p1 P1, handler ...HandleP1Func[P1]) (bool, error) {
	changed, err := HandleP1[P1](ctx, p1, handler...)
	return WithCondition(conditions, condition, changed, err)
}

func HandleP1Condition[P1 any](extract HandleP1ConditionExtract[P1], condition api.ConditionType, handler HandleP1ConditionFunc[P1]) HandleP1Func[P1] {
	return func(ctx context.Context, p1 P1) (bool, error) {
		c, changed, err := handler(ctx, p1)
		return WithConditionChange(extract(ctx, p1), condition, c, changed, err)
	}
}
