package main

import (
	"SFA/week11/Lecture30/grpc/pb"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"google.golang.org/grpc"
)

type StoryService struct {
	urlBase string
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
	Id    int    `json:"id"`
}

const port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, new(employeeService))
	log.Println("Starting server on port", port)
	s.Serve(lis)
}

func NewStoryService(url string) *StoryService {
	return &StoryService{urlBase: url}
}

type employeeService struct {
	pb.UnimplementedEmployeeServiceServer
}

func (s *employeeService) FetchTopStories(ctx context.Context,
	req *pb.TopCountRequest) (*pb.TopCountResponse, error) {

	ss := NewStoryService("https://hacker-news.firebaseio.com")

	res, err := http.Get(ss.urlBase + "/v0/topstories.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	temp := &pb.TopCountResponse{}
	err = json.Unmarshal(body, &temp.Arr)
	if err != nil {
		return nil, err
	}
	temp.Arr = temp.Arr[:req.MaxCount]
	return temp, nil
}

func (s *employeeService) FetchItems(ctx context.Context,
	resp *pb.TopCountResponse) (*pb.ResponseStories, error) {
	slice := make([]string, 0, 10)
	for _, e := range resp.Arr {
		slice = append(slice, strconv.Itoa(int(e)))
	}

	u, err := url.Parse("https://hacker-news.firebaseio.com")
	if err != nil {
		return nil, err
	}

	stories := &pb.ResponseStories{}

	for _, detail := range slice {
		u.Path = "/v0/item/" + detail + ".json"

		response, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		story := &pb.Story{}
		// var story Story
		err = json.Unmarshal(body, &story)
		if err != nil {
			return nil, err
		}

		stories.TopStories = append(stories.TopStories, story)
	}

	return stories, nil
}

func (s *employeeService) mustEmbedUnimplementedEmployeeServiceServer() {}
