package grpcclient

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	geo "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/geo"
	profile "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/profile"
	rate "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/rate"
	recommendation "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/recommendation"
	reservation "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/reservation"
	search "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/search"
	user "github.com/vhive-serverless/vSwarm-proto/proto/hotel_reserv/user"

	log "github.com/sirupsen/logrus"
)

//
// Hotel reservation -----
// 1. Geo service
type HotelGeoClient struct {
	ClientBase
	client geo.GeoClient
}

func (c *HotelGeoClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = geo.NewGeoClient(c.conn)
}

func (c *HotelGeoClient) Request(ctx context.Context, req Input) string {
	// Create a default forward request
	coordinates := strings.Split(req.Value, ",")

	lat, _ := strconv.ParseFloat(coordinates[0], 32)
	lon, _ := strconv.ParseFloat(coordinates[1], 32)

	fw_req := geo.Request{
		Lon: float32(lon),
		Lat: float32(lat),
	}

	fw_res, err := c.client.Nearby(ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Geo service: %v", err)
	}

	msg := fmt.Sprintf("req: { Lat:%f Lon:%f} resp: %+v", fw_req.Lon, fw_req.Lat, fw_res)
	// log.Println(msg)
	return msg
}

type HotelGeoGenerator struct {
	GeneratorBase
}

var hotelCoordinates = []string{
	"37.7867,-122.4112", "37.795,-122.5", "37.9,-122.3", "37.783,-122.41",
	"37.7,-129.4015", "37.7854,-122.4005", "37.7854,-122.4071",
	"37.7936,-122.3930",
}

func (g *HotelGeoGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = hotelCoordinates[0]
	case Linear:
		g.count += 1
		pkt.Value = hotelCoordinates[g.count%len(hotelCoordinates)]
	case Random:
		pkt.Value = hotelCoordinates[rand.Intn(len(hotelCoordinates))]
	}
	return pkt
}

func (c *HotelGeoClient) GetGenerator() Generator {
	return new(HotelGeoGenerator)
}

// 2. Profile service
type HotelProfileClient struct {
	ClientBase
	client profile.ProfileClient
}

func (c *HotelProfileClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = profile.NewProfileClient(c.conn)
}

func (c *HotelProfileClient) Request(ctx context.Context, req Input) string {
	payload := req.Value
	ids := strings.Split(payload, ",")
	// Create a forward request
	fw_req := profile.Request{
		HotelIds: ids,
		Locale:   "",
	}

	fw_res, err := c.client.GetProfiles(ctx, &fw_req)
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

func randIds(n int) string {
	s := "1"
	for i := 0; i < n; i++ {
		s += fmt.Sprintf(",%d", rand.Intn(50))
	}
	return s
}

func (g *HotelProfileGenerator) Next() Input {
	// For profile we generate hotel ids between 1 and 80
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "1,2"
	case Linear:
		g.count += 1
		pkt.Value = fmt.Sprintf("%d", g.count%80)
	case Random:
		n := rand.Intn(5)
		pkt.Value = randIds(n)
	}
	return pkt
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

func (c *HotelRateClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = rate.NewRateClient(c.conn)
}

func (c *HotelRateClient) Request(ctx context.Context, req Input) string {
	payload := req.Value
	ids := strings.Split(payload, ",")
	// Create a forward request
	fw_req := rate.Request{
		HotelIds: ids,
		InDate:   "2015-04-09",
		OutDate:  "2015-04-11",
	}
	fw_res, err := c.client.GetRates(ctx, &fw_req)
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
	// For profile we generate hotel ids between 1 and 80
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = "1,2"
	case Linear:
		g.count += 1
		pkt.Value = fmt.Sprintf("%d", g.count%80)
	case Random:
		n := rand.Intn(5)
		pkt.Value = randIds(n)
	}
	return pkt
}

func (c *HotelRateClient) GetGenerator() Generator {
	return new(HotelRateGenerator)
}

// 4. Recommendation service
type HotelRecommendationClient struct {
	ClientBase
	client recommendation.RecommendationClient
}

func (c *HotelRecommendationClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = recommendation.NewRecommendationClient(c.conn)
}

func (c *HotelRecommendationClient) Request(ctx context.Context, req Input) string {

	// Create a default forward request
	coordinates := strings.Split(req.Value, ",")

	lat, _ := strconv.ParseFloat(coordinates[0], 32)
	lon, _ := strconv.ParseFloat(coordinates[1], 32)

	payload := req.Value
	// Create a forward request
	fw_req := recommendation.Request{
		Require: "dis",
		Lon:     float64(lon),
		Lat:     float64(lat),
	}

	// If one of the require parameters is given as name we will use it
	if payload == "dis" || payload == "rate" || payload == "price" {
		fw_req.Require = req.Value
	}

	fw_res, err := c.client.GetRecommendations(ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke Recommendation service: %v", err)
	}

	msg := fmt.Sprintf("req: {Require: \"dis\", Lat:%f Lon:%f}, resp: %+v", fw_req.Lon, fw_req.Lat, fw_res)
	// log.Println(msg)
	return msg
}

type HotelRecommendationGenerator struct {
	GeneratorBase
}

func (g *HotelRecommendationGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = hotelCoordinates[0]
	case Linear:
		g.count += 1
		pkt.Value = hotelCoordinates[g.count%len(hotelCoordinates)]
	case Random:
		pkt.Value = hotelCoordinates[rand.Intn(len(hotelCoordinates))]
	}
	return pkt
}

func (c *HotelRecommendationClient) GetGenerator() Generator {
	return new(HotelRecommendationGenerator)
}

// 5. Reservation service
type HotelReservationClient struct {
	ClientBase
	client reservation.ReservationClient
}

func (c *HotelReservationClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = reservation.NewReservationClient(c.conn)
}

func (c *HotelReservationClient) Request(ctx context.Context, req Input) string {
	fw_method := req.Method
	payload := strings.Split(req.Value, ",")
	// Create a default forward request
	fw_req := reservation.Request{
		CustomerName: payload[0],
		HotelId:      []string{payload[1]},
		InDate:       "2015-04-09",
		OutDate:      "2015-04-11",
		RoomNumber:   1,
	}

	// Pass on to the real service function
	var fw_res *reservation.Result
	var err error

	switch fw_method {
	case "CheckAvailability", "0":
		// fw_req.HotelId[0] = "2"
		fw_res, err = c.client.CheckAvailability(ctx, &fw_req)

	case "MakeReservation", "1":
		fw_req.HotelId[0] = "3"
		fw_res, err = c.client.MakeReservation(ctx, &fw_req)

	default:
		log.Fatalf("Failed to understand requested method: %s", fw_method)
	}
	if err != nil {
		log.Fatalf("Fail to invoke Reservation service: %v", err)
	}

	msg := fmt.Sprintf("method: %s, req: {CustomerName: %s, HotelId: %v, InDate: \"2015-04-09\", OutDate: \"2015-04-11\", RoomNumber: 1,} resp: %+v", fw_method, fw_req.CustomerName, fw_req.HotelId, fw_res)
	// log.Println(msg)
	return msg
}

type HotelReservationGenerator struct {
	GeneratorBase
}

var hotelReservations = []string{
	"hello,2", "user_100,3", "user_100,10", "user_430,33",
	"adam,33", "drno,3", "user_ssd,10", "user_43,22",
	"peter,1", "test,3", "user_d,13", "user,23",
	"tom,4", "jack,3", "user_100,18", "hi,13",
}

func (g *HotelReservationGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = hotelReservations[0]
	case Linear:
		g.count += 1
		pkt.Value = hotelReservations[g.count%len(hotelReservations)]
	case Random:
		pkt.Value = hotelReservations[rand.Intn(len(hotelReservations))]
	}
	return pkt
}

func (c *HotelReservationClient) GetGenerator() Generator {
	return new(HotelReservationGenerator)
}

// 6. User service
type HotelUserClient struct {
	ClientBase
	client user.UserClient
}

func (c *HotelUserClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = user.NewUserClient(c.conn)
}

func (c *HotelUserClient) Request(ctx context.Context, req Input) string {
	payload := strings.Split(req.Value, ",")

	// Create a forward request
	fw_req := user.Request{
		Username: payload[0],
		Password: payload[1],
	}

	fw_res, err := c.client.CheckUser(ctx, &fw_req)
	if err != nil {
		log.Fatalf("Fail to invoke User service: %v", err)
	}

	msg := fmt.Sprintf("req: {Username: %s, Password: %s} resp: Correct:%v", fw_req.Username, fw_req.Password, fw_res.Correct)
	// log.Println(msg)
	return msg
}

type HotelUserGenerator struct {
	GeneratorBase
}

var hotelUsers = []string{
	"hello,hello", "user_100,pass_100", "user_100,pass_100", "user_430,pass_430",
	"hello,hello2", "user_ads,pass_asdf", "user_23,pass_100", "user_230,pass_5",
	"user_132,pass_132", "user_22,pass_22", "user_43,pass_43",
	"hello,hello2", "user_ads,pass_asdf", "user_23,pass_111", "user_233,pass_5",
}

func (g *HotelUserGenerator) Next() Input {
	var pkt = g.defaultInput
	switch g.GeneratorBase.generator {
	case Unique:
		pkt.Value = hotelUsers[0]
	case Linear:
		g.count += 1
		pkt.Value = hotelUsers[g.count%len(hotelUsers)]
	case Random:
		pkt.Value = hotelUsers[rand.Intn(len(hotelUsers))]
	}
	return pkt
}

func (c *HotelUserClient) GetGenerator() Generator {
	return new(HotelUserGenerator)
}

// 7. Search service
type HotelSearchClient struct {
	ClientBase
	client search.SearchClient
}

func (c *HotelSearchClient) Init(ctx context.Context, ip, port string) {
	c.Connect(ctx, ip, port)
	c.client = search.NewSearchClient(c.conn)
}

func (c *HotelSearchClient) Request(ctx context.Context, req Input) string {
	// Create a forward request
	fw_req := search.NearbyRequest{
		Lat:     37.7963,
		Lon:     -122.4015,
		InDate:  "2015-04-09",
		OutDate: "2015-04-11",
	}

	fw_res, err := c.client.Nearby(ctx, &fw_req)
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
