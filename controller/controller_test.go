package controller_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"league/main/controller"
)

func TestEchoHandler(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    string
	}{
		{
			name:        "Valid CSV",
			fileContent: "a,b,c\n1,2,3\n4,5,6\n",
			expected:    "a,b,c\n1,2,3\n4,5,6\n",
		},
		{
			name:        "Empty CSV",
			fileContent: "",
			expected:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/echo", bytes.NewBufferString(tt.fileContent))
			req.Header.Set("Content-Type", "multipart/form-data")

			// Create a form file to simulate file upload
			form := new(bytes.Buffer)
			writer := multipart.NewWriter(form)
			fileWriter, _ := writer.CreateFormFile("file", "test.csv")
			fileWriter.Write([]byte(tt.fileContent))
			writer.Close()
			req.Body = io.NopCloser(form)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			controller.EchoHandler(rr, req)

			// Check the response
			if rr.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", rr.Code)
			}
			if rr.Body.String() != tt.expected {
				t.Errorf("expected body %q; got %q", tt.expected, rr.Body.String())
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name          string
		setupRequest  func() (*http.Request, error)
		expectedError string
	}{
		{
			name: "missing file",
			setupRequest: func() (*http.Request, error) {
				// Create request without file
				return httptest.NewRequest(http.MethodPost, "/echo", nil), nil
			},
			expectedError: "error request Content-Type isn't multipart/form-data",
		},
		{
			name: "invalid CSV format",
			setupRequest: func() (*http.Request, error) {
				// Create multipart form with invalid CSV
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				part, err := writer.CreateFormFile("file", "test.csv")
				if err != nil {
					return nil, err
				}
				// Write invalid CSV with unclosed quote
				part.Write([]byte("a,\"unclosed,quote\nb,c"))
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/echo", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				return req, nil
			},
			expectedError: "error record on line 1",
		},
		{
			name: "wrong form field name",
			setupRequest: func() (*http.Request, error) {
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				part, err := writer.CreateFormFile("wrong_field", "test.csv")
				if err != nil {
					return nil, err
				}
				part.Write([]byte("a,b,c"))
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/echo", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				return req, nil
			},
			expectedError: "error http: no such file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.setupRequest()
			if err != nil {
				t.Fatalf("Failed to setup request: %v", err)
			}

			rec := httptest.NewRecorder()
			controller.EchoHandler(rec, req)

			// Check if response contains expected error message
			if !strings.Contains(rec.Body.String(), tt.expectedError) {
				t.Errorf("Expected error containing %q, got %q", tt.expectedError, rec.Body.String())
			}
		})
	}
}

// Additional test cases for other handlers
func TestInvertHandler(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    string
	}{
		{
			name:        "Valid Matrix",
			fileContent: "1,2,3\n4,5,6\n",
			expected:    "1,4\n2,5\n3,6\n",
		},
		{
			name:        "Invalid Matrix",
			fileContent: "invalid,data\n",
			expected:    "error invalid number at position [0,0]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/invert", bytes.NewBufferString(tt.fileContent))
			req.Header.Set("Content-Type", "multipart/form-data")

			// Create a form file to simulate file upload
			form := new(bytes.Buffer)
			writer := multipart.NewWriter(form)
			fileWriter, _ := writer.CreateFormFile("file", "test.csv")
			fileWriter.Write([]byte(tt.fileContent))
			writer.Close()
			req.Body = io.NopCloser(form)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			controller.InvertHandler(rr, req)

			// Check the response
			if rr.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", rr.Code)
			}
			if rr.Body.String() != tt.expected {
				t.Errorf("expected body %q; got %q", tt.expected, rr.Body.String())
			}
		})
	}
}

func TestFlattenHandler(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    string
	}{
		{
			name:        "Valid Matrix",
			fileContent: "1,2,3\n4,5,6\n",
			expected:    "1,2,3,4,5,6\n", // Expected output depends on the implementation of FlattenMatrix
		},
		{
			name:        "Invalid Matrix",
			fileContent: "invalid,data\n",
			expected:    "error invalid number at position [0,0]", // Adjust based on actual error handling
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/flatten", bytes.NewBufferString(tt.fileContent))
			req.Header.Set("Content-Type", "multipart/form-data")

			// Create a form file to simulate file upload
			form := new(bytes.Buffer)
			writer := multipart.NewWriter(form)
			fileWriter, _ := writer.CreateFormFile("file", "test.csv")
			fileWriter.Write([]byte(tt.fileContent))
			writer.Close()
			req.Body = io.NopCloser(form)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			controller.FlattenHandler(rr, req)

			// Check the response
			if rr.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", rr.Code)
			}
			if rr.Body.String() != tt.expected {
				t.Errorf("expected body %q; got %q", tt.expected, rr.Body.String())
			}
		})
	}
}

func TestSumHandler(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    string
	}{
		{
			name:        "Valid Numbers",
			fileContent: "1,2,3\n4,5,6\n",
			expected:    "21\n", // Assuming Reduce returns the sum
		},
		{
			name:        "Invalid Number",
			fileContent: "1,invalid,3\n",
			expected:    "error invalid number at position [0,1]", // Adjust based on actual error handling
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/sum", bytes.NewBufferString(tt.fileContent))
			req.Header.Set("Content-Type", "multipart/form-data")

			// Create a form file to simulate file upload
			form := new(bytes.Buffer)
			writer := multipart.NewWriter(form)
			fileWriter, _ := writer.CreateFormFile("file", "test.csv")
			fileWriter.Write([]byte(tt.fileContent))
			writer.Close()
			req.Body = io.NopCloser(form)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			controller.SumHandler(rr, req)

			// Check the response
			if rr.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", rr.Code)
			}
			if rr.Body.String() != tt.expected {
				t.Errorf("expected body %q; got %q", tt.expected, rr.Body.String())
			}
		})
	}
}

func TestMultiplyHandler(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    string
	}{
		{
			name:        "Valid Numbers",
			fileContent: "1,2,3\n4,5,6\n",
			expected:    "720\n", // Assuming Reduce returns the product
		},
		{
			name:        "Invalid Number",
			fileContent: "1,invalid,3\n",
			expected:    "error invalid number at position [0,1]", // Adjust based on actual error handling
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/multiply", bytes.NewBufferString(tt.fileContent))
			req.Header.Set("Content-Type", "multipart/form-data")

			// Create a form file to simulate file upload
			form := new(bytes.Buffer)
			writer := multipart.NewWriter(form)
			fileWriter, _ := writer.CreateFormFile("file", "test.csv")
			fileWriter.Write([]byte(tt.fileContent))
			writer.Close()
			req.Body = io.NopCloser(form)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			controller.MultiplyHandler(rr, req)

			// Check the response
			if rr.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", rr.Code)
			}
			if rr.Body.String() != tt.expected {
				t.Errorf("expected body %q; got %q", tt.expected, rr.Body.String())
			}
		})
	}
}
