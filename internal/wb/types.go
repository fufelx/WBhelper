package wb

import "time"

type SellerInfo struct {
	Name      string `json:"name"`
	Sid       string `json:"sid"`
	Tin       string `json:"tin"`
	TradeMark string `json:"tradeMark"`
}
type ChatEvent struct {
	Result struct {
		Next            int64     `json:"next"`
		NewestEventTime time.Time `json:"newestEventTime"`
		OldestEventTime time.Time `json:"oldestEventTime"`
		TotalEvents     int       `json:"totalEvents"`
		Events          []struct {
			ChatID    string `json:"chatID"`
			EventID   string `json:"eventID"`
			EventType string `json:"eventType"`
			IsNewChat bool   `json:"isNewChat,omitempty"`
			Message   struct {
				Attachments struct {
					GoodCard struct {
						NmID          int    `json:"nmID"`
						Price         int    `json:"price"`
						PriceCurrency string `json:"priceCurrency"`
						Rid           string `json:"rid"`
						Size          string `json:"size"`
					} `json:"goodCard"`
					Files []struct {
						ContentType string    `json:"contentType"`
						Date        time.Time `json:"date"`
						DownloadID  string    `json:"downloadID"`
						Name        string    `json:"name"`
						Url         string    `json:"url"`
						Size        int       `json:"size"`
					} `json:"files"`
					Images []struct {
						Date       time.Time `json:"date"`
						DownloadID string    `json:"downloadID"`
						Url        string    `json:"url"`
					} `json:"images"`
				} `json:"attachments,omitempty"`
				Text string `json:"text"`
			} `json:"message,omitempty"`
			Source       string    `json:"source,omitempty"`
			AddTimestamp int64     `json:"addTimestamp"`
			AddTime      time.Time `json:"addTime"`
			ReplySign    string    `json:"replySign,omitempty"`
			Sender       string    `json:"sender"`
			ClientName   string    `json:"clientName,omitempty"`
		} `json:"events"`
	} `json:"result"`
	Errors interface{} `json:"errors"`
}

type Chat struct {
	ChatID    string `json:"chatID"`
	IsNewChat bool   `json:"isNewChat,omitempty"`
	Message   []struct {
		Attachments struct {
			GoodCard struct {
				NmID          int    `json:"nmID"`
				Price         int    `json:"price"`
				PriceCurrency string `json:"priceCurrency"`
				Rid           string `json:"rid"`
				Size          string `json:"size"`
			} `json:"goodCard"`
			Files []struct {
				ContentType string    `json:"contentType"`
				Date        time.Time `json:"date"`
				DownloadID  string    `json:"downloadID"`
				Name        string    `json:"name"`
				Url         string    `json:"url"`
				Size        int       `json:"size"`
			} `json:"files"`
			Images []struct {
				Date       time.Time `json:"date"`
				DownloadID string    `json:"downloadID"`
				Url        string    `json:"url"`
			} `json:"images"`
		} `json:"attachments,omitempty"`
		Text string `json:"text"`
	} `json:"message,omitempty"`
	Source       string    `json:"source,omitempty"`
	AddTimestamp int64     `json:"addTimestamp"`
	AddTime      time.Time `json:"addTime"`
	ReplySign    string    `json:"replySign,omitempty"`
	Sender       string    `json:"sender"`
	ClientName   string    `json:"clientName,omitempty"`
}
