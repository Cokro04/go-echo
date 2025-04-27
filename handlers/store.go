package handlers

import (
	"net/http"
	"simple-store-app/config"
	"simple-store-app/models"

	"github.com/labstack/echo/v4"
)

func GetStores(c echo.Context) error {
	rows, err := config.DB.Query("SELECT id, name, address FROM stores")
	if err != nil {
		return err
	}
	defer rows.Close()

	var stores []models.Store
	for rows.Next() {
		var store models.Store
		if err := rows.Scan(&store.ID, &store.Name, &store.Address); err != nil {
			return err
		}
		stores = append(stores, store)
	}
	return c.JSON(http.StatusOK, stores)
}

func CreateStore(c echo.Context) error {
	s := new(models.Store)
	if err := c.Bind(s); err != nil {
		return err
	}

	_, err := config.DB.Exec("INSERT INTO stores (name, address) VALUES (?, ?)", s.Name, s.Address)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, s)
}

func UpdateStore(c echo.Context) error {
	id := c.Param("id")
	s := new(models.Store)
	if err := c.Bind(s); err != nil {
		return err
	}

	_, err := config.DB.Exec("UPDATE stores SET name = ?, address = ? WHERE id = ?", s.Name, s.Address, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, s)
}

func DeleteStore(c echo.Context) error {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM stores WHERE id = ?", id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
