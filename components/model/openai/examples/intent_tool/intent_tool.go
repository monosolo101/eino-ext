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
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cloudwego/eino/schema"

	"github.com/monosolo101/eino-ext/components/model/openai"
)

func main() {
	ctx := context.Background()
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  os.Getenv("OPENAI_API_KEY"),
		Model:   os.Getenv("OPENAI_MODEL"),
		BaseURL: os.Getenv("OPENAI_BASE_URL"),
		ByAzure: func() bool {
			if os.Getenv("OPENAI_BY_AZURE") == "true" {
				return true
			}
			return false
		}(),
	})
	if err != nil {
		log.Fatalf("NewChatModel of openai failed, err=%v", err)
	}
	err = chatModel.BindForcedTools([]*schema.ToolInfo{
		{
			Name: "user_company",
			Desc: "Retrieve the user's company and position based on their name and email.",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{
					"name":  {Type: "string", Desc: "user's name"},
					"email": {Type: "string", Desc: "user's email"}}),
		}, {
			Name: "user_salary",
			Desc: "Retrieve the user's salary based on their name and email.\n",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{
					"name":  {Type: "string", Desc: "user's name"},
					"email": {Type: "string", Desc: "user's email"},
				}),
		}})
	if err != nil {
		log.Fatalf("BindForcedTools of openai failed, err=%v", err)
	}
	resp, err := chatModel.Generate(ctx, []*schema.Message{{
		Role:    schema.System,
		Content: "As a real estate agent, provide relevant property information based on the user's salary and job using the user_company and user_salary APIs. An email address is required.",
	}, {
		Role:    schema.User,
		Content: "My name is John and my email is john@abc.com，Please recommend some houses that suit me.",
	}})
	if err != nil {
		log.Fatalf("Generate of openai failed, err=%v", err)
	}
	fmt.Printf("output: \n%v", resp)

	streamResp, err := chatModel.Stream(ctx, []*schema.Message{
		{
			Role:    schema.System,
			Content: "As a real estate agent, provide relevant property information based on the user's salary and job using the user_company and user_salary APIs. An email address is required.",
		}, {
			Role:    schema.User,
			Content: "My name is John and my email is john@abc.com，Please recommend some houses that suit me.",
		},
	})
	if err != nil {
		log.Fatalf("Stream of openai failed, err=%v", err)
	}
	var messages []*schema.Message
	for {
		chunk, err := streamResp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Recv of streamResp failed, err=%v", err)
		}
		messages = append(messages, chunk)
	}
	resp, err = schema.ConcatMessages(messages)
	if err != nil {
		log.Fatalf("ConcatMessages of openai failed, err=%v", err)
	}
	fmt.Printf("stream output: \n%v", resp)
}
