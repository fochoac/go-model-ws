package prueba

type TagExtra struct {
	Extra string `json:"extra" form:"name" binding:"required"`
}
