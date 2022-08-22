package canvas_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sketch/internal/canvas"
	mock_canvas "sketch/internal/canvas/mocks"
	. "sketch/tests"
	"sketch/tests/faker"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetById(t *testing.T) {
	type fields struct {
		expectedErr error
		canvas      *canvas.Canvas
	}

	type assertArgs struct {
		gotErr     error
		body       string
		statusCode int
	}

	fakeCanvas := canvas.NewCanvas("fake draw")
	fakeCanvasJSON, _ := json.Marshal(fakeCanvas)
	fakeErr := errors.New("fake")

	tests := []struct {
		name   string
		fields fields
		assert func(t *testing.T, args assertArgs)
	}{
		{
			name: "when there is no errors getting the canvas, should return a valid result",
			fields: fields{
				canvas: &fakeCanvas,
			},
			assert: func(t *testing.T, args assertArgs) {
				t.Helper()
				assert.NoError(t, args.gotErr)
				assert.JSONEq(t, string(fakeCanvasJSON), args.body)
				assert.Equal(t, http.StatusOK, args.statusCode)
			},
		},
		{
			name: "when there is an error getting a canvas, should return it",
			fields: fields{
				expectedErr: fakeErr,
			},
			assert: func(t *testing.T, args assertArgs) {
				t.Helper()
				assert.ErrorIs(t, args.gotErr, fakeErr)
			},
		},
		{
			name: "when there is no canvas, should return a 404",
			fields: fields{
				expectedErr: canvas.ErrNotFound,
			},
			assert: func(t *testing.T, args assertArgs) {
				t.Helper()
				assert.Equal(t, http.StatusNotFound, args.statusCode)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			serviceMock := mock_canvas.NewMockService(ctrl)
			handler := canvas.NewHandler(serviceMock)
			const id = "123"
			url := fmt.Sprintf("/%s", id)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			serviceMock.EXPECT().GetByID(gomock.Any(), id).
				Times(1).
				Return(tt.fields.canvas, tt.fields.expectedErr)

			params := httprouter.Params{{Key: "id", Value: id}}
			err := handler.GetById(w, req, params)

			args := assertArgs{
				body:       w.Body.String(),
				gotErr:     err,
				statusCode: w.Code,
			}
			tt.assert(t, args)
		})
	}
}

func TestHandler_Draw(t *testing.T) {
	type assertArgs struct {
		gotErr      error
		gotResponse string
	}
	type arrangeArgs struct {
		body             []byte
		called           int
		expectedResponse *canvas.DrawResponse
		expectedErr      error
	}
	fakeErr := errors.New("fake")
	fakeResponse := &canvas.DrawResponse{
		ID:      "id",
		Drawing: "🔥",
	}
	tests := []struct {
		name    string
		arrange arrangeArgs
		assert  func(t *testing.T, args assertArgs)
	}{
		{
			name: "when there is an error reading request body, should return it",
			arrange: arrangeArgs{
				body: []byte{},
			},
			assert: func(t *testing.T, args assertArgs) {
				assert.ErrorIs(t, args.gotErr, io.EOF)
			},
		},
		{
			name: "when there is an error creating the draw, should return it",
			arrange: arrangeArgs{
				called:      1,
				body:        ToJSON(faker.NewDrawRequests(t)),
				expectedErr: fakeErr,
			},
			assert: func(t *testing.T, args assertArgs) {
				assert.ErrorIs(t, args.gotErr, fakeErr)
			},
		},
		{
			name: "when there is an error validating the requests, should return it",
			arrange: arrangeArgs{
				body: ToJSON(faker.NewInvalidDrawRequests(t)),
			},
			assert: func(t *testing.T, args assertArgs) {
				assert.NotNil(t, args.gotErr)
			},
		},
		{
			name: "when there are no errors creating the draw, should create it successfully",
			arrange: arrangeArgs{
				called:           1,
				body:             ToJSON(faker.NewDrawRequests(t)),
				expectedResponse: fakeResponse,
			},
			assert: func(t *testing.T, args assertArgs) {
				assert.JSONEq(t, string(ToJSON(fakeResponse)), args.gotResponse)
				assert.NoError(t, args.gotErr)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			serviceMock := mock_canvas.NewMockService(ctrl)
			serviceMock.EXPECT().Save(gomock.Any(), gomock.Any()).
				Times(tc.arrange.called).
				Return(tc.arrange.expectedResponse, tc.arrange.expectedErr)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(tc.arrange.body))
			handler := canvas.NewHandler(serviceMock)
			err := handler.Draw(w, r, nil)

			tc.assert(t, assertArgs{gotErr: err, gotResponse: w.Body.String()})
		})
	}
}
