package usecase

import (
	"github.com/Azmi117/API-TV.git/internal/models"
	"github.com/Azmi117/API-TV.git/internal/pkg/apperror"
	"github.com/Azmi117/API-TV.git/internal/repository"
)

type TvUsecase struct {
	repo *repository.TvRepository
}

func NewTvUsecase(params *repository.TvRepository) *TvUsecase {
	return &TvUsecase{
		repo: params,
	}
}

func (u *TvUsecase) GetAll() ([]models.Tv, error) {
	data, err := u.repo.FindAll()

	if err != nil {
		return nil, apperror.Internal("Failed Get Data!")
	}

	return data, nil
}

func (u *TvUsecase) GetById(params int) (models.Tv, error) {
	data, err := u.repo.FindById(params)

	if err != nil {
		return models.Tv{}, apperror.NotFound("No data exist with this id!")
	}

	return data, nil
}

func (u *TvUsecase) Create(params models.Tv) (models.Tv, error) {
	exist, _ := u.repo.FindByName(params.Brand)

	if params.Brand == exist.Brand {
		return models.Tv{}, apperror.BadRequest("Brand is exist!")
	}

	if params.Brand == "" || params.Price == 0 || params.Quantity == 0 {
		return models.Tv{}, apperror.BadRequest("field can't be empty or zero!")
	}

	return params, nil
}

func (u *TvUsecase) Update(id int, params models.Tv) (models.Tv, error) {
	exist, err := u.repo.FindById(id)

	if err != nil {
		return models.Tv{}, apperror.NotFound("No data exist with this id!")
	}

	if params.Brand == "" {
		params.Brand = exist.Brand
	}

	if err := u.repo.Update(params); err != nil {
		return models.Tv{}, apperror.Internal("Failed update data!")
	}

	return params, nil
}

func (u *TvUsecase) Delete(id int) error {
	exist, err := u.repo.FindById(id)

	if err != nil {
		return apperror.NotFound("No data exist with this id!")
	}

	if exist.DeletedAt.Valid {
		return apperror.BadRequest("Data with this id already deleted")
	}

	return nil
}
