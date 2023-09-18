package database

import (
	"database/sql"

	"github.com/dihr/go-expert-final-challenge/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAll() ([]entity.Order, error) {
	query := "SELECT * FROM orders"
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]entity.Order, 0)
	for rows.Next() {
		currentRow := entity.Order{}
		if err := rows.Scan(
			&currentRow.ID,
			&currentRow.Price,
			&currentRow.Tax,
			&currentRow.FinalPrice,
		); err != nil {
			return nil, err
		}
		orders = append(orders, currentRow)
	}
	return orders, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
