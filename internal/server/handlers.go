package server

import (
	"net/http"
	"taskozon/internal/store"

	"github.com/gin-gonic/gin"
)

// Создание записи
func SetHandler(store *store.Store) gin.HandlerFunc {
	type request struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		TTL   int         `json:"ttl"`
	}

	return func(c *gin.Context) {
		req := &request{}
		if err := c.BindJSON(req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if req.Key == "" || req.Value == nil {
			c.Writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		if err := store.Set(req.Key, req.Value, req.TTL); err != nil {
			c.String(http.StatusMethodNotAllowed, err.Error())
			return
		}
		c.Writer.WriteHeader(http.StatusCreated)
	}
}

// Получение значения по ключу
func GetHandler(store *store.Store) gin.HandlerFunc {
	type request struct {
		Key string `json:"key"`
	}

	return func(c *gin.Context) {
		req := &request{}
		if err := c.BindJSON(req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if req.Key == "" {
			c.Writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		value := store.Get(req.Key)
		c.JSON(http.StatusOK, value)
	}
}

// Удаление записи по ключу
func DeleteHandler(store *store.Store) gin.HandlerFunc {
	type request struct {
		Key string `json:"key"`
	}

	return func(c *gin.Context) {
		req := &request{}
		if err := c.BindJSON(req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		store.Delete(req.Key)
		c.Writer.WriteHeader(http.StatusOK)
	}
}

// Получение всех существующих ключей
func GetAllKeysHandler(store *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		allKeys := store.GetAllKeys()
		c.JSON(http.StatusOK, allKeys)
	}
}
