package barrel_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-seidon/chariot/internal/barrel"
	"github.com/go-seidon/chariot/internal/repository"
	mock_repository "github.com/go-seidon/chariot/internal/repository/mock"
	mock_datetime "github.com/go-seidon/provider/datetime/mock"
	mock_identifier "github.com/go-seidon/provider/identifier/mock"
	"github.com/go-seidon/provider/system"
	mock_validation "github.com/go-seidon/provider/validation/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBarrel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Barrel Package")
}

var _ = Describe("Barrel Package", func() {

	Context("CreateBarrel function", Label("unit"), func() {

		var (
			ctx           context.Context
			currentTs     time.Time
			barrelService barrel.Barrel
			p             barrel.CreateBarrelParam
			validator     *mock_validation.MockValidator
			identifier    *mock_identifier.MockIdentifier
			clock         *mock_datetime.MockClock
			barrelRepo    *mock_repository.MockBarrel
			createParam   repository.CreateBarrelParam
			createRes     *repository.CreateBarrelResult
		)

		BeforeEach(func() {
			ctx = context.Background()
			currentTs = time.Now().UTC()
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			validator = mock_validation.NewMockValidator(ctrl)
			identifier = mock_identifier.NewMockIdentifier(ctrl)
			clock = mock_datetime.NewMockClock(ctrl)
			barrelRepo = mock_repository.NewMockBarrel(ctrl)
			barrelService = barrel.NewBarrel(barrel.BarrelParam{
				Validator:  validator,
				Identifier: identifier,
				Clock:      clock,
				BarrelRepo: barrelRepo,
			})
			p = barrel.CreateBarrelParam{
				Code:     "code",
				Name:     "name",
				Provider: "goseidon_hippo",
				Status:   "active",
			}
			createParam = repository.CreateBarrelParam{
				Id:        "id",
				Code:      p.Code,
				Name:      p.Name,
				Provider:  p.Provider,
				Status:    p.Status,
				CreatedAt: currentTs,
			}
			createRes = &repository.CreateBarrelResult{
				Id:        "id",
				Code:      p.Code,
				Name:      p.Name,
				Provider:  p.Provider,
				Status:    p.Status,
				CreatedAt: currentTs,
			}
		})

		When("there is invalid data", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(p)).
					Return(fmt.Errorf("invalid data")).
					Times(1)

				res, err := barrelService.CreateBarrel(ctx, p)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1002)))
				Expect(err.Message).To(Equal("invalid data"))
			})
		})

		When("failed generate id", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(p)).
					Return(nil).
					Times(1)

				identifier.
					EXPECT().
					GenerateId().
					Return("", fmt.Errorf("generate error")).
					Times(1)

				res, err := barrelService.CreateBarrel(ctx, p)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1001)))
				Expect(err.Message).To(Equal("generate error"))
			})
		})

		When("failed create barrel", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(p)).
					Return(nil).
					Times(1)

				identifier.
					EXPECT().
					GenerateId().
					Return("id", nil).
					Times(1)

				clock.
					EXPECT().
					Now().
					Return(currentTs).
					Times(1)

				barrelRepo.
					EXPECT().
					CreateBarrel(gomock.Eq(ctx), gomock.Eq(createParam)).
					Return(nil, fmt.Errorf("network error")).
					Times(1)

				res, err := barrelService.CreateBarrel(ctx, p)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1001)))
				Expect(err.Message).To(Equal("network error"))
			})
		})

		When("barrel is already exists", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(p)).
					Return(nil).
					Times(1)

				identifier.
					EXPECT().
					GenerateId().
					Return("id", nil).
					Times(1)

				clock.
					EXPECT().
					Now().
					Return(currentTs).
					Times(1)

				barrelRepo.
					EXPECT().
					CreateBarrel(gomock.Eq(ctx), gomock.Eq(createParam)).
					Return(nil, repository.ErrExists).
					Times(1)

				res, err := barrelService.CreateBarrel(ctx, p)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1002)))
				Expect(err.Message).To(Equal("barrel is already exists"))
			})
		})

		When("success create barrel", func() {
			It("should return result", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(p)).
					Return(nil).
					Times(1)

				identifier.
					EXPECT().
					GenerateId().
					Return("id", nil).
					Times(1)

				clock.
					EXPECT().
					Now().
					Return(currentTs).
					Times(1)

				barrelRepo.
					EXPECT().
					CreateBarrel(gomock.Eq(ctx), gomock.Eq(createParam)).
					Return(createRes, nil).
					Times(1)

				res, err := barrelService.CreateBarrel(ctx, p)

				Expect(err).To(BeNil())
				Expect(res.Success.Code).To(Equal(int32(1000)))
				Expect(res.Success.Message).To(Equal("success create barrel"))
				Expect(res.Id).To(Equal("id"))
				Expect(res.Code).To(Equal("code"))
				Expect(res.Name).To(Equal("name"))
				Expect(res.Status).To(Equal("active"))
				Expect(res.Provider).To(Equal("goseidon_hippo"))
				Expect(res.CreatedAt).To(Equal(currentTs))
			})
		})
	})

	Context("FindBarrelById function", Label("unit"), func() {

		var (
			ctx           context.Context
			currentTs     time.Time
			barrelService barrel.Barrel
			param         barrel.FindBarrelByIdParam
			result        *barrel.FindBarrelByIdResult
			validator     *mock_validation.MockValidator
			identifier    *mock_identifier.MockIdentifier
			clock         *mock_datetime.MockClock
			barrelRepo    *mock_repository.MockBarrel
			findParam     repository.FindBarrelParam
			findRes       *repository.FindBarrelResult
		)

		BeforeEach(func() {
			ctx = context.Background()
			currentTs = time.Now().UTC()
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			validator = mock_validation.NewMockValidator(ctrl)
			identifier = mock_identifier.NewMockIdentifier(ctrl)
			clock = mock_datetime.NewMockClock(ctrl)
			barrelRepo = mock_repository.NewMockBarrel(ctrl)
			barrelService = barrel.NewBarrel(barrel.BarrelParam{
				Validator:  validator,
				Identifier: identifier,
				Clock:      clock,
				BarrelRepo: barrelRepo,
			})

			param = barrel.FindBarrelByIdParam{
				Id: "id",
			}
			findParam = repository.FindBarrelParam{
				Id: param.Id,
			}
			findRes = &repository.FindBarrelResult{
				Id:        "id",
				Code:      "code",
				Name:      "name",
				Provider:  "goseidon_hippo",
				Status:    "active",
				CreatedAt: currentTs,
			}
			result = &barrel.FindBarrelByIdResult{
				Success: system.SystemSuccess{
					Code:    1000,
					Message: "success find barrel",
				},
				Id:        findRes.Id,
				Code:      findRes.Code,
				Name:      findRes.Name,
				Provider:  findRes.Provider,
				Status:    findRes.Status,
				CreatedAt: findRes.CreatedAt,
				UpdatedAt: findRes.UpdatedAt,
			}
		})

		When("there is invalid data", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(param)).
					Return(fmt.Errorf("invalid data")).
					Times(1)

				res, err := barrelService.FindBarrelById(ctx, param)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1002)))
				Expect(err.Message).To(Equal("invalid data"))
			})
		})

		When("failed find barrel", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(param)).
					Return(nil).
					Times(1)

				barrelRepo.
					EXPECT().
					FindBarrel(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(nil, fmt.Errorf("network error")).
					Times(1)

				res, err := barrelService.FindBarrelById(ctx, param)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1001)))
				Expect(err.Message).To(Equal("network error"))
			})
		})

		When("barrel is not available", func() {
			It("should return error", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(param)).
					Return(nil).
					Times(1)

				barrelRepo.
					EXPECT().
					FindBarrel(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(nil, repository.ErrNotFound).
					Times(1)

				res, err := barrelService.FindBarrelById(ctx, param)

				Expect(res).To(BeNil())
				Expect(err.Code).To(Equal(int32(1004)))
				Expect(err.Message).To(Equal("barrel is not available"))
			})
		})

		When("barrel is available", func() {
			It("should return result", func() {
				validator.
					EXPECT().
					Validate(gomock.Eq(param)).
					Return(nil).
					Times(1)

				barrelRepo.
					EXPECT().
					FindBarrel(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				res, err := barrelService.FindBarrelById(ctx, param)

				Expect(res).To(Equal(result))
				Expect(err).To(BeNil())
			})
		})
	})

})
