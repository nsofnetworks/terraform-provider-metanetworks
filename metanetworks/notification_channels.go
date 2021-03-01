package metanetworks

import (
	"errors"
	"log"
)

const (
	notificationChannelsEndpoint string = "/v1/notification_channels"
)

// NotificationChannels ...
type NotificationChannels struct {
	Name            string            `json:"name"`
	Description     string            `json:"description,omitempty"`
	Enabled         bool              `json:"enabled" type:"bool"`
	EmailConfig     []EmailConfig     `json:"email_config,omitempty"`
	SlackConfig     []SlackConfig     `json:"slack_config,omitempty"`
	PagerDutyConfig []PagerDutyConfig `json:"pagerduty_config,omitempty"`
	CreatedAt       string            `json:"created_at,omitempty" meta_api:"read_only"`
	ID              string            `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt      string            `json:"modified_at,omitempty" meta_api:"read_only"`
}

// EmailConfiguration ...
type EmailConfig struct {
	Recipients []string `json:"recipients,omitempty"`
}

// SlackConfiguration ...
type SlackConfig struct {
	Channel string `json:"channel,omitempty"`
	URL     string `json:"url,omitempty"`
}

// PagerDutyConfiguration ...
type PagerDutyConfig struct {
	APIKey []string `json:"api_key,omitempty"`
}

// GetNotificationChannels ...
func (c *Client) GetNotificationChannels(notificationChannelsID string) (*NotificationChannels, error) {
	var notificationChannels NotificationChannels
	err := c.Read(protocolGroupsEndpoint+"/"+notificationChannelsID, &notificationChannels)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Notification Channels Settings from Get: %s", notificationChannels.ID)
	return &notificationChannels, nil
}

// UpdateNotificationChannels ...
func (c *Client) UpdateNotificationChannels(notificationChannelsID string, notificationChannels *NotificationChannels) (*NotificationChannels, error) {
	resp, err := c.Update(notificationChannelsEndpoint+"/"+notificationChannelsID, *notificationChannels)
	if err != nil {
		return nil, err
	}
	updatedNotificationChannels, _ := resp.(*NotificationChannels)

	log.Printf("Returning Notification Channels Settings from Update: %s", updatedNotificationChannels.ID)
	return updatedNotificationChannels, nil
}

// CreateNotificationChannels ...
func (c *Client) CreateNotificationChannels(notificationChannels *NotificationChannels) (*NotificationChannels, error) {
	resp, err := c.Create(notificationChannelsEndpoint, *notificationChannels)
	if err != nil {
		return nil, err
	}

	createdNotificationChannels, ok := resp.(*NotificationChannels)
	if !ok {
		return nil, errors.New("Object returned from API was not a Notification Channels Pointer")
	}

	log.Printf("Returning NotificationChannels Settings from Create: %s", createdNotificationChannels.ID)
	return createdNotificationChannels, nil
}

// DeleteNotificationChannels ...
func (c *Client) DeleteNotificationChannels(notificationChannelsID string) error {
	err := c.Delete(notificationChannelsEndpoint + "/" + notificationChannelsID)
	if err != nil {
		return err
	}

	return nil
}
