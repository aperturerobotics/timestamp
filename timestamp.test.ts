import { timestamp } from './timestamp'
import { newTimestamp, timestampToTime } from './timestamp-util'

describe('timestamp', () => {
    it('should construct a timestamp', () => {
        let tb = new Date()
        let ts = newTimestamp(tb)
        expect(ts).not.toBeNull()
        let encTs = timestamp.Timestamp.encode(ts).finish()
        let ats = timestamp.Timestamp.decode(encTs)
        expect(ats).not.toBeNull()

        let ta = timestampToTime(ats)
        expect(ta).toEqual(tb)
    })
})
