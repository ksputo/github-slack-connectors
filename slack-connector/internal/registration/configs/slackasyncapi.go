package configs

//SlackAsyncAPI is a AsyncAPI 1.2.0 standard specification for Slack Events
var SlackAsyncAPI = []byte(`
{
    "asyncapi": "1.2.0",
    "baseTopic": "slack.events",
    "components": {
        "schemas": {
            "GenericEventWrapper": {
                "additionalProperties": true,
                "description": "Adapted from auto-generated content",
                "properties": {
                    "api_app_id": {
                        "description": " Use this to distinguish which app the event belongs to if you use multiple apps with the same Request URL.",
                        "title": "The unique identifier your installed Slack application.",
                        "type": "string",
                        "x-examples": [
                            "A2H9RFS1A"
                        ]
                    },
                    "authed_users": {
                        "items": {
                            "type": "string"
                        },
                        "minItems": 1,
                        "title": "An array of string-based User IDs. Each member of the collection represents a user that has installed your application/bot and indicates the described event would be visible to those users.",
                        "type": "array",
                        "uniqueItems": true
                    },
                    "event": {
                        "additionalProperties": true,
                        "properties": {
                            "event_ts": {
                                "title": "When the event was dispatched",
                                "type": "string"
                            },
                            "type": {
                                "title": "The specific name of the event",
                                "type": "string"
                            }
                        },
                        "required": [
                            "type",
                            "event_ts"
                        ],
                        "title": "The actual event, an object, that happened",
                        "type": "object",
                        "x-examples": [
                            {
                                "channel": "D0PNCRP9N",
                                "channel_type": "app_home",
                                "event_ts": "1525215129.000001",
                                "text": "How many cats did we herd yesterday?",
                                "ts": "1525215129.000001",
                                "type": "message",
                                "user": "U061F7AUR"
                            }
                        ]
                    },
                    "event_id": {
                        "title": "A unique identifier for this specific event, globally unique across all workspaces.",
                        "type": "string",
                        "x-examples": [
                            "Ev0PV52K25"
                        ]
                    },
                    "event_time": {
                        "title": "The epoch timestamp in seconds indicating when this event was dispatched.",
                        "type": "integer",
                        "x-examples": [
                            1525215129
                        ]
                    },
                    "team_id": {
                        "title": "The unique identifier of the workspace where the event occurred",
                        "type": "string",
                        "x-examples": [
                            "T1H9RESGL"
                        ]
                    },
                    "token": {
                        "title": "A verification token to validate the event originated from Slack",
                        "type": "string"
                    },
                    "type": {
                        "title": "Indicates which kind of event dispatch this is, usually 'event_callback'",
                        "type": "string",
                        "x-examples": [
                            "event_callback"
                        ]
                    }
                },
                "required": [
                    "token",
                    "team_id",
                    "api_app_id",
                    "event",
                    "type",
                    "event_id",
                    "event_time",
                    "authed_users"
                ],
                "title": "Standard event wrapper for the Events API",
                "type": "object",
                "x-examples": [
                    {
                        "api_app_id": "AXXXXXXXXX",
                        "authed_teams": [],
                        "event": {
                            "resources": [
                                {
                                    "resource": {
                                        "grant": {
                                            "resource_id": "DXXXXXXXX",
                                            "type": "specific"
                                        },
                                        "type": "im"
                                    },
                                    "scopes": [
                                        "chat:write:user",
                                        "im:read",
                                        "im:history",
                                        "commands"
                                    ]
                                }
                            ],
                            "type": "resources_added"
                        },
                        "event_id": "EvXXXXXXXX",
                        "event_time": 1234567890,
                        "team_id": "TXXXXXXXX",
                        "token": "XXYYZZ",
                        "type": "event_callback"
                    },
                    {
                        "api_app_id": "AXXXXXXXXX",
                        "authed_teams": [],
                        "event": {
                            "event_ts": "1360782804.083113",
                            "item": {
                                "channel": "C0G9QF9GZ",
                                "ts": "1360782400.498405",
                                "type": "message"
                            },
                            "item_user": "U0G9QF9C6",
                            "reaction": "thumbsup",
                            "type": "reaction_added",
                            "user": "U024BE7LH"
                        },
                        "event_id": "EvXXXXXXXX",
                        "event_time": 1234567890,
                        "team_id": "TXXXXXXXX",
                        "token": "XXYYZZ",
                        "type": "event_callback"
                    }
                ]
            }
        }
    },
    "externalDocs": {
        "description": "Slack Events API documentation",
        "url": "https://api.slack.com/events-api"
    },
    "info": {
        "contact": {
            "email": "developers@slack.com",
            "name": "Slack Platform Support",
            "url": "https://api.slack.com/support"
        },
        "description": "A webhook-based events bus using a subscription model for Slack apps",
        "termsOfService": "https://slack.com/terms-of-service/api",
        "title": "Slack Events API",
        "version": "1.0.0"
    },
    "servers": [
        {
            "description": "A pre-registered Request URL on servers you control where subscriptions will be delivered.",
            "scheme": "https",
            "url": "{request_url}",
            "variables": {
                "request_url": {
                    "description": "Your chosen request URL where events will be delivered."
                }
            }
        },
        {
            "description": "Your Slack app management console's event subscription configurator. Visit in a web browser and sign in to your workspace.",
            "scheme": "https",
            "url": "https://api.slack.com/apps/{slack_app_id}/event-subscriptions",
            "variables": {
                "slack_app_id": {
                    "description": "Your Slack app's unique identifier, found in the URL when editing your app on api.slack.com."
                }
            }
        }
    ],
    "tags": [
        {
            "description": "Events less about a workspace and more about your app",
            "name": "app_event"
        },
        {
            "description": "Part of workspace app's Permissions API",
            "name": "permissions_api"
        },
        {
            "description": "User token based apps can subscribe to this event",
            "name": "allows_user_tokens"
        },
        {
            "description": "Workspace token apps can subscribe to this event",
            "name": "allows_workspace_tokens"
        }
    ],
    "topics": {
        "app.mention": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for app_mention",
                    "url": "https://api.slack.com/events/app_mention"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Subscribe to only the message events that mention your app or bot",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "app.rate.limited": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for app_rate_limited",
                    "url": "https://api.slack.com/events/app_rate_limited"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Indicates your app's event subscriptions are being rate limited",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "app.uninstalled": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for app_uninstalled",
                    "url": "https://api.slack.com/events/app_uninstalled"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Your Slack app was uninstalled.",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "channel.archive": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_archive",
                    "url": "https://api.slack.com/events/channel_archive"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A channel was archived",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "channel.created": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_created",
                    "url": "https://api.slack.com/events/channel_created"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A channel was created",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "channel.deleted": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_deleted",
                    "url": "https://api.slack.com/events/channel_deleted"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A channel was deleted",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "channel.history.changed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_history_changed",
                    "url": "https://api.slack.com/events/channel_history_changed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Bulk updates were made to a channel's history",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:history",
                    "groups:history",
                    "mpim:history"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "channel.left": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_left",
                    "url": "https://api.slack.com/events/channel_left"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You left a channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "channel.rename": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_rename",
                    "url": "https://api.slack.com/events/channel_rename"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A channel was renamed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "channel.unarchive": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for channel_unarchive",
                    "url": "https://api.slack.com/events/channel_unarchive"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A channel was unarchived",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "dnd.updated": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for dnd_updated",
                    "url": "https://api.slack.com/events/dnd_updated"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Do not Disturb settings changed for the current user",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "dnd:read"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "dnd.updated.user": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for dnd_updated_user",
                    "url": "https://api.slack.com/events/dnd_updated_user"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Do not Disturb settings changed for a member",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "dnd:read"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "email.domain.changed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for email_domain_changed",
                    "url": "https://api.slack.com/events/email_domain_changed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "The workspace email domain has changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "team:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "emoji.changed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for emoji_changed",
                    "url": "https://api.slack.com/events/emoji_changed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A custom emoji has been added or changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "emoji:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.change": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_change",
                    "url": "https://api.slack.com/events/file_change"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file was changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.comment.added": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_comment_added",
                    "url": "https://api.slack.com/events/file_comment_added"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file comment was added",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.comment.deleted": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_comment_deleted",
                    "url": "https://api.slack.com/events/file_comment_deleted"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file comment was deleted",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.comment.edited": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_comment_edited",
                    "url": "https://api.slack.com/events/file_comment_edited"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file comment was edited",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.created": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_created",
                    "url": "https://api.slack.com/events/file_created"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file was created",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.deleted": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_deleted",
                    "url": "https://api.slack.com/events/file_deleted"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file was deleted",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.public": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_public",
                    "url": "https://api.slack.com/events/file_public"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file was made public",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.shared": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_shared",
                    "url": "https://api.slack.com/events/file_shared"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file was shared",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "file.unshared": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for file_unshared",
                    "url": "https://api.slack.com/events/file_unshared"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A file was unshared",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "files:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "grid.migration.finished": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for grid_migration_finished",
                    "url": "https://api.slack.com/events/grid_migration_finished"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "An enterprise grid migration has finished on this workspace.",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "grid.migration.started": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for grid_migration_started",
                    "url": "https://api.slack.com/events/grid_migration_started"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "An enterprise grid migration has started on this workspace.",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "group.archive": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_archive",
                    "url": "https://api.slack.com/events/group_archive"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A private channel was archived",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "group.close": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_close",
                    "url": "https://api.slack.com/events/group_close"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You closed a private channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "group.history.changed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_history_changed",
                    "url": "https://api.slack.com/events/group_history_changed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Bulk updates were made to a private channel's history",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:history"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "group.left": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_left",
                    "url": "https://api.slack.com/events/group_left"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You left a private channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "group.open": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_open",
                    "url": "https://api.slack.com/events/group_open"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You created a group DM",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "group.rename": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_rename",
                    "url": "https://api.slack.com/events/group_rename"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A private channel was renamed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "group.unarchive": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for group_unarchive",
                    "url": "https://api.slack.com/events/group_unarchive"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A private channel was unarchived",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "im.close": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for im_close",
                    "url": "https://api.slack.com/events/im_close"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You closed a DM",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "im:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "im.created": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for im_created",
                    "url": "https://api.slack.com/events/im_created"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A DM was created",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "im:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "im.history.changed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for im_history_changed",
                    "url": "https://api.slack.com/events/im_history_changed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Bulk updates were made to a DM's history",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "im:history"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "im.open": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for im_open",
                    "url": "https://api.slack.com/events/im_open"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You opened a DM",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "im:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "link.shared": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for link_shared",
                    "url": "https://api.slack.com/events/link_shared"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A message was posted containing one or more links relevant to your application",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "links:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "member.joined.channel": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for member_joined_channel",
                    "url": "https://api.slack.com/events/member_joined_channel"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A user joined a public or private channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read",
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "member.left.channel": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for member_left_channel",
                    "url": "https://api.slack.com/events/member_left_channel"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A user left a public or private channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:read",
                    "groups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "message": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for message",
                    "url": "https://api.slack.com/events/message"
                },
                "payload": {
                    "additionalProperties": true,
                    "description": "Generated from events/events-1512674163.json with shasum b83372d878e1276a26d9e4df3d4a2563a60440a6",
                    "properties": {
                        "api_app_id": {
                            "title": "Your Slack app's unique ID",
                            "type": "string"
                        },
                        "authed_users": {
                            "items": {
                                "type": "string"
                            },
                            "minItems": 1,
                            "type": "array",
                            "uniqueItems": true
                        },
                        "event": {
                            "additionalProperties": true,
                            "properties": {
                                "channel": {
                                    "type": "string"
                                },
                                "deleted_ts": {
                                    "type": "string"
                                },
                                "event_ts": {
                                    "type": "string"
                                },
                                "hidden": {
                                    "type": "boolean"
                                },
                                "previous_message": {
                                    "additionalProperties": true,
                                    "properties": {
                                        "comment": {},
                                        "file": {
                                            "additionalProperties": true,
                                            "properties": {
                                                "channels": {
                                                    "items": {
                                                        "type": "string"
                                                    },
                                                    "minItems": 1,
                                                    "type": "array",
                                                    "uniqueItems": true
                                                },
                                                "comments_count": {
                                                    "type": "integer"
                                                },
                                                "created": {
                                                    "type": "integer"
                                                },
                                                "display_as_bot": {
                                                    "type": "boolean"
                                                },
                                                "editable": {
                                                    "type": "boolean"
                                                },
                                                "editor": {
                                                    "type": "string"
                                                },
                                                "external_type": {
                                                    "type": "string"
                                                },
                                                "filetype": {
                                                    "type": "string"
                                                },
                                                "groups": {
                                                    "items": {},
                                                    "minItems": 0,
                                                    "type": "array",
                                                    "uniqueItems": true
                                                },
                                                "id": {
                                                    "type": "string"
                                                },
                                                "ims": {
                                                    "items": {},
                                                    "minItems": 0,
                                                    "type": "array",
                                                    "uniqueItems": true
                                                },
                                                "is_external": {
                                                    "type": "boolean"
                                                },
                                                "is_public": {
                                                    "type": "boolean"
                                                },
                                                "last_editor": {
                                                    "type": "string"
                                                },
                                                "mimetype": {
                                                    "type": "string"
                                                },
                                                "mode": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "permalink": {
                                                    "type": "string"
                                                },
                                                "permalink_public": {
                                                    "type": "string"
                                                },
                                                "pretty_type": {
                                                    "type": "string"
                                                },
                                                "preview": {},
                                                "public_url_shared": {
                                                    "type": "boolean"
                                                },
                                                "size": {
                                                    "type": "integer"
                                                },
                                                "state": {
                                                    "type": "string"
                                                },
                                                "timestamp": {
                                                    "type": "integer"
                                                },
                                                "title": {
                                                    "type": "string"
                                                },
                                                "updated": {
                                                    "type": "integer"
                                                },
                                                "url_private": {
                                                    "type": "string"
                                                },
                                                "url_private_download": {
                                                    "type": "string"
                                                },
                                                "user": {
                                                    "type": "string"
                                                },
                                                "username": {
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "id",
                                                "created",
                                                "timestamp",
                                                "name",
                                                "title",
                                                "mimetype",
                                                "filetype",
                                                "pretty_type",
                                                "user",
                                                "editable",
                                                "size",
                                                "mode",
                                                "is_external",
                                                "external_type",
                                                "is_public",
                                                "public_url_shared",
                                                "display_as_bot",
                                                "username",
                                                "url_private",
                                                "url_private_download",
                                                "permalink",
                                                "permalink_public",
                                                "preview",
                                                "updated",
                                                "editor",
                                                "last_editor",
                                                "state",
                                                "channels",
                                                "groups",
                                                "ims",
                                                "comments_count"
                                            ],
                                            "type": "object"
                                        },
                                        "subtype": {
                                            "type": "string"
                                        },
                                        "text": {
                                            "type": "string"
                                        },
                                        "ts": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "type",
                                        "subtype",
                                        "text",
                                        "file",
                                        "comment",
                                        "ts"
                                    ],
                                    "type": "object"
                                },
                                "subtype": {
                                    "type": "string"
                                },
                                "ts": {
                                    "type": "string"
                                },
                                "type": {
                                    "type": "string"
                                }
                            },
                            "required": [
                                "type",
                                "deleted_ts",
                                "subtype",
                                "hidden",
                                "channel",
                                "previous_message",
                                "event_ts",
                                "ts"
                            ],
                            "type": "object"
                        },
                        "event_id": {
                            "type": "string"
                        },
                        "event_time": {
                            "type": "integer"
                        },
                        "team_id": {
                            "title": "The unique identifier of the team/workspace where the event happened",
                            "type": "string"
                        },
                        "token": {
                            "title": "Verification token used to validate the origin of the event",
                            "type": "string"
                        },
                        "type": {
                            "type": "string"
                        }
                    },
                    "required": [
                        "token",
                        "team_id",
                        "api_app_id",
                        "event",
                        "type",
                        "event_id",
                        "event_time",
                        "authed_users"
                    ],
                    "title": "API method: events/events-1512674163.json",
                    "type": "object"
                },
                "summary": "A message was sent to a channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:history",
                    "groups:history",
                    "im:history",
                    "mpim:history"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "message.app.home": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for message.app_home",
                    "url": "https://api.slack.com/events/message.app_home"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A user sent a message to your Slack app",
                "tags": [
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "workspace"
                ]
            }
        },
        "message.channels": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for message.channels",
                    "url": "https://api.slack.com/events/message.channels"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A message was posted to a channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "channels:history"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "message.groups": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for message.groups",
                    "url": "https://api.slack.com/events/message.groups"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A message was posted to a private channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "groups:history"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "message.im": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for message.im",
                    "url": "https://api.slack.com/events/message.im"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A message was posted in a direct message channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "im:history"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "message.mpim": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for message.mpim",
                    "url": "https://api.slack.com/events/message.mpim"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A message was posted in a multiparty direct message channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "mpim:history"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "pin.added": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for pin_added",
                    "url": "https://api.slack.com/events/pin_added"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A pin was added to a channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "pins:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "pin.removed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for pin_removed",
                    "url": "https://api.slack.com/events/pin_removed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A pin was removed from a channel",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "pins:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "reaction.added": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for reaction_added",
                    "url": "https://api.slack.com/events/reaction_added"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A member has added an emoji reaction to an item",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "reactions:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "reaction.removed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for reaction_removed",
                    "url": "https://api.slack.com/events/reaction_removed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A member removed an emoji reaction",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "reactions:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "resources.added": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for resources_added",
                    "url": "https://api.slack.com/events/resources_added"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Access to a set of resources was granted for your app",
                "tags": [
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "workspace"
                ]
            }
        },
        "resources.removed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for resources_removed",
                    "url": "https://api.slack.com/events/resources_removed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Access to a set of resources was removed for your app",
                "tags": [
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "workspace"
                ]
            }
        },
        "scope.denied": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for scope_denied",
                    "url": "https://api.slack.com/events/scope_denied"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "OAuth scopes were denied to your app",
                "tags": [
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "workspace"
                ]
            }
        },
        "scope.granted": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for scope_granted",
                    "url": "https://api.slack.com/events/scope_granted"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "OAuth scopes were granted to your app",
                "tags": [
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "workspace"
                ]
            }
        },
        "star.added": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for star_added",
                    "url": "https://api.slack.com/events/star_added"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A member has starred an item",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "stars:read"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "star.removed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for star_removed",
                    "url": "https://api.slack.com/events/star_removed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A member removed a star",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "stars:read"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "subteam.created": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for subteam_created",
                    "url": "https://api.slack.com/events/subteam_created"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A User Group has been added to the workspace",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "usergroups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "subteam.members.changed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for subteam_members_changed",
                    "url": "https://api.slack.com/events/subteam_members_changed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "The membership of an existing User Group has changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "usergroups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "subteam.self.added": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for subteam_self_added",
                    "url": "https://api.slack.com/events/subteam_self_added"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You have been added to a User Group",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "usergroups:read"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "subteam.self.removed": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for subteam_self_removed",
                    "url": "https://api.slack.com/events/subteam_self_removed"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "You have been removed from a User Group",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    }
                ],
                "x-scopes-required": [
                    "usergroups:read"
                ],
                "x-tokens-allowed": [
                    "user"
                ]
            }
        },
        "subteam.updated": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for subteam_updated",
                    "url": "https://api.slack.com/events/subteam_updated"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "An existing User Group has been updated or its members changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "usergroups:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "team.domain.change": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for team_domain_change",
                    "url": "https://api.slack.com/events/team_domain_change"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "The workspace domain has changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "team:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "team.join": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for team_join",
                    "url": "https://api.slack.com/events/team_join"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A new member has joined",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "users:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "team.rename": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for team_rename",
                    "url": "https://api.slack.com/events/team_rename"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "The workspace name has changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "team:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "tokens.revoked": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for tokens_revoked",
                    "url": "https://api.slack.com/events/tokens_revoked"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "API tokens for your app were revoked.",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "url.verification": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for url_verification",
                    "url": "https://api.slack.com/events/url_verification"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "Verifies ownership of an Events API Request URL",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "app_event"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        },
        "user.change": {
            "subscribe": {
                "externalDocs": {
                    "description": "Event documentation for user_change",
                    "url": "https://api.slack.com/events/user_change"
                },
                "payload": {
                    "$ref": "#/components/schemas/GenericEventWrapper"
                },
                "summary": "A member's data has changed",
                "tags": [
                    {
                        "name": "allows_user_tokens"
                    },
                    {
                        "name": "allows_workspace_tokens"
                    }
                ],
                "x-scopes-required": [
                    "users:read"
                ],
                "x-tokens-allowed": [
                    "user",
                    "workspace"
                ]
            }
        }
    }
}
`)
