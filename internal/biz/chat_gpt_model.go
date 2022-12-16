package biz

import (
	"gorm.io/gorm"
)

type ChatGPT struct {
	gorm.Model
	Appid  string `gorm:"type:char(18);"`
	Openid string `gorm:"type:char(28);"`
	Code   string `gorm:"type:char(32);"`
	Prompt string `gorm:"type:varchar(1024);"`
	Result string `gorm:"type:longtext;"`
}

func (ChatGPT) TableName() string {
	return "chat_gpt"
}

type CompletionsResp struct {
	Error struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    interface{} `json:"code"`
	} `json:"error"`
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
