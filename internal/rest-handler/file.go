package rest_handler

import (
	"net/http"
	"strings"

	rest_app "github.com/go-seidon/chariot/generated/rest-app"
	"github.com/go-seidon/chariot/internal/file"
	"github.com/go-seidon/chariot/internal/storage/multipart"
	"github.com/go-seidon/provider/serialization"
	"github.com/go-seidon/provider/status"
	"github.com/labstack/echo/v4"
)

type fileHandler struct {
	fileClient file.File
	fileParser multipart.Parser
	serializer serialization.Serializer
}

func (h *fileHandler) UploadFile(ctx echo.Context) error {
	fileHeader, ferr := ctx.FormFile("file")
	if ferr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &rest_app.ResponseBodyInfo{
			Code:    status.INVALID_PARAM,
			Message: ferr.Error(),
		})
	}

	fileInfo, ferr := h.fileParser(fileHeader)
	if ferr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &rest_app.ResponseBodyInfo{
			Code:    status.INVALID_PARAM,
			Message: ferr.Error(),
		})
	}

	meta := map[string]string{}
	metas := strings.TrimSpace(ctx.FormValue("meta"))
	if metas != "" {
		err := h.serializer.Unmarshal([]byte(metas), &meta)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, &rest_app.ResponseBodyInfo{
				Code:    status.INVALID_PARAM,
				Message: err.Error(),
			})
		}
	}

	uploadFile, err := h.fileClient.UploadFile(ctx.Request().Context(), file.UploadFileParam{
		Data: fileInfo.Data,
		Info: file.UploadFileInfo{
			Name:      fileInfo.Name,
			Size:      fileInfo.Size,
			Mimetype:  fileInfo.Mimetype,
			Extension: fileInfo.Extension,
			Meta:      meta,
		},
		Setting: file.UploadFileSetting{
			Visibility: ctx.FormValue("visibility"),
			Barrels:    strings.Split(ctx.FormValue("barrels"), ","),
		},
	})
	if err != nil {
		switch err.Code {
		case status.INVALID_PARAM:
			return echo.NewHTTPError(http.StatusBadRequest, &rest_app.ResponseBodyInfo{
				Code:    err.Code,
				Message: err.Message,
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, &rest_app.ResponseBodyInfo{
			Code:    err.Code,
			Message: err.Message,
		})
	}

	return ctx.JSON(http.StatusCreated, &rest_app.UploadFileResponse{
		Code:    uploadFile.Success.Code,
		Message: uploadFile.Success.Message,
		Data: rest_app.UploadFileData{
			Id:         uploadFile.Id,
			Slug:       uploadFile.Slug,
			Name:       uploadFile.Name,
			Extension:  uploadFile.Extension,
			Size:       uploadFile.Size,
			Mimetype:   uploadFile.Mimetype,
			Visibility: rest_app.UploadFileDataVisibility(uploadFile.Visibility),
			Status:     rest_app.UploadFileDataStatus(uploadFile.Status),
			UploadedAt: uploadFile.UploadedAt.UnixMilli(),
			Meta: &rest_app.UploadFileData_Meta{
				AdditionalProperties: uploadFile.Meta,
			},
		},
	})
}

type FileParam struct {
	File       file.File
	FileParser multipart.Parser
	Serializer serialization.Serializer
}

func NewFile(p FileParam) *fileHandler {
	return &fileHandler{
		fileClient: p.File,
		fileParser: p.FileParser,
		serializer: p.Serializer,
	}
}
