package param

type JudgeRequest struct {
	Score int `json:"score"`
	Judgement string `json:"judgement"`
	ConversationID string `json:"conversation_id"`
}