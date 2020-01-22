package controllers

import (
	"strconv"
	"github.com/labstack/echo"
	"go-cleanarchitecture-restapi/domain"
	"go-cleanarchitecture-restapi/interfaces/database"
	"go-cleanarchitecture-restapi/usecase"
)

type MemoController struct {
	Interactor usecase.MemoInteractor
}

func NewMemoController (sqlHandler database.SqlHandler) *MemoController {
	return &MemoController {
		Interactor: usecase.MemoInteractor {
			MemoRepository: &database.MemoRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *MemoController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	memo, err := controller.Interactor.MemoById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, memo)
	return
}

func (controller *MemoController) Index(c echo.Context) (err error) {
	memos, err := controller.Interactor.Memos()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, memos)
	return
}

func (controller *MemoController) Create(c echo.Context) (err error) {
	m := domain.Memo{}
	c.Bind(&m)
	memo, err := controller.Interactor.Add(m)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, memo)
	return
}