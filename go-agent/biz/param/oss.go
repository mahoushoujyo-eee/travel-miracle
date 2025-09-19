package param

type GetUploadUrlRequest struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	ContentType string `json:"content_type"`
}

type GetDownloadUrlRequest struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	ContentType string `json:"content_type"`
}
