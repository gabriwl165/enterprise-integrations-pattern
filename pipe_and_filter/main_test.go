package pipe_and_filter

import (
	"context"
	"testing"
)

func TestPipeline_Use(t *testing.T) {
	pipeline := Pipeline{}

	if len(pipeline.Filters) != 0 {
		t.Error("expected empty pipeline, got filters")
	}

	pipeline.Use(RatingFilter)
	if len(pipeline.Filters) != 1 {
		t.Errorf("expected 1 filter, got %d", len(pipeline.Filters))
	}

	pipeline.Use(PriceFilter)
	if len(pipeline.Filters) != 2 {
		t.Errorf("expected 2 filters, got %d", len(pipeline.Filters))
	}
}

func TestPipeline_Execute(t *testing.T) {
	testProducts := []Product{
		{Name: "Test Product 1", Price: 50, Rating: 5, Availability: true},
		{Name: "Test Product 2", Price: 150, Rating: 4, Availability: true},
		{Name: "Test Product 3", Price: 75, Rating: 3, Availability: false},
		{Name: "Test Product 4", Price: 25, Rating: 5, Availability: true},
	}

	pipeline := Pipeline{}
	pipeline.Use(RatingFilter)
	pipeline.Use(PriceFilter)
	pipeline.Use(AvailabilityFilter)

	ctx := context.Background()
	ctx = context.WithValue(ctx, RatingFilterKey, 4)
	ctx = context.WithValue(ctx, PriceFilterKey, 100.0)

	result, err := pipeline.Execute(ctx, testProducts)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Should only have products with rating >= 4, price <= 100, and availability = true
	expected := 2 // Test Product 1 and Test Product 4
	if len(result) != expected {
		t.Errorf("expected %d products, got %d", expected, len(result))
	}

	for _, product := range result {
		if product.Rating < 4 {
			t.Errorf("product %s has rating %d, expected >= 4", product.Name, product.Rating)
		}
		if product.Price > 100.0 {
			t.Errorf("product %s has price %.2f, expected <= 100.0", product.Name, product.Price)
		}
		if !product.Availability {
			t.Errorf("product %s is not available, expected available", product.Name)
		}
	}
}

func TestRatingFilter(t *testing.T) {
	testProducts := []Product{
		{Name: "High Rating", Rating: 5},
		{Name: "Medium Rating", Rating: 3},
		{Name: "Low Rating", Rating: 1},
	}

	ctx := context.WithValue(context.Background(), RatingFilterKey, 3)

	result, err := RatingFilter(ctx, testProducts)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 products, got %d", len(result))
	}

	for _, product := range result {
		if product.Rating < 3 {
			t.Errorf("product %s has rating %d, expected >= 3", product.Name, product.Rating)
		}
	}
}

func TestPriceFilter(t *testing.T) {
	testProducts := []Product{
		{Name: "Cheap", Price: 50},
		{Name: "Medium", Price: 100},
		{Name: "Expensive", Price: 200},
	}

	ctx := context.WithValue(context.Background(), PriceFilterKey, 100.0)

	result, err := PriceFilter(ctx, testProducts)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 products, got %d", len(result))
	}

	for _, product := range result {
		if product.Price > 100.0 {
			t.Errorf("product %s has price %.2f, expected <= 100.0", product.Name, product.Price)
		}
	}
}

func TestAvailabilityFilter(t *testing.T) {
	testProducts := []Product{
		{Name: "Available", Availability: true},
		{Name: "Not Available", Availability: false},
		{Name: "Also Available", Availability: true},
	}

	result, err := AvailabilityFilter(context.Background(), testProducts)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 products, got %d", len(result))
	}

	for _, product := range result {
		if !product.Availability {
			t.Errorf("product %s is not available, expected available", product.Name)
		}
	}
}
