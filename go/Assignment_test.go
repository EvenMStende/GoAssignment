package main

import "testing"
import "net/http"
import "net/http/httptest"


func Test_handler(t *testing.T) {
  test := httptest.NewServer(http.HandlerFunc(handler))
  defer test.Close()

  r, err := http.Get(test.URL + "/project/v1/github.com/apache/kafka")

    if err != nil {
      t.Error("Didnt get the GET", err)
    }

    if r.StatusCode != http.StatusOK {
      t.Errorf("Expected internal server error got %d", r.StatusCode)
    }
}
