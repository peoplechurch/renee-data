package shared

// Member : Campus Member data structure
type Member struct {
	Account             string              `dynamo:"Account,omitempty" json:"account,omitempty"`
	AccountName         string              `dynamo:"AccountName,omitempty" json:"accountName,omitempty"`
	AgeGroup            string              `dynamo:"AgeGroup,omitempty" json:"ageGroup,omitempty"`
	Baptism             Baptism             `dynamo:"Baptism,omitempty" json:"baptism,omitempty"`
	Birthday            string              `dynamo:"Birthday,omitempty" json:"birthday,omitempty"`
	Campus              string              `dynamo:"Campus,omitempty" json:"campus,omitempty"`
	CampusName          string              `dynamo:"CampusName,omitempty" json:"campusName,omitempty"`
	CreatedAt           int64               `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	DiscoverYourPurpose DiscoverYourPurpose `dynamo:"DiscoverYourPurpose,omitempty" json:"discoverYourPurpose,omitempty"`
	Draw                DrawData            `dynamo:"Draw,omitempty" json:"draw,omitempty"`
	Email               string              `dynamo:"Email,omitempty" json:"email,omitemptyv"`
	Events              Group               `dynamo:"Events,omitempty" json:"events,omitempty"`
	FirstName           string              `dynamo:"FirstName,omitempty" json:"firstName,omitempty"`
	FullName            string              `dynamo:"FullName,omitempty" json:"fullName,omitempty"`
	Gender              string              `dynamo:"Gender,omitempty" json:"gender,omitempty"`
	GrowGroup           Group               `dynamo:"GrowGroup,omitempty" json:"growGroup,omitempty"`
	InterestGroup       Group               `dynamo:"InterestGroup,omitempty" json:"interestGroup,omitempty"`
	LastName            string              `dynamo:"LastName,omitempty" json:"lastName,omitempty"`
	LastUpdated         int64               `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	MembershipType      string              `dynamo:"MembershipType,omitempty" json:"membershipType,omitempty"`
	NewCreation         NewCreation         `dynamo:"NewCreation,omitempty" json:"newCreation,omitempty"`
	PhoneNumber         string              `dynamo:"PhoneNumber,omitempty" json:"phoneNumber,omitempty"`
	ReceiveEmail        bool                `dynamo:"ReceiveEmail,omitempty" json:"receiveEmail,omitempty"`
	ShirtSize           string              `dynamo:"ShirtSize,omitempty" json:"shirtSize,omitempty"`
	SlackID             string              `dynamo:"SlackID,omitempty" json:"slackID,omitempty"`
	SlackChannelID      string              `dynamo:"SlackChannelID,omitempty" json:"slackChannelID,omitempty"`
	SpiritualGifts      SpiritualGiftsData  `dynamo:"SpiritualGifts,omitempty" json:"spiritualGifts,omitempty"`
	ThisIsHome          ThisIsHome          `dynamo:"ThisIsHome,omitempty" json:"thisIsHome,omitempty"`
	UUID                string              `dynamo:"UUID,hash,omitempty" json:"UUID,omitempty"`
	Volunteer           Group               `dynamo:"Volunteer,omitempty" json:"volunteer,omitempty"`
}

// Account : Tenant Account data structure
type Account struct {
	Name        string `dynamo:"Name,omitempty" json:"name,omitempty"`
	UUID        string `dynamo:"UUID,hash,omitempty" json:"UUID,omitempty"`
	CreatedAt   int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	LastUpdated int64  `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
}

// Campus : Account Campus data structure
type Campus struct {
	Account                    string `dynamo:"Account,omitempty" json:"Account"`
	AccountName                string `dynamo:"AccountName,omitempty" json:"AccountName"`
	BibleClaimChannel          string `dynamo:"BibleClaimChannel,omitempty" json:"BibleClaimChannel"`
	ClerkNumber                string `dynamo:"ClerkNumber,omitempty" json:"ClerkNumber"`
	CurrentGrowGroupTier       []int  `dynamo:"CurrentGrowGroupTier,omitempty" json:"CurrentGrowGroupTier"`
	CreatedAt                  int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	DigitalChurchSlackResponse string `dynamo:"DigitalChurchSlackResponse,omitempty" json:"DigitalChurchSlackResponse"`
	DigitalChurchCommand       string `dynamo:"DigitalChurchCommand,omitempty" json:"DigitalChurchCommand"`
	EventCommand               string `dynamo:"EventsCommand,omitempty" json:"EventsCommand"`
	EventsTeam                 string `dynamo:"EventsTeam,omitempty" json:"EventsTeam"`
	GiveCommand                string `dynamo:"GiveCommand,omitempty" json:"GiveCommand"`
	GiveTextResponse           string `dynamo:"GiveTextResponse,omitempty" json:"GiveTextResponse"`
	GiveTrigger                string `dynamo:"GiveTrigger,omitempty" json:"GiveTrigger"`
	GiveSlackResponse          string `dynamo:"GiveSlackResponse,omitempty" json:"GiveSlackResponse"`
	GrowthTrackCommand         string `dynamo:"GrowthTrackCommand,omitempty" json:"GrowthTrackCommand"`
	GrowthTrackName            string `dynamo:"GrowthTrackName,omitempty" json:"GrowthTrackName"`
	GrowthTrackTeam            string `dynamo:"GrowthTrackTeam,omitempty" json:"GrowthTrackTeam"`
	GrowthTrackW1Name          string `dynamo:"GrowthTrackW1Name,omitempty" json:"GrowthTrackW1Name"`
	GrowthTrackW2Name          string `dynamo:"GrowthTrackW2Name,omitempty" json:"GrowthTrackW2Name"`
	GrowthTrackW1Command       string `dynamo:"GrowthTrackW1Command,omitempty" json:"GrowthTrackW1Command"`
	GrowthTrackW2Command       string `dynamo:"GrowthTrackW2Command,omitempty" json:"GrowthTrackW2Command"`
	// Message to send when someone has already completed week 1 of Growth Track
	GrowthTrackW1Redundant      string   `dynamo:"GrowthTrackW1Redundant,omitempty" json:"GrowthTrackW1Redundant"`
	GrowGroupCommand            string   `dynamo:"GrowGroupCommand,omitempty" json:"GrowGroupCommand"`
	GrowGroupDescription        string   `dynamo:"GrowGroupDescription,omitempty" json:"GrowGroupDescription"`
	GrowGroupName               string   `dynamo:"GrowGroupName,omitempty" json:"GrowGroupName"`
	GrowGroupStudyKeyWord       string   `dynamo:"GrowGroupStudyKeyWord,omitempty" json:"GrowGroupStudyKeyWord"`
	GrowGroupTeam               string   `dynamo:"GrowGroupTeam,omitempty" json:"GrowGroupTeam"`
	HealthTeam                  string   `dynamo:"HealthTeam,omitempty" json:"HealthTeam"`
	HomeTrigger                 string   `dynamo:"HomeTrigger,omitempty" json:"HomeTrigger"`
	InterestGroupCommand        string   `dynamo:"InterestGroupCommand,omitempty" json:"InterestGroupCommand"`
	InterestGroupDescription    string   `dynamo:"InterestGroupDescription,omitempty" json:"InterestGroupDescription"`
	InterestGroupName           string   `dynamo:"InterestGroupName,omitempty" json:"InterestGroupName"`
	InterestGroupTeam           string   `dynamo:"InterestGroupTeam,omitempty" json:"InterestGroupTeam"`
	LastUpdated                 int64    `dynamo:"LastUpdated,omitempty" json:"lastUpdated"`
	Name                        string   `dynamo:"Name,omitempty" json:"Name"`
	NewPersonChannel            string   `dynamo:"NewPersonChannel,omitempty" json:"NewPersonChannel"`
	NewPersonReneeChannel       string   `dynamo:"NewPersonReneeChannel,omitempty" json:"NewPersonReneeChannel"`
	NewPersonContactChannel     string   `dynamo:"NewPersonContactChannel,omitempty" json:"NewPersonContactChannel"`
	NCTrigger                   string   `dynamo:"NCTrigger,omitempty" json:"NCTrigger"`
	NCChannel                   string   `dynamo:"NCChannel,omitempty" json:"NCChannel"`
	NewCreationName             string   `dynamo:"newCreationName,omitempty" json:"newCreationName,omitempty"`
	NewCreationsDinnerName      string   `dynamo:"NewCreationsDinnerName,omitempty" json:"NewCreationsDinnerName"`
	NewCreationsTeam            string   `dynamo:"NewCreationsTeam,omitempty" json:"NewCreationsTeam"`
	MarketingPlatforms          []string `dynamo:"MarketingPlatforms,omitempty" json:"MarketingPlatforms"`
	OnboardChatWithTeam         string   `dynamo:"OnboardChatWithTeam,omitempty" json:"OnboardChatWithTeam"`
	OnboardGrowGroups           string   `dynamo:"OnboardGrowGroups,omitempty" json:"OnboardGrowGroups"`
	OnboardGrowthTrack          string   `dynamo:"OnboardGrowthTrack,omitempty" json:"OnboardGrowthTrack"`
	OnboardLookingAround        string   `dynamo:"OnboardLookingAround,omitempty" json:"OnboardLookingAround"`
	OurCodeCommand              string   `dynamo:"OurCodeCommand,omitempty" json:"OurCodeCommand"`
	OurCodeSlackResponse        string   `dynamo:"OurCodeSlackResponse,omitempty" json:"OurCodeSlackResponse"`
	ServeCommand                string   `dynamo:"ServeCommand,omitempty" json:"ServeCommand"`
	ServeTeam                   string   `dynamo:"ServeTeam,omitempty" json:"ServeTeam"`
	ServeTeamDescription        string   `dynamo:"ServeTeamDescription,omitempty" json:"ServeTeamDescription"`
	ServeDescriptionGrowthTrack string   `dynamo:"ServeDescriptionGrowthTrack,omitempty" json:"ServeDescriptionGrowthTrack"`
	ServeTeamName               string   `dynamo:"ServeTeamName,omitempty" json:"ServeTeamName"`
	SlackAppID                  string   `dynamo:"SlackAppID,omitempty" json:"SlackAppID"`
	SlackBotToken               string   `dynamo:"SlackBotToken,omitempty" json:"SlackBotToken"`
	SlackClientID               string   `dynamo:"SlackClientID,omitempty" json:"SlackClientID"`
	SlackClientSecret           string   `dynamo:"SlackClientSecret,omitempty" json:"SlackClientSecret"`
	SlackGeneralChannelID       string   `dynamo:"SlackGeneralChannelID,omitempty" json:"SlackGeneralChannelID"`
	SlackInviteToken            string   `dynamo:"SlackInviteToken,omitempty" json:"SlackInviteToken"`
	SlackOAuthToken             string   `dynamo:"SlackOAuthToken,omitempty" json:"SlackOAuthToken"`
	SlackSigningSecret          string   `dynamo:"SlackSigningSecret,omitempty" json:"SlackSigningSecret"`
	SlackTeamID                 string   `dynamo:"SlackTeamID,omitempty" json:"SlackTeamID"`
	SlackURL                    string   `dynamo:"SlackURL,omitempty" json:"SlackURL"`
	SlackVerificationToken      string   `dynamo:"SlackVerificationToken,omitempty" json:"SlackVerificationToken"`
	SlashRenee                  string   `dynamo:"SlashRenee,omitempty" json:"SlashRenee"`
	TimeZone                    string   `dynamo:"TimeZone,omitempty" json:"TimeZone"`
	UUID                        string   `dynamo:"UUID,omitempty" json:"UUID"`
	WelcomeSlackMessage         string   `dynamo:"WelcomeSlackMessage,omitempty" json:"WelcomeSlackMessage"`
}

type SpiritualGiftsData struct {
	Assessment []map[string]string `dynamo:"Assessment,omitempty" json:"assessment,omitempty"`
	Gifts      map[string]int      `dynamo:"Gifts,omitempty" json:"gifts,omitempty"`
}

// Baptism data of a Campus Member
type Baptism struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// Group groups a Campus Member belongs to
type Group struct {
	Groups []MemberGroupData `dynamo:"Groups,omitempty" json:"groups,omitempty"`
	Status bool              `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// MemberGroupData individual group a Campus Member belongs to
type MemberGroupData struct {
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	GroupName string `dynamo:"GroupName,omitempty" json:"groupName,omitempty"`
	UUID      string `dynamo:"UUID" json:"UUID,omitempty"`
}

// NewCreation data of a Campus Member
type NewCreation struct {
	FollowUp      bool           `dynamo:"FollowUp,omitempty" json:"followUp,omitempty"`
	FirstDecision NCFirstTime    `dynamo:"FirstDecision,omitempty" json:"firstDecision,omitempty"`
	Rededication  NCRededication `dynamo:"Rededication,omitempty" json:"rededication,omitempty"`
	Status        bool           `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// NCFirstTime data of a NewCreation
type NCFirstTime struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// NCRededication data of a NewCreation
type NCRededication struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// ThisIsHome data of a Campus Member
type ThisIsHome struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// DiscoverYourPurpose data of a Campus Member
type DiscoverYourPurpose struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// DrawData of a Campus Member
type DrawData struct {
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Type      string `dynamo:"Type,omitemtpy" json:"type,omitempty"`
}

// Groups data structure
type Groups struct {
	Account         string                       `dynamo:"Account,omitempty" json:"account,omitempty"`
	AccountName     string                       `dynamo:"AccountName,omitempty" json:"accountName,omitempty"`
	Active          bool                         `dynamo:"Active,omitempty" json:"active,omitempty"`
	AgeGroup        string                       `dynamo:"AgeGroup,omitempty" json:"ageGroup,omitempty"`
	Attendance      map[string][]GroupAttendance `dynamo:"Attendance,omitempty" json:"attendance,omitempty"`
	Campus          string                       `dynamo:"Campus,omitempty" json:"campus,omitempty"`
	CampusName      string                       `dynamo:"CampusName,omitempty" json:"campusName,omitempty"`
	Capacity        GroupCapacity                `dynamo:"Capacity,omitempty" json:"capacity,omitempty"`
	ConfirmResponse string                       `dynamo:"ConfirmResponse,omitempty" json:"confirmResponse,omitempty"`
	CreatedAt       int64                        `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Details         GroupDetails                 `dynamo:"Details,omitempty" json:"details,omitempty"`
	GroupName       string                       `dynamo:"GroupName,omitempty" json:"groupName,omitempty"`
	LastUpdated     int64                        `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Leaders         []Member                     `dynamo:"Leaders,omitempty" json:"leaders,omitempty"`
	UUID            string                       `dynamo:"UUID,hash,omitempty" json:"UUID,omitempty"`
	Oversights      []Member                     `dynamo:"Oversights,omitempty" json:"oversights,omitempty"`
	Perpetual       bool                         `dynamo:"Perpetual,omitempty" json:"perpetual,omitempty"`
	Private         bool                         `dynamo:"Private,omitempty" json:"private,omitempty"`
	Registrations   []Member                     `dynamo:"Registrations,omitempty" json:"registrations,omitempty"`
	RegistrationURL string                       `dynamo:"RegistrationURL,omitempty" json:"registrationURL,omitempty"`
	RequestFormLink string                       `dynamo:"RequestFormLink,omitempty" json:"requestFormLink,omitempty"`
	StudyKey        string                       `dynamo:"StudyKey,omitempty" json:"studyKey,omitempty"`
	Type            string                       `dynamo:"Type,omitempty" json:"type,omitempty"`
}

// GroupDetails of a Groups
type GroupDetails struct {
	ChannelID   string `dynamo:"ChannelID,omitempty" json:"channelID,omitempty"`
	Description string `dynamo:"Description,omitempty" json:"description,omitempty"`
	EndTime     int64  `dynamo:"EndTime,omitempty" json:"endTime,omitempty"`
	Location    string `dynamo:"Location,omitempty" json:"location,omitempty"`
	StartTime   int64  `dynamo:"StartTime,omitempty" json:"startTime,omitempty"`
	Tier        int    `dynamo:"Tier,omitempty" json:"tier,omitempty"`
}

// GroupCapacity of a Groups
type GroupCapacity struct {
	AtCapacity bool `dynamo:"AtCapacity,omitempty" json:"atCapacity,omitempty"`
	Limit      int  `dynamo:"Limit,omitempty" json:"limit,omitempty"`
	Openings   int  `dynamo:"Openings,omitempty" json:"openings,omitempty"`
	Unlimited  bool `dynamo:"Unlimited,omitempty" json:"unlimited,omitempty"`
}

// GroupAttendance of a Groups
type GroupAttendance struct {
	Account   string `dynamo:"Account,omitempty" json:"account,omitempty"`
	Campus    string `dynamo:"Campus,omitempty" json:"campus,omitempty"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	FullName  string `dynamo:"FullName,omitempty" json:"fullName,omitempty"`
	UUID      string `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
}

// NewCreationQR ...
type NewCreationQR struct {
	UUID      string `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
	Account   string `dynamo:"Account,omitempty" json:"account,omitempty"`
	Campus    string `dynamo:"Campus,omitempty" json:"campus,omitempty"`
	FirstName string `dynamo:"FirstName,omitempty" json:"firstName,omitempty"`
	LastName  string `dynamo:"LastName,omitempty" json:"lastName,omitempty"`
	FullName  string `dynamo:"FullName,omitempty" json:"fullName,omitempty"`
	SentTo    string `dynamo:"SentTo,omitempty" json:"sentTo,omitempty"`
	Status    string `dynamo:"Status,omitempty" json:"status,omitempty"`
	SentAt    int64  `dynamo:"SentAt,omitempty" json:"sentAt,omitempty"`
	ExpiredAt int64  `dynamo:"ExpiredAt,omitempty" json:"expiredAt,omitempty"`
}
