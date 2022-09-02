package handlers

import (
	"net/http"
	"strconv"
	"time"
	"vendors/interfaces"
	"vendors/mb"
	"vendors/models"

	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	IVendor interfaces.IVendor
	Conn    []string
}

func (vh *VendorHandler) Create() func(*gin.Context) {
	return func(ctx *gin.Context) {
		vendor := new(models.Vendor)
		// var buf []byte
		// _, err := ctx.Request.Body.Read(buf)
		// err = json.Unmarshal(buf, vendor)
		err := ctx.Bind(vendor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}

		err = vendor.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}

		vendor.Status = "created"
		vendor.LastModified = time.Now().Unix()

		v, err := vh.IVendor.Add(vendor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}
		err = mb.Publish(vh.Conn, "vendors.created", *vendor)
		ctx.JSON(http.StatusCreated, v)
		ctx.Abort()
	}
}

func (vh *VendorHandler) GetBy() func(*gin.Context) {
	return func(ctx *gin.Context) {

		_id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid parameter"})
			ctx.Abort()
			return
		}

		id, err := strconv.Atoi(_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid parameter:" + err.Error()})
			ctx.Abort()
			return
		}
		v, err := vh.IVendor.GetBy(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, v)
		ctx.Abort()
	}
}

func (vh *VendorHandler) GetAll() func(*gin.Context) {
	return func(ctx *gin.Context) {
		vendors, err := vh.IVendor.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, vendors)
		ctx.Abort()
	}
}
func (vh *VendorHandler) UpdateBy() func(*gin.Context) {
	return func(ctx *gin.Context) {
		_id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid parameter"})
			ctx.Abort()
			return
		}
		id, err := strconv.Atoi(_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid parameter:" + err.Error()})
			ctx.Abort()
			return
		}
		data := make(map[string]interface{})
		err = ctx.Bind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}
		v, err := vh.IVendor.UpdateBy(id, data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, v)
		ctx.Abort()
	}
}

func (vh *VendorHandler) DeleteBy() func(*gin.Context) {
	return func(ctx *gin.Context) {

		_id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid parameter"})
			ctx.Abort()
			return
		}

		id, err := strconv.Atoi(_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid parameter:" + err.Error()})
			ctx.Abort()
			return
		}

		v, err := vh.IVendor.DeleteBy(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, v)
		ctx.Abort()
	}
}
