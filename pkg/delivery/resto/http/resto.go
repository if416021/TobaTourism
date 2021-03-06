package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/TobaTourism/pkg/models"
)

func (d *resto) GetAllRestoWithKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := d.restoUsecase.GetAllRestoWithKuliner()
	if err != nil {
		log.Println(err)
	}
	resp.Data = data

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) GetAllResto(c echo.Context) error {
	// user := c.Request().Context().Value("user").(uint) //Grab the id of the user that send the request
	// log.Println(user)
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := d.restoUsecase.GetAllResto()
	if err != nil {
		log.Println(err)
	}
	resp.Data = data

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) GetDetailResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	restoID, _ := strconv.ParseInt(c.Param("restaurantID"), 10, 64)

	data, err := d.restoUsecase.GetDetailResto(restoID)
	if err != nil {
		log.Println(err)
	}
	resp.Data = data

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) InsertResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	name := c.FormValue("restaurantName")
	location := c.FormValue("restaurantLocation")
	contact := c.FormValue("restaurantContact")

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Restoran][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	files := form.File["restaurantImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileRestoran, models.RestoranTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Restoran][InsertAttachment] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = d.restoUsecase.CreateResto(name, location, contact, attachmentID)
	if err != nil {
		log.Println("[Delivery][Restoran][CreateResto] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) UpdateImageResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	restoID := c.Param("restaurantID")

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Restoran][MultipartForm update image] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	files := form.File["restaurantImage"]
	log.Println("Files ", files)
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileRestoran, models.RestoranTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Restoran][InsertAttachment for Update] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = d.restoUsecase.UpdateImageRestoran(restoID, attachmentID)
	if err != nil {
		log.Println("[Delivery][Restoran][UpdateImageResto] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) UpdateResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	restoID := c.Param("restaurantID")
	restoName := c.FormValue("restaurantName")
	restoContact := c.FormValue("restaurantContact")
	restoLocation := c.FormValue("restaurantLocation")

	err := d.restoUsecase.UpdateRestoran(restoID, restoName, restoContact, restoLocation)
	if err != nil {
		log.Println("[Delivery][Restoran][UpadateRespo] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) DeleteResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	restoID := c.Param("restaurantID")
	log.Println(restoID)
	err := d.restoUsecase.DeleteResto(restoID)
	if err != nil {
		log.Println("[Delivery][Restoran][Delete] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
