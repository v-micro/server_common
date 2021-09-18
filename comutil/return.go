package comutil

// 定义返回变量
type Response struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Details    interface{} `json:"details"`
	TotalCount int         `json:"total_count"`
}
