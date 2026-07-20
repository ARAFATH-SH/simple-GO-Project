package repo

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProduct(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.Id = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.Id == productId {
			return product, nil
		}
	}

	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.Id != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.Id == product.Id {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func generateInitialProduct(r *productRepo) {
	prd1 := &Product{
		Id:          1,
		Title:       "Wireless Headphones",
		Description: "High-quality noise-cancelling headphones.",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	prd2 := &Product{
		Id:          2,
		Title:       "Smart Watch",
		Description: "Stylish smart watch with health tracking.",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	prd3 := &Product{
		Id:          3,
		Title:       "Running Shoes",
		Description: "Lightweight shoes for everyday running.",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
}
