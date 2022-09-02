package DisGOAuth

// Constant OAuth URL Scopes
const (
	// Non-Whitelist
	ScopeIdentify                   = "identify"
	ScopeBot                        = "bot"
	ScopeEmail                      = "email"
	ScopeGuilds                     = "guilds"
	ScopeGuildsJoin                 = "guilds.join"
	ScopeConnections                = "connections"
	ScopeGroupDMJoin                = "gdm.join"
	ScopeMessagesRead               = "messages.read"
	ScopeWebhookIncoming            = "webhook.Incoming"
	ScopeApplicationsBuildsRead     = "applications.builds.read"
	ScopeApplicationsStoreUpdate    = "applications.store.update"
	ScopeApplicationsEntitlements   = "applications.entitlements"
	ScopeApplicationsCommands       = "applications.commands"
	ScopeApplicationsCommandsUpdate = "applications.commands.update"

	// Whitelist Only
	ScopeRPC                      = "rpc"
	ScopeRPCAPI                   = "rpc.api"
	ScopeRPCNotificationsRead     = "rpc.notifications.read"
	ScopeApplicationsBuildsUpload = "applications.builds.upload"
	ScopeRelationshipsRead        = "relationships.read"
	ScopeActivitiesRead           = "activities.read"
	ScopeActivitiesWrite          = "activities.write"
)
