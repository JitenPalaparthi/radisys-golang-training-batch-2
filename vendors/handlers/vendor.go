package handlers

import (
	"net/http"
	"time"
	"vendors/interfaces"
	"vendors/models"

	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	IVendor interfaces.IVendor
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

		ctx.JSON(http.StatusOK, v)
		ctx.Abort()
	}
}
