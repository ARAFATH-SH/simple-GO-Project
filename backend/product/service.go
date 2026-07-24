package product

import "ecommerce/domain"

type service struct {
	prdRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Create(p)
}
func (svc *service) Get(productID int) (*domain.Product, error) {
	return svc.prdRepo.Get(productID)
}
func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdRepo.List()
}
func (svc *service) Delete(productID int) error {
	return svc.prdRepo.Delete(productID)
}
func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Update(product)
}
