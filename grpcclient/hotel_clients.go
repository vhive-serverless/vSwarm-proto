package grpcclient

import (
	"fmt"
	"strings"

	geo "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/geo"
	profile "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/profile"
	rate "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/rate"
	recommendation "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/recommendation"
	reservation "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/reservation"
	search "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/search"
	user "github.com/ease-lab/vSwarm-proto/proto/hotel_reserv/user"

	log "github.com/sirupsen/logrus"
)

//
// Hotel reservation -----
// 1. Geo service
type HotelGeoClient struct {
	ClientBase
	client geo.GeoClient
}

func (c *HotelGeoClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = geo.NewGeoClient(c.conn)
}

func (c *HotelGeoClient) Request(req Input) string {
	// Create a default forward request
	fw_req := geo.Request{
		Lat: 37.7963,
		Lon: -122.4015,
	}

	fw_res, err := c.client.Nearby(c.ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Geo service: %v", err)
	}

	msg := fmt.Sprintf("req: {Lat: 37.7963, Lon: -122.4015,} resp: %+v", fw_res)
	// log.Println(msg)
	return msg
}

type HotelGeoGenerator struct {
	GeneratorBase
}

func (g *HotelGeoGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelGeoClient) GetGenerator() Generator {
	return new(HotelGeoGenerator)
}

// 2. Profile service
type HotelProfileClient struct {
	ClientBase
	client profile.ProfileClient
}

func (c *HotelProfileClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = profile.NewProfileClient(c.conn)
}

func (c *HotelProfileClient) Request(req Input) string {
	payload := req.Value
	ids := strings.Split(payload, ",")
	// Create a forward request
	fw_req := profile.Request{
		HotelIds: ids,
		Locale:   "",
	}

	fw_res, err := c.client.GetProfiles(c.ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Profile service: %v", err)
	}

	msg := fmt.Sprintf("req: {HotelIds: %+v, Locale: \"\"} resp: %+v", ids, fw_res)
	// log.Println(msg)
	return msg
}

type HotelProfileGenerator struct {
	GeneratorBase
}

func (g *HotelProfileGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelProfileClient) GetGenerator() Generator {
	return new(HotelProfileGenerator)
}

// ------------------------------------------------------------
// 3. Rate service
type HotelRateClient struct {
	ClientBase
	client rate.RateClient
}

func (c *HotelRateClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = rate.NewRateClient(c.conn)
}

func (c *HotelRateClient) Request(req Input) string {
	payload := req.Value
	ids := strings.Split(payload, ",")
	// Create a forward request
	fw_req := rate.Request{
		HotelIds: ids,
		InDate:   "2015-04-09",
		OutDate:  "2015-04-11",
	}
	fw_res, err := c.client.GetRates(c.ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Rate service: %v", err)
	}

	msg := fmt.Sprintf("req: {HotelIds: %+v, InDate: \"2015-04-09\", OutDate: \"2015-04-11\"}, resp: %+v", ids, fw_res)
	// log.Println(msg)
	return msg
}

type HotelRateGenerator struct {
	GeneratorBase
}

func (g *HotelRateGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelRateClient) GetGenerator() Generator {
	return new(HotelRateGenerator)
}

// 4. Recommendation service
type HotelRecommendationClient struct {
	ClientBase
	client recommendation.RecommendationClient
}

func (c *HotelRecommendationClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = recommendation.NewRecommendationClient(c.conn)
}

func (c *HotelRecommendationClient) Request(req Input) string {
	payload := req.Value
	// Create a forward request
	fw_req := recommendation.Request{
		Require: "dis",
		Lat:     37.7834,
		Lon:     -122.4081,
	}

	// If one of the require parameters is given as name we will use it
	if payload == "dis" || payload == "rate" || payload == "price" {
		fw_req.Require = req.Value
	}

	fw_res, err := c.client.GetRecommendations(c.ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Recommendation service: %v", err)
	}

	msg := fmt.Sprintf("req: {Require: \"dis\", Lat: 37.7834, Lon: -122.4081}, resp: %+v", fw_res)
	// log.Println(msg)
	return msg
}

type HotelRecommendationGenerator struct {
	GeneratorBase
}

func (g *HotelRecommendationGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelRecommendationClient) GetGenerator() Generator {
	return new(HotelRecommendationGenerator)
}

// 5. Reservation service
type HotelReservationClient struct {
	ClientBase
	client reservation.ReservationClient
}

func (c *HotelReservationClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = reservation.NewReservationClient(c.conn)
}

func (c *HotelReservationClient) Request(req Input) string {
	fw_method, payload := req.Method, req.Value
	// Create a default forward request
	fw_req := reservation.Request{
		CustomerName: payload,
		HotelId:      []string{"2"},
		InDate:       "2015-04-09",
		OutDate:      "2015-04-11",
		RoomNumber:   1,
	}

	// Pass on to the real service function
	var fw_res *reservation.Result
	var err error

	switch fw_method {
	case "CheckAvailability", "0":
		fw_req.HotelId[0] = "2"
		fw_res, err = c.client.CheckAvailability(c.ctx, &fw_req)

	case "MakeReservation", "1":
		fw_req.HotelId[0] = "3"
		fw_res, err = c.client.MakeReservation(c.ctx, &fw_req)

	default:
		log.Fatalf("Failed to understand requested method: %s", fw_method)
	}
	if err != nil {
		log.Fatalf("Fail to invoke Reservation service: %v", err)
	}

	msg := fmt.Sprintf("method: %s, req: {CustomerName: %v, HotelId: []string{\"2\"}, InDate: \"2015-04-09\", OutDate: \"2015-04-11\", RoomNumber: 1,} resp: %+v", fw_method, payload, fw_res)
	// log.Println(msg)
	return msg
}

type HotelReservationGenerator struct {
	GeneratorBase
}

func (g *HotelReservationGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelReservationClient) GetGenerator() Generator {
	return new(HotelReservationGenerator)
}

// 6. User service
type HotelUserClient struct {
	ClientBase
	client user.UserClient
}

func (c *HotelUserClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = user.NewUserClient(c.conn)
}

func (c *HotelUserClient) Request(req Input) string {
	payload := req.Value
	// Create a forward request
	fw_req := user.Request{
		Username: payload,
		Password: payload,
	}

	fw_res, err := c.client.CheckUser(c.ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke User service: %v", err)
	}

	msg := fmt.Sprintf("req: {Username: %s, Password: %s} resp: %+v", payload, payload, fw_res)
	// log.Println(msg)
	return msg
}

type HotelUserGenerator struct {
	GeneratorBase
}

func (g *HotelUserGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelUserClient) GetGenerator() Generator {
	return new(HotelUserGenerator)
}

// 7. Search service
type HotelSearchClient struct {
	ClientBase
	client search.SearchClient
}

func (c *HotelSearchClient) Init(ip, port string) {
	c.Connect(ip, port)
	c.client = search.NewSearchClient(c.conn)
}

func (c *HotelSearchClient) Request(req Input) string {
	// Create a forward request
	fw_req := search.NearbyRequest{
		Lat:     37.7963,
		Lon:     -122.4015,
		InDate:  "2015-04-09",
		OutDate: "2015-04-11",
	}

	fw_res, err := c.client.Nearby(c.ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Search service: %v", err)
	}

	msg := fmt.Sprintf("req: {Lat: 37.7963, Lon: -122.4015, InDate: \"2015-04-09\", OutDate: \"2015-04-11\"}, resp: %+v", fw_res)
	// log.Println(msg)
	return msg
}

type HotelSearchGenerator struct {
	GeneratorBase
}

func (g *HotelSearchGenerator) Next() Input {
	return g.defaultInput
}

func (c *HotelSearchClient) GetGenerator() Generator {
	return new(HotelSearchGenerator)
}
