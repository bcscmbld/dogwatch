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
		kind   dognzb.Type
		size   int
		status int
	}{
		{
			desc:   "4_movies",
			kind:   dognzb.Movies,
			status: http.StatusOK,
			size:   4,
		}, {
			desc:   "0_movies",
			kind:   dognzb.Movies,
			status: http.StatusOK,
			size:   0,
		}, {
			desc:   "2_series",
			kind:   dognzb.TV,
			status: http.StatusOK,
			size:   2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			d := NewMockGetter("list", tC.desc, tC.status)

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
		status int
		errMsg string
	}{
		{
			desc:   "404_response",
			status: http.StatusNotFound,
			errMsg: "bad response: 404",
		}, {
			desc:   "xml_error_code",
			status: http.StatusOK,
			errMsg: "Incorrect user credentials",
		}, {
			desc:   "empty_response",
			status: http.StatusOK,
			errMsg: "EOF",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			d := NewMockGetter("list", tC.desc, tC.status)

			// act
			_, err := d.List(dognzb.Movies)

			// assert
			if err == nil {
				t.Errorf("expected err to be '%v', got '%v'", nil, err)
			}

			if err.Error() != tC.errMsg {
				t.Errorf(
					"expected error message to be '%v', got '%v'",
					tC.errMsg,
					err.Error(),
				)
			}
		})
	}
}

func TestAddHappyPath(t *testing.T) {
	testCases := []struct {
		desc string
		kind dognzb.Type
		id   string
	}{
		{
			desc: "series_added", kind: dognzb.TV, id: "247808",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			d := NewMockGetter("add", tC.desc, http.StatusOK)

			// act
			q, err := d.Add(tC.kind, tC.id)

			// assert
			if err != nil {
				t.Errorf("expected err to be '%v', got '%v'", nil, err)
			}
			if q == nil {
				t.Errorf("expected q not to be '%v', got '%v'", nil, q)
			}
		})
	}
}

func TestAddSadPath(t *testing.T) {
	testCases := []struct {
		desc   string
		kind   dognzb.Type
		id     string
		errMsg string
		status int
	}{
		{
			desc:   "already_exists",
			kind:   dognzb.TV,
			id:     "247808",
			status: http.StatusOK,
			errMsg: "Game of Thrones already exists in your TV Watchlist.",
		}, {
			desc:   "404_response",
			kind:   dognzb.TV,
			id:     "234234",
			status: http.StatusNotFound,
			errMsg: "bad response: 404",
		}, {
			desc:   "empty_response",
			kind:   dognzb.Movies,
			status: http.StatusOK,
			id:     "5637",
			errMsg: "EOF",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			d := NewMockGetter("add", tC.desc, tC.status)

			// act
			q, err := d.Add(tC.kind, tC.id)

			// assert
			if err == nil {
				t.Errorf("expected err to not be '%v', got '%v'", nil, err)
			}

			if err.Error() != tC.errMsg {
				t.Errorf("expected err to not be '%v', got '%v'", tC.errMsg, err)
			}

			if q != nil {
				t.Errorf("expected q to be '%v', got '%v'", nil, q)
			}
		})
	}
}

func TestRemoveHappyPath(t *testing.T) {
	testCases := []struct {
		desc string
		kind dognzb.Type
		id   string
	}{
		{
			desc: "series_removed", kind: dognzb.TV, id: "247808",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			d := NewMockGetter("remove", tC.desc, http.StatusOK)

			// act
			q, err := d.Remove(tC.kind, tC.id)

			// assert
			if err != nil {
				t.Errorf("expected err to be '%v', got '%v'", nil, err)
			}
			if q.Code == "" {
				t.Errorf("expected Code to exists, got '%v'", q.Code)
			}
			if q.Description == "" {
				t.Errorf("expected Description to exist, got '%v'", q.Description)
			}
			if q.ErrorCode != "" {
				t.Errorf("expected no ErrorCode, got '%v'", q.ErrorCode)
			}
			if q.ErrorDesc != "" {
				t.Errorf("expected no ErrorDesc, got '%v'", q.ErrorDesc)
			}
		})
	}
}

func TestRemoveSadPath(t *testing.T) {
	testCases := []struct {
		desc   string
		kind   dognzb.Type
		id     string
		status int
		errMsg string
	}{
		{
			desc:   "failed",
			kind:   dognzb.TV,
			id:     "247808",
			status: http.StatusOK,
			errMsg: "Unable to find something...",
		}, {
			desc:   "404_response",
			kind:   dognzb.TV,
			id:     "234234",
			status: http.StatusNotFound,
			errMsg: "bad response: 404",
		}, {
			desc:   "empty_response",
			kind:   dognzb.Movies,
			status: http.StatusOK,
			id:     "5637",
			errMsg: "EOF",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			d := NewMockGetter("remove", tC.desc, tC.status)

			// act
			q, err := d.Remove(tC.kind, tC.id)

			// assert
			if err == nil {
				t.Errorf("expected err to be '%v', got '%v'", nil, err)
			}
			if q != nil {
				t.Errorf("expected q to be '%v', got '%v'", nil, q)
			}
			if err.Error() != tC.errMsg {
				t.Errorf("expected errMsg to be '%v', got '%v'", tC.errMsg, err)
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

func NewMockGetter(kind, name string, status int) *(dognzb.DogNZB) {
	body, _ := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s/%s.xml", kind, name))
	return dognzb.New("another-api", &mockGetter{
		response: &http.Response{
			StatusCode: status,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		},
	})

}
