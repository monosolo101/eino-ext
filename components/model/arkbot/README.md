# Volcengine Ark Bot

A Volcengine Ark Bot implementation for [Eino](https://github.com/cloudwego/eino) that implements the `ToolCallingChatModel` interface. This enables seamless integration with Eino's LLM capabilities for enhanced natural language processing and generation.

## Features

- Implements `github.com/cloudwego/eino/components/model.ToolCallingChatModel`
- Easy integration with Eino's model system
- Configurable model parameters
- Support for chat completion
- Support for streaming responses
- Custom response parsing support
- Flexible model configuration

## Installation

```bash
go get github.com/monosolo101/eino-ext/components/model/arkbot@latest
```

## Quick Start

Here's a quick example of how to use the Ark Bot:

```go
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
	"encoding/json"
	"log"

	"github.com/cloudwego/eino/schema"

	"github.com/monosolo101/eino-ext/components/model/arkbot"
)

func main() {
	ctx := context.Background()

	// Get ARK_API_KEY and ARK_MODEL_ID: https://www.volcengine.com/docs/82379/1399008
	chatModel, err := arkbot.NewChatModel(ctx, &arkbot.ChatModelConfig{
		APIKey: "xxx",
		Model:  "yyy",
	})

	if err != nil {
		log.Fatalf("NewChatModel failed, err=%v", err)
	}

	inMsgs := []*schema.Message{
		{
			Role:    schema.User,
			Content: "What's the weather in Beijing?",
		},
	}

	msg, err := chatModel.Generate(ctx, inMsgs)
	if err != nil {
		log.Fatalf("Generate failed, err=%v", err)
	}

	log.Printf("generate output: \n")
	log.Printf("  request_id: %s\n", arkbot.GetArkRequestID(msg))
	if bu, ok := arkbot.GetBotUsage(msg); ok {
		bbu, _ := json.Marshal(bu)
		log.Printf("  bot_usage: %s\n", string(bbu))
	}
	if ref, ok := arkbot.GetBotChatResultReference(msg); ok {
		bRef, _ := json.Marshal(ref)
		log.Printf("  bot_chat_result_reference: %s\n", bRef)
	}
	respBody, _ := json.MarshalIndent(msg, "  ", "  ")
	log.Printf("  body: %s\n", string(respBody))
}

```

## Configuration

The model can be configured using the `arkbot.Config` struct:

```go
type Config struct {
    // Timeout specifies the maximum duration to wait for API responses
    // Optional. Default: 10 minutes
    Timeout *time.Duration `json:"timeout"`

    // RetryTimes specifies the number of retry attempts for failed API calls
    // Optional. Default: 2
    RetryTimes *int `json:"retry_times"`

    // BaseURL specifies the base URL for Ark service
    // Optional. Default: "https://ark.cn-beijing.volces.com/api/v3"
    BaseURL string `json:"base_url"`
    // Region specifies the region where Ark service is located
    // Optional. Default: "cn-beijing"
    Region string `json:"region"`

    // The following three fields are about authentication - either APIKey or AccessKey/SecretKey pair is required
    // For authentication details, see: https://www.volcengine.com/docs/82379/1298459
    // APIKey takes precedence if both are provided
    APIKey    string `json:"api_key"`
    AccessKey string `json:"access_key"`
    SecretKey string `json:"secret_key"`

    // The following fields correspond to Ark's chat completion API parameters
    // Ref: https://www.volcengine.com/docs/82379/1298454

    // Model specifies the ID of endpoint on ark platform
    // Required
    Model string `json:"model"`

    // MaxTokens limits the maximum number of tokens that can be generated in the chat completion and the range of values is [0, 4096]
    // Optional. Default: 4096
    MaxTokens *int `json:"max_tokens,omitempty"`

    // Temperature specifies what sampling temperature to use
    // Generally recommend altering this or TopP but not both
    // Range: 0.0 to 1.0. Higher values make output more random
    // Optional. Default: 1.0
    Temperature *float32 `json:"temperature,omitempty"`

    // TopP controls diversity via nucleus sampling
    // Generally recommend altering this or Temperature but not both
    // Range: 0.0 to 1.0. Lower values make output more focused
    // Optional. Default: 0.7
    TopP *float32 `json:"top_p,omitempty"`

    // Stop sequences where the API will stop generating further tokens
    // Optional. Example: []string{"\n", "User:"}
    Stop []string `json:"stop,omitempty"`

    // FrequencyPenalty prevents repetition by penalizing tokens based on frequency
    // Range: -2.0 to 2.0. Positive values decrease likelihood of repetition
    // Optional. Default: 0
    FrequencyPenalty *float32 `json:"frequency_penalty,omitempty"`

    // LogitBias modifies likelihood of specific tokens appearing in completion
    // Optional. Map token IDs to bias values from -100 to 100
    LogitBias map[string]int `json:"logit_bias,omitempty"`

    // PresencePenalty prevents repetition by penalizing tokens based on presence
    // Range: -2.0 to 2.0. Positive values increase likelihood of new topics
    // Optional. Default: 0
    PresencePenalty *float32 `json:"presence_penalty,omitempty"`

    // CustomHeader the http header passed to model when requesting model
    CustomHeader map[string]string `json:"custom_header"`
}
```

## Request Options

The Ark model supports various request options to customize the behavior of API calls. Here are the available options:

```go
// WithCustomHeader sets custom headers for a single request
// the headers will override all the headers given in ChatModelConfig.CustomHeader
func WithCustomHeader(m map[string]string) model.Option {}
```

## For More Details

- [Eino Documentation](https://github.com/cloudwego/eino)
- [Volcengine Ark Model Documentation](https://www.volcengine.com/docs/82379/1263272)
