package shared

// this is a data representation of how each individual member is stored
type Member struct {
	Account             string              `dynamo:"Account" json:"Account,omitempty"`
	AccountName         string              `dynamo:"AccountName,omitempty" json:"AccountName"`
	AgeGroup            string              `dynamo:"AgeGroup,omitempty" json:"AgeGroup,omitempty"`
	Baptism             Baptism             `dynamo:"Baptism,omitempty" json:"Baptism,omitempty"`
	Birthday            string              `dynamo:"Birthday,omitempty" json:"Birthday,omitempty"`
	Campus              string              `dynamo:"Campus" json:"Campus,omitempty"`
	CampusName          string              `dynamo:"CampusName" json:"CampusName,omitempty"`
	CreatedAt           int64               `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	DiscoverYourPurpose DiscoverYourPurpose `dynamo:"DiscoverYourPurpose,omitempty" json:"DiscoverYourPurpose,omitempty"`
	Draw                DrawData            `dynamo:"Draw,omitempty" json:"Draw,omitempty"`
	Email               string              `dynamo:"Email,omitempty" json:"Email,omitemptyv"`
	Events              Group               `dynamo:"Events,omitempty" json:"Events,omitempty"`
	FirstName           string              `dynamo:"FirstName,omitempty" json:"FirstName,omitempty"`
	FullName            string              `dynamo:"FullName,omitempty" json:"FullName,omitempty"`
	Gender              string              `dynamo:"Gender,omitempty" json:"Gender,omitempty"`
	GrowGroup           Group               `dynamo:"GrowGroup,omitempty" json:"GrowGroup,omitempty"`
	InterestGroup       Group               `dynamo:"InterestGroup,omitempty" json:"InterestGroup,omitempty"`
	LastName            string              `dynamo:"LastName,omitempty" json:"LastName,omitempty"`
	LastUpdated         int64               `dynamo:"LastUpdated,omitempty" json:"LastUpdated,omitempty"`
	MembershipType      string              `dynamo:"MembershipType,omitempty" json:"MembershipType,omitempty"`
	NewCreation         NewCreation         `dynamo:"NewCreation,omitempty" json:"NewCreation,omitempty"`
	PhoneNumber         string              `dynamo:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
	ReceiveEmail        bool                `dynamo:"ReceiveEmail,omitempty" json:"ReceiveEmail,omitempty"`
	ShirtSize           string              `dynamo:"ShirtSize,omitempty" json:"ShirtSize,omitempty"`
	SlackID             string              `dynamo:"SlackID,omitempty" json:"SlackID,omitempty"`
	SlackChannelID      string              `dynamo:"SlackChannelID,omitempty" json:"SlackChannelID,omitempty"`
	SpiritualGifts      SpiritualGiftsData  `dynamo:"SpiritualGifts,omitempty" json:"SpiritualGifts,omitempty"`
	ThisIsHome          ThisIsHome          `dynamo:"ThisIsHome,omitempty" json:"ThisIsHome,omitempty"`
	UUID                string              `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
	Volunteer           Group               `dynamo:"Volunteer,omitempty" json:"Volunteer,omitempty"`
}

// this is how a tenant account is stored
type Account struct {
	Name string `dynamo:"Name,omitempty" json:"Name"`
	UUID string `dynamo:"UUID,omitempty" json:"UUID"`
}

/// this is how an individual campus data is stored to determine what account renee belongs to
type Campus struct {
	Account              string `dynamo:"Account,omitempty" json:"Account"`
	AccountName          string `dynamo:"AccountName,omitempty" json:"AccountName"`
	BibleClaimChannel    string `dynamo:"BibleClaimChannel,omitempty" json:"BibleClaimChannel"`
	ClerkNumber          string `dynamo:"ClerkNumber,omitempty" json:"ClerkNumber"`
	CurrentGrowGroupTier []int  `dynamo:"CurrentGrowGroupTier,omitempty" json:"CurrentGrowGroupTier"`

	DigitalChurchSlackResponse string `dynamo:"DigitalChurchSlackResponse,omitempty" json:"DigitalChurchSlackResponse"`
	DigitalChurchCommand       string `dynamo:"DigitalChurchCommand,omitempty" json:"DigitalChurchCommand"`

	EventCommand string `dynamo:"EventsCommand,omitempty" json:"EventsCommand"`
	EventsTeam   string `dynamo:"EventsTeam,omitempty" json:"EventsTeam"`

	GiveCommand       string `dynamo:"GiveCommand,omitempty" json:"GiveCommand"`
	GiveTextResponse  string `dynamo:"GiveTextResponse,omitempty" json:"GiveTextResponse"`
	GiveTrigger       string `dynamo:"GiveTrigger,omitempty" json:"GiveTrigger"`
	GiveSlackResponse string `dynamo:"GiveSlackResponse,omitempty" json:"GiveSlackResponse"`

	GrowthTrackCommand      string `dynamo:"GrowthTrackCommand,omitempty" json:"GrowthTrackCommand"`
	GrowthTrackName         string `dynamo:"GrowthTrackName,omitempty" json:"GrowthTrackName"`
	GrowthTrackTeam         string `dynamo:"GrowthTrackTeam,omitempty" json:"GrowthTrackTeam"`
	GrowthTrackWeek1Name    string `dynamo:"GrowthTrackWeek1Name,omitempty" json:"GrowthTrackWeek1Name"`
	GrowthTrackWeek2Name    string `dynamo:"GrowthTrackWeek2Name,omitempty" json:"GrowthTrackWeek2Name"`
	GrowthTrackWeek1Command string `dynamo:"GrowthTrackWeek1Command,omitempty" json:"GrowthTrackWeek1Command"`
	GrowthTrackWeek2Command string `dynamo:"GrowthTrackWeek2Command,omitempty" json:"GrowthTrackWeek2Command"`

	GrowGroupCommand      string `dynamo:"GrowGroupCommand,omitempty" json:"GrowGroupCommand"`
	GrowGroupDescription  string `dynamo:"GrowGroupDescription,omitempty" json:"GrowGroupDescription"`
	GrowGroupName         string `dynamo:"GrowGroupName,omitempty" json:"GrowGroupName"`
	GrowGroupStudyKeyWord string `dynamo:"GrowGroupStudyKeyWord,omitempty" json:"GrowGroupStudyKeyWord"`
	GrowGroupTeam         string `dynamo:"GrowGroupTeam,omitempty" json:"GrowGroupTeam"`

	HealthTeam  string `dynamo:"HealthTeam,omitempty" json:"HealthTeam"`
	HomeTrigger string `dynamo:"HomeTrigger,omitempty" json:"HomeTrigger"`

	InterestGroupDescription string `dynamo:"InterestGroupDescription,omitempty" json:"InterestGroupDescription"`
	InterestGroupName        string `dynamo:"InterestGroupName,omitempty" json:"InterestGroupName"`
	InterestGroupTeam        string `dynamo:"InterestGroupTeam,omitempty" json:"InterestGroupTeam"`

	Name string `dynamo:"Name,omitempty" json:"Name"`

	NewPersonChannel        string `dynamo:"NewPersonChannel,omitempty" json:"NewPersonChannel"`
	NewPersonContactChannel string `dynamo:"NewPersonContactChannel,omitempty" json:"NewPersonContactChannel"`
	NCTrigger               string `dynamo:"NCTrigger,omitempty" json:"NCTrigger"`
	NCChannel               string `dynamo:"NCChannel,omitempty" json:"NCChannel"`
	NewCreationsDinnerName  string `dynamo:"NewCreationsDinnerName,omitempty" json:"NewCreationsDinnerName"`
	NewCreationsTeam        string `dynamo:"NewCreationsTeam,omitempty" json:"NewCreationsTeam"`

	MarketingPlatforms   []string `dynamo:"MarketingPlatforms,omitempty" json:"MarketingPlatforms"`
	OnboardChatWithTeam  string   `dynamo:"OnboardChatWithTeam,omitempty" json:"OnboardChatWithTeam"`
	OnboardGrowGroups    string   `dynamo:"OnboardGrowGroups,omitempty" json:"OnboardGrowGroups"`
	OnboardGrowthTrack   string   `dynamo:"OnboardGrowthTrack,omitempty" json:"OnboardGrowthTrack"`
	OnboardLookingAround string   `dynamo:"OnboardLookingAround,omitempty" json:"OnboardLookingAround"`

	OurCodeCommand       string `dynamo:"OurCodeCommand,omitempty" json:"OurCodeCommand"`
	OurCodeSlackResponse string `dynamo:"OurCodeSlackResponse,omitempty" json:"OurCodeSlackResponse"`

	ServeCommand                string `dynamo:"ServeCommand,omitempty" json:"ServeCommand"`
	ServeTeam                   string `dynamo:"ServeTeam,omitempty" json:"ServeTeam"`
	ServeTeamDescription        string `dynamo:"ServeTeamDescription,omitempty" json:"ServeTeamDescription"`
	ServeDescriptionGrowthTrack string `dynamo:"ServeDescriptionGrowthTrack,omitempty" json:"ServeDescriptionGrowthTrack"`
	ServeTeamName               string `dynamo:"ServeTeamName,omitempty" json:"ServeTeamName"`

	SlackAppID             string `dynamo:"SlackAppID,omitempty" json:"SlackAppID"`
	SlackBotToken          string `dynamo:"SlackBotToken,omitempty" json:"SlackBotToken"`
	SlackClientID          string `dynamo:"SlackClientID,omitempty" json:"SlackClientID"`
	SlackClientSecret      string `dynamo:"SlackClientSecret,omitempty" json:"SlackClientSecret"`
	SlackGeneralChannelID  string `dynamo:"SlackGeneralChannelID,omitempty" json:"SlackGeneralChannelID"`
	SlackInviteToken       string `dynamo:"SlackInviteToken,omitempty" json:"SlackInviteToken"`
	SlackOAuthToken        string `dynamo:"SlackOAuthToken,omitempty" json:"SlackOAuthToken"`
	SlackSigningSecret     string `dynamo:"SlackSigningSecret,omitempty" json:"SlackSigningSecret"`
	SlackTeamID            string `dynamo:"SlackTeamID,omitempty" json:"SlackTeamID"`
	SlackURL               string `dynamo:"SlackURL,omitempty" json:"SlackURL"`
	SlackVerificationToken string `dynamo:"SlackVerificationToken,omitempty" json:"SlackVerificationToken"`

	SlashRenee string `dynamo:"SlashRenee,omitempty" json:"SlashRenee"`

	TimeZone            string `dynamo:"TimeZone,omitempty" json:"TimeZone"`
	UUID                string `dynamo:"UUID,omitempty" json:"UUID"`
	WelcomeSlackMessage string `dynamo:"WelcomeSlackMessage,omitempty" json:"WelcomeSlackMessage"`
}

type SpiritualGiftsData struct {
	Assessment []map[string]string `dynamo:"Assessment,omitempty" json:"Assessment,omitempty"`
	Gifts      map[string]int      `dynamo:"Gifts,omitempty" json:"Gifts,omitempty"`
}

type Baptism struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type Group struct {
	Groups []MemberGroupData `dynamo:"Groups,omitempty" json:"Groups,omitempty"`
	Status bool              `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type MemberGroupData struct {
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	GroupName string `dynamo:"GroupName,omitempty" json:"GroupName,omitempty"`
	UUID      string `dynamo:"UUID" json:"UUID,omitempty"`
}

type NewCreation struct {
	FollowUp      bool           `dynamo:"FollowUp,omitempty" json:"FollowUp,omitempty"`
	FirstDecision NCFirstTime    `dynamo:"FirstDecision,omitempty" json:"FirstDecision,omitempty"`
	Rededication  NCRededication `dynamo:"Rededication,omitempty" json:"Rededication,omitempty"`
	Status        bool           `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type NCFirstTime struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type NCRededication struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type ThisIsHome struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type DiscoverYourPurpose struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
}

type DrawData struct {
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	Type      string `dynamo:"Type,omitemtpy" json:"Type,omitempty"`
}

type Groups struct {
	Account         string                       `dynamo:"Account" json:"Account"`
	AccountName     string                       `dynamo:"AccountName" json:"AccountName"`
	Active          bool                         `dynamo:"Active,omitempty" json:"Active,omitempty"`
	AgeGroup        string                       `dynamo:"AgeGroup,omitempty" json:"AgeGroup,omitempty"`
	Attendance      map[string][]GroupAttendance `dynamo:"Attendance,omitempty" json:"Attendance,omitempty"`
	Campus          string                       `dynamo:"Campus" json:"Campus,omitempty"`
	CampusName      string                       `dynamo:"CampusName" json:"CampusName,omitempty"`
	Capacity        GroupCapacity                `dynamo:"Capacity, omitempty" json:"Capacity,omitempty"`
	ConfirmResponse string                       `dynamo:"ConfirmResponse, omitempty" json:"ConfirmResponse,omitempty"`
	CreatedAt       int64                        `dynamo:"CreatedAt" json:"CreatedAt,omitempty"`
	Details         GroupDetails                 `dynamo:"Details,omitempty" json:"Details,omitempty"`
	GroupName       string                       `dynamo:"GroupName" json:"GroupName,omitempty"`
	LastUpdated     int64                        `dynamo:"LastUpdated" json:"LastUpdated,omitempty"`
	Leaders         []Member                     `dynamo:"Leaders,omitempty" json:"Leaders,omitempty"`
	UUID            string                       `dynamo:"UUID" json:"UUID"`
	Oversights      []Member                     `dynamo:"Oversights,omitempty" json:"Oversights,omitempty"`
	Perpetual       bool                         `dynamo:"Perpetual,omitempty" json:"Perpetual,omitempty"`
	Private         bool                         `dynamo:"Private,omitempty" json:"Private,omitempty"`
	Registrations   []Member                     `dynamo:"Registrations,omitempty" json:"Registrations,omitempty"`
	RegistrationURL string                       `dynamo:"RegistrationURL,omitempty" json:"RegistrationURL,omitempty"`
	RequestFormLink string                       `dynamo:"RequestFormLink,omitempty" json:"RequestFormLink,omitempty"`
	StudyKey        string                       `dynamo:"StudyKey,omitempty" json:"StudyKey,omitempty"`
	Type            string                       `dynamo:"Type,omitempty" json:"Type,omitempty"`
}

type GroupDetails struct {
	ChannelID   string `dynamo:"ChannelID,omitempty" json:"ChannelID,omitempty"`
	Description string `dynamo:"Description,omitempty" json:"Description,omitempty"`
	EndTime     int64  `dynamo:"EndTime,omitempty" json:"EndTime,omitempty"`
	Location    string `dynamo:"Location,omitempty" json:"Location,omitempty"`
	StartTime   int64  `dynamo:"StartTime,omitempty" json:"StartTime,omitempty"`
	Tier        int    `dynamo:"Tier,omitempty" json:"Tier,omitempty"`
}

type GroupCapacity struct {
	AtCapacity bool `dynamo:"AtCapacity" json:"AtCapacity,omitempty"`
	Limit      int  `dynamo:"Limit" json:"Limit,omitempty"`
	Openings   int  `dynamo:"Openings" json:"Openings,omitempty"`
	Unlimited  bool `dynamo:"Unlimited" json:"Unlimited,omitempty"`
}

type GroupAttendance struct {
	Account   string `dynamo:"Account,omitempty" json:"Account"`
	Campus    string `dynamo:"Campus,omitempty" json:"Campus"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt"`
	FullName  string `dynamo:"FullName,omitempty" json:"FullName"`
	UUID      string `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
}
