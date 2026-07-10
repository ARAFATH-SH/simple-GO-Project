package database

var productList []Product

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.Id = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if productID == product.Id {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.Id == product.Id {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product = make([]Product, 0)

	for _, p := range productList {
		if p.Id != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	productList = []Product{
		{Id: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		{Id: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		{Id: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
	}
}
