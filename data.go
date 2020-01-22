package shared

// this is a data representation of how each individual member is stored
type Member struct {
	Account             string              `dynamo:"Account" json:"Account,omitempty"`
	AccountName         string              `dynamo:"AccountName,omitempty" json:"AccountName"`
	Campus              string              `dynamo:"Campus" json:"Campus,omitempty"`
	CampusName          string              `dynamo:"CampusName" json:"CampusName,omitempty"`
	AgeGroup            string              `dynamo:"AgeGroup,omitempty" json:"AgeGroup,omitempty"`
	ShirtSize           string              `dynamo:"ShirtSize,omitempty" json:"ShirtSize,omitempty"`
	Draw                DrawData            `dynamo:"Draw,omitempty" json:"Draw,omitempty"`
	Baptism             Baptism             `dynamo:"Baptism,omitempty" json:"Baptism,omitempty"`
	Birthday            string              `dynamo:"Birthday,omitempty" json:"Birthday,omitempty"`
	Email               string              `dynamo:"Email,omitempty" json:"Email,omitemptyv"`
	FirstName           string              `dynamo:"FirstName,omitempty" json:"FirstName,omitempty"`
	FullName            string              `dynamo:"FullName,omitempty" json:"FullName,omitempty"`
	Gender              string              `dynamo:"Gender,omitempty" json:"Gender,omitempty"`
	LastName            string              `dynamo:"LastName,omitempty" json:"LastName,omitempty"`
	CreatedAt           int64               `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	LastUpdated         int64               `dynamo:"LastUpdated,omitempty" json:"LastUpdated,omitempty"`
	MembershipType      string              `dynamo:"MembershipType,omitempty" json:"MembershipType,omitempty"`
	PhoneNumber         string              `dynamo:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
	ReceiveEmail        bool                `dynamo:"ReceiveEmail,omitempty" json:"ReceiveEmail,omitempty"`
	SlackID             string              `dynamo:"SlackID,omitempty" json:"SlackID,omitempty"`
	SlackChannelID      string              `dynamo:"SlackChannelID,omitempty" json:"SlackChannelID,omitempty"`
	UUID                string              `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
	NewCreation         NewCreation         `dynamo:"NewCreation,omitempty" json:"NewCreation,omitempty"`
	Volunteer           Group               `dynamo:"Volunteer,omitempty" json:"Volunteer,omitempty"`
	InterestGroup       Group               `dynamo:"InterestGroup,omitempty" json:"InterestGroup,omitempty"`
	GrowGroup           Group               `dynamo:"GrowGroup,omitempty" json:"GrowGroup,omitempty"`
	Events              Group               `dynamo:"Events,omitempty" json:"Events,omitempty"`
	ThisIsHome          ThisIsHome          `dynamo:"ThisIsHome,omitempty" json:"ThisIsHome,omitempty"`
	DiscoverYourPurpose DiscoverYourPurpose `dynamo:"DiscoverYourPurpose,omitempty" json:"DiscoverYourPurpose,omitempty"`
	SpiritualGifts      SpiritualGiftsData  `dynamo:"SpiritualGifts,omitempty" json:"SpiritualGifts,omitempty"`
}

// this is how a tenant account is stored
type Account struct {
	UUID string `dynamo:"UUID,omitempty" json:"UUID"`
	Name string `dynamo:"Name,omitempty" json:"Name"`
}

// this is how an individual campus data is stored to determine what account renee belongs to
type Campus struct {
	UUID                       string `dynamo:"UUID,omitempty" json:"UUID"`
	Account                    string `dynamo:"Account,omitempty" json:"Account"`
	AccountName                string `dynamo:"AccountName,omitempty" json:"AccountName"`
	Name                       string `dynamo:"Name,omitempty" json:"Name"`
	ClerkNumber                string `dynamo:"ClerkNumber,omitempty" json:"ClerkNumber"`
	GiveTextResponse           string `dynamo:"GiveTextResponse,omitempty" json:"GiveTextResponse"`
	GiveSlackResponse          string `dynamo:"GiveSlackResponse,omitempty" json:"GiveSlackResponse"`
	DigitalChurchSlackResponse string `dynamo:"DigitalChurchSlackResponse,omitempty" json:"DigitalChurchSlackResponse"`
	GiveTrigger                string `dynamo:"GiveTrigger,omitempty" json:"GiveTrigger"`
	GrowthTrackTeam            string `dynamo:"GrowthTrackTeam,omitempty" json:"GrowthTrackTeam"`
	GrowGroupTeam              string `dynamo:"GrowGroupTeam,omitempty" json:"GrowGroupTeam"`
	HealthTeam                 string `dynamo:"HealthTeam,omitempty" json:"HealthTeam"`
	HomeTrigger                string `dynamo:"HomeTrigger,omitempty" json:"HomeTrigger"`
	NCTrigger                  string `dynamo:"NCTrigger,omitempty" json:"NCTrigger"`
	NewCreationsTeam           string `dynamo:"NewCreationsTeam,omitempty" json:"NewCreationsTeam"`
	CurrentGrowGroupTier       []int  `dynamo:"CurrentGrowGroupTier,omitempty" json:"CurrentGrowGroupTier"`
	OurCodeSlackResponse       string `dynamo:"OurCodeSlackResponse,omitempty" json:"OurCodeSlackResponse"`
	SlackAppID                 string `dynamo:"SlackAppID,omitempty" json:"SlackAppID"`
	SlackBotToken              string `dynamo:"SlackBotToken,omitempty" json:"SlackBotToken"`
	SlackClientID              string `dynamo:"SlackClientID,omitempty" json:"SlackClientID"`
	SlackGeneralChannelID      string `dynamo:"SlackGeneralChannelID,omitempty" json:"SlackGeneralChannelID"`
	SlackClientSecret          string `dynamo:"SlackClientSecret,omitempty" json:"SlackClientSecret"`
	SlackInviteToken           string `dynamo:"SlackInviteToken,omitempty" json:"SlackInviteToken"`
	SlackOAuthToken            string `dynamo:"SlackOAuthToken,omitempty" json:"SlackOAuthToken"`
	SlackSigningSecret         string `dynamo:"SlackSigningSecret,omitempty" json:"SlackSigningSecret"`
	SlackTeamID                string `dynamo:"SlackTeamID,omitempty" json:"SlackTeamID"`
	SlackURL                   string `dynamo:"SlackURL,omitempty" json:"SlackURL"`
	SlackVerificationToken     string `dynamo:"SlackVerificationToken,omitempty" json:"SlackVerificationToken"`
	WelcomeSlackMessage        string `dynamo:"WelcomeSlackMessage,omitempty" json:"WelcomeSlackMessage"`
}

type SpiritualGiftsData struct {
	Assessment []map[string]string `dynamo:"Assessment,omitempty" json:"Assessment,omitempty"`
	Gifts      map[string]int      `dynamo:"Gifts,omitempty" json:"Gifts,omitempty"`
}

type Baptism struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type Group struct {
	Status bool              `dynamo:"Status,omitempty" json:"Status,omitempty"`
	Groups []MemberGroupData `dynamo:"Groups,omitempty" json:"Groups,omitempty"`
}

type MemberGroupData struct {
	UUID      string `dynamo:"UUID" json:"UUID,omitempty"`
	GroupName string `dynamo:"GroupName,omitempty" json:"GroupName,omitempty"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type NewCreation struct {
	FollowUp      bool           `dynamo:"FollowUp,omitempty" json:"FollowUp,omitempty"`
	Status        bool           `dynamo:"Status,omitempty" json:"Status,omitempty"`
	FirstDecision NCFirstTime    `dynamo:"FirstDecision,omitempty" json:"FirstDecision,omitempty"`
	Rededication  NCRededication `dynamo:"Rededication,omitempty" json:"Rededication,omitempty"`
}

type NCFirstTime struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type NCRededication struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type ThisIsHome struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type DiscoverYourPurpose struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type DrawData struct {
	Type      string `dynamo:"Type,omitemtpy" json:"Type,omitempty"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type Groups struct {
	Account         string                       `dynamo:"Account" json:"Account"`
	AccountName     string                       `dynamo:"AccountName" json:"AccountName"`
	Campus          string                       `dynamo:"Campus" json:"Campus,omitempty"`
	CampusName      string                       `dynamo:"CampusName" json:"CampusName,omitempty"`
	UUID            string                       `dynamo:"UUID" json:"UUID"`
	Leaders         []Member                     `dynamo:"Leaders,omitempty" json:"Leaders,omitempty"`
	Registrations   []Member                     `dynamo:"Registrations,omitempty" json:"Registrations,omitempty"`
	Oversights      []Member                     `dynamo:"Oversights,omitempty" json:"Oversights,omitempty"`
	GroupName       string                       `dynamo:"GroupName" json:"GroupName,omitempty"`
	CreatedAt       int64                        `dynamo:"CreatedAt" json:"CreatedAt,omitempty"`
	LastUpdated     int64                        `dynamo:"LastUpdated" json:"LastUpdated,omitempty"`
	AgeGroup        string                       `dynamo:"AgeGroup,omitempty" json:"AgeGroup,omitempty"`
	Details         GroupDetails                 `dynamo:"Details,omitempty" json:"Details,omitempty"`
	Capacity        GroupCapacity                `dynamo:"Capacity, omitempty" json:"Capacity,omitempty"`
	StudyKey        string                       `dynamo:"StudyKey,omitempty" json:"StudyKey,omitempty"`
	Type            string                       `dynamo:"Type,omitempty" json:"Type,omitempty"`
	RegistrationURL string                       `dynamo:"RegistrationURL,omitempty" json:"RegistrationURL,omitempty"`
	RequestFormLink string                       `dynamo:"RequestFormLink,omitempty" json:"RequestFormLink,omitempty"`
	Active          bool                         `dynamo:"Active,omitempty" json:"Active,omitempty"`
	Private         bool                         `dynamo:"Private,omitempty" json:"Private,omitempty"`
	Attendance      map[string][]GroupAttendance `dynamo:"Attendance,omitempty" json:"Attendance,omitempty"`
}

type GroupDetails struct {
	ChannelID   string `dynamo:"ChannelID,omitempty" json:"ChannelID,omitempty"`
	Tier        int    `dynamo:"Tier,omitempty" json:"Tier,omitempty"`
	Location    string `dynamo:"Location,omitempty" json:"Location,omitempty"`
	StartTime   int64  `dynamo:"StartTime,omitempty" json:"StartTime,omitempty"`
	EndTime     int64  `dynamo:"EndTime,omitempty" json:"EndTime,omitempty"`
	Description string `dynamo:"Description,omitempty" json:"Description,omitempty"`
}

type GroupCapacity struct {
	AtCapacity bool `dynamo:"AtCapacity" json:"AtCapacity,omitempty"`
	Limit      int  `dynamo:"Limit" json:"Limit,omitempty"`
	Openings   int  `dynamo:"Openings" json:"Openings,omitempty"`
	Unlimited  bool `dynamo:"Unlimited" json:"Unlimited,omitempty"`
}

type GroupAttendance struct {
	UUID      string `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
	FullName  string `dynamo:"FullName,omitempty" json:"FullName"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt"`
	Account   string `dynamo:"Account,omitempty" json:"Account"`
	Campus    string `dynamo:"Campus,omitempty" json:"Campus"`
}
