package shared

import "errors"

// TODOs: THERE ARE SEVEN (7) TODOs IN THIS FILE

// QUESTION: Could this file be moved to renee-data-api or go-renee and be
// imported from there to other projects? I (Noah) would like to define methods
// on these types, but since they are imported, Go won't let you define methods
// on them in other projects. I'm thinking of methods for group signups, for
// sending messages, etc.

// ------------------------------------------------------------------------------------------------------------------------------------------------------------
// METADATA ---------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------------

// Metadata CreatedAt, LastUpdated, and UUID
type Metadata struct {
	CreatedAt   int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	LastUpdated int64  `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	UUID        string `dynamo:"UUID,hash,omitempty" json:"UUID,omitempty"`
}

// AccountMetadata AccountID and AccountName
type AccountMetadata struct {
	// TODO: rename Account -> AccountID in DB
	AccountID   string `dynamo:"AccountID,omitempty" json:"accountID,omitempty"`
	AccountName string `dynamo:"AccountName,omitempty" json:"accountName,omitempty"`
}

// CampusMetadata CampusID and CampusName
type CampusMetadata struct {
	// TODO: rename Campus -> CampusID in DB
	CampusID   string `dynamo:"CampusID,omitempty" json:"campusID,omitempty"`
	CampusName string `dynamo:"CampusName,omitempty" json:"campusName,omitempty"`
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------
// ACCOUNT ----------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------------

// Account a tenant account
type Account struct {
	Metadata
	Name string `dynamo:"Name,omitempty" json:"name,omitempty"`
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------
// CAMPUS -----------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------------

// Campus a tenant account campus
type Campus struct {
	AccountMetadata
	Metadata

	// Campus Info
	Name                string   `dynamo:"Name,omitempty" json:"name,omitempty"`
	TimeZone            string   `dynamo:"TimeZone,omitempty" json:"timeZone,omitempty"`
	MarketingPlatforms  []string `dynamo:"MarketingPlatforms,omitempty" json:"marketingPlatforms,omitempty"`
	WelcomeSlackMessage string   `dynamo:"WelcomeSlackMessage,omitempty" json:"welcomeSlackMessage,omitempty"`

	// Clerk
	ClerkNumber      string `dynamo:"ClerkNumber,omitempty" json:"clerkNumber,omitempty"`
	GiveTrigger      string `dynamo:"GiveTrigger,omitempty" json:"giveTrigger,omitempty"`
	GiveTextResponse string `dynamo:"GiveTextResponse,omitempty" json:"giveTextResponse,omitempty"`
	HomeTrigger      string `dynamo:"HomeTrigger,omitempty" json:"homeTrigger,omitempty"`
	NCTrigger        string `dynamo:"NCTrigger,omitempty" json:"NCTrigger,omitempty"`

	// Slack Data
	SlackAppID             string `dynamo:"SlackAppID,omitempty" json:"slackAppID,omitempty"`
	SlackBotToken          string `dynamo:"SlackBotToken,omitempty" json:"slackBotToken,omitempty"`
	SlackClientID          string `dynamo:"SlackClientID,omitempty" json:"slackClientID,omitempty"`
	SlackClientSecret      string `dynamo:"SlackClientSecret,omitempty" json:"slackClientSecret,omitempty"`
	SlackGeneralChannelID  string `dynamo:"SlackGeneralChannelID,omitempty" json:"slackGeneralChannelID,omitempty"`
	SlackInviteToken       string `dynamo:"SlackInviteToken,omitempty" json:"slackInviteToken,omitempty"`
	SlackOAuthToken        string `dynamo:"SlackOAuthToken,omitempty" json:"slackOAuthToken,omitempty"`
	SlackSigningSecret     string `dynamo:"SlackSigningSecret,omitempty" json:"slackSigningSecret,omitempty"`
	SlackTeamID            string `dynamo:"SlackTeamID,omitempty" json:"slackTeamID,omitempty"`
	SlackURL               string `dynamo:"SlackURL,omitempty" json:"slackURL,omitempty"`
	SlackVerificationToken string `dynamo:"SlackVerificationToken,omitempty" json:"slackVerificationToken,omitempty"`

	// Notification Channels
	NotificationChannelBibleClaim            string `dynamo:"NotificationChannelBibleClaim,omitempty" json:"notificationChannelBibleClaim,omitempty"`
	NotificationChannelEvents                string `dynamo:"NotificationChannelEvents,omitempty" json:"notificationChannelEvents,omitempty"`
	NotificationChannelGrowGroups            string `dynamo:"NotificationChannelGrowGroups,omitempty" json:"notificationChannelGrowGroups,omitempty"`
	NotificationChannelGrowthTrack           string `dynamo:"NotificationChannelGrowthTrack,omitempty" json:"notificationChannelGrowthTrack,omitempty"`
	NotificationChannelGrowthTrackCompletion string `dynamo:"NotificationChannelGrowthTrackCompletion,omitempty" json:"notificationChannelGrowthTrackCompletion,omitempty"`
	NotificationChannelInterestGroups        string `dynamo:"NotificationChannelInterestGroups,omitempty" json:"notificationChannelInterestGroups,omitempty"`
	NotificationChannelNewCreation           string `dynamo:"NotificationChannelNewCreation,omitempty" json:"notificationChannelNewCreation,omitempty"`
	NotificationChannelNewPerson             string `dynamo:"NotificationChannelNewPerson,omitempty" json:"notificationChannelNewPerson,omitempty"`
	NotificationChannelNewPersonRenee        string `dynamo:"NotificationChannelNewPersonRenee,omitempty" json:"notificationChannelNewPersonRenee,omitempty"`
	NotificationChannelNewPersonContact      string `dynamo:"NotificationChannelNewPersonContact,omitempty" json:"notificationChannelNewPersonContact,omitempty"`
	NotificationChannelServeTeams            string `dynamo:"NotificationChannelServeTeams,omitempty" json:"notificationChannelServeTeams,omitempty"`

	// Onboarding FLow
	OnboardChatWithTeam  string `dynamo:"OnboardChatWithTeam,omitempty" json:"onboardChatWithTeam,omitempty"`
	OnboardGrowGroups    string `dynamo:"OnboardGrowGroups,omitempty" json:"onboardGrowGroups,omitempty"`
	OnboardGrowthTrack   string `dynamo:"OnboardGrowthTrack,omitempty" json:"onboardGrowthTrack,omitempty"`
	OnboardLookingAround string `dynamo:"OnboardLookingAround,omitempty" json:"onboardLookingAround,omitempty"`

	// Slash Commands and Responses
	DigitalChurchCommand       string `dynamo:"DigitalChurchCommand,omitempty" json:"digitalChurchCommand,omitempty"`
	DigitalChurchSlackResponse string `dynamo:"DigitalChurchSlackResponse,omitempty" json:"digitalChurchSlackResponse,omitempty"`
	EventCommand               string `dynamo:"EventsCommand,omitempty" json:"eventsCommand,omitempty"`
	GrowthTrackCommand         string `dynamo:"GrowthTrackCommand,omitempty" json:"growthTrackCommand,omitempty"`
	GiveCommand                string `dynamo:"GiveCommand,omitempty" json:"giveCommand,omitempty"`
	GiveSlackResponse          string `dynamo:"GiveSlackResponse,omitempty" json:"giveSlackResponse,omitempty"`
	GrowthTrackW1Command       string `dynamo:"GrowthTrackW1Command,omitempty" json:"growthTrackW1Command,omitempty"`
	GrowthTrackW2Command       string `dynamo:"GrowthTrackW2Command,omitempty" json:"growthTrackW2Command,omitempty"`
	GrowGroupCommand           string `dynamo:"GrowGroupCommand,omitempty" json:"growGroupCommand,omitempty"`
	InterestGroupCommand       string `dynamo:"InterestGroupCommand,omitempty" json:"interestGroupCommand,omitempty"`
	OurCodeCommand             string `dynamo:"OurCodeCommand,omitempty" json:"ourCodeCommand,omitempty"`
	OurCodeSlackResponse       string `dynamo:"OurCodeSlackResponse,omitempty" json:"ourCodeSlackResponse,omitempty"`
	ServeCommand               string `dynamo:"ServeCommand,omitempty" json:"serveCommand,omitempty"`
	SlashRenee                 string `dynamo:"SlashRenee,omitempty" json:"slashRenee,omitempty"`

	// Team IDs
	EventsTeam        string `dynamo:"EventsTeam,omitempty" json:"eventsTeam,omitempty"`
	GrowthTrackTeam   string `dynamo:"GrowthTrackTeam,omitempty" json:"growthTrackTeam,omitempty"`
	GrowGroupTeam     string `dynamo:"GrowGroupTeam,omitempty" json:"growGroupTeam,omitempty"`
	InterestGroupTeam string `dynamo:"InterestGroupTeam,omitempty" json:"interestGroupTeam,omitempty"`
	HealthTeam        string `dynamo:"HealthTeam,omitempty" json:"healthTeam,omitempty"`
	NewCreationsTeam  string `dynamo:"NewCreationsTeam,omitempty" json:"newCreationsTeam,omitempty"`
	ServeTeam         string `dynamo:"ServeTeam,omitempty" json:"serveTeam,omitempty"`

	// Growth Track
	GrowthTrackW1Redundant string `dynamo:"GrowthTrackW1Redundant,omitempty" json:"growthTrackW1Redundant,omitempty"` // Message to send when someone has already completed week 1 of Growth Track
	GrowthTrackName        string `dynamo:"GrowthTrackName,omitempty" json:"growthTrackName,omitempty"`
	GrowthTrackW1Name      string `dynamo:"GrowthTrackW1Name,omitempty" json:"growthTrackW1Name,omitempty"`
	GrowthTrackW2Name      string `dynamo:"GrowthTrackW2Name,omitempty" json:"growthTrackW2Name,omitempty"`

	// Grow Group
	GrowGroupDescription  string `dynamo:"GrowGroupDescription,omitempty" json:"growGroupDescription,omitempty"`
	GrowGroupName         string `dynamo:"GrowGroupName,omitempty" json:"growGroupName,omitempty"`
	GrowGroupStudyKeyWord string `dynamo:"GrowGroupStudyKeyWord,omitempty" json:"growGroupStudyKeyWord,omitempty"`
	GrowGroupTiers        []int  `dynamo:"GrowGroupTiers,omitempty" json:"growGroupTiers,omitempty"`
	CurrentGrowGroupTier  []int  `dynamo:"CurrentGrowGroupTier,omitempty" json:"currentGrowGroupTier,omitempty"`

	// Interest Group
	InterestGroupDescription string `dynamo:"InterestGroupDescription,omitempty" json:"interestGroupDescription,omitempty"`
	InterestGroupName        string `dynamo:"InterestGroupName,omitempty" json:"interestGroupName,omitempty"`

	// New Creations
	NewCreationName        string `dynamo:"NewCreationName,omitempty" json:"newCreationName,omitempty"`
	NewCreationsDinnerName string `dynamo:"NewCreationsDinnerName,omitempty" json:"newCreationsDinnerName,omitempty"`

	// Serve Teams
	ServeTeamDescription        string `dynamo:"ServeTeamDescription,omitempty" json:"serveTeamDescription,omitempty"`
	ServeDescriptionGrowthTrack string `dynamo:"ServeDescriptionGrowthTrack,omitempty" json:"serveDescriptionGrowthTrack,omitempty"`
	ServeTeamName               string `dynamo:"ServeTeamName,omitempty" json:"serveTeamName,omitempty"`
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------
// MEMBER -----------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------------

// Member a tenant account campus member
type Member struct {
	AccountMetadata
	CampusMetadata
	Metadata

	// Personal Data
	Birthday    string `dynamo:"Birthday,omitempty" json:"birthday,omitempty"`
	Email       string `dynamo:"Email,omitempty" json:"email,omitemptyv"`
	FirstName   string `dynamo:"FirstName,omitempty" json:"firstName,omitempty"`
	FullName    string `dynamo:"FullName,omitempty" json:"fullName,omitempty"`
	Gender      string `dynamo:"Gender,omitempty" json:"gender,omitempty"`
	LastName    string `dynamo:"LastName,omitempty" json:"lastName,omitempty"`
	PhoneNumber string `dynamo:"PhoneNumber,omitempty" json:"phoneNumber,omitempty"`

	// Organizational Data
	AgeGroup       string `dynamo:"AgeGroup,omitempty" json:"ageGroup,omitempty"`
	MembershipType string `dynamo:"MembershipType,omitempty" json:"membershipType,omitempty"`
	ReceiveEmail   bool   `dynamo:"ReceiveEmail,omitempty" json:"receiveEmail,omitempty"`
	ShirtSize      string `dynamo:"ShirtSize,omitempty" json:"shirtSize,omitempty"`
	SlackID        string `dynamo:"SlackID,omitempty" json:"slackID,omitempty"`
	SlackChannelID string `dynamo:"SlackChannelID,omitempty" json:"slackChannelID,omitempty"`
	SlackOnboard   int64  `dynamo:"SlackOnboard,omitempty" json:"slackOnboard,omitempty"`

	// Group Data
	// TODO: Append "Data" onto the names of these fields in the DB
	// and rename the TIH and DYP fields to match what is defined below.
	BaptismData        StatusData         `dynamo:"BaptismData,omitempty" json:"baptismData,omitempty"`
	DrawData           DrawData           `dynamo:"DrawData,omitempty" json:"drawData,omitempty"`
	EventsData         GroupData          `dynamo:"EventsData,omitempty" json:"eventsData,omitempty"`
	GrowGroupData      GroupData          `dynamo:"GrowGroupData,omitempty" json:"growGroupData,omitempty"`
	GrowthTrackW1Data  StatusData         `dynamo:"GrowthTrackW1Data,omitempty" json:"growthTrackW1Data,omitempty"`
	GrowthTrackW2Data  StatusData         `dynamo:"GrowthTrackW2Data,omitempty" json:"growthTrackW2Data,omitempty"`
	InterestGroupData  GroupData          `dynamo:"InterestGroupData,omitempty" json:"interestGroupData,omitempty"`
	NewCreationData    NewCreationData    `dynamo:"NewCreationData,omitempty" json:"newCreationData,omitempty"`
	SpiritualGiftsData SpiritualGiftsData `dynamo:"SpiritualGiftsData,omitempty" json:"spiritualGiftsData,omitempty"`
	VolunteerData      GroupData          `dynamo:"VolunteerData,omitempty" json:"volunteerData,omitempty"`
}

// ----------------------------------------------------------------------------
// MEMBER COMPOSITIONAL TYPES -------------------------------------------------
// ----------------------------------------------------------------------------

// GroupData the MemberGroupData for each Group a Member has signed up for
type GroupData struct {
	Groups []MemberGroupData `dynamo:"Groups,omitempty" json:"groups,omitempty"`
	Status bool              `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// IncentiveClaimData for the qr code for an incentive of a new creation
type IncentiveClaimData struct {
	Active    bool   `dynamo:"Active,omitempty" json:"active,omitempty"`
	SentAt    int64  `dynamo:"SentAt,omitempty" json:"sentAt,omitempty"`
	ExpiredAt int64  `dynamo:"ExpiredAt,omitempty" json:"expiredAt,omitempty"`
	QRCodeURL string `dynamo:"QRCodeURL,omitempty" json:"qrCodeURL,omitempty"`
}

// MemberGroupData the time (CreatedAt) a Member joined a Group, the GroupName,
// and Group UUID
type MemberGroupData struct {
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	GroupName string `dynamo:"GroupName,omitempty" json:"groupName,omitempty"`
	UUID      string `dynamo:"UUID" json:"UUID,omitempty"`
}

// NewCreationData Status indicating if a Member is a New Creation, whether they wish to
// be followed up with, whether they are making this decision for the first time or
// rededicating their life, and whether they have claimed their incentive
type NewCreationData struct {
	FollowUp       bool               `dynamo:"FollowUp,omitempty" json:"followUp,omitempty"`
	FirstDecision  StatusData         `dynamo:"FirstDecision,omitempty" json:"firstDecision,omitempty"`
	Rededication   StatusData         `dynamo:"Rededication,omitempty" json:"rededication,omitempty"`
	Status         bool               `dynamo:"Status,omitempty" json:"status,omitempty"`
	IncentiveClaim IncentiveClaimData `dynamo:"IncentiveClaim,omitempty" json:"incentiveClaim,omitempty"`
}

// SpiritualGiftsData a Member's Spiritual Gifts Assessment answers
// along with the resulting calculated Gifts
type SpiritualGiftsData struct {
	Assessment []map[string]string `dynamo:"Assessment,omitempty" json:"assessment,omitempty"`
	Gifts      map[string]int      `dynamo:"Gifts,omitempty" json:"gifts,omitempty"`
}

// StatusData holds the time the status was CreatedAt and the bool value of the Status
type StatusData struct {
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Status    bool  `dynamo:"Status,omitempty" json:"status,omitempty"`
}

// DrawData of a Campus Member (TODO: Document this object in detail)
type DrawData struct {
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Type      string `dynamo:"Type,omitemtpy" json:"type,omitempty"`
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------
// GROUP ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------------

// Group a tenant account campus group
type Group struct {
	AccountMetadata
	CampusMetadata
	Metadata

	// Root Data
	// TODO: Capacity -> CapacityData in the DB
	GroupName       string       `dynamo:"GroupName,omitempty" json:"groupName,omitempty"`
	Active          bool         `dynamo:"Active,omitempty" json:"active,omitempty"`
	Private         bool         `dynamo:"Private,omitempty" json:"private,omitempty"`
	Details         Details      `dynamo:"Details,omitempty" json:"details,omitempty"`
	CapacityData    CapacityData `dynamo:"CapacityData,omitempty" json:"capacitydata,omitempty"`
	AgeGroup        string       `dynamo:"AgeGroup,omitempty" json:"ageGroup,omitempty"`
	ConfirmResponse string       `dynamo:"ConfirmResponse,omitempty" json:"confirmResponse,omitempty"`

	// Member Data
	Leaders       []GroupMemberData            `dynamo:"Leaders,omitempty" json:"leaders,omitempty"`
	Oversights    []GroupMemberData            `dynamo:"Oversights,omitempty" json:"oversights,omitempty"`
	Registrations []GroupMemberData            `dynamo:"Registrations,omitempty" json:"registrations,omitempty"`
	Attendance    map[string][]GroupMemberData `dynamo:"Attendance,omitempty" json:"attendance,omitempty"`
}

// Event a Group that has the Type "Event"
type Event struct {
	Group
	Type            string `default:"Event" dynamo:"Type,omitempty" json:"type,omitempty"`
	RegistrationURL string `dynamo:"RegistrationURL,omitempty" json:"registrationURL,omitempty"`
	Perpetual       bool   `dynamo:"Perpetual,omitempty" json:"perpetual,omitempty"`
}

// GrowGroup a Group that has the Type "Grow Group"
type GrowGroup struct {
	Group
	StudyKey string `dynamo:"StudyKey,omitempty" json:"studyKey,omitempty"`
	Tier     int    `dynamo:"Tier,omitempty" json:"tier,omitempty"`
	Type     string `default:"Grow Group" dynamo:"Type,omitempty" json:"type,omitempty"`
}

// GrowthTrack a Group that has the Type "Growth Track"
type GrowthTrack struct {
	Group
	Type string `default:"Growth Track" dynamo:"Type,omitempty" json:"type,omitempty"`
}

// InterestGroup a Group that has the Type "Interest Group"
type InterestGroup struct {
	Group
	Type string `default:"Interest Group" dynamo:"Type,omitempty" json:"type,omitempty"`
}

// ServeTeam a Group that has the Type "Serve Team"
type ServeTeam struct {
	Group
	BypassGrowthTrack bool   `dynamo:"BypassGrowthTrack,omitempty" json:"bypassGrowthTrack,omitempty"`
	RequestFormLink   string `dynamo:"RequestFormLink,omitempty" json:"requestFormLink,omitempty"`
	Type              string `default:"Serve Team" dynamo:"Type,omitempty" json:"type,omitempty"`
}

// ----------------------------------------------------------------------------
// GROUP CONSTS ---------------------------------------------------------------
// ----------------------------------------------------------------------------

// GroupType a type of Group (to be assigned to the Type field)
type GroupType string

const (
	// GroupTypeBaptism baptism group type
	GroupTypeBaptism GroupType = "Baptism"
	// GroupTypeEvent event group type
	GroupTypeEvent = "Event"
	// GroupTypeInterestGroup interest group group type
	GroupTypeInterestGroup = "Interest Group"
	// GroupTypeGrowGroup grow group group type
	GroupTypeGrowGroup = "Grow Group"
	// GroupTypeGrowthTrack growth track group type
	GroupTypeGrowthTrack = "Growth Track"
	// GroupTypeNewCreation new creation group type
	GroupTypeNewCreation = "New Creation Dinner"
	// GroupTypeServeTeam serve team group type
	GroupTypeServeTeam = "Serve Team"
)

// IsValid will return an error if gt is not a valid GroupType
func (gt GroupType) IsValid() error {
	switch gt {
	case GroupTypeBaptism, GroupTypeEvent,
		GroupTypeInterestGroup, GroupTypeGrowGroup,
		GroupTypeGrowthTrack, GroupTypeNewCreation,
		GroupTypeServeTeam:
		return nil
	}
	return errors.New("Invalid group type")
}

// TODO: Use these GroupTypes when constructing any Group and assigning its
// type and use IsValid in the API to check incoming group types.

// ----------------------------------------------------------------------------
// GROUP COMPOSITIONAL TYPES --------------------------------------------------
// ----------------------------------------------------------------------------

// CapacityData of a Group
type CapacityData struct {
	AtCapacity bool `dynamo:"AtCapacity,omitempty" json:"atCapacity,omitempty"`
	Limit      int  `dynamo:"Limit,omitempty" json:"limit,omitempty"`
	Openings   int  `dynamo:"Openings,omitempty" json:"openings,omitempty"`
	Unlimited  bool `dynamo:"Unlimited,omitempty" json:"unlimited,omitempty"`
}

// Details of a Group
type Details struct {
	ChannelID   string   `dynamo:"ChannelID,omitempty" json:"channelID,omitempty"`
	Description string   `dynamo:"Description,omitempty" json:"description,omitempty"`
	EndTime     int64    `dynamo:"EndTime,omitempty" json:"endTime,omitempty"`
	Location    Location `dynamo:"Location,omitempty" json:"location,omitempty"`
	StartTime   int64    `dynamo:"StartTime,omitempty" json:"startTime,omitempty"`
	// TODO: move tier to root object for grow groups in DB
	// Tier int `dynamo:"Tier,omitempty" json:"tier,omitempty"`
}

// GroupMemberData represents a member of a Group
type GroupMemberData struct {
	AccountID string `dynamo:"Account,omitempty" json:"account,omitempty"`
	CampusID  string `dynamo:"Campus,omitempty" json:"campus,omitempty"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	FullName  string `dynamo:"FullName,omitempty" json:"fullName,omitempty"`
	SlackID   string `dynamo:"SlackID,omitempty" json:"slackID,omitempty"`
}

// Location of a Group
type Location struct {
	StreetAddress   string `dynamo:"StreetAddress,omitempty" json:"startTime,omitempty"`
	ExtendedAddress string `dynamo:"ExtendedAddress,omitempty" json:"extendedAddress,omitempty"`
	Locality        string `dynamo:"Locality,omitempty" json:"locality,omitempty"`
	PostalCode      string `dynamo:"PostalCode,omitempty" json:"postalCode,omitempty"`
	Region          string `dynamo:"Region,omitempty" json:"region,omitempty"`
	Country         string `dynamo:"Country,omitempty" json:"country,omitempty"`
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------
// MISC/UNUSED ------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------------

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
