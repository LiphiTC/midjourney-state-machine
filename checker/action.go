package checker

import (
	"fmt"

	"github.com/goapi-ai/midjourney-state-machine/model"
	"golang.org/x/exp/slices"
)

/*
 * check whether the action is already triggered or not
 * action: user input action
 * actionToTaskIdMap: taskId map for actions, key is action, value is taskId
 */
func CheckActionTriggered(action string, actionToTaskIdMap map[string]string) (taskId string, err error) {
	// reroll action can be triggered multiple times, so it is not recorded in the map
	if action == model.ActionReroll {
		return "", nil
	}
	if taskId, ok := actionToTaskIdMap[action]; ok {
		return taskId, nil
	}
	return "", fmt.Errorf("invalid action: %s", action)
}

/*
 * create an empty map for action to taskId, used to record triggered taskId
 */
func CreateActionToTaskIdMap(actions []string) map[string]string {
	actionToTaskIdMap := make(map[string]string, len(actions))
	for _, action := range actions {
		if action == model.ActionReroll {
			continue
		}
		actionToTaskIdMap[action] = ""
	}
	return actionToTaskIdMap
}

func GetAvailableActionsFromMap(parentAction string, actionToTaskIdMap map[string]string) (actions []string) {
	for action, taskId := range actionToTaskIdMap {
		if taskId == "" {
			actions = append(actions, action)
		}
	}
	// all action results can be rerolled, except upscael
	if !slices.Contains(model.ActionsUpscale, parentAction) {
		actions = append(actions, model.ActionReroll)
	}
	return
}

/*
 * get available action list
 * panState: verticle if pan up/down was executed in history, horizontal if pan left/right was executed, cleared at outpaint action
 */
func GetActions(parentAction, action, panState string) ([]string, string) {
	switch action {
	case model.ActionImagine, model.ActionBlend:
		fallthrough
	case model.ActionVariation1, model.ActionVariation2, model.ActionVariation3, model.ActionVariation4:
		fallthrough
	case model.ActionVaryStrong, model.ActionVarySubtle:
		fallthrough
	case model.ActionInpaint:
		fallthrough
	case model.ActionOutpaint50, model.ActionOutpaint75, model.ActionOutpaintCustom, model.ActionOutpaintMakeSquare:
		fallthrough
	case model.ActionDescribe1, model.ActionDescribe2, model.ActionDescribe3, model.ActionDescribe4:
		return model.State1Actions, ""
	case model.ActionDescribe:
		return model.State2Actions, ""
	case model.ActionReroll:
		if parentAction == model.ActionDescribe {
			return model.State2Actions, ""
		} else if slices.Contains(model.ActionsPan, parentAction) {
			return model.State4Actions, panState
		} else {
			return model.State1Actions, ""
		}
	case model.ActionUpscale1, model.ActionUpscale2, model.ActionUpscale3, model.ActionUpscale4:
		if panState == model.PanStateVertical {
			return model.State5Actions, panState
		} else if panState == model.PanStateHorizontal {
			return model.State6Actions, panState
		} else {
			return model.State3Actions, panState
		}
	case model.ActionPanUp, model.ActionPanDown, model.ActionPanLeft, model.ActionPanRight:
		if action == model.ActionPanUp || action == model.ActionPanDown {
			return model.State4Actions, model.PanStateVertical
		} else {
			return model.State4Actions, model.PanStateHorizontal
		}
	default:
		return []string{}, ""
	}
}
