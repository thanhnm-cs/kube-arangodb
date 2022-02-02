//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package reconcile

import (
	"context"

	core "k8s.io/api/core/v1"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil"
	inspectorInterface "github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector"
	"github.com/rs/zerolog"
)

// createRotateMemberPlan creates a plan to rotate (stop-recreate-start) an existing
// member.
func createRotateMemberPlan(log zerolog.Logger, member api.MemberStatus,
	group api.ServerGroup, reason string) api.Plan {
	log.Debug().
		Str("id", member.ID).
		Str("role", group.AsRole()).
		Str("reason", reason).
		Msg("Creating rotation plan")
	plan := api.Plan{
		api.NewAction(api.ActionTypeCleanTLSKeyfileCertificate, group, member.ID, "Remove server keyfile and enforce renewal/recreation"),
		api.NewAction(api.ActionTypeResignLeadership, group, member.ID, reason),
		api.NewAction(api.ActionTypeKillMemberPod, group, member.ID, reason),
		api.NewAction(api.ActionTypeRotateMember, group, member.ID, reason),
		api.NewAction(api.ActionTypeWaitForMemberUp, group, member.ID),
		api.NewAction(api.ActionTypeWaitForMemberInSync, group, member.ID),
	}
	return plan
}

func emptyPlanBuilder(ctx context.Context,
	log zerolog.Logger, apiObject k8sutil.APIObject,
	spec api.DeploymentSpec, status api.DeploymentStatus,
	cachedStatus inspectorInterface.Inspector, context PlanBuilderContext) api.Plan {
	return nil
}

func removeConditionActionV2(actionReason string, conditionType api.ConditionType) api.Action {
	return api.NewAction(api.ActionTypeSetConditionV2, api.ServerGroupUnknown, "", actionReason).
		AddParam(setConditionActionV2KeyAction, string(conditionType)).
		AddParam(setConditionActionV2KeyType, setConditionActionV2KeyTypeRemove)
}

func updateConditionActionV2(actionReason string, conditionType api.ConditionType, status bool, reason, message, hash string) api.Action {
	statusBool := core.ConditionTrue
	if !status {
		statusBool = core.ConditionFalse
	}

	return api.NewAction(api.ActionTypeSetConditionV2, api.ServerGroupUnknown, "", actionReason).
		AddParam(setConditionActionV2KeyAction, string(conditionType)).
		AddParam(setConditionActionV2KeyType, setConditionActionV2KeyTypeAdd).
		AddParam(setConditionActionV2KeyStatus, string(statusBool)).
		AddParam(setConditionActionV2KeyReason, reason).
		AddParam(setConditionActionV2KeyMessage, message).
		AddParam(setConditionActionV2KeyHash, hash)
}

func removeMemberConditionActionV2(actionReason string, conditionType api.ConditionType, group api.ServerGroup, member string) api.Action {
	return api.NewAction(api.ActionTypeSetMemberConditionV2, group, member, actionReason).
		AddParam(setConditionActionV2KeyAction, string(conditionType)).
		AddParam(setConditionActionV2KeyType, setConditionActionV2KeyTypeRemove)
}

func updateMemberConditionActionV2(actionReason string, conditionType api.ConditionType, group api.ServerGroup, member string, status bool, reason, message, hash string) api.Action {
	statusBool := core.ConditionTrue
	if !status {
		statusBool = core.ConditionFalse
	}

	return api.NewAction(api.ActionTypeSetMemberConditionV2, group, member, actionReason).
		AddParam(setConditionActionV2KeyAction, string(conditionType)).
		AddParam(setConditionActionV2KeyType, setConditionActionV2KeyTypeAdd).
		AddParam(setConditionActionV2KeyStatus, string(statusBool)).
		AddParam(setConditionActionV2KeyReason, reason).
		AddParam(setConditionActionV2KeyMessage, message).
		AddParam(setConditionActionV2KeyHash, hash)
}