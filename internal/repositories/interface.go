package repositories

import "github.com/Fabimhso/go-categories-msvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
}