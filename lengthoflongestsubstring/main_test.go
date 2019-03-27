package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLengthOfLongestSubString(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcabcd", 4},
		{"我是是是李正哈哈", 4},
	}

	for _, tt := range tests {
		if actual := LengthOfLongestSubString(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s, expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkLengthOfLongestSubString(b *testing.B) {
	s := "我是是是李正哈哈"
	ans := 4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := LengthOfLongestSubString(s); actual != ans {
			b.Errorf("got %d for input %s, expected %d", actual, s, ans)
		}
	}
}

func onError(resp http.ResponseWriter, req *http.Request) error {
	return nil
	//return errors.New("error la")
}

func errWrapper(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func TestHttp(t *testing.T) {
	f := errWrapper(onError)
	server := httptest.NewServer(f)
	fmt.Println(server.URL)
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	body := strings.Trim(string(bytes), "\n")
	if resp.StatusCode != 200 {
		t.Errorf("出错啦: %s", body)
	}
}
