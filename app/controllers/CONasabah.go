package controllers

import (
	"net/http"
	"time"

	"github.com/farridkun/go-echo-challenge/app/models"
	"github.com/farridkun/go-echo-challenge/app/responses"
	"github.com/farridkun/go-echo-challenge/infra/database"
	"github.com/farridkun/go-echo-challenge/pkg/auth"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var CONasabah *mongo.Collection = database.GetCollection(database.DB, "nasabah")
var validate = validator.New()

func CreateDataNasabah(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	var nasabah models.Nasabah
	defer cancel()

	hash, err := auth.EncryptPassword(nasabah.Password)

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
		Id:       primitive.NewObjectID(),
		Cif:      nasabah.Cif,
		Nama:     nasabah.Nama,
		NoHp:     nasabah.NoHp,
		Email:    nasabah.Email,
		Password: string(hash),
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

func GetDataNasabah(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	nasabahId := c.Param("nasabahId")
	var nasabah models.Nasabah
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(nasabahId)

	err := CONasabah.FindOne(ctx, bson.M{"id": objId}).Decode(&nasabah)

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
		Status:  http.StatusOK,
		Message: "Granted!, Successfully to getting data Nasabah",
		Data: &echo.Map{
			"data": nasabah,
		},
	})
}

func UpdateDataNasabah(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	nasabahId := c.Param("nasabahId")
	var nasabah models.Nasabah
	defer cancel()

	hash, err := auth.EncryptPassword(nasabah.Password)
	objId, _ := primitive.ObjectIDFromHex(nasabahId)

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

	update := bson.M{
		"cif":      nasabah.Cif,
		"nama":     nasabah.Nama,
		"noHp":     nasabah.NoHp,
		"email":    nasabah.Email,
		"password": string(hash),
	}

	result, err := CONasabah.UpdateOne(
		ctx,
		bson.M{"id": objId},
		bson.M{"$set": update},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.RENasabah{
			Status:  http.StatusInternalServerError,
			Message: "Reject!, 500 - Internal Server Error",
			Data: &echo.Map{
				"data": err.Error(),
			},
		})
	}

	var updatedNasabah models.Nasabah
	if result.MatchedCount == 1 {
		err := CONasabah.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedNasabah)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.RENasabah{
				Status:  http.StatusInternalServerError,
				Message: "Reject!, 500 - Internal Server Error",
				Data: &echo.Map{
					"data": err.Error(),
				},
			})
		}
	}

	return c.JSON(http.StatusOK, responses.RENasabah{
		Status:  http.StatusOK,
		Message: "Granted!, Successfully to updating data Nasabah",
		Data: &echo.Map{
			"data": updatedNasabah,
		},
	})
}

func DeleteDataNasabah(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	nasabahId := c.Param("nasabahId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(nasabahId)

	result, err := CONasabah.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.RENasabah{
			Status:  http.StatusInternalServerError,
			Message: "Reject!, 500 - Internal Server Error",
			Data: &echo.Map{
				"data": err.Error(),
			},
		})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.RENasabah{
			Status:  http.StatusInternalServerError,
			Message: "Reject!, 404 - Not Found",
			Data: &echo.Map{
				"data": "Nasabah ID was not found!",
			},
		})
	}

	return c.JSON(http.StatusOK, responses.RENasabah{
		Status:  http.StatusOK,
		Message: "Granted!, Successfully to updating data Nasabah",
		Data: &echo.Map{
			"data": "Granted!, Nasabah successfully deleted",
		},
	})
}

func GetAllNasabah(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	var nasabahAll []models.Nasabah
	defer cancel()

	results, err := CONasabah.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.RENasabah{
			Status:  http.StatusInternalServerError,
			Message: "Reject!, 500 - Internal Server Error",
			Data: &echo.Map{
				"data": err.Error(),
			},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var nasabah models.Nasabah
		if err = results.Decode(&nasabah); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.RENasabah{
				Status:  http.StatusInternalServerError,
				Message: "Reject!, 500 - Internal Server Error",
				Data: &echo.Map{
					"data": err.Error(),
				},
			})
		}

		nasabahAll = append(nasabahAll, nasabah)
	}

	return c.JSON(http.StatusOK, responses.RENasabah{
		Status:  http.StatusOK,
		Message: "Granted!, Successfully to getting all data Nasabah",
		Data: &echo.Map{
			"data": nasabahAll,
		},
	})
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hi, this is up ⚡")
}
