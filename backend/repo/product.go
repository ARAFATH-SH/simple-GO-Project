package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"image_url" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	//productList []*Product //in memory
	db *sqlx.DB
}

// constructor
func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	// generateInitialProduct(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		insert into products(
			title,
			description,
			price,
			image_url
		) values(
		  	$1,
		  	$2,
		  	$3,
			$4
		)
		RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL)
	err := row.Scan(&p.Id)

	if err != nil {
		return nil, err
	}

	return &p, nil

}
func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product

	query := `
		SELECT 
		id,
    	title,
    	description,
    	price,
    	image_url
		FROM products
		WHERE id = $1
	`

	err := r.db.Get(&prd, query, id)

	if err != nil {

		if err == sql.ErrNoRows {

			return nil, nil
		}
		return nil, err
	}

	return &prd, nil

}
func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product

	querry := `
		SELECT 
		id,
    	title,
    	description,
    	price,
    	image_url 
		FROM products
	`

	err := r.db.Select(&prdList, querry)

	if err != nil {
		return nil, err
	}

	return prdList, nil
}
func (r *productRepo) Delete(productID int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	_, err := r.db.Exec(query, productID)

	if err != nil {
		return err
	}
	return nil
}
func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
		UPDATE products
		SET 
			title=$1,
			description=$2,
			price=$3,
			image_url=$4
		WHERE id=$5
		RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL, p.Id)
	err := row.Err()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &p, nil

}

// func generateInitialProduct(r *productRepo) {
// 	prd1 := &Product{
// 		Id:          1,
// 		Title:       "Wireless Headphones",
// 		Description: "High-quality noise-cancelling headphones.",
// 		Price:       129.99,
// 		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
// 	}
// 	prd2 := &Product{
// 		Id:          2,
// 		Title:       "Smart Watch",
// 		Description: "Stylish smart watch with health tracking.",
// 		Price:       199.99,
// 		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
// 	}
// 	prd3 := &Product{
// 		Id:          3,
// 		Title:       "Running Shoes",
// 		Description: "Lightweight shoes for everyday running.",
// 		Price:       89.50,
// 		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
// 	}

// 	r.productList = append(r.productList, prd1)
// 	r.productList = append(r.productList, prd2)
// 	r.productList = append(r.productList, prd3)
// }
