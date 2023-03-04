package main

import (
	"context"
	"errors"
	"gofetch/ghapp"
	"gofetch/postedimage"
	"image"
	"net/http"
	"testing"
	"time"
)

type mockCompliantChecker struct {
	returnBool bool
	returnErr  error
}

func (m mockCompliantChecker) IsCompliant(_ image.Image) (bool, error) {
	return m.returnBool, m.returnErr
}

type mockGithubHandler struct {
	link string
	err  error
}

func (m mockGithubHandler) MakePullRequest(_ context.Context, _ ghapp.FileToCommit) (string, error) {
	return m.link, m.err

}

type mockBuilder struct {
	err error
}

func (m mockBuilder) Build(_, _, _ string) (postedimage.Image, error) {
	returnErr := m.err
	return postedimage.Image{}, returnErr
}

func TestProcessImageReturnsInternalServerErrorIfComplianceCheckerErrors(t *testing.T) {
	// given an image upload handler with a compliance checker that will error
	imgHandler := ImageUploadHandler{
		complianceHandler: mockCompliantChecker{returnErr: errors.New("error")}}
	// when the handler processes an image, and the compliance checker returns an error
	httpReturn, _ := imgHandler.processImage(context.Background(), postedimage.Image{})
	// then an internal server error will be returned.
	if httpReturn != http.StatusInternalServerError {
		t.Fail()
	}
}

func TestProcessImageReturnsPreconditionFailedIfNotCompliant(t *testing.T) {
	// given an image upload handler with a compliance checker that fails the image
	imgHandler := ImageUploadHandler{
		complianceHandler: mockCompliantChecker{returnBool: false}}
	// when the handler processes an image, and the compliance checker returns it as not compliant
	httpReturn, _ := imgHandler.processImage(context.Background(), postedimage.Image{})
	// then a precondition failed error will be returned.
	if httpReturn != http.StatusPreconditionFailed {
		t.Fail()
	}
}

func TestProcessImageReturnsServerErrorIfCannotPostToGithub(t *testing.T) {
	// given an image upload handler with a compliance checker that passes the image, but a github handler that errors
	imgHandler := ImageUploadHandler{
		complianceHandler: mockCompliantChecker{returnBool: true},
		app:               mockGithubHandler{err: errors.New("error")}}
	// when the handler processes an image, and the github handler returns an error
	httpReturn, _ := imgHandler.processImage(context.Background(), postedimage.Image{})
	// then an internal server error will be returned.
	if httpReturn != http.StatusInternalServerError {
		t.Fail()
	}
}

func TestProcessImageHappyPath(t *testing.T) {
	// given an image upload handler with a compliance checker that passes the image, and a github handler that does not error

	expectedLink := "localhost:1337/YourPullRequest"
	imgHandler := ImageUploadHandler{
		complianceHandler: mockCompliantChecker{returnBool: true},
		app:               mockGithubHandler{link: expectedLink}}
	// when the handler processes an image, and the github handler does not error
	httpReturn, link := imgHandler.processImage(context.Background(), postedimage.Image{})
	// then http accepted will be returned, alongside the expected link to the pull request.
	if httpReturn != http.StatusAccepted || link != expectedLink {
		t.Fail()
	}
}

type httpContextMock struct {
	wasAbortedWithCode int
	jsonCalled         bool
}

func (h *httpContextMock) Deadline() (deadline time.Time, ok bool) {
	return
}

func (h *httpContextMock) Done() <-chan struct{} {
	return nil
}

func (h *httpContextMock) Err() error {
	return nil
}

func (h *httpContextMock) Value(_ any) any {
	return nil
}

func (h *httpContextMock) BindJSON(_ any) error {
	return nil
}

func (h *httpContextMock) AbortWithStatus(code int) {
	h.wasAbortedWithCode = code
}

func (h *httpContextMock) JSON(_ int, _ any) {
	h.jsonCalled = true
}

func TestPostedErrorsIfBuilderErrors(t *testing.T) {
	// given an image upload handler that has a builder that always throws an error, and a mock http context
	mockContext := &httpContextMock{wasAbortedWithCode: 0}
	imgHandler := ImageUploadHandler{
		builder: mockBuilder{err: errors.New("error")},
	}
	// when a request comes in, and the builder throws an error
	imgHandler.HandleImageUpload(mockContext)
	// then the mock http context will have been aborted with the bad request code
	if mockContext.wasAbortedWithCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestReturnsJSONIfNoError(t *testing.T) {
	// given an image upload handler that has a builder that does not throw an error, and a mock http context
	mockContext := &httpContextMock{jsonCalled: false, wasAbortedWithCode: 0}
	imgHandler := ImageUploadHandler{
		builder:           mockBuilder{},
		complianceHandler: mockCompliantChecker{returnBool: false},
	}
	// when a request comes in, and the builder throws an error
	imgHandler.HandleImageUpload(mockContext)
	// then the mock http context will not been aborted
	if mockContext.wasAbortedWithCode != 0 {
		t.Fail()
	}
	// and will have called JSON
	if mockContext.jsonCalled != true {
		t.Fail()
	}
}
