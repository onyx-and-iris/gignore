package gignore

// Option is a function that configures a GignoreClient.
type Option func(*Client)

// WithTemplateDirectory sets the template directory for the GignoreClient.
func WithTemplateDirectory(directory string) Option {
	return func(c *Client) {
		c.registry.Directory = directory
	}
}
