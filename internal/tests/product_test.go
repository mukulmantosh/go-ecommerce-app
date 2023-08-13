package tests

//import (
//	"bytes"
//	"fmt"
//	"github.com/stretchr/testify/suite"
//	"io"
//	"net/http"
//	"testing"
//)
//
//type EndToEndSuite struct {
//	suite.Suite
//}
//
//func TestEndToEndSuite(t *testing.T) {
//	suite.Run(t, new(EndToEndSuite))
//}
//
//func (s *EndToEndSuite) TestAddNewProduct() {
//	c := http.Client{}
//	jsonBody := []byte(`{"client_message": "hello, server!"}`)
//	bodyReader := bytes.NewReader(jsonBody)
//	r, _ := c.Post("http://localhost:8080/products", "application/json", bodyReader)
//	defer r.Body.Close()
//	_, err := io.ReadAll(r.Body)
//	if err != nil {
//		fmt.Printf("server: could not read request body: %s\n", err)
//	}
//	s.Equal(http.StatusCreated, r.StatusCode)
//
//}
