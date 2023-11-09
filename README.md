# midjourney-state-machine
This repo represents the generation task state machine of Midjourney, as late as November, 2023.

## This repo contains three components:
- A state machine graph build on [Plantuml](http://www.plantuml.com), stored in state_graph.iuml
- A data model written in Golang, as the data structure of the states of the state machine, see /model/state.go
- A set of checker functions written in Golang, it's used in [GoAPI's Midjourney API](https://www.goapi.ai/midjourney-api) for task state checking.

## How to use
Run the test program by executing `go run main.go`.

The program will need you to enter an input. You can input one of the following:
- `action`: This will execute the action on the last task.
- `taskId::action`: This will execute the action on the specified task.
- `quit`: This will exit the program.

After entering an input, the program will process the action and display the results:
- `parent task`: taskId, panState, action, actrionTree of parent task.
- `new task`/`repeat task`: taskId and available actions of input action.

You can input any action from the avilable actions list to continue.

## Action state graph
![action-state-graph](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/goapi-ai/midjourney-state-machine/main/state_graph.iuml)
