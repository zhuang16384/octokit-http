package services

import (
	"time"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type RequestOptions struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type ResponseData struct {
	Status     string            `json:"status"`
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Time       int64             `json:"time"` // in milliseconds
	Size       int64             `json:"size"` // in bytes
}

type RequestService struct {
	client *resty.Client
}

func NewRequestService() *RequestService {
	return &RequestService{
		client: resty.New(),
	}
}

func (s *RequestService) SendRequest(options RequestOptions) (*ResponseData, error) {
	req := s.client.R()

	// Set Headers
	if options.Headers != nil {
		req.SetHeaders(options.Headers)
	}

	// Set Body
	if options.Body != "" {
		req.SetBody(options.Body)
	}

	startTime := time.Now()
	
	// Execute Request
	var resp *resty.Response
	var err error

	switch options.Method {
	case http.MethodGet:
		resp, err = req.Get(options.URL)
	case http.MethodPost:
		resp, err = req.Post(options.URL)
	case http.MethodPut:
		resp, err = req.Put(options.URL)
	case http.MethodDelete:
		resp, err = req.Delete(options.URL)
	case http.MethodPatch:
		resp, err = req.Patch(options.URL)
	case http.MethodHead:
		resp, err = req.Head(options.URL)
	case http.MethodOptions:
		resp, err = req.Options(options.URL)
	default:
		// Default to GET if unknown, or handle error
		resp, err = req.Get(options.URL)
	}

	if err != nil {
		return nil, err
	}

	duration := time.Since(startTime).Milliseconds()

	// Convert Headers
	headers := make(map[string]string)
	for k, v := range resp.Header() {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	return &ResponseData{
		Status:     resp.Status(),
		StatusCode: resp.StatusCode(),
		Headers:    headers,
		Body:       string(resp.Body()),
		Time:       duration,
		Size:       resp.Size(),
	}, nil
}
