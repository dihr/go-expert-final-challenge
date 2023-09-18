package usecase

import "github.com/dihr/go-expert-final-challenge/internal/entity"

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}
type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := l.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	ordersDTOList := make([]ListOrderOutputDTO, 0)
	for _, order := range orders {
		ordersDTOList = append(ordersDTOList, ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersDTOList, nil
}
