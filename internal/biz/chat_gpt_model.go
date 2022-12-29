package biz

import (
	"gorm.io/gorm"
)

type ChatGPTType uint8

const (
	ChatGPTTypeReply ChatGPTType = iota + 1
	ChatGPTTypeQuote
)

type ChatGPT struct {
	gorm.Model
	Appid  string      `gorm:"type:char(18);"`
	Openid string      `gorm:"type:char(28);"`
	Code   string      `gorm:"type:char(32);"`
	Prompt string      `gorm:"type:varchar(1024);"`
	Result string      `gorm:"type:longtext;"`
	Type   ChatGPTType `gorm:"type:tinyint;"`
}

func (ChatGPT) TableName() string {
	return "chat_gpt"
}

type ChatGPTQuota struct {
	gorm.Model
	Openid      string `gorm:"type:char(28);"`
	Date        uint64 `gorm:"type:int(10);"`
	UseCount    uint64 `gorm:"type:int(10);"`
	UnusedCount uint64 `gorm:"type:int(10);"`
}

func (ChatGPTQuota) TableName() string {
	return "chat_gpt_quota"
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
