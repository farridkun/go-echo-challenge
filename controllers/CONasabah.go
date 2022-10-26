package controllers

import (
	"net/http"
	"time"

	"github.com/farridkun/go-echo-challenge/configs"
	"github.com/farridkun/go-echo-challenge/models"
	"github.com/farridkun/go-echo-challenge/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var CONasabah *mongo.Collection = configs.GetCollection(configs.DB, "nasabah")
var validate = validator.New()

func COCreateNasabah(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	var nasabah models.Nasabah
	defer cancel()

	if err := c.Bind(&nasabah); err != nil {
		return c.JSON(http.StatusBadRequest, responses.RENasabah{
			Status:  http.StatusBadRequest,
			Message: "Reject!, 400 - Bad Request",
			Data: &echo.Map{
				"data": err.Error(),
			},
		})
	}

	if validationErr := validate.Struct(&nasabah); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.RENasabah{
			Status:  http.StatusBadRequest,
			Message: "Reject!, 400 - Try to check the validation",
			Data: &echo.Map{
				"data": validationErr.Error(),
			},
		})
	}

	addNasabah := models.Nasabah{
		Id:   primitive.NewObjectID(),
		Cif:  nasabah.Cif,
		Nama: nasabah.Nama,
		NoHp: nasabah.NoHp,
	}

	result, err := CONasabah.InsertOne(ctx, addNasabah)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.RENasabah{
			Status:  http.StatusInternalServerError,
			Message: "Reject!, 500 - Internal Server Error",
			Data: &echo.Map{
				"data": err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, responses.RENasabah{
		Status:  http.StatusCreated,
		Message: "Granted!, Successfully to add new Nasabah",
		Data: &echo.Map{
			"data": result,
		},
	})
}
