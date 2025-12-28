package integration_setup

import (
	"context"
	"testing"

	order_service "alves.com/backend/modules/order/service"
	stock_service "alves.com/backend/modules/stock/service"
	user_service "alves.com/backend/modules/user/service"
)

type services struct {
	User  user_service.IService
	Stock stock_service.IService
	Order order_service.IService
}

type Harness struct {
	Ctx      context.Context
	DB       *TestDB
	Services *services
}

func NewHarness(t *testing.T) *Harness {
	t.Helper()

	testDB, err := NewTestDB(t.Name())
	if err != nil {
		t.Fatalf("failed to setup mongo: %v", err)
	}

	t.Cleanup(func() {
		_ = testDB.Teardown()
	})

	db := testDB.Database
	ctx := context.Background()

	testRedis := NewTestRedis(t)

	userService := user_service.New(db.Collection("user"), testRedis, "")
	stockService := stock_service.New(db.Collection("stock"))
	orderService := order_service.New(
		db.Collection("order"),
		userService,
		stockService,
	)

	services := services{
		User:  userService,
		Stock: stockService,
		Order: orderService,
	}

	return &Harness{
		Ctx:      ctx,
		DB:       testDB,
		Services: &services,
	}
}
