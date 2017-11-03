package dognzb_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		desc string
		api  string
	}{
		{
			desc: "with api", api: "some-api",
		},
		{
			desc: "empty api", api: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			d := dognzb.New(tC.api, &http.Client{})
			if d == nil {
				t.Errorf("Expected a dognzb object, got nil")
			}
		})
	}
}

func TestListHappyPath(t *testing.T) {
	testCases := []struct {
		desc   string
		api    string
		kind   dognzb.Type
		size   int
		status int
	}{
		{
			desc:   "4_movies",
			api:    "a-valid-api",
			kind:   dognzb.Movies,
			status: http.StatusOK,
			size:   4,
		}, {
			desc:   "0_movies",
			api:    "a-valid-api",
			kind:   dognzb.Movies,
			status: http.StatusOK,
			size:   0,
		}, {
			desc:   "2_series",
			api:    "a-valid-api",
			kind:   dognzb.TV,
			status: http.StatusOK,
			size:   2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			body, _ := ioutil.ReadFile(fmt.Sprintf("./fixtures/list_%s.xml", tC.desc))
			d := dognzb.New(tC.api, &mockGetter{
				response: &http.Response{
					StatusCode: tC.status,
					Body:       ioutil.NopCloser(bytes.NewReader(body)),
				},
			})

			// act
			items, err := d.List(tC.kind)

			// assert
			if err != nil {
				t.Errorf("expected err to be '%v', got '%v'", nil, err)
			}
			if size := len(items); size != tC.size {
				t.Errorf("expected size to be %v, got %v", tC.size, size)
			}
			for _, item := range items {
				if tC.kind == dognzb.TV && item.TVdbID == "" {
					t.Error("expected tvdbid to not be \"\"")
				}
				if tC.kind == dognzb.Movies && item.ImdbID == "" {
					t.Error("expected imdbid to not be \"\"")
				}
			}
		})
	}
}

func TestListSadPath(t *testing.T) {
	testCases := []struct {
		desc   string
		api    string
		size   int
		status int
		errMsg string
	}{
		{
			desc:   "404_response",
			api:    "some-api",
			status: http.StatusNotFound,
			errMsg: "bad response: 404",
		}, {
			desc:   "xml_error_code",
			api:    "a-valid-api",
			status: http.StatusOK,
			errMsg: "Incorrect user credentials",
		}, {
			desc:   "empty_response",
			api:    "a-valid-api",
			status: http.StatusOK,
			errMsg: "EOF",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			body, _ := ioutil.ReadFile(fmt.Sprintf("./fixtures/list_%s.xml", tC.desc))
			d := dognzb.New(tC.api, &mockGetter{
				response: &http.Response{
					StatusCode: tC.status,
					Body:       ioutil.NopCloser(bytes.NewReader(body)),
				},
			})

			// act
			_, err := d.List(dognzb.Movies)

			// assert
			if err == nil {
				t.Errorf("expected err to be '%v', got '%v'", nil, err)
			}

			if err.Error() != tC.errMsg {
				t.Errorf("expected error message to be '%v', got '%v'", tC.errMsg, err.Error())
			}
		})
	}
}

type mockGetter struct {
	response *http.Response
}

func (m *mockGetter) Get(url string) (*http.Response, error) {
	return m.response, nil
}
