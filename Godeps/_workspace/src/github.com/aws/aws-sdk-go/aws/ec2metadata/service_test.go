package ec2metadata_test

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/DaveBlooman/api-common/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/DaveBlooman/api-common/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/DaveBlooman/api-common/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/session"
	"github.com/DaveBlooman/api-common/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestClientOverrideDefaultHTTPClientTimeout(t *testing.T) {
	svc := ec2metadata.New(session.New())

	assert.NotEqual(t, http.DefaultClient, svc.Config.HTTPClient)
	assert.Equal(t, 5*time.Second, svc.Config.HTTPClient.Timeout)
}

func TestClientNotOverrideDefaultHTTPClientTimeout(t *testing.T) {
	origClient := *http.DefaultClient
	http.DefaultClient.Transport = &http.Transport{}
	defer func() {
		http.DefaultClient = &origClient
	}()

	svc := ec2metadata.New(session.New())

	assert.Equal(t, http.DefaultClient, svc.Config.HTTPClient)

	tr, ok := svc.Config.HTTPClient.Transport.(*http.Transport)
	assert.True(t, ok)
	assert.NotNil(t, tr)
	assert.Nil(t, tr.Dial)
}

func TestClientDisableOverrideDefaultHTTPClientTimeout(t *testing.T) {
	svc := ec2metadata.New(session.New(aws.NewConfig().WithEC2MetadataDisableTimeoutOverride(true)))

	assert.Equal(t, http.DefaultClient, svc.Config.HTTPClient)
}

func TestClientOverrideDefaultHTTPClientTimeoutRace(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("us-east-1a"))
	}))

	cfg := aws.NewConfig().WithEndpoint(server.URL)
	runEC2MetadataClients(t, cfg, 100)
}

func TestClientOverrideDefaultHTTPClientTimeoutRaceWithTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("us-east-1a"))
	}))

	cfg := aws.NewConfig().WithEndpoint(server.URL).WithHTTPClient(&http.Client{
		Transport: http.DefaultTransport,
	})

	runEC2MetadataClients(t, cfg, 100)
}

func runEC2MetadataClients(t *testing.T, cfg *aws.Config, atOnce int) {
	var wg sync.WaitGroup
	wg.Add(atOnce)
	for i := 0; i < atOnce; i++ {
		go func() {
			svc := ec2metadata.New(session.New(), cfg)
			_, err := svc.Region()
			assert.NoError(t, err)
			wg.Done()
		}()
	}
	wg.Wait()
}
