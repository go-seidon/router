package resthandler_test

import (
	"bytes"
	encoding_json "encoding/json"
	"fmt"
	"io"
	mime_multipart "mime/multipart"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/go-seidon/chariot/generated/restapp"
	"github.com/go-seidon/chariot/internal/file"
	mock_file "github.com/go-seidon/chariot/internal/file/mock"
	"github.com/go-seidon/chariot/internal/resthandler"
	"github.com/go-seidon/chariot/internal/storage/multipart"
	mock_io "github.com/go-seidon/provider/io/mock"
	"github.com/go-seidon/provider/serialization/json"
	serialization "github.com/go-seidon/provider/serialization/mock"
	"github.com/go-seidon/provider/system"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Handler", func() {
	Context("UploadFile function", Label("unit", "slow"), func() {
		var (
			ctx         echo.Context
			h           func(ctx echo.Context) error
			rec         *httptest.ResponseRecorder
			serializer  *serialization.MockSerializer
			fileClient  *mock_file.MockFile
			uploadParam file.UploadFileParam
			uploadRes   *file.UploadFileResult
		)

		BeforeEach(func() {
			body := new(bytes.Buffer)
			writer := mime_multipart.NewWriter(body)
			meta, err := writer.CreateFormField("meta")
			if err != nil {
				AbortSuite("failed create meta field: " + err.Error())
			}

			_, err = meta.Write([]byte(`{"user_id": "123", "feature": "profile"}`))
			if err != nil {
				AbortSuite("failed write meta field: " + err.Error())
			}

			visibility, err := writer.CreateFormField("visibility")
			if err != nil {
				AbortSuite("failed create visibility field: " + err.Error())
			}

			_, err = visibility.Write([]byte("public"))
			if err != nil {
				AbortSuite("failed write visibility field: " + err.Error())
			}

			barrels, err := writer.CreateFormField("barrels")
			if err != nil {
				AbortSuite("failed create barrels field: " + err.Error())
			}

			_, err = barrels.Write([]byte("hippo1,hippo2"))
			if err != nil {
				AbortSuite("failed write barrels field: " + err.Error())
			}

			_, err = writer.CreateFormFile("file", "file.go")
			if err != nil {
				AbortSuite("failed create file mock: " + err.Error())
			}

			err = writer.Close()
			if err != nil {
				AbortSuite("failed close writer: " + err.Error())
			}

			req := httptest.NewRequest(http.MethodPost, "/", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec = httptest.NewRecorder()

			e := echo.New()
			ctx = e.NewContext(req, rec)

			t := GinkgoT()
			ctrl := gomock.NewController(t)
			serializer = serialization.NewMockSerializer(ctrl)
			jsonSerializer := json.NewSerializer()
			fileClient = mock_file.NewMockFile(ctrl)
			fileHandler := resthandler.NewFile(resthandler.FileParam{
				Serializer: jsonSerializer,
				File:       fileClient,
				FileParser: func(h *mime_multipart.FileHeader) (*multipart.FileInfo, error) {
					return &multipart.FileInfo{
						Data:      nil,
						Name:      "dolphin 22",
						Size:      23342,
						Extension: "jpg",
						Mimetype:  "image/jpeg",
					}, nil
				},
			})
			h = fileHandler.UploadFile

			uploadParam = file.UploadFileParam{
				Data: nil,
				Info: file.UploadFileInfo{
					Name:      "dolphin 22",
					Mimetype:  "image/jpeg",
					Extension: "jpg",
					Size:      23342,
					Meta: map[string]string{
						"user_id": "123",
						"feature": "profile",
					},
				},
				Setting: file.UploadFileSetting{
					Visibility: "public",
					Barrels:    []string{"hippo1", "hippo2"},
				},
			}
			uploadRes = &file.UploadFileResult{
				Success: system.SystemSuccess{
					Code:    1000,
					Message: "success upload file",
				},
				Id:         "id",
				Slug:       "dolphine-22",
				Name:       "dolphin 22",
				Mimetype:   "image/jpeg",
				Extension:  "jpg",
				Size:       23342,
				Visibility: "public",
				Status:     "available",
				Meta: map[string]string{
					"user_id": "123",
					"feature": "profile",
				},
				UploadedAt: time.Now().UTC(),
			}
		})

		When("failed bind multipart file", func() {
			It("should return error", func() {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
				rec = httptest.NewRecorder()

				e := echo.New()
				ctx := e.NewContext(req, rec)

				err := h(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 400,
					Message: &restapp.ResponseBodyInfo{
						Code:    1002,
						Message: "no multipart boundary param in Content-Type",
					},
				}))
			})
		})

		When("failed parse file", func() {
			It("should return error", func() {
				fileHandler := resthandler.NewFile(resthandler.FileParam{
					Serializer: serializer,
					File:       fileClient,
					FileParser: func(h *mime_multipart.FileHeader) (*multipart.FileInfo, error) {
						return nil, fmt.Errorf("disk error")
					},
				})
				err := fileHandler.UploadFile(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 400,
					Message: &restapp.ResponseBodyInfo{
						Code:    1002,
						Message: "disk error",
					},
				}))
			})
		})

		When("failed parse meta", func() {
			It("should return error", func() {
				fileHandler := resthandler.NewFile(resthandler.FileParam{
					Serializer: serializer,
					File:       fileClient,
					FileParser: func(h *mime_multipart.FileHeader) (*multipart.FileInfo, error) {
						return &multipart.FileInfo{
							Data:      nil,
							Name:      "dolphin 22",
							Size:      23342,
							Extension: "jpg",
							Mimetype:  "image/jpeg",
						}, nil
					},
				})

				serializer.
					EXPECT().
					Unmarshal(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("invalid data type")).
					Times(1)

				err := fileHandler.UploadFile(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 400,
					Message: &restapp.ResponseBodyInfo{
						Code:    1002,
						Message: "invalid data type",
					},
				}))
			})
		})

		When("there is invalid data", func() {
			It("should return error", func() {
				fileClient.
					EXPECT().
					UploadFile(gomock.Eq(ctx.Request().Context()), gomock.Eq(uploadParam)).
					Return(nil, &system.SystemError{
						Code:    1002,
						Message: "invalid param",
					}).
					Times(1)

				err := h(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 400,
					Message: &restapp.ResponseBodyInfo{
						Code:    1002,
						Message: "invalid param",
					},
				}))
			})
		})

		When("failed upload file", func() {
			It("should return error", func() {
				fileClient.
					EXPECT().
					UploadFile(gomock.Eq(ctx.Request().Context()), gomock.Eq(uploadParam)).
					Return(nil, &system.SystemError{
						Code:    1001,
						Message: "network error",
					}).
					Times(1)

				err := h(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 500,
					Message: &restapp.ResponseBodyInfo{
						Code:    1001,
						Message: "network error",
					},
				}))
			})
		})

		When("success upload file", func() {
			It("should return result", func() {
				fileClient.
					EXPECT().
					UploadFile(gomock.Eq(ctx.Request().Context()), gomock.Eq(uploadParam)).
					Return(uploadRes, nil).
					Times(1)

				err := h(ctx)

				res := &restapp.UploadFileResponse{}
				encoding_json.Unmarshal(rec.Body.Bytes(), res)

				Expect(err).To(BeNil())
				Expect(res.Code).To(Equal(uploadRes.Success.Code))
				Expect(res.Message).To(Equal(uploadRes.Success.Message))
				Expect(res.Data).To(Equal(restapp.UploadFileData{
					Id:         uploadRes.Id,
					Slug:       uploadRes.Slug,
					Name:       uploadRes.Name,
					Extension:  uploadRes.Extension,
					Mimetype:   uploadRes.Mimetype,
					Size:       uploadRes.Size,
					Status:     restapp.UploadFileDataStatus(uploadRes.Status),
					Visibility: restapp.UploadFileDataVisibility(uploadRes.Visibility),
					UploadedAt: uploadRes.UploadedAt.Local().UnixMilli(),
					Meta: &restapp.UploadFileData_Meta{
						AdditionalProperties: uploadRes.Meta,
					},
				}))
			})
		})
	})

	Context("RetrieveFileBySlug function", Label("unit"), func() {
		var (
			currentTs  time.Time
			ctx        echo.Context
			h          func(ctx echo.Context) error
			rec        *httptest.ResponseRecorder
			fileClient *mock_file.MockFile
			findParam  file.RetrieveFileBySlugParam
			findRes    *file.RetrieveFileBySlugResult
			fileData   *mock_io.MockReadCloser
		)

		BeforeEach(func() {
			currentTs = time.Now().UTC()

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec = httptest.NewRecorder()

			e := echo.New()
			ctx = e.NewContext(req, rec)
			ctx.SetParamNames("slug")
			ctx.SetParamValues("lumba.jpg")

			t := GinkgoT()
			ctrl := gomock.NewController(t)
			fileClient = mock_file.NewMockFile(ctrl)
			fileHandler := resthandler.NewFile(resthandler.FileParam{
				File: fileClient,
			})
			h = fileHandler.RetrieveFileBySlug
			findParam = file.RetrieveFileBySlugParam{
				Slug: "lumba.jpg",
			}
			fileData = mock_io.NewMockReadCloser(ctrl)
			findRes = &file.RetrieveFileBySlugResult{
				Success: system.SystemSuccess{
					Code:    1000,
					Message: "success retrieve file",
				},
				Data:       fileData,
				Id:         "id",
				Slug:       "lumba.jpg",
				Name:       "Lumba",
				Mimetype:   "image/jpeg",
				Extension:  "jpg",
				Size:       23343,
				Visibility: "public",
				Status:     "available",
				Meta:       map[string]string{},
				UploadedAt: currentTs,
			}
		})

		When("there is invalid data", func() {
			It("should return error", func() {
				fileClient.
					EXPECT().
					RetrieveFileBySlug(gomock.Eq(ctx.Request().Context()), gomock.Eq(findParam)).
					Return(nil, &system.SystemError{
						Code:    1002,
						Message: "invalid data",
					}).
					Times(1)

				err := h(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 400,
					Message: &restapp.ResponseBodyInfo{
						Code:    1002,
						Message: "invalid data",
					},
				}))
			})
		})

		When("file is not available", func() {
			It("should return error", func() {
				fileClient.
					EXPECT().
					RetrieveFileBySlug(gomock.Eq(ctx.Request().Context()), gomock.Eq(findParam)).
					Return(nil, &system.SystemError{
						Code:    1004,
						Message: "file is not available",
					}).
					Times(1)

				err := h(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 404,
					Message: &restapp.ResponseBodyInfo{
						Code:    1004,
						Message: "file is not available",
					},
				}))
			})
		})

		When("failed find file", func() {
			It("should return error", func() {
				fileClient.
					EXPECT().
					RetrieveFileBySlug(gomock.Eq(ctx.Request().Context()), gomock.Eq(findParam)).
					Return(nil, &system.SystemError{
						Code:    1001,
						Message: "network error",
					}).
					Times(1)

				err := h(ctx)

				Expect(err).To(Equal(&echo.HTTPError{
					Code: 500,
					Message: &restapp.ResponseBodyInfo{
						Code:    1001,
						Message: "network error",
					},
				}))
			})
		})

		When("success retrieve file", func() {
			It("should return error", func() {
				fileData.
					EXPECT().
					Read(gomock.Any()).
					Return(0, io.EOF).
					Times(1)

				fileClient.
					EXPECT().
					RetrieveFileBySlug(gomock.Eq(ctx.Request().Context()), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				err := h(ctx)

				Expect(err).To(BeNil())
			})
		})
	})
})
