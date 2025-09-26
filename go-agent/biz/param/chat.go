package param

type ChatRequest struct {
	Prompt         string `json:"prompt"`
	ImgUrl         string `json:"img_url"`
	UserId         int64  `json:"user_id"`
	ConversationId string `json:"conversation_id"`
}

type UploadFileRequest struct {
	Type string `json:"type"`
	FileName string `json:"file_name"`
	Size int64 `json:"size"`
	ContentType string `json:"content_type"`
}
