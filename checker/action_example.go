package checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/goapi-ai/midjourney-state-machine/model"
	"github.com/google/uuid"
)

var initMessage = `action checker example starts
- input ` + "`action`" + ` such as ` + "`upscale1`" + `/` + "`outpaint_custom`" + ` based on last task
- input ` + "`taskId::action`" + ` such as ` + "`sampleId::pan_left`" + ` to execute action on target task
- input ` + "`quit`" + ` to exit
init available actions: [imagine, blend, describe]
`

// simplified TaskRuntime for example
type TaskRuntime struct {
	TaskId     string
	PanState   string
	Action     string
	ActionTree map[string]string
}

func RunActionCheckerExample() {
	// local storage for test, redis etc. can be used in production
	var taskRuntimes = make(map[string]*TaskRuntime)
	// simulate a init task
	parentTaskId := "init"
	taskRuntimes[parentTaskId] = &TaskRuntime{
		TaskId:     parentTaskId,
		PanState:   "",
		Action:     "",
		ActionTree: CreateActionToTaskIdMap([]string{model.ActionImagine, model.ActionBlend, model.ActionDescribe}),
	}
	fmt.Println(initMessage)

	// read user input in loop, simulate multi-round user operations
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("please input `action` or `taskId::action` or `quit`:")
		scanner.Scan()
		input := scanner.Text()
		taskId, action := ParseInput(input)

		if action == "quit" {
			break
		}

		// execute action on target task if taskId is provided, otherwise execute action on last task
		if taskId != "" {
			parentTaskId = taskId
		}
		parentTaskRuntime := taskRuntimes[parentTaskId]
		taskId, err := CheckActionTriggered(action, parentTaskRuntime.ActionTree)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("available actions: %v\n\n", GetAvailableActionsFromMap(parentTaskRuntime.Action, parentTaskRuntime.ActionTree))
			continue
		}
		if taskId != "" {
			parentTaskId = taskId
			parentTaskRuntime = taskRuntimes[taskId]
			fmt.Printf("[repeat task] taskId: [%s], actions: %v\n\n", taskId, GetAvailableActionsFromMap(parentTaskRuntime.Action, parentTaskRuntime.ActionTree))
			continue
		}

		taskId = uuid.New().String()
		actions, panState := GetActions(parentTaskRuntime.Action, action, parentTaskRuntime.PanState)
		taskRuntime := &TaskRuntime{
			TaskId:     taskId,
			PanState:   panState,
			Action:     action,
			ActionTree: CreateActionToTaskIdMap(actions),
		}

		// reroll action does not change action, thus doesn't need to be recorded in ActionTree
		if action == model.ActionReroll {
			taskRuntime.Action = taskRuntimes[parentTaskId].Action
		} else {
			parentTaskRuntime.ActionTree[action] = taskId
		}
		taskRuntimes[taskId] = taskRuntime
		fmt.Printf("[parent task] taskId: [%s], panState: [%s], action: [%s], actionTree: %v\n",
			parentTaskRuntime.TaskId, parentTaskRuntime.PanState, parentTaskRuntime.Action, parentTaskRuntime.ActionTree)
		fmt.Printf("[new task] taskId: [%s], actions: %v\n\n", taskId, actions)
		parentTaskId = taskId
	}
}

func ParseInput(input string) (taskId string, action string) {
	params := strings.Split(input, "::")
	if len(params) == 2 {
		return params[0], params[1]
	}
	return "", params[0]
}
