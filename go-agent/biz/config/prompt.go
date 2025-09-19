package config

import (
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// 创建模板
var Template = prompt.FromMessages(schema.FString,
	schema.SystemMessage("你是一个{role}。"),
	schema.MessagesPlaceholder("history_key", false),
	&schema.Message{
		Role:    schema.User,
		Content: "{prompt}。",
	},
	// 多模态消息（包含图片）
	&schema.Message{
		Role: schema.User,
		MultiContent: []schema.ChatMessagePart{
			{
				Type: schema.ChatMessagePartTypeText,
				Text: "这张图片是什么？",
			},
			{
				Type: schema.ChatMessagePartTypeImageURL,
				ImageURL: &schema.ChatMessageImageURL{
					URL:    "https://example.com/image.jpg",
					Detail: schema.ImageURLDetailAuto,
				},
			},
		},
	},
)

//This is a method to use template
// 准备变量
// variables := map[string]any{
//     "role": "专业的助手",
//     "prompt": "写一首诗",
//     "history_key": []*schema.Message{{Role: schema.User, Content: "告诉我油画是什么?"}, {Role: schema.Assistant, Content: "油画是xxx"}},
// }

// 格式化模板
// messages, err := template.Format(context.Background(), variables)
