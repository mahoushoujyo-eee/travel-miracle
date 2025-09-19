package param

type ChatRequest struct {
	Prompt         string `json:"prompt"`
	ImgUrl         string `json:"img_url"`
	UserId         int64  `json:"user_id"`
	ConversationId string `json:"conversation_id"`
}
