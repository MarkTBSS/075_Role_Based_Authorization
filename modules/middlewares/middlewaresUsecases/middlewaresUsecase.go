package middlewaresUsecases

import (
	_pkgMiddlewares "github.com/MarkTBSS/075_Role_Based_Authorization/modules/middlewares"
	_pkgMiddlewaresMiddlewaresRepositories "github.com/MarkTBSS/075_Role_Based_Authorization/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*_pkgMiddlewares.Role, error)
}

type middlewaresUsecase struct {
	middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}

func (u *middlewaresUsecase) FindRole() ([]*_pkgMiddlewares.Role, error) {
	return u.middlewaresRepository.FindRole()
}
