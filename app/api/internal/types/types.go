// Code generated by goctl. DO NOT EDIT.
package types

type CheckRequest struct {
	StudentID string `json:"student_id"`
}

type CheckResponse struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Token     string `json:"token"`
}

type StartRequest struct {
}

type StartResponse struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Account   string `json:"account"`
	Password  string `json:"password"`
}
