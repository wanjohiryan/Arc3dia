export interface Message {
	init?: MessageInit
	segment?: MessageSegment
	beat?: MessageBeat
}

export interface MessageInit {
	id: string
}

export interface MessageSegment {
	init: string // id of the init segment
	timestamp: number // presentation timestamp in milliseconds of the first sample
	// TODO track would be nice
}

export interface MessageBeat {
	timestamp: number // presentation timestamp in milliseconds
}

export interface Debug {
	max_bitrate: number
}