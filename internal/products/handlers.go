package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jack/ecom/internal/json"
	"github.com/jackc/pgx/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1.Call the service -> ListProducts
	// 2.Return JSON in an HTTP ResponseWriter

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)
}

func (h *handler) FindProductById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.service.FindProductById(r.Context(), id)
	if err != nil {
		// 检查是否是"记录不存在"的错误
		if err == pgx.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		// 其他错误返回 500
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}

type CreateProductRequest struct {
	Name          string `json:"name"`
	PriceInCenter int32  `json:"price_in_center"`
	Quantity      int32  `json:"quantity"`
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求体
	var req CreateProductRequest
	if err := json.Read(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. 验证必填字段
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if req.PriceInCenter < 0 {
		http.Error(w, "price_in_center must be >= 0", http.StatusBadRequest)
		return
	}

	// 3. 调用 service
	product, err := h.service.CreateProduct(r.Context(), req.Name, req.PriceInCenter, req.Quantity)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. 返回响应
	json.Write(w, http.StatusCreated, product)
}
