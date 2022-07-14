package testutils

type DummyWriter struct{}

func (d DummyWriter) Write(p []byte) (int, error) { return 0, nil }
