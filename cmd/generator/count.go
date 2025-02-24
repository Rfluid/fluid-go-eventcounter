package main

import (
	common_service "github.com/reb-felipe/eventcounter/cmd/common/service"
	eventcounter "github.com/reb-felipe/eventcounter/pkg"
)

func CountMessages(msgs []*eventcounter.Message) map[eventcounter.EventType]map[string]int {
	output := make(map[eventcounter.EventType]map[string]int)

	for _, v := range msgs {
		if _, ok := output[v.EventType]; !ok {
			output[v.EventType] = make(map[string]int)
		}
		if _, ok := output[v.EventType][v.UserID]; !ok {
			output[v.EventType][v.UserID] = 0
		}
		output[v.EventType][v.UserID] += 1
	}

	return output
}

func Write(path string, msgs []*eventcounter.Message) {
	for i, v := range CountMessages(msgs) {
		if err := common_service.CreateAndWriteFile(path, string(i), v); err != nil {
			continue
		}
	}
}
