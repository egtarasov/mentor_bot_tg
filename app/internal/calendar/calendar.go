package calendar

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"telegrambot_new_emploee/internal/models"
	"time"
)

const (
	userIdParam = "id"
)

type Calendar interface {
	GetMeetingsById(user *models.User) ([]models.Meeting, error)
}

type calendar struct {
	url     string
	timeout time.Duration
}

func NewCalendar(path string, timeout time.Duration) Calendar {
	return &calendar{
		url: path,
	}
}

func (c *calendar) GetMeetingsById(user *models.User) ([]models.Meeting, error) {
	// Initiate a http client.
	client := http.Client{
		Timeout: c.timeout,
	}

	// Create a request.
	req, err := c.getMeetingsByIdReq(user)
	if err != nil {
		return nil, err
	}

	// Get data.
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("calendar is not responding")
	}

	return c.getMeetingsById(resp)
}

func (c *calendar) getMeetingsByIdReq(user *models.User) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add(userIdParam, strconv.FormatInt(user.TelegramId, 10))
	req.URL.RawQuery = q.Encode()

	log.Printf("Request: %s", req.URL.String())

	return req, nil
}

func (c *calendar) getMeetingsById(resp *http.Response) ([]models.Meeting, error) {
	var meetings []models.Meeting
	defer resp.Body.Close()

	d := json.NewDecoder(resp.Body)
	err := d.Decode(&meetings)
	if err != nil {
		return nil, err
	}

	return meetings, nil
}
