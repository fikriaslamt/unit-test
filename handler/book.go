package handler

import (
	"sesi_8/helper"
	"sesi_8/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) GetAllBook(c *gin.Context) {
	res, err := h.app.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	if len(res) == 0 {
		helper.NoContent(c)
		return
	}

	helper.Ok(c, res)
}
func (h HttpServer) GetBookById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetBookById(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}
func (h HttpServer) AddBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.AddBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
	}

	helper.Ok(c, res)

}
func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.UpdateBook(int64(idInt), in)
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}
func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Book Deleted Successfully")
}
