package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/model"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/remote"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
)

type MercadoLivreRepositoryImpl struct {
	ds remote.MercadoLivreDataSource
}

func NewMercadoLivreRepository(ds remote.MercadoLivreDataSource) ports.MercadoLivreRepository {
	return &MercadoLivreRepositoryImpl{
		ds: ds,
	}
}

func (repo *MercadoLivreRepositoryImpl) Generate(ctx context.Context, token string, form domain.Order, orderID int) (domain.QRCodeDataResponse, error) {
	items := make([]model.Item, 0)

	totalAmount := 0

	for _, value := range form.OrderProduct {
		productId := strconv.Itoa(int(value.ProductID))
		totalAmount += 1

		items = append(items, model.Item{
			Description: fmt.Sprintf("Product: %v", productId),
			SkuNumber:   productId,
			Title:       productId,
			UnitMeasure: "unit",
			Quantity:    1,
			UnitPrice:   value.ProductPrice,
			TotalAmount: 1,
		})
	}

	// expirationDate := time.Now().Local().Add(time.Hour * 86)

	input := model.QRCodeInput{
		Description:       fmt.Sprintf("Order: %v", orderID),
		TotalAmount:       totalAmount,
		ExpirationDate:    "2024-06-28T22:34:56.559-04:00",
		ExternalReference: strconv.Itoa(orderID),
		Items:             items,
		Title:             strconv.Itoa(form.TicketNumber),
	}

	qrData, err := repo.ds.Generate(ctx, token, input)

	if err != nil {
		return domain.QRCodeDataResponse{}, err
	}

	return domain.QRCodeDataResponse{
		Data: qrData,
	}, nil
}

func (repo *MercadoLivreRepositoryImpl) GetMercadoLivrePaymentData(ctx context.Context, token string, endpoint string) error {
	err := repo.ds.GetPaymentData(ctx, token, endpoint)

	if err != nil {
		return err
	}

	return nil
}