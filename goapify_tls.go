package goapifytls

import (
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/data-harvesters/goapify"
)

type TlsClient struct {
	tls_client.HttpClient
	actor *goapify.Actor
}

func NewTlsClient(actor *goapify.Actor, options []tls_client.HttpClientOption) (*TlsClient, error) {
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}

	return &TlsClient{
		HttpClient: client,
		actor:      actor,
	}, nil
}

func DefaultOptions() []tls_client.HttpClientOption {
	return []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_124),
		tls_client.WithNotFollowRedirects(),
	}
}

// ProxiedClient gives a given http client with a proxy if available
func (t *TlsClient) ProxiedClient() tls_client.HttpClient {
	if t.actor.ProxyConfiguration == nil {
		return t.HttpClient
	}
	proxyUrl, err := t.actor.ProxyConfiguration.Proxy()
	if err != nil {
		return nil
	}
	client := t.HttpClient

	err = client.SetProxy(proxyUrl.String())
	if err != nil {
		return nil
	}

	return client
}
