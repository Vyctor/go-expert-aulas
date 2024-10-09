package product

type ProductUsecase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{repository}
}

func (uc *ProductUsecase) GetProduct(id int) (*Product, error) {
	return uc.repository.GetProduct(id)
}
