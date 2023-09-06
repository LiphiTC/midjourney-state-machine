# midjourney-state-machine
This repo represents the current task state machine of Midjourney, as late as September, 2023.

## This repo contains three component
- A state machine graph build on [Plantuml](http://www.plantuml.com), stored in state_graph.iuml
- A data model written in Golang, as the data structure of the states of the state machine, see /model/state.go
- A set of checker functions written in Golang, it's used in [GoAPI.ai Midjourney API](https://goapi.ai/midjourney-api) for task state checking.

## How to use

## Action state graph
![action-state-graph](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/goapi-ai/midjourney-state-machine/main/state_graph.iuml)
