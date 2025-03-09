package gignore

type Option func(*GignoreClient)

func WithTemplateDirectory(directory string) Option {
	return func(g *GignoreClient) {
		g.registry.Directory = directory
	}
}
