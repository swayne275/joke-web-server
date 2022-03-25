package randomuser

import "testing"

func TestParseNames(t *testing.T) {
	cases := []struct {
		apiResponse   string
		expectedFirst string
		expectedLast  string
		expectErr     bool
	}{
		{
			apiResponse:   "{\"results\":[{\"gender\":\"female\",\"name\":{\"title\":\"Miss\",\"first\":\"Claudia\",\"last\":\"Johnson\"},\"location\":{\"street\":{\"number\":8261,\"name\":\"Locust Rd\"},\"city\":\"El Paso\",\"state\":\"Connecticut\",\"country\":\"United States\",\"postcode\":29950,\"coordinates\":{\"latitude\":\"84.3100\",\"longitude\":\"43.6867\"},\"timezone\":{\"offset\":\"+7:00\",\"description\":\"Bangkok, Hanoi, Jakarta\"}},\"email\":\"claudia.johnson@example.com\",\"login\":{\"uuid\":\"370cd742-3522-40b1-bd46-4c972f476622\",\"username\":\"greenfish810\",\"password\":\"doughnut\",\"salt\":\"vklcSfQZ\",\"md5\":\"75b2b5ea7413e810a79bfe08282c803f\",\"sha1\":\"9178ad8a6d99c004ae34e3277b06f7a4663dd404\",\"sha256\":\"e9542397296af773692be1f7bd93e2d82bd5a3da89d46de017fbaa32bd2f7696\"},\"dob\":{\"date\":\"1975-07-10T02:53:06.125Z\",\"age\":45},\"registered\":{\"date\":\"2010-08-08T12:58:31.814Z\",\"age\":10},\"phone\":\"(924)-441-1650\",\"cell\":\"(228)-947-1304\",\"id\":{\"name\":\"SSN\",\"value\":\"059-39-3854\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/13.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/13.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/13.jpg\"},\"nat\":\"US\"}],\"info\":{\"seed\":\"5e5d50e564f979f9\",\"results\":1,\"page\":1,\"version\":\"1.3\"}}",
			expectedFirst: "Claudia",
			expectedLast:  "Johnson",
			expectErr:     false,
		},
		{
			apiResponse:   "{\"results\":[{\"gender\":\"female\",\"name\":{\"title\":\"Miss\",\"first1\":\"Claudia\",\"last\":\"Johnson\"},\"location\":{\"street\":{\"number\":8261,\"name\":\"Locust Rd\"},\"city\":\"El Paso\",\"state\":\"Connecticut\",\"country\":\"United States\",\"postcode\":29950,\"coordinates\":{\"latitude\":\"84.3100\",\"longitude\":\"43.6867\"},\"timezone\":{\"offset\":\"+7:00\",\"description\":\"Bangkok, Hanoi, Jakarta\"}},\"email\":\"claudia.johnson@example.com\",\"login\":{\"uuid\":\"370cd742-3522-40b1-bd46-4c972f476622\",\"username\":\"greenfish810\",\"password\":\"doughnut\",\"salt\":\"vklcSfQZ\",\"md5\":\"75b2b5ea7413e810a79bfe08282c803f\",\"sha1\":\"9178ad8a6d99c004ae34e3277b06f7a4663dd404\",\"sha256\":\"e9542397296af773692be1f7bd93e2d82bd5a3da89d46de017fbaa32bd2f7696\"},\"dob\":{\"date\":\"1975-07-10T02:53:06.125Z\",\"age\":45},\"registered\":{\"date\":\"2010-08-08T12:58:31.814Z\",\"age\":10},\"phone\":\"(924)-441-1650\",\"cell\":\"(228)-947-1304\",\"id\":{\"name\":\"SSN\",\"value\":\"059-39-3854\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/13.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/13.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/13.jpg\"},\"nat\":\"US\"}],\"info\":{\"seed\":\"5e5d50e564f979f9\",\"results\":1,\"page\":1,\"version\":\"1.3\"}}",
			expectedFirst: "",
			expectedLast:  "",
			expectErr:     true,
		},
		{
			apiResponse:   "{\"results\":[{\"gender\":\"female\",\"name\":{\"title\":\"Miss\",\"first\":\"Claudia\",\"last1\":\"Johnson\"},\"location\":{\"street\":{\"number\":8261,\"name\":\"Locust Rd\"},\"city\":\"El Paso\",\"state\":\"Connecticut\",\"country\":\"United States\",\"postcode\":29950,\"coordinates\":{\"latitude\":\"84.3100\",\"longitude\":\"43.6867\"},\"timezone\":{\"offset\":\"+7:00\",\"description\":\"Bangkok, Hanoi, Jakarta\"}},\"email\":\"claudia.johnson@example.com\",\"login\":{\"uuid\":\"370cd742-3522-40b1-bd46-4c972f476622\",\"username\":\"greenfish810\",\"password\":\"doughnut\",\"salt\":\"vklcSfQZ\",\"md5\":\"75b2b5ea7413e810a79bfe08282c803f\",\"sha1\":\"9178ad8a6d99c004ae34e3277b06f7a4663dd404\",\"sha256\":\"e9542397296af773692be1f7bd93e2d82bd5a3da89d46de017fbaa32bd2f7696\"},\"dob\":{\"date\":\"1975-07-10T02:53:06.125Z\",\"age\":45},\"registered\":{\"date\":\"2010-08-08T12:58:31.814Z\",\"age\":10},\"phone\":\"(924)-441-1650\",\"cell\":\"(228)-947-1304\",\"id\":{\"name\":\"SSN\",\"value\":\"059-39-3854\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/13.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/13.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/13.jpg\"},\"nat\":\"US\"}],\"info\":{\"seed\":\"5e5d50e564f979f9\",\"results\":1,\"page\":1,\"version\":\"1.3\"}}",
			expectedFirst: "",
			expectedLast:  "",
			expectErr:     true,
		},
	}

	for _, c := range cases {
		first, last, err := parseNames([]byte(c.apiResponse))
		gotErr := (err != nil)

		if first != c.expectedFirst {
			t.Errorf("incorrect first name. Expected '%s', got '%s'", c.expectedFirst, first)
		}
		if last != c.expectedLast {
			t.Errorf("incorrect last name. Expected '%s', got '%s'", c.expectedLast, last)
		}
		if gotErr != c.expectErr {
			t.Errorf("incorrect error presense. Expected %t, got %t", c.expectErr, gotErr)
		}
	}
}
