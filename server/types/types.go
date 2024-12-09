package types

type CrawlListQuestionResponse struct {
	Data struct {
		Problems struct {
			Nodes []Question `json:"nodes"`
		} `json:"problems"`
	} `json:"data"`
	Extensions struct {
		Debug []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"debug"`
	} `json:"extensions"`
}

type Question struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	Slug          string `json:"slug"`
	Date          string `json:"date"`
	FeaturedImage any    `json:"featuredImage"`
	Info          struct {
		Company  []string `json:"company"`
		Category []string `json:"category"`
		Type     []string `json:"type"`
	} `json:"info"`
	Common struct {
		Difficulty      []string `json:"difficulty"`
		VideoLink       any      `json:"videoLink"`
		ComingSoon      bool     `json:"comingSoon"`
		Subtitle        any      `json:"subtitle"`
		PremiumQuestion bool     `json:"premiumQuestion"`
		StudyPlan       []string `json:"studyPlan"`
	} `json:"common"`
}

type CrawQuestionDetailResponse struct {
	PageProps struct {
		Problem  QuestionDetail `json:"problem"`
		AuthUser struct {
			PhotoURL            string `json:"photoURL"`
			ProviderID          string `json:"providerId"`
			VerifiedPhone       bool   `json:"verifiedPhone"`
			FirstName           string `json:"firstName"`
			OnboardingCompleted bool   `json:"onboardingCompleted"`
			DisplayName         string `json:"displayName"`
			CreatedAt           int64  `json:"createdAt"`
			SolvedProblems      []any  `json:"solvedProblems"`
			StarredProblems     []any  `json:"starredProblems"`
			VerifiedEmail       bool   `json:"verifiedEmail"`
			Role                string `json:"role"`
			LastName            string `json:"lastName"`
			PhoneNumber         string `json:"phoneNumber"`
			Source              string `json:"source"`
			IsTeachableUser     bool   `json:"isTeachableUser"`
			HasPassword         bool   `json:"hasPassword"`
			DislikedProblems    []any  `json:"dislikedProblems"`
			UpdatedAt           int64  `json:"updatedAt"`
			IsEnrolled          bool   `json:"isEnrolled"`
			UID                 string `json:"uid"`
			Name                string `json:"name"`
			LikedProblems       []any  `json:"likedProblems"`
			CourseID            int    `json:"courseId"`
			Email               string `json:"email"`
		} `json:"authUser"`
		Country         string `json:"country"`
		SentryTraceData string `json:"_sentryTraceData"`
		SentryBaggage   string `json:"_sentryBaggage"`
	} `json:"pageProps"`
	NSsp bool `json:"__N_SSP"`
}
type QuestionDetail struct {
	Type             string     `json:"type"`
	ID               string     `json:"id"`
	Title            string     `json:"title"`
	Difficulty       []string   `json:"difficulty"`
	Link             string     `json:"link"`
	ProblemStatement string     `json:"problemStatement"`
	PremiumQuestion  bool       `json:"premiumQuestion"`
	TestCases        []TestCase `json:"testCases"`
	Examples         []any      `json:"examples"`
	Constraints      []struct {
		Constraint any `json:"constraint"`
	} `json:"constraints"`
	Order               any      `json:"order"`
	StarterJSCode       string   `json:"starterJSCode"`
	StarterHTMLCode     any      `json:"starterHTMLCode"`
	StarterCSSCode      any      `json:"starterCSSCode"`
	SolutionCode        string   `json:"solutionCode"`
	StarterFunctionName any      `json:"starterFunctionName"`
	Solution            string   `json:"solution"`
	Category            []string `json:"category"`
	Likes               int      `json:"likes"`
	Dislikes            any      `json:"dislikes"`
	VideoLink           string   `json:"videoLink"`
	Company             []string `json:"company"`
	Subtitle            any      `json:"subtitle"`
	ComingSoon          any      `json:"comingSoon"`
	Seo                 struct {
		PageTitle               string `json:"pageTitle"`
		PageDescription         string `json:"pageDescription"`
		PageURL                 string `json:"pageURL"`
		IsCodingPage            bool   `json:"isCodingPage"`
		UseGeneralPageSD        bool   `json:"useGeneralPageSD"`
		Index                   bool   `json:"index"`
		OgImage                 string `json:"ogImage"`
		TwitterImage            string `json:"twitterImage"`
		Keywords                string `json:"keywords"`
		SecondaryStructuredData struct {
			Context     string `json:"@context"`
			Type        string `json:"@type"`
			Name        string `json:"name"`
			Text        string `json:"text"`
			DateCreated string `json:"dateCreated"`
			Author      struct {
				Type string `json:"@type"`
				Name string `json:"name"`
			} `json:"author"`
			AcceptedAnswer struct {
				Type        string `json:"@type"`
				Text        string `json:"text"`
				DateCreated string `json:"dateCreated"`
				Author      struct {
					Type string `json:"@type"`
					Name string `json:"name"`
				} `json:"author"`
			} `json:"acceptedAnswer"`
		} `json:"secondaryStructuredData"`
	} `json:"seo"`
}

type TestCase struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Test        string `json:"test"`
	Passed      any    `json:"passed"`
	Output      any    `json:"output"`
	Console     any    `json:"console"`
}
