#!/bin/bash

# Script to run 'go mod tidy' in all directories containing go.mod files

set -e

echo "Running go mod tidy in all Go module directories..."

# List of directories with go.mod files
directories=(
    "./devops"
    "./libs/acl/langfuse"
    "./libs/acl/opentelemetry"
    "./libs/acl/openai"
    "./components/document/transformer/splitter/html"
    "./components/document/transformer/splitter/recursive"
    "./components/document/transformer/splitter/semantic"
    "./components/document/transformer/splitter/markdown"
    "./components/document/transformer/reranker/score"
    "./components/document/parser/html"
    "./components/document/parser/docx"
    "./components/document/parser/xlsx"
    "./components/document/parser/pdf"
    "./components/document/loader/url"
    "./components/document/loader/file"
    "./components/document/loader/s3"
    "./components/retriever/dify"
    "./components/retriever/milvus"
    "./components/retriever/redis"
    "./components/retriever/volc_vikingdb"
    "./components/retriever/es8"
    "./components/retriever/volc_knowledge"
    "./components/indexer/milvus"
    "./components/indexer/redis"
    "./components/indexer/volc_vikingdb"
    "./components/indexer/es8"
    "./components/tool/searxng"
    "./components/tool/googlesearch"
    "./components/tool/httprequest"
    "./components/tool/bingsearch"
    "./components/tool/commandline"
    "./components/tool/duckduckgo/v2"
    "./components/tool/duckduckgo"
    "./components/tool/wikipedia"
    "./components/tool/sequentialthinking"
    "./components/tool/browseruse"
    "./components/tool/mcp"
    "./components/model/arkbot"
    "./components/model/deepseek"
    "./components/model/qwen"
    "./components/model/gemini"
    "./components/model/qianfan"
    "./components/model/ollama"
    "./components/model/openai"
    "./components/model/claude"
    "./components/model/ark"
    "./components/prompt/mcp"
    "./components/embedding/gemini"
    "./components/embedding/qianfan"
    "./components/embedding/ollama"
    "./components/embedding/dashscope"
    "./components/embedding/openai"
    "./components/embedding/cache/examples"
    "./components/embedding/cache/redis"
    "./components/embedding/cache"
    "./components/embedding/tencentcloud"
    "./components/embedding/ark"
    "./callbacks/apmplus"
    "./callbacks/langfuse"
    "./callbacks/cozeloop"
)

# Counter for tracking progress
total=${#directories[@]}
current=0

# Run go mod tidy in each directory
for dir in "${directories[@]}"; do
    current=$((current + 1))
    echo "[$current/$total] Processing $dir..."
    
    if [ -d "$dir" ]; then
        cd "$dir"
        if [ -f "go.mod" ]; then
            # Remove go.sum to resolve potential checksum mismatches
            if [ -f "go.sum" ]; then
                rm go.sum
            fi
            GOPROXY=direct go mod tidy
            echo "✓ go mod tidy completed in $dir"
        else
            echo "⚠ go.mod not found in $dir"
        fi
        cd - > /dev/null
    else
        echo "⚠ Directory $dir does not exist"
    fi
done

echo "All go mod tidy operations completed!"