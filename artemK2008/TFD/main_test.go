package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	inputData int
	Answer    int
	Expected  int
}

var Cases []TestCase = []TestCase{
	{
		inputData: 0,
		Expected:  1,
	},
	{
		inputData: 1,
		Expected:  1,
	},
	{
		inputData: 3,
		Expected:  6,
	},
	{
		inputData: 5,
		Expected:  120,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if test.Answer = factorial(test.inputData); test.Answer != test.Expected {
			t.Errorf("test case %d failed, result %v, expected %v", id, test.Answer, test.Expected)
		}
	}
}

type HttpTestCase struct {
	Name     string
	Numeric  int
	Expected []byte
}

var HttpCases = []HttpTestCase{
	{
		Name:     "first test",
		Numeric:  1,
		Expected: []byte("1"),
	},
	{
		Name:     "second test",
		Numeric:  3,
		Expected: []byte("6"),
	},
	{
		Name:     "third test",
		Numeric:  6,
		Expected: []byte("720"),
	},
}

func TestHandleFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandleFactorial)
	for _, test := range HttpCases {
		t.Run(test.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder() // куда писать ответ
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, _ := http.NewRequest("GET", handlerData, nil) // какой будет запрос
			/*
				data := io.Reader(".................")
				request,err := http.Post("http://localhost:8080/factorial?num=1", "application/json",data)
			*/
			handler.ServeHTTP(recorder, request) // выполняем запрос. Ответ записываем в рекордер
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("Test %s failed: input: %v, result %v, expected %v",
					test.Name,
					test.Numeric,
					string(recorder.Body.Bytes()),
					string(test.Expected))
			}
		}) // Под-тестовый ранер
	}
}
