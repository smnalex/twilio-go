package chat

import (
	"context"
	"io"
)

type mockHTTPClient struct {
	GetInvoked    bool
	GetFunc       func(context.Context, string) ([]byte, error)
	PostInvoked   bool
	PostFunc      func(context.Context, string, io.Reader) ([]byte, error)
	DeleteInvoked bool
	DeleteFunc    func(context.Context, string) ([]byte, error)
}

func (m *mockHTTPClient) Get(ctx context.Context, path string) ([]byte, error) {
	m.GetInvoked = true
	return m.GetFunc(ctx, path)
}

func (m *mockHTTPClient) Post(ctx context.Context, path string, body io.Reader) ([]byte, error) {
	m.PostInvoked = true
	return m.PostFunc(ctx, path, body)
}

func (m *mockHTTPClient) Delete(ctx context.Context, path string) ([]byte, error) {
	m.DeleteInvoked = true
	return m.DeleteFunc(ctx, path)
}
