package httpHelper

import (
	// "bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dorajistyle/goyangi/util/log"
)

// PostToTarget send a post request to a target address.
func PostToTarget(req *http.Request) (string, int, error) {
	client := &http.Client{Timeout: time.Second * 10}
	return PostToTargeWithClient(client, req)
}

// PostToTargetSSLSelfSigned send a post request to a target.
func PostToTargetSSLSelfSigned(req *http.Request) (string, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr,
		Timeout: time.Second * 10}
	return PostToTargeWithClient(client, req)
}

// PostToTargeWithClient send a post request to a target with custom client.
func PostToTargeWithClient(client *http.Client, req *http.Request) (string, int, error) {
	var responseBody string
	resp, err := client.Do(req)
	if err != nil {
		return responseBody, http.StatusBadRequest, err
		// panic(err)
	}
	defer resp.Body.Close()
	//
	log.Debugf("response Status: %s\n", resp.Status)
	log.Debugf("response Headers: %s\n", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseBody, http.StatusBadRequest, err
		// panic(err)
	}
	responseBody = string(body)
	log.Debugf("response Body:\n", responseBody)
	return responseBody, http.StatusOK, nil
}
