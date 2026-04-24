package grpcclient

import (
	log "github.com/sirupsen/logrus"
)

func FindServiceName(functionName string) string {
	switch functionName {
	case "aes-go", "aes-python", "aes-nodejs":
		return "aes"
	case "auth-go", "auth-python", "auth-nodejs":
		return "auth"
	case "fibonacci-go", "fibonacci-python", "fibonacci-nodejs", "fibonacci-cpp":
		return "fibonacci"
	case "compression-python":
		return "compression"
	case "graph-bfs-python":
		return "graph-bfs"
	case "gptj-python":
		return "gptj"
	case "image-rotate-python", "image-rotate-nodejs", "image-rotate-go", "image-rotate-cython":
		return "image-rotate"
	case "video-processing-python":
		return "video-processing"
	case "rnn-serving-python":
		return "rnn-serving"
	case "video-analytics-standalone-python":
		return "video-analytics-standalone"
	case "spright-parking-python":
		return "spright-parking"
	default:
		return functionName
	}
}

// Helper to find a grpc client for the given service
func FindGrpcClient(service_name string) GrpcClient {
	switch service_name {
	case "helloworld":
		log.Trace("Found Helloworld client")
		return new(HelloWorldClient)

	case "aes":
		log.Trace("Found AES client")
		return new(AesClient)
	case "auth":
		log.Trace("Found Auth client")
		return new(AuthClient)
	case "fibonacci":
		log.Trace("Found Fibonacci client")
		return new(FibonacciClient)
	case "compression":
		log.Trace("Found Compression client")
		return new(FileCompressClient)
	case "graph-bfs":
		log.Trace("Found Graph BFS client")
		return new(GraphBFSClient)
	case "gptj":
		log.Trace("Found gptj client")
		return new(GptJClient)
	case "image-rotate":
		log.Trace("Found image rotate client")
		return new(ImageRotateClient)
	case "video-processing":
		log.Trace("Found video processing client")
		return new(VideoProcessingClient)
	case "rnn-serving":
		log.Trace("Found rnn serving client")
		return new(RNNServingClient)
	case "video-analytics-standalone":
		log.Trace("Found video analytics standalone client")
		return new(VideoAnalyticsClient)

		// Hotel reservation ---
	case "Geo", "geo":
		log.Trace("Found geo client")
		return new(HotelGeoClient)
	case "Profile", "profile":
		log.Trace("Found profile client")
		return new(HotelProfileClient)
	case "Rate", "rate":
		log.Trace("Found rate client")
		return new(HotelRateClient)
	case "Recommendation", "recommendation":
		log.Trace("Found recommendation client")
		return new(HotelRecommendationClient)
	case "Reservation", "reservation":
		log.Trace("Found reservation client")
		return new(HotelReservationClient)
	case "User", "user":
		log.Trace("Found user client")
		return new(HotelUserClient)
	case "Search", "search":
		log.Trace("Found search client")
		return new(HotelSearchClient)

	// Hipster shop ---
	case "Ad", "adservice":
		log.Trace("Found Ad client for online shop")
		return new(ShopAdServiceClient)
	case "Cart", "cartservice":
		log.Trace("Found Cart client for online shop")
		return new(ShopCartServiceClient)
	case "Checkout", "checkoutservice":
		log.Trace("Found Checkout client for online shop")
		return new(ShopCheckoutServiceClient)
	case "Currency", "currencyservice":
		log.Trace("Found Currency client for online shop")
		return new(ShopCurrencyServiceClient)
	case "Email", "emailservice":
		log.Trace("Found Email client for online shop")
		return new(ShopEmailServiceClient)
	case "Payment", "paymentservice":
		log.Trace("Found Payment client for online shop")
		return new(ShopPaymentServiceClient)
	case "ProductCatalog", "productcatalogservice":
		log.Trace("Found ProductCatalog client for online shop")
		return new(ShopProductCatalogServiceClient)
	case "recommendationservice":
		log.Trace("Found Recommendation client for online shop")
		return new(ShopRecommendationServiceClient)
	case "Shipping", "shippingservice":
		log.Trace("Found Shipping client for online shop")
		return new(ShopShippingServiceClient)

	// Spright parking ---
	case "spright-parking":
		log.Trace("Found Spright client for spright parking")
		return new(ParkingClient)

	// Default ---------
	default:
		log.Warnf("Did not find a matching client for %s... Will use the default Hello world client. \n", service_name)
		return new(HelloWorldClient)
	}
}
