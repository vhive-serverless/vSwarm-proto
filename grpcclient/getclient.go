// MIT License

// Copyright (c) 2022 EASE lab

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
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
	case "bert-python":
		return "bert"
	case "fibonacci-go", "fibonacci-python", "fibonacci-nodejs", "fibonacci-cpp":
		return "fibonacci"
	default:
		return functionName
	}
}

// Helper to find a grpc client for the given service
func FindGrpcClient(service_name string) GrpcClient {
	switch service_name {
	case "helloworld":
		log.Debug("Found Helloworld client")
		return new(HelloWorldClient)

	case "aes":
		log.Debug("Found AES client")
		return new(AesClient)
	case "auth":
		log.Debug("Found Auth client")
		return new(AuthClient)
	case "bert":
		log.Debug("Found Bert client")
		return new(BertClient)
	case "fibonacci":
		log.Debug("Found Fibonacci client")
		return new(FibonacciClient)

		// Hotel reservation ---
	case "Geo", "geo":
		log.Debug("Found geo client")
		return new(HotelGeoClient)
	case "Profile", "profile":
		log.Debug("Found profile client")
		return new(HotelProfileClient)
	case "Rate", "rate":
		log.Debug("Found rate client")
		return new(HotelRateClient)
	case "Recommendation", "recommendation":
		log.Debug("Found recommendation client")
		return new(HotelRecommendationClient)
	case "Reservation", "reservation":
		log.Debug("Found reservation client")
		return new(HotelReservationClient)
	case "User", "user":
		log.Debug("Found user client")
		return new(HotelUserClient)
	case "Search", "search":
		log.Debug("Found search client")
		return new(HotelSearchClient)

	// Hipster shop ---
	case "Ad", "adservice":
		log.Debug("Found Ad client for online shop")
		return new(ShopAdServiceClient)
	case "Cart", "cartservice":
		log.Debug("Found Cart client for online shop")
		return new(ShopCartServiceClient)
	case "Checkout", "checkoutservice":
		log.Debug("Found Checkout client for online shop")
		return new(ShopCheckoutServiceClient)
	case "Currency", "currencyservice":
		log.Debug("Found Currency client for online shop")
		return new(ShopCurrencyServiceClient)
	case "Email", "emailservice":
		log.Debug("Found Email client for online shop")
		return new(ShopEmailServiceClient)
	case "Payment", "paymentservice":
		log.Debug("Found Payment client for online shop")
		return new(ShopPaymentServiceClient)
	case "ProductCatalog", "productcatalogservice":
		log.Debug("Found ProductCatalog client for online shop")
		return new(ShopProductCatalogServiceClient)
	case "recommendationservice":
		log.Debug("Found Recommendation client for online shop")
		return new(ShopRecommendationServiceClient)
	case "Shipping", "shippingservice":
		log.Debug("Found Shipping client for online shop")
		return new(ShopShippingServiceClient)

	// Default ---------
	default:
		log.Warnf("Did not find a matching client for %s... Will use the default Hello world client. \n", service_name)
		return new(HelloWorldClient)
	}
}
