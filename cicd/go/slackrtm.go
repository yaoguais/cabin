package main

import (
	"errors"

	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
)

// SetupRTM setup real time messaging
func (s *SlackBot) SetupRTM() error {
	if s.rtm != nil {
		return errors.New("slack rtm alreay running")
	}

	s.rtm = s.client.NewRTM()

	go s.rtm.ManageConnection()
	go s.manageRTMIncomingEvents()

	return nil
}

func (s *SlackBot) manageRTMIncomingEvents() {
	for msg := range s.rtm.IncomingEvents {
		logrus.WithField("message", msg).Debug("slack rtm message")

		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
		case *slack.ConnectedEvent:
		case *slack.MessageEvent:
		case *slack.PresenceChangeEvent:
		case *slack.LatencyReport:
		case *slack.RTMError:
			logrus.WithError(errors.New(ev.Error())).Info("slack rtm error message")
		case *slack.InvalidAuthEvent:
			logrus.Error("invalid slack rtm credentials")
			return
		default:
		}
	}
}

// ShutdownRTM close real time messaging connection
func (s *SlackBot) ShutdownRTM() error {
	if s.rtm == nil {
		return nil
	}

	err := s.rtm.Disconnect()
	s.rtm = nil
	return err
}
