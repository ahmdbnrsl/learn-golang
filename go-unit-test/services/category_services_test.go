package services

import "testing"
import "unit-test/repository"
import "unit-test/entity"
import "github.com/stretchr/testify/mock"
import "github.com/stretchr/testify/assert"

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryServices = CategoryServices{Repository: categoryRepository}

func TestCategoryServices_GetSuccess(t *testing.T) {
    category := entity.Category{
        Id : "2",
        Name : "Shoes",
    }
    categoryRepository.Mock.On("FindById", "2").Return(category)
    
    result, err := categoryServices.Get("2")
    assert.Nil(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, category.Id, result.Id)
    assert.Equal(t, category.Name, result.Name)
}

func TestCategoryServices_GetNotFound(t *testing.T) {
    categoryRepository.Mock.On("FindById", "1").Return(nil)
    
    result, err := categoryServices.Get("1")
    assert.Nil(t, result)
    assert.NotNil(t, err)
    
}