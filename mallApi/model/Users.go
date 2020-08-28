package model

//User 结构体
type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"-"` // "-"不会导出到JSON中
}
