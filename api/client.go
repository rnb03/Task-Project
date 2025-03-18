package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://jsonplaceholder.typicode.com"
)

// Client is a JSONPlaceholder API client
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new JSONPlaceholder API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// makeRequest makes an HTTP request to the JSONPlaceholder API
func (c *Client) makeRequest(method, path string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.httpClient.Do(req)
}

// GetPosts retrieves all posts
func (c *Client) GetPosts() ([]Post, error) {
	resp, err := c.makeRequest(http.MethodGet, "/posts", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPost retrieves a specific post by ID
func (c *Client) GetPost(id int) (*Post, error) {
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/posts/%d", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}
	return &post, nil
}

// GetComments retrieves all comments
func (c *Client) GetComments() ([]Comment, error) {
	resp, err := c.makeRequest(http.MethodGet, "/comments", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// GetCommentsForPost retrieves comments for a specific post
func (c *Client) GetCommentsForPost(postID int) ([]Comment, error) {
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/posts/%d/comments", postID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// GetUsers retrieves all users
func (c *Client) GetUsers() ([]User, error) {
	resp, err := c.makeRequest(http.MethodGet, "/users", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser retrieves a specific user by ID
func (c *Client) GetUser(id int) (*User, error) {
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/users/%d", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAlbums retrieves all albums
func (c *Client) GetAlbums() ([]Album, error) {
	resp, err := c.makeRequest(http.MethodGet, "/albums", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var albums []Album
	if err := json.NewDecoder(resp.Body).Decode(&albums); err != nil {
		return nil, err
	}
	return albums, nil
}

// GetPhotos retrieves all photos
func (c *Client) GetPhotos() ([]Photo, error) {
	resp, err := c.makeRequest(http.MethodGet, "/photos", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var photos []Photo
	if err := json.NewDecoder(resp.Body).Decode(&photos); err != nil {
		return nil, err
	}
	return photos, nil
}

// GetTodos retrieves all todos
func (c *Client) GetTodos() ([]Todo, error) {
	resp, err := c.makeRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}
