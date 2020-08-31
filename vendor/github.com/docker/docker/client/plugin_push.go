package client // import "github.com/docker/docker/client"

import (
	"io"

	"golang.org/x/net/context"
)

// PluginPush pushes a plugin to a registry
func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error) {
	headers := map[string][]string{"X-Registry-Auth": {registryAuth}}
	resp, err := cli.post(ctx, "/plugins/"+name+"/push", nil, nil, headers)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
