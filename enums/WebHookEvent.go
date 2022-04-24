package enums

type WebHookEvent string

const (
	WebHookEvent_event_channel_terminated             WebHookEvent = "event_channel_terminated"
	WebHookEvent_event_outoing_call                   WebHookEvent = "event_outoing_call"
	WebHookEvent_event_incoming_call                  WebHookEvent = "event_incoming_call"
	WebHookEvent_event_call_answered                  WebHookEvent = "event_call_answered"
	WebHookEvent_event_phone_disconnected_from_webapp WebHookEvent = "event_phone_disconnected_from_webapp"
	WebHookEvent_event_phone_connected_to_webapp      WebHookEvent = "event_phone_connected_to_webapp"
	WebHookEvent_event_click_to_dial_chrome_extension WebHookEvent = "event_click_to_dial_chrome_extension"
	WebHookEvent_event_phone_created                  WebHookEvent = "event_phone_created"
	WebHookEvent_event_phone_updated                  WebHookEvent = "event_phone_updated"
	WebHookEvent_event_phone_deleted                  WebHookEvent = "event_phone_deleted"
	WebHookEvent_event_phone_line_created             WebHookEvent = "event_phone_line_created"
	WebHookEvent_event_phone_line_deleted             WebHookEvent = "event_phone_line_deleted"
	WebHookEvent_event_extension_created              WebHookEvent = "event_extension_created"
	WebHookEvent_event_extension_updated              WebHookEvent = "event_extension_updated"
	WebHookEvent_event_extension_deleted              WebHookEvent = "event_extension_deleted"
	WebHookEvent_event_line_key_group_updated         WebHookEvent = "event_line_key_group_updated"
	WebHookEvent_event_phone_configuration_updated    WebHookEvent = "event_phone_configuration_updated"
	WebHookEvent_event_phone_configuration_deleted    WebHookEvent = "event_phone_configuration_deleted"
	WebHookEvent_event_phone_connected                WebHookEvent = "event_phone_connected"
	WebHookEvent_event_phone_disconnected             WebHookEvent = "event_phone_disconnected"
	WebHookEvent_event_extension_dial_not_answered    WebHookEvent = "event_extension_dial_not_answered"
	WebHookEvent_event_extension_queue_not_answered   WebHookEvent = "event_extension_queue_not_answered"
	WebHookEvent_event_power_dialer_created           WebHookEvent = "event_power_dialer_created"
	WebHookEvent_event_power_dialer_changed           WebHookEvent = "event_power_dialer_changed"
)
