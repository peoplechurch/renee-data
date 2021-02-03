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
	GrowthTrack         Group               `dynamo:"GrowthTrack,omitempty" json:"growthTrack,omitempty"`
	LastName            string              `dynamo:"LastName,omitempty" json:"lastName,omitempty"`
	LastUpdated         int64               `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	MembershipType      string              `dynamo:"MembershipType,omitempty" json:"membershipType,omitempty"`
	NewCreation         NewCreation         `dynamo:"NewCreation,omitempty" json:"newCreation,omitempty"`
	PhoneNumber         string              `dynamo:"PhoneNumber,omitempty" json:"phoneNumber,omitempty"`
	ReceiveEmail        bool                `dynamo:"ReceiveEmail,omitempty" json:"receiveEmail,omitempty"`
	ShirtSize           string              `dynamo:"ShirtSize,omitempty" json:"shirtSize,omitempty"`
	SlackID             string              `dynamo:"SlackID,omitempty" json:"slackID,omitempty"`
	SlackChannelID      string              `dynamo:"SlackChannelID,omitempty" json:"slackChannelID,omitempty"`
	SlackOnboard        int64               `dynamo:"SlackOnboard,omitempty" json:"slackOnboard,omitempty"`
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
	Account                                  string   `dynamo:"Account,omitempty" json:"account,omitempty"`
	AccountName                              string   `dynamo:"AccountName,omitempty" json:"accountName,omitempty"`
	ClerkNumber                              string   `dynamo:"ClerkNumber,omitempty" json:"clerkNumber,omitempty"`
	CurrentGrowGroupTier                     []int    `dynamo:"CurrentGrowGroupTier,omitempty" json:"currentGrowGroupTier,omitempty"`
	CreatedAt                                int64    `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	DigitalChurchSlackResponse               string   `dynamo:"DigitalChurchSlackResponse,omitempty" json:"digitalChurchSlackResponse,omitempty"`
	DigitalChurchCommand                     string   `dynamo:"DigitalChurchCommand,omitempty" json:"digitalChurchCommand,omitempty"`
	EventCommand                             string   `dynamo:"EventsCommand,omitempty" json:"eventsCommand,omitempty"`
	GiveCommand                              string   `dynamo:"GiveCommand,omitempty" json:"giveCommand,omitempty"`
	GiveTextResponse                         string   `dynamo:"GiveTextResponse,omitempty" json:"giveTextResponse,omitempty"`
	GiveTrigger                              string   `dynamo:"GiveTrigger,omitempty" json:"giveTrigger,omitempty"`
	GiveSlackResponse                        string   `dynamo:"GiveSlackResponse,omitempty" json:"giveSlackResponse,omitempty"`
	GrowthTrackCommand                       string   `dynamo:"GrowthTrackCommand,omitempty" json:"growthTrackCommand,omitempty"`
	GrowthTrackName                          string   `dynamo:"GrowthTrackName,omitempty" json:"growthTrackName,omitempty"`
	GrowthTrackW1Name                        string   `dynamo:"GrowthTrackW1Name,omitempty" json:"growthTrackW1Name,omitempty"`
	GrowthTrackW2Name                        string   `dynamo:"GrowthTrackW2Name,omitempty" json:"growthTrackW2Name,omitempty"`
	GrowthTrackW1Command                     string   `dynamo:"GrowthTrackW1Command,omitempty" json:"growthTrackW1Command,omitempty"`
	GrowthTrackW2Command                     string   `dynamo:"GrowthTrackW2Command,omitempty" json:"growthTrackW2Command,omitempty"`
	GrowthTrackW1Redundant                   string   `dynamo:"GrowthTrackW1Redundant,omitempty" json:"growthTrackW1Redundant,omitempty"` // Message to send when someone has already completed week 1 of Growth Track
	GrowGroupCommand                         string   `dynamo:"GrowGroupCommand,omitempty" json:"growGroupCommand,omitempty"`
	GrowGroupDescription                     string   `dynamo:"GrowGroupDescription,omitempty" json:"growGroupDescription,omitempty"`
	GrowGroupName                            string   `dynamo:"GrowGroupName,omitempty" json:"growGroupName,omitempty"`
	GrowGroupStudyKeyWord                    string   `dynamo:"GrowGroupStudyKeyWord,omitempty" json:"growGroupStudyKeyWord,omitempty"`
	GrowGroupTiers                           []int    `dynamo:"GrowGroupTiers,omitempty" json:"growGroupTiers,omitempty"`
	HomeTrigger                              string   `dynamo:"HomeTrigger,omitempty" json:"homeTrigger,omitempty"`
	InterestGroupCommand                     string   `dynamo:"InterestGroupCommand,omitempty" json:"interestGroupCommand,omitempty"`
	InterestGroupDescription                 string   `dynamo:"InterestGroupDescription,omitempty" json:"interestGroupDescription,omitempty"`
	InterestGroupName                        string   `dynamo:"InterestGroupName,omitempty" json:"interestGroupName,omitempty"`
	LastUpdated                              int64    `dynamo:"LastUpdated,omitempty" json:"lastUpdated"`
	Name                                     string   `dynamo:"Name,omitempty" json:"name,omitempty"`
	NCTrigger                                string   `dynamo:"NCTrigger,omitempty" json:"NCTrigger,omitempty"`
	NewCreationName                          string   `dynamo:"NewCreationName,omitempty" json:"newCreationName,omitempty"`
	NewCreationsDinnerName                   string   `dynamo:"NewCreationsDinnerName,omitempty" json:"newCreationsDinnerName,omitempty"`
	NotificationChannelAccountAlerts         string   `dynamo:"NotificationChannelAccountAlerts,omitempty" json:"notificationChannelAccountAlerts,omitempty"`
	NotificationChannelBibleClaim            string   `dynamo:"NotificationChannelBibleClaim,omitempty" json:"notificationChannelBibleClaim,omitempty"`
	NotificationChannelEvents                string   `dynamo:"NotificationChannelEvents,omitempty" json:"notificationChannelEvents,omitempty"`
	NotificationChannelGrowGroups            string   `dynamo:"NotificationChannelGrowGroups,omitempty" json:"notificationChannelGrowGroups,omitempty"`
	NotificationChannelGrowthTrack           string   `dynamo:"NotificationChannelGrowthTrack,omitempty" json:"notificationChannelGrowthTrack,omitempty"`
	NotificationChannelGrowthTrackCompletion string   `dynamo:"NotificationChannelGrowthTrackCompletion,omitempty" json:"notificationChannelGrowthTrackCompletion,omitempty"`
	NotificationChannelInterestGroups        string   `dynamo:"NotificationChannelInterestGroups,omitempty" json:"notificationChannelInterestGroups,omitempty"`
	NotificationChannelNewCreation           string   `dynamo:"NotificationChannelNewCreation,omitempty" json:"notificationChannelNewCreation,omitempty"`
	NotificationChannelNewPerson             string   `dynamo:"NotificationChannelNewPerson,omitempty" json:"notificationChannelNewPerson,omitempty"`
	NotificationChannelNewPersonRenee        string   `dynamo:"NotificationChannelNewPersonRenee,omitempty" json:"notificationChannelNewPersonRenee,omitempty"`
	NotificationChannelNewPersonContact      string   `dynamo:"NotificationChannelNewPersonContact,omitempty" json:"notificationChannelNewPersonContact,omitempty"`
	NotificationChannelServeTeams            string   `dynamo:"NotificationChannelServeTeams,omitempty" json:"notificationChannelServeTeams,omitempty"`
	MailchimpAudienceID                      string   `dynamo:"MailchimpAudienceID,omitempty" json:"mailchimpAudienceID,omitempty"`
	MailchimpAPIKey                          string   `dynamo:"MailchimpAPIKey,omitempty" json:"mailchimpAPIKey,omitempty"`
	MarketingPlatforms                       []string `dynamo:"MarketingPlatforms,omitempty" json:"marketingPlatforms,omitempty"`
	OnboardChatWithTeam                      string   `dynamo:"OnboardChatWithTeam,omitempty" json:"onboardChatWithTeam,omitempty"`
	OnboardGrowGroups                        string   `dynamo:"OnboardGrowGroups,omitempty" json:"onboardGrowGroups,omitempty"`
	OnboardGrowthTrack                       string   `dynamo:"OnboardGrowthTrack,omitempty" json:"onboardGrowthTrack,omitempty"`
	OnboardLookingAround                     string   `dynamo:"OnboardLookingAround,omitempty" json:"onboardLookingAround,omitempty"`
	OurCodeCommand                           string   `dynamo:"OurCodeCommand,omitempty" json:"ourCodeCommand,omitempty"`
	OurCodeSlackResponse                     string   `dynamo:"OurCodeSlackResponse,omitempty" json:"ourCodeSlackResponse,omitempty"`
	ServeCommand                             string   `dynamo:"ServeCommand,omitempty" json:"serveCommand,omitempty"`
	ReporterNewPeopleTeam                    string   `dynamo:"ReporterNewPeopleTeam,omitempty" json:"reporterNewPeopleTeam,omitempty"`
	ReporterNewCreationsTeam                 string   `dynamo:"ReporterNewCreationsTeam,omitempty" json:"reporterNewCreationsTeam,omitempty"`
	ReporterGrowthTrackTeam                  string   `dynamo:"ReporterGrowthTrackTeam,omitempty" json:"reporterGrowthTrackTeam,omitempty"`
	ReporterVolunteersTeam                   string   `dynamo:"ReporterVolunteersTeam,omitempty" json:"reporterVolunteersTeam,omitempty"`
	ReporterGrowGroupsTeam                   string   `dynamo:"ReporterGrowGroupsTeam,omitempty" json:"reporterGrowGroupsTeam,omitempty"`
	ReporterInterestGroupsTeam               string   `dynamo:"ReporterInterestGroupsTeam,omitempty" json:"reporterInterestGroupsTeam,omitempty"`
	ReporterEventsTeam                       string   `dynamo:"ReporterEventsTeam,omitempty" json:"reporterEventsTeam,omitempty"`
	ReporterMarketingTeam                    string   `dynamo:"ReporterMarketingTeam,omitempty" json:"reporterMarketingTeam,omitempty"`
	ServeTeamDescription                     string   `dynamo:"ServeTeamDescription,omitempty" json:"serveTeamDescription,omitempty"`
	ServeDescriptionGrowthTrack              string   `dynamo:"ServeDescriptionGrowthTrack,omitempty" json:"serveDescriptionGrowthTrack,omitempty"`
	ServeTeamName                            string   `dynamo:"ServeTeamName,omitempty" json:"serveTeamName,omitempty"`
	SlackAppID                               string   `dynamo:"SlackAppID,omitempty" json:"slackAppID,omitempty"`
	SlackBotToken                            string   `dynamo:"SlackBotToken,omitempty" json:"slackBotToken,omitempty"`
	SlackClientID                            string   `dynamo:"SlackClientID,omitempty" json:"slackClientID,omitempty"`
	SlackClientSecret                        string   `dynamo:"SlackClientSecret,omitempty" json:"slackClientSecret,omitempty"`
	SlackGeneralChannelID                    string   `dynamo:"SlackGeneralChannelID,omitempty" json:"slackGeneralChannelID,omitempty"`
	SlackInviteToken                         string   `dynamo:"SlackInviteToken,omitempty" json:"slackInviteToken,omitempty"`
	SlackOAuthToken                          string   `dynamo:"SlackOAuthToken,omitempty" json:"slackOAuthToken,omitempty"`
	SlackSigningSecret                       string   `dynamo:"SlackSigningSecret,omitempty" json:"slackSigningSecret,omitempty"`
	SlackTeamID                              string   `dynamo:"SlackTeamID,omitempty" json:"slackTeamID,omitempty"`
	SlackURL                                 string   `dynamo:"SlackURL,omitempty" json:"slackURL,omitempty"`
	SlackVerificationToken                   string   `dynamo:"SlackVerificationToken,omitempty" json:"slackVerificationToken,omitempty"`
	SlashRenee                               string   `dynamo:"SlashRenee,omitempty" json:"slashRenee,omitempty"`
	TimeZone                                 string   `dynamo:"TimeZone,omitempty" json:"timeZone,omitempty"`
	UUID                                     string   `dynamo:"UUID,hash,omitempty" json:"UUID"`
	WelcomeSlackMessage                      string   `dynamo:"WelcomeSlackMessage,omitempty" json:"welcomeSlackMessage,omitempty"`
}

// SpiritualGiftsData of a Campus Member
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
	FollowUp       bool           `dynamo:"FollowUp,omitempty" json:"followUp,omitempty"`
	FirstDecision  NCFirstTime    `dynamo:"FirstDecision,omitempty" json:"firstDecision,omitempty"`
	Rededication   NCRededication `dynamo:"Rededication,omitempty" json:"rededication,omitempty"`
	Status         bool           `dynamo:"Status,omitempty" json:"status,omitempty"`
	IncentiveClaim IncentiveClaim `dynamo:"IncentiveClaim,omitempty" json:"incentiveClaim,omitempty"`
}

// IncentiveClaim is the data structure for the data involving the qr code for an incentive of a new creation
type IncentiveClaim struct {
	Active    bool   `dynamo:"Active,omitempty" json:"active,omitempty"`
	SentAt    int64  `dynamo:"SentAt,omitempty" json:"sentAt,omitempty"`
	ExpiredAt int64  `dynamo:"ExpiredAt,omitempty" json:"expiredAt,omitempty"`
	QRCodeURL string `dynamo:"QRCodeURL,omitempty" json:"qrCodeURL,omitempty"`
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
	Account           string                       `dynamo:"Account,omitempty" json:"account,omitempty"`
	AccountName       string                       `dynamo:"AccountName,omitempty" json:"accountName,omitempty"`
	Active            bool                         `dynamo:"Active,omitempty" json:"active,omitempty"`
	AgeGroup          string                       `dynamo:"AgeGroup,omitempty" json:"ageGroup,omitempty"`
	Attendance        map[string][]GroupAttendance `dynamo:"Attendance,omitempty" json:"attendance,omitempty"`
	Campus            string                       `dynamo:"Campus,omitempty" json:"campus,omitempty"`
	CampusName        string                       `dynamo:"CampusName,omitempty" json:"campusName,omitempty"`
	Capacity          GroupCapacity                `dynamo:"Capacity,omitempty" json:"capacity,omitempty"`
	ConfirmResponse   string                       `dynamo:"ConfirmResponse,omitempty" json:"confirmResponse,omitempty"`
	CreatedAt         int64                        `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	Details           GroupDetails                 `dynamo:"Details,omitempty" json:"details,omitempty"`
	GroupName         string                       `dynamo:"GroupName,omitempty" json:"groupName,omitempty"`
	BypassGrowthTrack bool                         `dynamo:"BypassGrowthTrack,omitempty" json:"bypassGrowthTrack,omitempty"`
	LastUpdated       int64                        `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Leaders           []Member                     `dynamo:"Leaders,omitempty" json:"leaders,omitempty"`
	UUID              string                       `dynamo:"UUID,hash,omitempty" json:"UUID,omitempty"`
	Oversights        []Member                     `dynamo:"Oversights,omitempty" json:"oversights,omitempty"`
	Perpetual         bool                         `dynamo:"Perpetual,omitempty" json:"perpetual,omitempty"`
	Private           bool                         `dynamo:"Private,omitempty" json:"private,omitempty"`
	Registrations     []Member                     `dynamo:"Registrations,omitempty" json:"registrations,omitempty"`
	RegistrationURL   string                       `dynamo:"RegistrationURL,omitempty" json:"registrationURL,omitempty"`
	RequestFormLink   string                       `dynamo:"RequestFormLink,omitempty" json:"requestFormLink,omitempty"`
	StudyKey          string                       `dynamo:"StudyKey,omitempty" json:"studyKey,omitempty"`
	Type              string                       `dynamo:"Type,omitempty" json:"type,omitempty"`
}

// GroupDetails of a Groups
type GroupDetails struct {
	ChannelID       string   `dynamo:"ChannelID,omitempty" json:"channelID,omitempty"`
	Description     string   `dynamo:"Description,omitempty" json:"description,omitempty"`
	EndTime         int64    `dynamo:"EndTime,omitempty" json:"endTime,omitempty"`
	Location        Location `dynamo:"Location,omitempty" json:"location,omitempty"`
	StartTime       int64    `dynamo:"StartTime,omitempty" json:"startTime,omitempty"`
	Tier            int      `dynamo:"Tier,omitempty" json:"tier,omitempty"`
	NewCreationName string   `dynamo:"newCreationName,omitempty" json:"newCreationName,omitempty"`
}

// Location of a group
type Location struct {
	StreetAddress   string `dynamo:"StreetAddress,omitempty" json:"startTime,omitempty"`
	ExtendedAddress string `dynamo:"ExtendedAddress,omitempty" json:"extendedAddress,omitempty"`
	Locality        string `dynamo:"Locality,omitempty" json:"locality,omitempty"`
	PostalCode      string `dynamo:"PostalCode,omitempty" json:"postalCode,omitempty"`
	Region          string `dynamo:"Region,omitempty" json:"region,omitempty"`
	Country         string `dynamo:"Country,omitempty" json:"country,omitempty"`
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
