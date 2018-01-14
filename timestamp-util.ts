import { timestamp } from './timestamp'

// newTimestamp builds a new timestamp.
export function newTimestamp(date?: Date): timestamp.Timestamp {
    date = date || new Date()
    let timeUnixMs = date.valueOf()
    return timestamp.Timestamp.create({ timeUnixMs })
}

// timestampToTime converts a timestamp to a date.
export function timestampToTime(timestamp: timestamp.Timestamp): Date {
    return new Date(timestamp.timeUnixMs.toNumber())
}
