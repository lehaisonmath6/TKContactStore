package models

type TKContactResult struct {
	ErrorCode int       `json:"errorCode"`
	Data      TKContact `json:"data`
}

type TKSynContactResult struct {
	ErroCode int             `json:"errorCode"`
	Data     []TKContactItem `json:"data"`
}

type ErrorResult struct {
	Result int `json:"errorCode"`
}
