/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"log"

	"github.com/cloudwego/eino/schema"

	"github.com/monosolo101/eino-ext/components/model/ollama"
)

func main() {

	ctx := context.Background()
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434",
		Model:   "llama3.2:1b",
	})
	if err != nil {
		log.Printf("NewChatModel failed, err=%v", err)
		return
	}

	err = chatModel.BindTools([]*schema.ToolInfo{
		{
			Name: "user_company",
			Desc: "根据用户的姓名和邮箱，查询用户的公司和职位信息",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{
					"name": {
						Type: "string",
						Desc: "用户的姓名",
					},
					"email": {
						Type: "string",
						Desc: "用户的邮箱",
					},
				}),
		},
		{
			Name: "user_salary",
			Desc: "根据用户的姓名和邮箱，查询用户的薪酬信息",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{
					"name": {
						Type: "string",
						Desc: "用户的姓名",
					},
					"email": {
						Type: "string",
						Desc: "用户的邮箱",
					},
				}),
		},
	})
	if err != nil {
		log.Printf("BindForcedTools failed, err=%v", err)
		return
	}

	resp, err := chatModel.Generate(ctx, []*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一名房产经纪人，结合用户的薪酬和工作，使用 user_company、user_salary 两个 API，为其提供相关的房产信息。邮箱是必须的",
		},
		{
			Role:    schema.User,
			Content: "我的姓名是 zhangsan，我的邮箱是 zhangsan@bytedance.com，请帮我推荐一些适合我的房子。",
		},
	})

	if err != nil {
		log.Printf("Generate failed, err=%v", err)
		return
	}

	log.Printf("output: \n%+v", resp)
}
