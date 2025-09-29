package pipe_and_filter

import (
	"context"
	"fmt"
)

const (
	RatingFilterKey = "filter.rating"
	PriceFilterKey  = "filter.price"
)

type Product struct {
	Name         string
	Price        float64
	Rating       int
	Discount     float64
	NetPrice     float64
	Availability bool
}

type Filter func(context.Context, []Product) ([]Product, error)

type Pipeline struct {
	Filters []Filter
}

func (p *Pipeline) Use(filter Filter) {
	p.Filters = append(p.Filters, filter)
}

func (p *Pipeline) Execute(ctx context.Context, input []Product) ([]Product, error) {
	var (
		output = input
		err    error
	)

	for _, filter := range p.Filters {
		output, err = filter(ctx, output)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

var RatingFilter Filter = func(ctx context.Context, products []Product) ([]Product, error) {
	rating := ctx.Value(RatingFilterKey).(int)

	var output []Product
	for _, product := range products {
		if product.Rating >= rating {
			output = append(output, product)
		}
	}

	return output, nil
}

var PriceFilter Filter = func(ctx context.Context, products []Product) ([]Product, error) {
	price := ctx.Value(PriceFilterKey).(float64)

	var output []Product
	for _, product := range products {
		if product.Price <= price {
			output = append(output, product)
		}
	}

	return output, nil

}

var AvailabilityFilter Filter = func(ctx context.Context, products []Product) ([]Product, error) {
	var output []Product
	for _, product := range products {
		if product.Availability {
			output = append(output, product)
		}
	}
	return output, nil
}

var productList []Product

func init() {
	productList = []Product{
		{Name: "Product 1", Price: 100, Rating: 4, Discount: 0.1, NetPrice: 90, Availability: true},
		{Name: "Product 2", Price: 200, Rating: 5, Discount: 0.2, NetPrice: 160, Availability: true},
		{Name: "Product 3", Price: 300, Rating: 3, Discount: 0.3, NetPrice: 210, Availability: false},
		{Name: "Product 4", Price: 400, Rating: 2, Discount: 0.4, NetPrice: 240, Availability: true},
		{Name: "Product 5", Price: 500, Rating: 1, Discount: 0.5, NetPrice: 250, Availability: false},
	}
}

func main() {
	pipeline := Pipeline{}
	pipeline.Use(RatingFilter)
	pipeline.Use(PriceFilter)
	pipeline.Use(AvailabilityFilter)

	ctx := context.Background()
	ctx = context.WithValue(ctx, RatingFilterKey, 4)
	ctx = context.WithValue(ctx, PriceFilterKey, 100.5)

	products, err := pipeline.Execute(ctx, productList)
	if err != nil {
		fmt.Printf("Error executing pipeline: %v\n", err)
	}

	fmt.Printf("Products: %v\n", products)
}
