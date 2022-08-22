package canvas_test

import (
	"context"
	"sketch/internal/canvas"
	mock_canvas "sketch/internal/canvas/mocks"
	"sketch/tests/faker"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetByID(t *testing.T) {

	type repositoryMock struct {
		canvas canvas.Canvas
		error  error
	}

	fakeCanvas := canvas.NewCanvas("fake draw")

	testCases := []struct {
		name       string
		repository repositoryMock
		assert     func(t *testing.T, canvas *canvas.Canvas, err error)
	}{
		{
			name: "when repository returns an error, should return it",
			repository: repositoryMock{
				error: faker.NewError(),
			},
			assert: func(t *testing.T, result *canvas.Canvas, err error) {
				assert.Nil(t, result)
				assert.ErrorIs(t, err, faker.NewError())
			},
		},
		{
			name: "when theres is no errors, should return the canvas successfully",
			repository: repositoryMock{
				canvas: fakeCanvas,
			},
			assert: func(t *testing.T, canvas *canvas.Canvas, err error) {
				assert.EqualValues(t, fakeCanvas, *canvas)
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repositoryMock := mock_canvas.NewMockRepository(ctrl)
			service := canvas.NewService(repositoryMock, nil)
			ctx := context.Background()
			const id = "fake-id"
			repositoryMock.EXPECT().GetByID(ctx, id).
				Times(1).
				Return(tc.repository.canvas, tc.repository.error)

			result, err := service.GetByID(ctx, id)

			tc.assert(t, result, err)
		})
	}
}

func TestService_Save(t *testing.T) {

	type repositoryMock struct {
		err    error
		called int
	}

	type drawerMock struct {
		draw string
		err  error
	}

	type setupMocks struct {
		drawerMock drawerMock
		repository repositoryMock
	}

	requests := faker.NewDrawRequests(t)

	testCases := []struct {
		name   string
		mocks  setupMocks
		assert func(t *testing.T, canvas *canvas.DrawResponse, err error)
	}{
		{
			name: "when drawing fails, should return error",
			mocks: setupMocks{
				drawerMock: drawerMock{
					err: faker.NewError(),
				},
			},
			assert: func(t *testing.T, canvas *canvas.DrawResponse, err error) {
				assert.Nil(t, canvas)
				assert.ErrorIs(t, err, faker.NewError())
			},
		},
		{
			name: "when saving the canvas fails, should return an error",
			mocks: setupMocks{
				drawerMock: drawerMock{
					draw: ":)",
				},
				repository: repositoryMock{
					err:    faker.NewError(),
					called: 1,
				},
			},
			assert: func(t *testing.T, canvas *canvas.DrawResponse, err error) {
				assert.ErrorIs(t, err, faker.NewError())
				assert.Nil(t, canvas)
			},
		},
		{
			name: "when saving the canvas succeed, should return the draw response",
			mocks: setupMocks{
				drawerMock: drawerMock{
					draw: ":)",
				},
				repository: repositoryMock{
					called: 1,
				},
			},
			assert: func(t *testing.T, canvas *canvas.DrawResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, canvas)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repositoryMock := mock_canvas.NewMockRepository(ctrl)
			drawerMock := mock_canvas.NewMockDrawer(ctrl)

			service := canvas.NewService(repositoryMock, drawerMock)
			ctx := context.Background()

			drawerMock.EXPECT().Draw(requests).
				Times(1).
				Return(tc.mocks.drawerMock.draw, tc.mocks.drawerMock.err)

			repositoryMock.EXPECT().Save(ctx, gomock.Any()).
				Times(tc.mocks.repository.called).
				Return(tc.mocks.repository.err)

			result, err := service.Save(ctx, requests)

			tc.assert(t, result, err)
		})
	}
}
