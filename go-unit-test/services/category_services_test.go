package services

import "testing"
import "unit-test/repository"
import "github.com/stretchr/testify/mock"
import "github.com/stretchr/testify/assert"

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryServices = CategoryServices{Repository: categoryRepository}

func TestCategoryServices_Get(t *testing.T) {
    categoryRepository.Mock.On("FindById", "1").Return(nil)
    
    category, err := categoryServices.Get("1")
    assert.Nil(t, category)
    assert.NotNil(t, err)
}