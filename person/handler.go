package person

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zacscoding/gotools-example/logging"
	"github.com/zacscoding/gotools-example/person/database"
	"github.com/zacscoding/gotools-example/person/model"
	"log"
	"net/http"
)

// Handler is responsible for serving apis such as "POST /api/person", "GET /api/person/:email"
type Handler struct {
	personDB database.PersonDB
}

// NewHandler returns a new Handler with given database.PersonDB.
//
// To route apis, call Handler.Route function.
func NewHandler(db database.PersonDB) (*Handler, error) {
	return &Handler{personDB: db}, nil
}

// Route routes person apis from given echo.Echo
func (h *Handler) Route(e *echo.Echo) {
	p := e.Group("/api/person")
	p.GET("/:email", h.HandleGetPersonByEmail)
	p.POST("", h.HandleSavePerson)
	p.GET("/logging", h.HandleLoggingCheckFail)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	log.Println(string(data))
}

func (h *Handler) HandleSavePerson(c echo.Context) error {
	var p model.Person
	if err := c.Bind(&p); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if err := h.personDB.Save(ctx, &p); err != nil {
		if err == database.ErrKeyConflict {
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "duplicate email",
			})
		}
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "created",
	})
}

func (h *Handler) HandleGetPersonByEmail(c echo.Context) error {
	email := c.Param("email")
	ctx := c.Request().Context()
	person, err := h.personDB.FindByEmail(ctx, email)
	if err != nil {
		if err == database.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": fmt.Sprintf("person %s not found", email),
			})
		}
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, person)
}

func (h *Handler) HandleLoggingCheckFail(c echo.Context) error {
	logging.DefaultLogger().Warnw("normal usage", "key1", "value1")
	logging.DefaultLogger().Warnw("Mismatch key values", "mismatch") // panic occur
	logging.DefaultLogger().Infow("Invalid key type", 32)            // panic occur
	return nil
}
