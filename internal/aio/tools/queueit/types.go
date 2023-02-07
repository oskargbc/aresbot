package queueit

import "time"

type QueueEnterResponse struct {
	QueueID                      string      `json:"queueId"`
	RedirectURL                  interface{} `json:"redirectUrl"`
	ChallengeFailed              bool        `json:"challengeFailed"`
	MissingCustomDataKey         bool        `json:"missingCustomDataKey"`
	CustomDataUniqueKeyViolation bool        `json:"customDataUniqueKeyViolation"`
	InvalidQueueitEnqueueToken   bool        `json:"invalidQueueitEnqueueToken"`
	ServerIsBusy                 bool        `json:"serverIsBusy"`
}

type QueueItReponse struct {
	IsBeforeOrIdle bool   `json:"isBeforeOrIdle"`
	PageID         string `json:"pageId"`
	PageClass      string `json:"pageClass"`
	QueueState     int    `json:"QueueState"`
	Ticket         struct {
		QueueNumber             interface{} `json:"queueNumber"`
		UsersInLineAheadOfYou   interface{} `json:"usersInLineAheadOfYou"`
		ExpectedServiceTime     interface{} `json:"expectedServiceTime"`
		QueuePaused             bool        `json:"queuePaused"`
		LastUpdatedUTC          time.Time   `json:"lastUpdatedUTC"`
		WhichIsIn               interface{} `json:"whichIsIn"`
		LastUpdated             string      `json:"lastUpdated"`
		Progress                int         `json:"progress"`
		TimeZonePostfix         interface{} `json:"timeZonePostfix"`
		ExpectedServiceTimeUTC  time.Time   `json:"expectedServiceTimeUTC"`
		CustomURLParams         string      `json:"customUrlParams"`
		SdkVersion              interface{} `json:"sdkVersion"`
		WindowStartTimeUTC      interface{} `json:"windowStartTimeUTC"`
		WindowStartTime         string      `json:"windowStartTime"`
		SecondsToStart          int         `json:"secondsToStart"`
		UsersInQueue            int         `json:"usersInQueue"`
		EventStartTimeFormatted string      `json:"eventStartTimeFormatted"`
		EventStartTimeUTC       time.Time   `json:"eventStartTimeUTC"`
	} `json:"ticket"`
	Message interface{} `json:"message"`
	Texts   struct {
		CountdownFinishedText        string        `json:"countdownFinishedText"`
		QueueBody                    interface{}   `json:"queueBody"`
		QueueHeader                  interface{}   `json:"queueHeader"`
		Header                       string        `json:"header"`
		Body                         string        `json:"body"`
		DisclaimerText               interface{}   `json:"disclaimerText"`
		StyleSheets                  string        `json:"styleSheets"`
		Javascripts                  []interface{} `json:"javascripts"`
		LogoSrc                      string        `json:"logoSrc"`
		ToppanelIFrameSrc            string        `json:"toppanelIFrameSrc"`
		SidepanelIFrameSrc           interface{}   `json:"sidepanelIFrameSrc"`
		LeftpanelIFrameSrc           interface{}   `json:"leftpanelIFrameSrc"`
		RightpanelIFrameSrc          interface{}   `json:"rightpanelIFrameSrc"`
		MiddlepanelIFrameSrc         interface{}   `json:"middlepanelIFrameSrc"`
		BottompanelIFrameSrc         interface{}   `json:"bottompanelIFrameSrc"`
		FaviconURL                   string        `json:"faviconUrl"`
		Languages                    []interface{} `json:"languages"`
		WhatIsThisURL                string        `json:"whatIsThisUrl"`
		QueueItLogoPointsToURL       string        `json:"queueItLogoPointsToUrl"`
		WelcomeSoundUrls             []string      `json:"welcomeSoundUrls"`
		CookiesAllowedInfoText       string        `json:"cookiesAllowedInfoText"`
		CookiesNotAllowedInfoText    string        `json:"cookiesNotAllowedInfoText"`
		CookiesAllowedInfoTooltip    string        `json:"cookiesAllowedInfoTooltip"`
		CookiesNotAllowedInfoTooltip string        `json:"cookiesNotAllowedInfoTooltip"`
		AppleTouchIconURL            string        `json:"AppleTouchIconUrl"`
		DocumentTitle                string        `json:"DocumentTitle"`
		Tags                         []interface{} `json:"tags"`
		MessageUpdatedTimeAgo        string        `json:"messageUpdatedTimeAgo"`
	} `json:"texts"`
	Layout struct {
		LanguageSelectorVisible       bool `json:"languageSelectorVisible"`
		LogoVisible                   bool `json:"logoVisible"`
		DynamicMessageVisible         bool `json:"dynamicMessageVisible"`
		ReminderEmailVisible          bool `json:"reminderEmailVisible"`
		ExpectedServiceTimeVisible    bool `json:"expectedServiceTimeVisible"`
		QueueNumberVisible            bool `json:"queueNumberVisible"`
		UsersInLineAheadOfYouVisible  bool `json:"usersInLineAheadOfYouVisible"`
		WhichIsInVisible              bool `json:"whichIsInVisible"`
		SidePanelVisible              bool `json:"sidePanelVisible"`
		TopPanelVisible               bool `json:"topPanelVisible"`
		LeftPanelVisible              bool `json:"leftPanelVisible"`
		RightPanelVisible             bool `json:"rightPanelVisible"`
		MiddlePanelVisible            bool `json:"middlePanelVisible"`
		BottomPanelVisible            bool `json:"bottomPanelVisible"`
		UsersInQueueVisible           bool `json:"usersInQueueVisible"`
		QueueIsPausedVisible          bool `json:"queueIsPausedVisible"`
		ReminderVisible               bool `json:"reminderVisible"`
		ServicedSoonVisible           bool `json:"servicedSoonVisible"`
		FirstInLineVisible            bool `json:"firstInLineVisible"`
		QueueNumberLoadingVisible     bool `json:"queueNumberLoadingVisible"`
		ProgressVisible               bool `json:"progressVisible"`
		IsRedirectPromptDialogEnabled bool `json:"isRedirectPromptDialogEnabled"`
		IsQueueitFooterHidden         bool `json:"isQueueitFooterHidden"`
	} `json:"layout"`
	ForecastStatus string `json:"forecastStatus"`
	LayoutName     string `json:"layoutName"`
	LayoutVersion  int64  `json:"layoutVersion"`
	UpdateInterval int    `json:"updateInterval"`
}

type QueueFinishedResponse struct {
	RedirectURL        string `json:"redirectUrl"`
	IsRedirectToTarget bool   `json:"isRedirectToTarget"`
}

type LoadedQueueResponse struct {
	CustomerID   string `json:"customerId"`
	Integrations []struct {
		Name            string `json:"name"`
		ActionType      string `json:"actionType"`
		EventID         string `json:"eventId"`
		CookieDomain    string `json:"cookieDomain"`
		QueueDomain     string `json:"queueDomain"`
		ForcedTargetURL string `json:"forcedTargetUrl"`
		Culture         string `json:"culture"`
		ExtendValidity  bool   `json:"extendValidity"`
		Validity        int    `json:"validity"`
		RedirectLogic   string `json:"redirectLogic"`
		Triggers        []struct {
			TriggerParts []struct {
				Operator       string `json:"operator"`
				ValueToCompare string `json:"valueToCompare"`
				URLPart        string `json:"urlPart"`
				ValidatorType  string `json:"validatorType"`
				IsNegative     bool   `json:"isNegative"`
				IsIgnoreCase   bool   `json:"isIgnoreCase"`
			} `json:"triggerParts"`
			LogicalOperator string `json:"logicalOperator"`
		} `json:"triggers"`
		OnVerified         string `json:"onVerified"`
		OnDisabled         string `json:"onDisabled"`
		OnTimeout          string `json:"onTimeout"`
		RemoveTokenFromURL bool   `json:"removeTokenFromUrl"`
	} `json:"integrations"`
}

type PowResponse struct {
	Function   string `json:"function"`
	SessionID  string `json:"sessionId"`
	Meta       string `json:"meta"`
	Parameters struct {
		Input     string `json:"input"`
		ZeroCount int    `json:"zeroCount"`
	} `json:"parameters"`
}
type PowSessionId struct {
	UserID     string     `json:"userId"`
	Meta       string     `json:"meta"`
	SessionID  string     `json:"sessionId"`
	Solution   Solution   `json:"solution"`
	Tags       []string   `json:"tags"`
	Stats      Stats      `json:"stats"`
	Parameters Parameters `json:"parameters"`
}
type Solution struct {
	Postfix int    `json:"postfix"`
	Hash    string `json:"hash"`
}
type Stats struct {
	Duration       int    `json:"duration"`
	Tries          int    `json:"tries"`
	UserAgent      string `json:"userAgent"`
	Screen         string `json:"screen"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browserVersion"`
	IsMobile       bool   `json:"isMobile"`
	Os             string `json:"os"`
	OsVersion      string `json:"osVersion"`
	CookiesEnabled bool   `json:"cookiesEnabled"`
}
type Parameters struct {
	Input     string `json:"input"`
	ZeroCount int    `json:"zeroCount"`
}

type PowSolveData struct {
	ChallengeType string `json:"challengeType"`
	SessionID     string `json:"sessionId"`
	CustomerID    string `json:"customerId"`
	EventID       string `json:"eventId"`
	Version       int    `json:"version"`
}

type PowSolveResponse struct {
	IsVerified  bool      `json:"isVerified"`
	Timestamp   time.Time `json:"timestamp"`
	SessionInfo struct {
		SessionID     string    `json:"sessionId"`
		Timestamp     time.Time `json:"timestamp"`
		Checksum      string    `json:"checksum"`
		SourceIP      string    `json:"sourceIp"`
		ChallengeType string    `json:"challengeType"`
		Version       int       `json:"version"`
	} `json:"sessionInfo"`
}

type RecaptchaSolveData struct {
	ChallengeType string `json:"challengeType"`
	SessionID     string `json:"sessionId"`
	CustomerID    string `json:"customerId"`
	EventID       string `json:"eventId"`
	Version       int    `json:"version"`
}

type RecaptchaSolveResponse struct {
	IsVerified  bool      `json:"isVerified"`
	Timestamp   time.Time `json:"timestamp"`
	SessionInfo struct {
		SessionID     string    `json:"sessionId"`
		Timestamp     time.Time `json:"timestamp"`
		Checksum      string    `json:"checksum"`
		SourceIP      string    `json:"sourceIp"`
		ChallengeType string    `json:"challengeType"`
		Version       int       `json:"version"`
	} `json:"sessionInfo"`
}

type QueueItEnterRequest struct {
	ChallengeSessions []ChallengeSessions `json:"challengeSessions"`
	LayoutName        string              `json:"layoutName"`
	CustomURLParams   string              `json:"customUrlParams"`
	TargetURL         string              `json:"targetUrl"`
	Referrer          string              `json:"Referrer"`
	LayoutVersion     string              `json:"layoutVersion"`
}
type ChallengeSessions struct {
	SessionID     string    `json:"sessionId"`
	Timestamp     time.Time `json:"timestamp"`
	Checksum      string    `json:"checksum"`
	SourceIP      string    `json:"sourceIp"`
	ChallengeType string    `json:"challengeType"`
	Version       int       `json:"version"`
}

type ApiReponse struct {
	Test []Test `json:"test"`
	Drop []Drop `json:"drop"`
}
type Params struct {
	C             string `json:"c"`
	E             string `json:"e"`
	T             string `json:"t"`
	TargetURL     string `json:"targetUrl"`
	Sitekey       string `json:"sitekey"`
	LayoutName    string `json:"layoutName"`
	LayoutVersion string `json:"layoutVersion"`
	Referrer      string `json:"referrer"`
}
type Test struct {
	Shop   string `json:"shop"`
	Params Params `json:"params"`
}
type Drop struct {
	Shop   string `json:"shop"`
	Params Params `json:"params"`
}
