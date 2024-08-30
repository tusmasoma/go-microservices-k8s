package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/tusmasoma/go-microservice-k8s/microservice-k8s-demo/order/entity"
	repo_mock "github.com/tusmasoma/go-microservice-k8s/microservice-k8s-demo/order/repository/mock"
	service_mock "github.com/tusmasoma/go-microservice-k8s/microservice-k8s-demo/order/service/mock"
)

func TestOrderUseCase_GetOrderCreationResources(t *testing.T) {
	t.Parallel()

	customers := []entity.Customer{
		{
			ID:   uuid.New().String(),
			Name: "customer1",
		},
	}

	items := []entity.CatalogItem{
		{
			ID:    uuid.New().String(),
			Name:  "item1",
			Price: 1000,
		},
	}

	patterns := []struct {
		name  string
		setup func(
			m *repo_mock.MockCustomerRepository,
			m1 *repo_mock.MockCatalogItemRepository,
			m2 *repo_mock.MockOrderRepository,
			m3 *repo_mock.MockOrderLineRepository,
			m4 *service_mock.MockOrderService,
		)
		arg struct {
			ctx context.Context
		}
		want struct {
			customers []entity.Customer
			items     []entity.CatalogItem
			err       error
		}
	}{
		{
			name: "success",
			setup: func(
				cr *repo_mock.MockCustomerRepository,
				cir *repo_mock.MockCatalogItemRepository,
				or *repo_mock.MockOrderRepository,
				olr *repo_mock.MockOrderLineRepository,
				os *service_mock.MockOrderService,
			) {
				cr.EXPECT().List(gomock.Any()).Return(customers, nil)
				cir.EXPECT().List(gomock.Any()).Return(items, nil)
			},
			arg: struct {
				ctx context.Context
			}{
				ctx: context.Background(),
			},
			want: struct {
				customers []entity.Customer
				items     []entity.CatalogItem
				err       error
			}{
				customers: customers,
				items:     items,
				err:       nil,
			},
		},
	}

	for _, tt := range patterns {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			cr := repo_mock.NewMockCustomerRepository(ctrl)
			cir := repo_mock.NewMockCatalogItemRepository(ctrl)
			or := repo_mock.NewMockOrderRepository(ctrl)
			olr := repo_mock.NewMockOrderLineRepository(ctrl)
			os := service_mock.NewMockOrderService(ctrl)

			if tt.setup != nil {
				tt.setup(cr, cir, or, olr, os)
			}

			ouc := NewOrderUseCase(cr, cir, or, olr, os)

			gotCustomers, gotItems, err := ouc.GetOrderCreationResources(tt.arg.ctx)
			if (err != nil) != (tt.want.err != nil) {
				t.Errorf("GetOrderCreationResources() error = %v, wantErr %v", err, tt.want.err)
			} else if err != nil && tt.want.err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("GetOrderCreationResources() error = %v, wantErr %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(gotCustomers, tt.want.customers) {
				t.Errorf("GetOrderCreationResources() got = %v, want %v", gotCustomers, tt.want.customers)
			}
			if !reflect.DeepEqual(gotItems, tt.want.items) {
				t.Errorf("GetOrderCreationResources() got = %v, want %v", gotItems, tt.want.items)
			}
		})
	}
}

func TestOrderUseCase_GetOrder(t *testing.T) {
	t.Parallel()

	orderID := uuid.New().String()
	customerID := uuid.New().String()
	catalogItemID := uuid.New().String()
	orderDate := time.Now()

	order := entity.Order{
		ID: orderID,
		Customer: entity.Customer{
			ID:   customerID,
			Name: "customer1",
		},
		OrderDate: orderDate,
		OrderLines: []entity.OrderLine{
			{
				Count: 1,
				CatalogItem: entity.CatalogItem{
					ID:    catalogItemID,
					Name:  "item1",
					Price: 1000,
				},
			},
		},
		TotalPrice: 1000,
	}

	patterns := []struct {
		name  string
		setup func(
			m *repo_mock.MockCustomerRepository,
			m1 *repo_mock.MockCatalogItemRepository,
			m2 *repo_mock.MockOrderRepository,
			m3 *repo_mock.MockOrderLineRepository,
			m4 *service_mock.MockOrderService,
		)
		arg struct {
			ctx context.Context
			id  string
		}
		want struct {
			order *entity.Order
			err   error
		}
	}{
		{
			name: "success",
			setup: func(
				cr *repo_mock.MockCustomerRepository,
				cir *repo_mock.MockCatalogItemRepository,
				or *repo_mock.MockOrderRepository,
				olr *repo_mock.MockOrderLineRepository,
				os *service_mock.MockOrderService,
			) {
				or.EXPECT().Get(gomock.Any(), orderID).Return(
					&entity.OrderModel{
						ID:         orderID,
						CustomerID: customerID,
						OrderDate:  orderDate,
					},
					nil,
				)
				olr.EXPECT().List(gomock.Any(), orderID).Return(
					[]entity.OrderLineModel{
						{
							OrderID:       orderID,
							CatalogItemID: catalogItemID,
							Count:         1,
						},
					}, nil)
				cr.EXPECT().Get(gomock.Any(), customerID).Return(
					&entity.Customer{
						ID:   customerID,
						Name: "customer1",
					}, nil)
				cir.EXPECT().Get(gomock.Any(), catalogItemID).Return(
					&entity.CatalogItem{
						ID:    catalogItemID,
						Name:  "item1",
						Price: 1000,
					}, nil)
			},
			arg: struct {
				ctx context.Context
				id  string
			}{
				ctx: context.Background(),
				id:  orderID,
			},
			want: struct {
				order *entity.Order
				err   error
			}{
				order: &order,
				err:   nil,
			},
		},
	}

	for _, tt := range patterns {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			cr := repo_mock.NewMockCustomerRepository(ctrl)
			cir := repo_mock.NewMockCatalogItemRepository(ctrl)
			or := repo_mock.NewMockOrderRepository(ctrl)
			olr := repo_mock.NewMockOrderLineRepository(ctrl)
			os := service_mock.NewMockOrderService(ctrl)

			if tt.setup != nil {
				tt.setup(cr, cir, or, olr, os)
			}

			ouc := NewOrderUseCase(cr, cir, or, olr, os)

			gotOrder, err := ouc.GetOrder(tt.arg.ctx, tt.arg.id)
			if (err != nil) != (tt.want.err != nil) {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.want.err)
			} else if err != nil && tt.want.err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(gotOrder, tt.want.order) {
				t.Errorf("GetOrder() got = %v, want %v", gotOrder, tt.want.order)
			}
		})
	}
}

func TestOrderUseCase_ListOrder(t *testing.T) {
	t.Parallel()

	orderID := uuid.New().String()
	customerID := uuid.New().String()
	catalogItemID := uuid.New().String()
	orderDate := time.Now()

	orders := []entity.Order{
		{
			ID: orderID,
			Customer: entity.Customer{
				ID:   customerID,
				Name: "customer1",
			},
			OrderDate: orderDate,
			OrderLines: []entity.OrderLine{
				{
					Count: 1,
					CatalogItem: entity.CatalogItem{
						ID:    catalogItemID,
						Name:  "item1",
						Price: 1000,
					},
				},
			},
			TotalPrice: 1000,
		},
	}

	patterns := []struct {
		name  string
		setup func(
			m *repo_mock.MockCustomerRepository,
			m1 *repo_mock.MockCatalogItemRepository,
			m2 *repo_mock.MockOrderRepository,
			m3 *repo_mock.MockOrderLineRepository,
			m4 *service_mock.MockOrderService,
		)
		arg struct {
			ctx context.Context
		}
		want struct {
			orders []entity.Order
			err    error
		}
	}{
		{
			name: "success",
			setup: func(
				cr *repo_mock.MockCustomerRepository,
				cir *repo_mock.MockCatalogItemRepository,
				or *repo_mock.MockOrderRepository,
				olr *repo_mock.MockOrderLineRepository,
				os *service_mock.MockOrderService,
			) {
				or.EXPECT().List(gomock.Any()).Return(
					[]entity.OrderModel{
						{
							ID:         orderID,
							CustomerID: customerID,
							OrderDate:  orderDate,
						},
					},
					nil,
				)
				olr.EXPECT().List(gomock.Any(), orderID).Return(
					[]entity.OrderLineModel{
						{
							OrderID:       orderID,
							CatalogItemID: catalogItemID,
							Count:         1,
						},
					}, nil)
				cr.EXPECT().Get(gomock.Any(), customerID).Return(
					&entity.Customer{
						ID:   customerID,
						Name: "customer1",
					}, nil)
				cir.EXPECT().Get(gomock.Any(), catalogItemID).Return(
					&entity.CatalogItem{
						ID:    catalogItemID,
						Name:  "item1",
						Price: 1000,
					}, nil)
			},
			arg: struct {
				ctx context.Context
			}{
				ctx: context.Background(),
			},
			want: struct {
				orders []entity.Order
				err    error
			}{
				orders: orders,
				err:    nil,
			},
		},
	}

	for _, tt := range patterns {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			cr := repo_mock.NewMockCustomerRepository(ctrl)
			cir := repo_mock.NewMockCatalogItemRepository(ctrl)
			or := repo_mock.NewMockOrderRepository(ctrl)
			olr := repo_mock.NewMockOrderLineRepository(ctrl)
			os := service_mock.NewMockOrderService(ctrl)

			if tt.setup != nil {
				tt.setup(cr, cir, or, olr, os)
			}

			ouc := NewOrderUseCase(cr, cir, or, olr, os)

			gotOrders, err := ouc.ListOrder(tt.arg.ctx)
			if (err != nil) != (tt.want.err != nil) {
				t.Errorf("ListOrder() error = %v, wantErr %v", err, tt.want.err)
			} else if err != nil && tt.want.err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("ListOrder() error = %v, wantErr %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(gotOrders, tt.want.orders) {
				t.Errorf("ListOrder() got = %v, want %v", gotOrders, tt.want.orders)
			}
		})
	}
}

func TestOrderUseCase_CreateOrder(t *testing.T) {
	t.Parallel()

	customerID := uuid.New().String()
	catalogItemID := uuid.New().String()

	patterns := []struct {
		name  string
		setup func(
			m *repo_mock.MockCustomerRepository,
			m1 *repo_mock.MockCatalogItemRepository,
			m2 *repo_mock.MockOrderRepository,
			m3 *repo_mock.MockOrderLineRepository,
			m4 *service_mock.MockOrderService,
		)
		arg struct {
			ctx    context.Context
			params *CreateOrderParams
		}
		wantErr error
	}{
		{
			name: "success",
			setup: func(
				cr *repo_mock.MockCustomerRepository,
				cir *repo_mock.MockCatalogItemRepository,
				or *repo_mock.MockOrderRepository,
				olr *repo_mock.MockOrderLineRepository,
				os *service_mock.MockOrderService,
			) {
				cr.EXPECT().Get(gomock.Any(), customerID).Return(
					&entity.Customer{
						ID:   customerID,
						Name: "customer1",
					}, nil)
				cir.EXPECT().Get(gomock.Any(), catalogItemID).Return(
					&entity.CatalogItem{
						ID:    catalogItemID,
						Name:  "item1",
						Price: 1000,
					}, nil)
				os.EXPECT().CreateOrder(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, order entity.Order) {
					if order.Customer.ID != customerID {
						t.Errorf("unexpected customerID: got %v, want %v", order.Customer.ID, customerID)
					}
					if order.OrderLines[0].CatalogItem.ID != catalogItemID {
						t.Errorf("unexpected catalogItemID: got %v, want %v", order.OrderLines[0].CatalogItem.ID, catalogItemID)
					}
					if order.TotalPrice != 1000 {
						t.Errorf("unexpected totalPrice: got %v, want %v", order.TotalPrice, 1000)
					}
				}).Return(nil)
			},
			arg: struct {
				ctx    context.Context
				params *CreateOrderParams
			}{
				ctx: context.Background(),
				params: &CreateOrderParams{
					CustomerID: customerID,
					OrderLine: []struct {
						CatalogItemID string
						Count         int
					}{
						{
							CatalogItemID: catalogItemID,
							Count:         1,
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range patterns {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			cr := repo_mock.NewMockCustomerRepository(ctrl)
			cir := repo_mock.NewMockCatalogItemRepository(ctrl)
			or := repo_mock.NewMockOrderRepository(ctrl)
			olr := repo_mock.NewMockOrderLineRepository(ctrl)
			os := service_mock.NewMockOrderService(ctrl)

			if tt.setup != nil {
				tt.setup(cr, cir, or, olr, os)
			}

			ouc := NewOrderUseCase(cr, cir, or, olr, os)

			err := ouc.CreateOrder(tt.arg.ctx, tt.arg.params)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
