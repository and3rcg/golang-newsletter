package models

type EmailContent struct {
	Subject     string   `json:"subject" form:"subject" validate:"required"`
	ContentText string   `json:"text" form:"text" validate:"required"`
	ContentHTML string   `json:"html" form:"html"`
	Tags        []string `json:"tags" form:"tags"`
}
