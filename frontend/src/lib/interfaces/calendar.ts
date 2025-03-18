import type { DateTime } from 'luxon';
import type { DateString } from './date';
export interface CalendarItem<T> {
    date: DateString;
    data: T;
}

export interface MinimalCalendarUIItemMark {
    reserved?: boolean;
    unavailable?: boolean;
    past?: boolean;
    today?: boolean;
    // selected?: boolean;
    // use selected props on calendar instaed.
}
export interface MinimalCalendarUIItem {
    date: DateTime;
    mark?: MinimalCalendarUIItemMark;
}

export interface MinimalCalendarUIItemWithHref extends MinimalCalendarUIItem {
    href?: string;
}
