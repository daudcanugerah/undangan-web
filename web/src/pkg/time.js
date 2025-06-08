import dayjs from 'dayjs'
import utc from "dayjs/plugin/utc"
import timezone from "dayjs/plugin/timezone"

const tz = "Asia/Jakarta";

dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.tz.setDefault(tz)

export default dayjs
