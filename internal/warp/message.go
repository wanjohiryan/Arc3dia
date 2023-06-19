package warp

type Message struct {
	Init      *MessageInit      `json:"init,omitempty"`
	Segment   *MessageSegment   `json:"segment,omitempty"`
	Beat      *MessageHeartBeat `json:"beat,omitempty"`
	Bandwidth *MessageBandwidth `json:"bandwidth,omitempty"`
}

type MessageInit struct {
	Id string `json:"id"` // ID of the init segment
}

type MessageSegment struct {
	Init      string `json:"init"`      // ID of the init segment to use for this segment
	Timestamp int    `json:"timestamp"` // PTS of the first frame in milliseconds
}

type MessageBandwidth struct {
	MaxBitrate int `json:"max_bitrate"` // Artificially limit the QUIC max bitrate
}

type MessageHeartBeat struct {
	Timestamp int `json:"timestamp"` // current timestamp in milliseconds
}
