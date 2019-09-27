package domain

type FileData struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Etag string `json:"etag"`
}

type FileInformation struct {
	Key             string   `json:"key"`
	Filename        string   `json:"filename"`
	ContentType     string   `json:"contentType"`
	File            FileData `json:"file"`
	FileSize        int32    `json:"fileSize"`
	DispositionType string   `json:"dispositionType"`
}

type FileMetaData struct {
	ETag     string `json:"eTag"`
	Fid      string `json:"fid"`
	FileName string `json:"fileName"`
	FileUrl  string `json:"fileUrl"`
	Size     int    `json:"size"`
}

type FileSize struct {
	Size int `json:"size"`
}
