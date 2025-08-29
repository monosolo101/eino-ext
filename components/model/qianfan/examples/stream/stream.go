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

	"github.com/cloudwego/eino/schema"
	"github.com/monosolo101/eino-ext/components/model/qianfan"
)

func main() {
	ctx := context.Background()
	qcfg := qianfan.GetQianfanSingletonConfig()
	// How to get Access Key/Secret Key: https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	qcfg.AccessKey = "your_access_key"
	qcfg.SecretKey = "your_secret_key"

	cm, err := qianfan.NewChatModel(ctx, &qianfan.ChatModelConfig{
		Model:               "ernie-3.5-8k",
		Temperature:         of(float32(0.7)),
		TopP:                of(float32(0.7)),
		MaxCompletionTokens: of(1024),
	})
	if err != nil {
		log.Fatalf("NewChatModel of qianfan failed, err=%v", err)
	}

	sr, err := cm.Stream(ctx, []*schema.Message{
		schema.UserMessage("你好"),
	})
	if err != nil {
		log.Fatalf("Stream of qianfan failed, err=%v", err)
	}

	var ms []*schema.Message
	for {
		m, err := sr.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Stream of qianfan failed, err=%v", err)
		}

		fmt.Println(m)
		ms = append(ms, m)
	}

	// assistant: 您好！
	// assistant: 我是文心
	// assistant: 一言，
	// assistant: 很高兴与您
	// assistant: 交流。
	// assistant: 请问有什么
	// assistant: 我可以帮助
	// assistant: 您的吗
	// assistant: ？
	// assistant: 无论是知识
	// assistant: 问答、
	// assistant: 文本创作
	// assistant: 还是其他
	// assistant: 任何问题，
	// assistant: 我都会尽力
	// assistant: 为您提供帮助
	// assistant: 。

	sm, err := schema.ConcatMessages(ms)
	if err != nil {
		log.Fatalf("ConcatMessages failed, err=%v", err)
	}

	fmt.Println(sm)
	// assistant: 您好！我是文心一言，很高兴与您交流。请问有什么我可以帮助您的吗？无论是知识问答、文本创作还是其他任何问题，我都会尽力为您提供帮助。
	// usage: &{1 32 33}
}

func of[T any](t T) *T {
	return &t
}
