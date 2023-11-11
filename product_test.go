package woocommerce

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_All(t *testing.T) {
	params := ProductsQueryParams{}
	items, _, _, _, err := wooClient.Services.Product.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.Product.All: %s", err.Error())
	} else {
		if len(items) > 0 {
			mainId = items[0].ID
		}
	}
}

func TestProductService_One(t *testing.T) {
	t.Run("TestProductService_All", TestProductService_All)
	product, err := wooClient.Services.Product.One(mainId)
	if err != nil {
		t.Errorf("wooClient.Services.Product.One: %s", err.Error())
	} else {
		assert.Equal(t, mainId, product.ID, "product id")
	}
}

func TestProductService_CreateUpdateDelete(t *testing.T) {
	name := gofakeit.Word()
	req := CreateProductRequest{
		Name: name,
	}
	item, err := wooClient.Services.Product.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.Product.Create error: %s", err.Error())
	}
	productId := item.ID
	assert.Equal(t, name, item.Name, "product name")
	name = gofakeit.Word()
	updateReq := UpdateProductRequest{
		Name: name,
	}
	_, err = wooClient.Services.Product.Update(productId, updateReq)
	if err != nil {
		t.Fatalf("wooClient.Services.Product.Update error: %s", err.Error())
	}
	item, err = wooClient.Services.Product.One(productId)
	if err != nil {
		t.Fatalf("wooClient.Services.Product.One error: %s", err.Error())
	}
	assert.Equal(t, name, item.Name, "product name")

	// Delete
	_, err = wooClient.Services.Product.Delete(productId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.Product.Delete error: %s", err.Error())
	}
	_, err = wooClient.Services.Product.One(productId)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("wooClient.Services.Product.Delete(%d) failed", productId)
	}

	// Batch
	_, err = wooClient.Services.Product.Batch(productId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.Product.Delete error: %s", err.Error())
	}
	_, err = wooClient.Services.Product.One(productId)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("wooClient.Services.Product.Delete(%d) failed", productId)
	}
}


func TestProductService_Batch(t *testing.T) {
	n := 3
	createRequests := make([]BatchProductCreateItem, n)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		req := BatchProductCreateItem{
			Name:        gofakeit.Word(),
			Description: gofakeit.Address().Address,
		}
		createRequests[i] = req
		names[i] = req.Name
	}
	batchReq := BatchProductRequest{
		Create: createRequests,
	}
	result, err := wooClient.Services.Product.Batch(batchReq)
	if err != nil {
		t.Fatalf("wooClient.Services.Product.Batch() error: %s", err.Error())
	}
	assert.Equal(t, n, len(result.Create), "Batch create return len")
	returnNames := make([]string, 0)
	for _, d := range result.Create {
		returnNames = append(returnNames, d.Name)
	}
	assert.Equal(t, names, returnNames, "check names is equal")
}