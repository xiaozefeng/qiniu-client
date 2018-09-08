package model

type Account struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket string `json:"bucket"`
	RootPath string `json:"rootPath"`
	Prefix string `json:"prefix"`
}
