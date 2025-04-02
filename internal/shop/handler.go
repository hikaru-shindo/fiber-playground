package shop

type Handler struct {
	productStore ProductStore
}

func NewHandler(productStore ProductStore) Handler {
	return Handler{
		productStore: productStore,
	}
}
