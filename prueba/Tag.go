package prueba

type Tag struct {
	Name        string   `json:"name" form:"name" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	ExtraData   TagExtra `json:"extra"}`
}
