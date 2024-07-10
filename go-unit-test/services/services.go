package services 

import "unit-test/entity"
import "unit-test/repository"
import "errors"

type CategoryServices struct {
    Repository repository.CategoryRepository
}

func (service CategoryServices) Get(id string) (*entity.Category, error) {
    category := service.Repository.FindById(id)
    if category == nil {
        return nil, errors.New("Category not found")
    } else {
        return category, nil
    }
}