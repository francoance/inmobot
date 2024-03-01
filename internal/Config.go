package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

var Configuration InmobotConfiguration

type InmobotConfiguration struct {
	TelegramUserId          int    `json:"TelegramUserId,omitempty"`
	Url                     string `json:"InmoclickUrl,omitempty"`
	RequestTimeoutInSeconds int    `json:"RequestTimeoutInSeconds,omitempty"`
	ScrapPeriodInMinutes    int    `json:"ScrapPeriodInMinutes,omitempty"`
	TelegramBotToken        string `json:"TelegramBotToken,omitempty"`
}

func InitConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic("No configuration file.")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	Configuration = InmobotConfiguration{}
	err = decoder.Decode(&Configuration)
	if err != nil {
		panic(fmt.Sprintf("Error decoding configuration file: %v", err))
	}
}
