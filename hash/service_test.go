package hash

import (
	"context"
	"crypto/md5"
	"errors"
	"testing"
)

func TestService_MD5_Err(t *testing.T) {
	// given
	mock := newResponseGetterMock(nil, errors.New("not nil error"))
	s := NewService(mock)

	// when
	_, actualErr := s.MD5(context.Background(), "any url")

	// then
	if actualErr == nil {
		t.Error("missed error")
	}
}

func TestService_MD5_Exec(t *testing.T) {
	// given
	respBody := []byte("random string")
	mock := newResponseGetterMock(respBody, nil)
	s := NewService(mock)

	// when
	actual, actualErr := s.MD5(context.Background(), "any url")

	// expected
	expected := md5.Sum(respBody)

	// then
	if actualErr != nil {
		t.Error("unexpected error")
		t.FailNow()
	}
	if actual != expected {
		t.Errorf("expected = %x, got = %x", expected, actual)
	}
}
