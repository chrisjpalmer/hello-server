// HelloServer containerised tooling

package main

import (
	"context"
	"dagger/htmlx-test/internal/dagger"
	"fmt"

	"golang.org/x/mod/modfile"
)

type HelloServer struct {
	Src *dagger.Directory
}

func New(
	// +defaultPath="/"
	// +ignore=[".dagger/"]
	src *dagger.Directory,
) *HelloServer {
	return &HelloServer{Src: src}
}

func (m *HelloServer) Run(ctx context.Context) (*dagger.Service, error) {
	ctr, err := m.Build(ctx)
	if err != nil {
		return nil, fmt.Errorf("error building source: %w", err)
	}

	return ctr.AsService(), nil
}

// +check
func (m *HelloServer) Build(ctx context.Context) (*dagger.Container, error) {
	ctr, err := m.buildCtr(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting build container: %w", err)
	}

	app := ctr.
		WithExec([]string{"go", "build", "-o", "app"}).
		File("app")

	return dag.Container().
		From("alpine:latest").
		WithWorkdir("/app").
		WithFile("app", app).
		WithExposedPort(8080).
		WithEntrypoint([]string{"/app/app"}), nil
}

// +check
func (m *HelloServer) CheckGenerated(ctx context.Context) error {
	chgset, err := m.Generate(ctx)
	if err != nil {
		return fmt.Errorf("error generating templates: %w", err)
	}

	empty, err := chgset.IsEmpty(ctx)
	if err != nil {
		return fmt.Errorf("error calling is empty: %w", err)
	}

	if !empty {
		return fmt.Errorf("templates are not up to date")
	}

	return nil
}

// +generate
func (m *HelloServer) Generate(ctx context.Context) (*dagger.Changeset, error) {
	ctr, err := m.buildCtr(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting build container: %w", err)
	}

	src := withTemplGenerate(ctr).Directory(".")

	return m.Src.WithDirectory(".", src).Changes(m.Src), nil
}

func withTemplGenerate(ctr *dagger.Container) *dagger.Container {
	return ctr.WithExec([]string{"go", "tool", "templ", "generate"})
}

func (m *HelloServer) buildCtr(ctx context.Context) (*dagger.Container, error) {
	ver, err := m.goVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting go version: %w", err)
	}

	return dag.Container().
		From("golang:"+ver).
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-cache")).
		WithEnvVariable("CGO_ENABLED", "0").
		WithWorkdir("/src").
		WithDirectory("/src", m.Src), nil

}

func (m *HelloServer) goVersion(ctx context.Context) (string, error) {
	s, err := m.Src.File("go.mod").Contents(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting file contents: %w", err)
	}

	f, err := modfile.Parse("go.mod", []byte(s), nil)
	if err != nil {
		return "", fmt.Errorf("error parsing go.mod file: %w", err)
	}

	return f.Go.Version, nil
}
