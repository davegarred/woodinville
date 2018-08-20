package web

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"net/http/cookiejar"
)

func TestServer(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(defaultPathResolver())
	defer ts.Close()

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       0,
	}

	res, err := client.Get(ts.URL + "/user")
	assert.Nil(err)
	assert.Equal(401, res.StatusCode)

	res, err = client.Get(ts.URL + "/?im=DG")
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)
	res, err = client.Get(ts.URL + "/user")
	assert.Nil(err)
	assert.Equal(403, res.StatusCode)

	res, err = client.Get(ts.URL + "/?im=DAV")
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)
	res, err = client.Get(ts.URL + "/user")
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)


	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)
	assert.Equal(`{"id":"DAV","name":"Dave","admin":true,"visits":{}}`, string(body))
}
