package service

import (
	"context"

	"github.com/dihr/go-expert-final-challenge/internal/infra/grpc/pb"
	"github.com/dihr/go-expert-final-challenge/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) ListOrders(context.Context, *pb.Blank) (*pb.OrderList, error) {
	orderListOutput, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	response := make([]*pb.Order, 0)
	for _, item := range orderListOutput {
		response = append(response, &pb.Order{
			Id:         item.ID,
			Price:      float32(item.Price),
			Tax:        float32(item.Tax),
			FinalPrice: float32(item.FinalPrice),
		})
	}
	return &pb.OrderList{
		Orders: response,
	}, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
