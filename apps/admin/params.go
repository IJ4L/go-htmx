package admin

type AddProduct struct {
	Name     string  `form:"productName" binding:"required"`
	Category string  `form:"productCategory" binding:"required"`
	Price    float64 `form:"productPrice" binding:"required"`
}
