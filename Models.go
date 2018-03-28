package HotspotSystemGoClient

type Error struct {
	Error string `json:"error"`
}

type Metadata struct {
	TotalCount int `json:"total_count"`
}

type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ArrayOfLocations struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}

type Customer struct {
	ID                   string `json:"id"`
	UserName             string `json:"user_name"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	CompanyName          string `json:"company_name"`
	Address              string `json:"address"`
	City                 string `json:"city"`
	State                string `json:"state"`
	Zip                  string `json:"zip"`
	CountryCode          string `json:"country_code"`
	Phone                string `json:"phone"`
	SocialNetwork        string `json:"social_network"`
	SocialID             string `json:"social_id"`
	SocialUsername       string `json:"social_username"`
	SocialLink           string `json:"social_link"`
	SocialGender         string `json:"social_gender"`
	SocialAgeRange       string `json:"social_age_range"`
	SocialFollowersCount string `json:"social_followers_count"`
	RegisteredAt         string `json:"registered_at"`
}

type ArrayOfCustomers struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID                   string `json:"id"`
		UserName             string `json:"user_name"`
		Name                 string `json:"name"`
		Email                string `json:"email"`
		CompanyName          string `json:"company_name"`
		Address              string `json:"address"`
		City                 string `json:"city"`
		State                string `json:"state"`
		Zip                  string `json:"zip"`
		CountryCode          string `json:"country_code"`
		Phone                string `json:"phone"`
		SocialNetwork        string `json:"social_network"`
		SocialID             string `json:"social_id"`
		SocialUsername       string `json:"social_username"`
		SocialLink           string `json:"social_link"`
		SocialGender         string `json:"social_gender"`
		SocialAgeRange       string `json:"social_age_range"`
		SocialFollowersCount string `json:"social_followers_count"`
		RegisteredAt         string `json:"registered_at"`
	} `json:"items"`
}

type Subscriber struct {
	ID                   string `json:"id"`
	UserName             string `json:"user_name"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	CompanyName          string `json:"company_name"`
	Address              string `json:"address"`
	City                 string `json:"city"`
	State                string `json:"state"`
	Zip                  string `json:"zip"`
	CountryCode          string `json:"country_code"`
	Phone                string `json:"phone"`
	SocialNetwork        string `json:"social_network"`
	SocialID             string `json:"social_id"`
	SocialUsername       string `json:"social_username"`
	SocialLink           string `json:"social_link"`
	SocialGender         string `json:"social_gender"`
	SocialAgeRange       string `json:"social_age_range"`
	SocialFollowersCount string `json:"social_followers_count"`
	RegisteredAt         string `json:"registered_at"`
}

type ArrayOfSubscribers struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID                   string `json:"id"`
		UserName             string `json:"user_name"`
		Name                 string `json:"name"`
		Email                string `json:"email"`
		CompanyName          string `json:"company_name"`
		Address              string `json:"address"`
		City                 string `json:"city"`
		State                string `json:"state"`
		Zip                  string `json:"zip"`
		CountryCode          string `json:"country_code"`
		Phone                string `json:"phone"`
		SocialNetwork        string `json:"social_network"`
		SocialID             string `json:"social_id"`
		SocialUsername       string `json:"social_username"`
		SocialLink           string `json:"social_link"`
		SocialGender         string `json:"social_gender"`
		SocialAgeRange       string `json:"social_age_range"`
		SocialFollowersCount string `json:"social_followers_count"`
		RegisteredAt         string `json:"registered_at"`
	} `json:"items"`
}

type Voucher struct {
	Serial          string `json:"serial"`
	VoucherCode     string `json:"voucher_code"`
	LimitTl         int    `json:"limit_tl"`
	SimultaneousUse int    `json:"simultaneous_use"`
	LimitDl         int    `json:"limit_dl"`
	LimitUl         int    `json:"limit_ul"`
	UsageExp        string `json:"usage_exp"`
	Validity        int    `json:"validity"`
	PriceEnduser    int    `json:"price_enduser"`
	Currency        string `json:"currency"`
}

type ArrayOfVouchers struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		Serial          string `json:"serial"`
		VoucherCode     string `json:"voucher_code"`
		LimitTl         int    `json:"limit_tl"`
		SimultaneousUse int    `json:"simultaneous_use"`
		LimitDl         int    `json:"limit_dl"`
		LimitUl         int    `json:"limit_ul"`
		UsageExp        string `json:"usage_exp"`
		Validity        int    `json:"validity"`
		PriceEnduser    int    `json:"price_enduser"`
		Currency        string `json:"currency"`
	} `json:"items"`
}

type MacTransaction struct {
	ID            int    `json:"id"`
	Operator      string `json:"operator"`
	LocationID    int    `json:"location_id"`
	UserName      string `json:"user_name"`
	ActionDateGmt string `json:"action_date_gmt"`
	PackageID     string `json:"package_id"`
	UserAgent     string `json:"user_agent"`
	Customer      string `json:"customer"`
	Newsletter    int    `json:"newsletter"`
	CompanyName   string `json:"company_name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	CountryCode   string `json:"country_code"`
	Phone         string `json:"phone"`
	Q1            string `json:"q1"`
	A1            string `json:"a1"`
	Q2            string `json:"q2"`
	A2            string `json:"a2"`
	Q3            string `json:"q3"`
	A3            string `json:"a3"`
	Q4            string `json:"q4"`
	A4            string `json:"a4"`
	Q5            string `json:"q5"`
	A5            string `json:"a5"`
}

type ArrayOfMacTransactions struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID            int    `json:"id"`
		Operator      string `json:"operator"`
		LocationID    int    `json:"location_id"`
		UserName      string `json:"user_name"`
		ActionDateGmt string `json:"action_date_gmt"`
		PackageID     string `json:"package_id"`
		UserAgent     string `json:"user_agent"`
		Customer      string `json:"customer"`
		Newsletter    int    `json:"newsletter"`
		CompanyName   string `json:"company_name"`
		Email         string `json:"email"`
		Address       string `json:"address"`
		City          string `json:"city"`
		State         string `json:"state"`
		Zip           string `json:"zip"`
		CountryCode   string `json:"country_code"`
		Phone         string `json:"phone"`
		Q1            string `json:"q1"`
		A1            string `json:"a1"`
		Q2            string `json:"q2"`
		A2            string `json:"a2"`
		Q3            string `json:"q3"`
		A3            string `json:"a3"`
		Q4            string `json:"q4"`
		A4            string `json:"a4"`
		Q5            string `json:"q5"`
		A5            string `json:"a5"`
	} `json:"items"`
}

type VoucherTransaction struct {
	ID            int    `json:"id"`
	Operator      string `json:"operator"`
	LocationID    int    `json:"location_id"`
	UserName      string `json:"user_name"`
	Customer      string `json:"customer"`
	ActionDateGmt string `json:"action_date_gmt"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	UserAgent     string `json:"user_agent"`
	Newsletter    int    `json:"newsletter"`
	CompanyName   string `json:"company_name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	CountryCode   string `json:"country_code"`
	Phone         string `json:"phone"`
	Language      string `json:"language"`
	Smscountry    string `json:"smscountry"`
}

type ArrayOfVoucherTransactions struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID            int    `json:"id"`
		Operator      string `json:"operator"`
		LocationID    int    `json:"location_id"`
		UserName      string `json:"user_name"`
		Customer      string `json:"customer"`
		ActionDateGmt string `json:"action_date_gmt"`
		Amount        int    `json:"amount"`
		Currency      string `json:"currency"`
		UserAgent     string `json:"user_agent"`
		Newsletter    int    `json:"newsletter"`
		CompanyName   string `json:"company_name"`
		Email         string `json:"email"`
		Address       string `json:"address"`
		City          string `json:"city"`
		State         string `json:"state"`
		Zip           string `json:"zip"`
		CountryCode   string `json:"country_code"`
		Phone         string `json:"phone"`
		Language      string `json:"language"`
		Smscountry    string `json:"smscountry"`
	} `json:"items"`
}

type SocialTranaction struct {
	ID                   int    `json:"id"`
	Operator             string `json:"operator"`
	LocationID           int    `json:"location_id"`
	UserName             string `json:"user_name"`
	ActionDateGmt        string `json:"action_date_gmt"`
	PackageID            string `json:"package_id"`
	UserAgent            string `json:"user_agent"`
	Customer             string `json:"customer"`
	Newsletter           int    `json:"newsletter"`
	CompanyName          string `json:"company_name"`
	Email                string `json:"email"`
	Address              string `json:"address"`
	City                 string `json:"city"`
	State                string `json:"state"`
	Zip                  string `json:"zip"`
	CountryCode          string `json:"country_code"`
	Phone                string `json:"phone"`
	SocialID             string `json:"social_id"`
	SocialUsername       string `json:"social_username"`
	SocialLink           string `json:"social_link"`
	SocialGender         string `json:"social_gender"`
	SocialAgeRange       string `json:"social_age_range"`
	SocialFollowersCount string `json:"social_followers_count"`
	SocialNetwork        string `json:"social_network"`
}

type ArrayOfSocialTransactions struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID                   int    `json:"id"`
		Operator             string `json:"operator"`
		LocationID           int    `json:"location_id"`
		UserName             string `json:"user_name"`
		ActionDateGmt        string `json:"action_date_gmt"`
		PackageID            string `json:"package_id"`
		UserAgent            string `json:"user_agent"`
		Customer             string `json:"customer"`
		Newsletter           int    `json:"newsletter"`
		CompanyName          string `json:"company_name"`
		Email                string `json:"email"`
		Address              string `json:"address"`
		City                 string `json:"city"`
		State                string `json:"state"`
		Zip                  string `json:"zip"`
		CountryCode          string `json:"country_code"`
		Phone                string `json:"phone"`
		SocialID             string `json:"social_id"`
		SocialUsername       string `json:"social_username"`
		SocialLink           string `json:"social_link"`
		SocialGender         string `json:"social_gender"`
		SocialAgeRange       string `json:"social_age_range"`
		SocialFollowersCount string `json:"social_followers_count"`
		SocialNetwork        string `json:"social_network"`
	} `json:"items"`
}

type PaidTransaction struct {
	ID            int    `json:"id"`
	Operator      string `json:"operator"`
	LocationID    int    `json:"location_id"`
	UserName      string `json:"user_name"`
	Customer      string `json:"customer"`
	ActionDateGmt string `json:"action_date_gmt"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	UserAgent     string `json:"user_agent"`
	Newsletter    int    `json:"newsletter"`
	CompanyName   string `json:"company_name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	CountryCode   string `json:"country_code"`
	Phone         string `json:"phone"`
	Language      string `json:"language"`
	Smscountry    string `json:"smscountry"`
}

type ArrayOfPaidTransactions struct {
	Metadata struct {
		TotalCount int `json:"total_count"`
	} `json:"metadata"`
	Items []struct {
		ID            int    `json:"id"`
		Operator      string `json:"operator"`
		LocationID    int    `json:"location_id"`
		UserName      string `json:"user_name"`
		Customer      string `json:"customer"`
		ActionDateGmt string `json:"action_date_gmt"`
		Amount        int    `json:"amount"`
		Currency      string `json:"currency"`
		UserAgent     string `json:"user_agent"`
		Newsletter    int    `json:"newsletter"`
		CompanyName   string `json:"company_name"`
		Email         string `json:"email"`
		Address       string `json:"address"`
		City          string `json:"city"`
		State         string `json:"state"`
		Zip           string `json:"zip"`
		CountryCode   string `json:"country_code"`
		Phone         string `json:"phone"`
		Language      string `json:"language"`
		Smscountry    string `json:"smscountry"`
	} `json:"items"`
}

type Pong struct {
	Ping string `json:"ping"`
}

type Owner struct {
	UserID   int    `json:"userId"`
	Operator string `json:"operator"`
}

type ArrayOfLocationsAsOptions struct {
	Items []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}
