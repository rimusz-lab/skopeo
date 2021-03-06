package main

import (
	"github.com/containers/image/transports"
	"github.com/containers/image/types"
	"github.com/urfave/cli"
)

// contextFromGlobalOptions returns a types.SystemContext depending on c.
func contextFromGlobalOptions(c *cli.Context) *types.SystemContext {
	tlsVerify := c.GlobalBoolT("tls-verify")
	return &types.SystemContext{
		RegistriesDirPath:           c.GlobalString("registries.d"),
		DockerCertPath:              c.GlobalString("cert-path"),
		DockerInsecureSkipTLSVerify: !tlsVerify,
	}
}

// ParseImage converts image URL-like string to an initialized handler for that image.
// The caller must call .Close() on the returned Image.
func parseImage(c *cli.Context) (types.Image, error) {
	imgName := c.Args().First()
	ref, err := transports.ParseImageName(imgName)
	if err != nil {
		return nil, err
	}
	return ref.NewImage(contextFromGlobalOptions(c))
}

// parseImageSource converts image URL-like string to an ImageSource.
// requestedManifestMIMETypes is as in types.ImageReference.NewImageSource.
// The caller must call .Close() on the returned ImageSource.
func parseImageSource(c *cli.Context, name string, requestedManifestMIMETypes []string) (types.ImageSource, error) {
	ref, err := transports.ParseImageName(name)
	if err != nil {
		return nil, err
	}
	return ref.NewImageSource(contextFromGlobalOptions(c), requestedManifestMIMETypes)
}
