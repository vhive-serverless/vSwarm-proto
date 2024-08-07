# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/hipstershop/demo.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1cproto/hipstershop/demo.proto\x12\x0bhipstershop\"0\n\x08\x43\x61rtItem\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\"F\n\x0e\x41\x64\x64ItemRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12#\n\x04item\x18\x02 \x01(\x0b\x32\x15.hipstershop.CartItem\"#\n\x10\x45mptyCartRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"!\n\x0eGetCartRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"=\n\x04\x43\x61rt\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12$\n\x05items\x18\x02 \x03(\x0b\x32\x15.hipstershop.CartItem\"\x07\n\x05\x45mpty\"B\n\x1aListRecommendationsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x13\n\x0bproduct_ids\x18\x02 \x03(\t\"2\n\x1bListRecommendationsResponse\x12\x13\n\x0bproduct_ids\x18\x01 \x03(\t\"\x84\x01\n\x07Product\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x0f\n\x07picture\x18\x04 \x01(\t\x12%\n\tprice_usd\x18\x05 \x01(\x0b\x32\x12.hipstershop.Money\x12\x12\n\ncategories\x18\x06 \x03(\t\">\n\x14ListProductsResponse\x12&\n\x08products\x18\x01 \x03(\x0b\x32\x14.hipstershop.Product\"\x1f\n\x11GetProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\"&\n\x15SearchProductsRequest\x12\r\n\x05query\x18\x01 \x01(\t\"?\n\x16SearchProductsResponse\x12%\n\x07results\x18\x01 \x03(\x0b\x32\x14.hipstershop.Product\"^\n\x0fGetQuoteRequest\x12%\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x14.hipstershop.Address\x12$\n\x05items\x18\x02 \x03(\x0b\x32\x15.hipstershop.CartItem\"8\n\x10GetQuoteResponse\x12$\n\x08\x63ost_usd\x18\x01 \x01(\x0b\x32\x12.hipstershop.Money\"_\n\x10ShipOrderRequest\x12%\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x14.hipstershop.Address\x12$\n\x05items\x18\x02 \x03(\x0b\x32\x15.hipstershop.CartItem\"(\n\x11ShipOrderResponse\x12\x13\n\x0btracking_id\x18\x01 \x01(\t\"a\n\x07\x41\x64\x64ress\x12\x16\n\x0estreet_address\x18\x01 \x01(\t\x12\x0c\n\x04\x63ity\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x01(\t\x12\x0f\n\x07\x63ountry\x18\x04 \x01(\t\x12\x10\n\x08zip_code\x18\x05 \x01(\x05\"<\n\x05Money\x12\x15\n\rcurrency_code\x18\x01 \x01(\t\x12\r\n\x05units\x18\x02 \x01(\x03\x12\r\n\x05nanos\x18\x03 \x01(\x05\"8\n\x1eGetSupportedCurrenciesResponse\x12\x16\n\x0e\x63urrency_codes\x18\x01 \x03(\t\"N\n\x19\x43urrencyConversionRequest\x12 \n\x04\x66rom\x18\x01 \x01(\x0b\x32\x12.hipstershop.Money\x12\x0f\n\x07to_code\x18\x02 \x01(\t\"\x90\x01\n\x0e\x43reditCardInfo\x12\x1a\n\x12\x63redit_card_number\x18\x01 \x01(\t\x12\x17\n\x0f\x63redit_card_cvv\x18\x02 \x01(\x05\x12#\n\x1b\x63redit_card_expiration_year\x18\x03 \x01(\x05\x12$\n\x1c\x63redit_card_expiration_month\x18\x04 \x01(\x05\"e\n\rChargeRequest\x12\"\n\x06\x61mount\x18\x01 \x01(\x0b\x32\x12.hipstershop.Money\x12\x30\n\x0b\x63redit_card\x18\x02 \x01(\x0b\x32\x1b.hipstershop.CreditCardInfo\"(\n\x0e\x43hargeResponse\x12\x16\n\x0etransaction_id\x18\x01 \x01(\t\"R\n\tOrderItem\x12#\n\x04item\x18\x01 \x01(\x0b\x32\x15.hipstershop.CartItem\x12 \n\x04\x63ost\x18\x02 \x01(\x0b\x32\x12.hipstershop.Money\"\xbf\x01\n\x0bOrderResult\x12\x10\n\x08order_id\x18\x01 \x01(\t\x12\x1c\n\x14shipping_tracking_id\x18\x02 \x01(\t\x12)\n\rshipping_cost\x18\x03 \x01(\x0b\x32\x12.hipstershop.Money\x12.\n\x10shipping_address\x18\x04 \x01(\x0b\x32\x14.hipstershop.Address\x12%\n\x05items\x18\x05 \x03(\x0b\x32\x16.hipstershop.OrderItem\"V\n\x1cSendOrderConfirmationRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\'\n\x05order\x18\x02 \x01(\x0b\x32\x18.hipstershop.OrderResult\"\xa3\x01\n\x11PlaceOrderRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x15\n\ruser_currency\x18\x02 \x01(\t\x12%\n\x07\x61\x64\x64ress\x18\x03 \x01(\x0b\x32\x14.hipstershop.Address\x12\r\n\x05\x65mail\x18\x05 \x01(\t\x12\x30\n\x0b\x63redit_card\x18\x06 \x01(\x0b\x32\x1b.hipstershop.CreditCardInfo\"=\n\x12PlaceOrderResponse\x12\'\n\x05order\x18\x01 \x01(\x0b\x32\x18.hipstershop.OrderResult\"!\n\tAdRequest\x12\x14\n\x0c\x63ontext_keys\x18\x01 \x03(\t\"*\n\nAdResponse\x12\x1c\n\x03\x61\x64s\x18\x01 \x03(\x0b\x32\x0f.hipstershop.Ad\"(\n\x02\x41\x64\x12\x14\n\x0credirect_url\x18\x01 \x01(\t\x12\x0c\n\x04text\x18\x02 \x01(\t2\xca\x01\n\x0b\x43\x61rtService\x12<\n\x07\x41\x64\x64Item\x12\x1b.hipstershop.AddItemRequest\x1a\x12.hipstershop.Empty\"\x00\x12;\n\x07GetCart\x12\x1b.hipstershop.GetCartRequest\x1a\x11.hipstershop.Cart\"\x00\x12@\n\tEmptyCart\x12\x1d.hipstershop.EmptyCartRequest\x1a\x12.hipstershop.Empty\"\x00\x32\x83\x01\n\x15RecommendationService\x12j\n\x13ListRecommendations\x12\'.hipstershop.ListRecommendationsRequest\x1a(.hipstershop.ListRecommendationsResponse\"\x00\x32\x83\x02\n\x15ProductCatalogService\x12G\n\x0cListProducts\x12\x12.hipstershop.Empty\x1a!.hipstershop.ListProductsResponse\"\x00\x12\x44\n\nGetProduct\x12\x1e.hipstershop.GetProductRequest\x1a\x14.hipstershop.Product\"\x00\x12[\n\x0eSearchProducts\x12\".hipstershop.SearchProductsRequest\x1a#.hipstershop.SearchProductsResponse\"\x00\x32\xaa\x01\n\x0fShippingService\x12I\n\x08GetQuote\x12\x1c.hipstershop.GetQuoteRequest\x1a\x1d.hipstershop.GetQuoteResponse\"\x00\x12L\n\tShipOrder\x12\x1d.hipstershop.ShipOrderRequest\x1a\x1e.hipstershop.ShipOrderResponse\"\x00\x32\xb7\x01\n\x0f\x43urrencyService\x12[\n\x16GetSupportedCurrencies\x12\x12.hipstershop.Empty\x1a+.hipstershop.GetSupportedCurrenciesResponse\"\x00\x12G\n\x07\x43onvert\x12&.hipstershop.CurrencyConversionRequest\x1a\x12.hipstershop.Money\"\x00\x32U\n\x0ePaymentService\x12\x43\n\x06\x43harge\x12\x1a.hipstershop.ChargeRequest\x1a\x1b.hipstershop.ChargeResponse\"\x00\x32h\n\x0c\x45mailService\x12X\n\x15SendOrderConfirmation\x12).hipstershop.SendOrderConfirmationRequest\x1a\x12.hipstershop.Empty\"\x00\x32\x62\n\x0f\x43heckoutService\x12O\n\nPlaceOrder\x12\x1e.hipstershop.PlaceOrderRequest\x1a\x1f.hipstershop.PlaceOrderResponse\"\x00\x32H\n\tAdService\x12;\n\x06GetAds\x12\x16.hipstershop.AdRequest\x1a\x17.hipstershop.AdResponse\"\x00\x42<Z:github.com/vhive-serverless/vSwarm-proto/proto/hipstershopb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.hipstershop.demo_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z:github.com/vhive-serverless/vSwarm-proto/proto/hipstershop'
  _globals['_CARTITEM']._serialized_start=45
  _globals['_CARTITEM']._serialized_end=93
  _globals['_ADDITEMREQUEST']._serialized_start=95
  _globals['_ADDITEMREQUEST']._serialized_end=165
  _globals['_EMPTYCARTREQUEST']._serialized_start=167
  _globals['_EMPTYCARTREQUEST']._serialized_end=202
  _globals['_GETCARTREQUEST']._serialized_start=204
  _globals['_GETCARTREQUEST']._serialized_end=237
  _globals['_CART']._serialized_start=239
  _globals['_CART']._serialized_end=300
  _globals['_EMPTY']._serialized_start=302
  _globals['_EMPTY']._serialized_end=309
  _globals['_LISTRECOMMENDATIONSREQUEST']._serialized_start=311
  _globals['_LISTRECOMMENDATIONSREQUEST']._serialized_end=377
  _globals['_LISTRECOMMENDATIONSRESPONSE']._serialized_start=379
  _globals['_LISTRECOMMENDATIONSRESPONSE']._serialized_end=429
  _globals['_PRODUCT']._serialized_start=432
  _globals['_PRODUCT']._serialized_end=564
  _globals['_LISTPRODUCTSRESPONSE']._serialized_start=566
  _globals['_LISTPRODUCTSRESPONSE']._serialized_end=628
  _globals['_GETPRODUCTREQUEST']._serialized_start=630
  _globals['_GETPRODUCTREQUEST']._serialized_end=661
  _globals['_SEARCHPRODUCTSREQUEST']._serialized_start=663
  _globals['_SEARCHPRODUCTSREQUEST']._serialized_end=701
  _globals['_SEARCHPRODUCTSRESPONSE']._serialized_start=703
  _globals['_SEARCHPRODUCTSRESPONSE']._serialized_end=766
  _globals['_GETQUOTEREQUEST']._serialized_start=768
  _globals['_GETQUOTEREQUEST']._serialized_end=862
  _globals['_GETQUOTERESPONSE']._serialized_start=864
  _globals['_GETQUOTERESPONSE']._serialized_end=920
  _globals['_SHIPORDERREQUEST']._serialized_start=922
  _globals['_SHIPORDERREQUEST']._serialized_end=1017
  _globals['_SHIPORDERRESPONSE']._serialized_start=1019
  _globals['_SHIPORDERRESPONSE']._serialized_end=1059
  _globals['_ADDRESS']._serialized_start=1061
  _globals['_ADDRESS']._serialized_end=1158
  _globals['_MONEY']._serialized_start=1160
  _globals['_MONEY']._serialized_end=1220
  _globals['_GETSUPPORTEDCURRENCIESRESPONSE']._serialized_start=1222
  _globals['_GETSUPPORTEDCURRENCIESRESPONSE']._serialized_end=1278
  _globals['_CURRENCYCONVERSIONREQUEST']._serialized_start=1280
  _globals['_CURRENCYCONVERSIONREQUEST']._serialized_end=1358
  _globals['_CREDITCARDINFO']._serialized_start=1361
  _globals['_CREDITCARDINFO']._serialized_end=1505
  _globals['_CHARGEREQUEST']._serialized_start=1507
  _globals['_CHARGEREQUEST']._serialized_end=1608
  _globals['_CHARGERESPONSE']._serialized_start=1610
  _globals['_CHARGERESPONSE']._serialized_end=1650
  _globals['_ORDERITEM']._serialized_start=1652
  _globals['_ORDERITEM']._serialized_end=1734
  _globals['_ORDERRESULT']._serialized_start=1737
  _globals['_ORDERRESULT']._serialized_end=1928
  _globals['_SENDORDERCONFIRMATIONREQUEST']._serialized_start=1930
  _globals['_SENDORDERCONFIRMATIONREQUEST']._serialized_end=2016
  _globals['_PLACEORDERREQUEST']._serialized_start=2019
  _globals['_PLACEORDERREQUEST']._serialized_end=2182
  _globals['_PLACEORDERRESPONSE']._serialized_start=2184
  _globals['_PLACEORDERRESPONSE']._serialized_end=2245
  _globals['_ADREQUEST']._serialized_start=2247
  _globals['_ADREQUEST']._serialized_end=2280
  _globals['_ADRESPONSE']._serialized_start=2282
  _globals['_ADRESPONSE']._serialized_end=2324
  _globals['_AD']._serialized_start=2326
  _globals['_AD']._serialized_end=2366
  _globals['_CARTSERVICE']._serialized_start=2369
  _globals['_CARTSERVICE']._serialized_end=2571
  _globals['_RECOMMENDATIONSERVICE']._serialized_start=2574
  _globals['_RECOMMENDATIONSERVICE']._serialized_end=2705
  _globals['_PRODUCTCATALOGSERVICE']._serialized_start=2708
  _globals['_PRODUCTCATALOGSERVICE']._serialized_end=2967
  _globals['_SHIPPINGSERVICE']._serialized_start=2970
  _globals['_SHIPPINGSERVICE']._serialized_end=3140
  _globals['_CURRENCYSERVICE']._serialized_start=3143
  _globals['_CURRENCYSERVICE']._serialized_end=3326
  _globals['_PAYMENTSERVICE']._serialized_start=3328
  _globals['_PAYMENTSERVICE']._serialized_end=3413
  _globals['_EMAILSERVICE']._serialized_start=3415
  _globals['_EMAILSERVICE']._serialized_end=3519
  _globals['_CHECKOUTSERVICE']._serialized_start=3521
  _globals['_CHECKOUTSERVICE']._serialized_end=3619
  _globals['_ADSERVICE']._serialized_start=3621
  _globals['_ADSERVICE']._serialized_end=3693
# @@protoc_insertion_point(module_scope)
