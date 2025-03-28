import type { HourString } from '$lib/interfaces/date';

export const splitArray = <T>(arr: T[], size: number) => {
    const result = [];
    for (let i = 0; i < arr.length; i += size) {
        result.push(arr.slice(i, i + size));
    }
    return result;
};
export const zeroPad = (text: string) => text.padStart(2, '0');
export const generateHours = () => Array.from(Array(24).keys());
export const generateHourStrings = () =>
    generateHours().map((i) => zeroPad(i.toString()) as HourString);

export const getHourRangeString = (hours: number[]) => {
    const is_sequence = [...hours] // clone
        .sort((a, b) => a - b)
        .every((value, index, array) => index == 0 || array[index - 1] + 1 == value);

    if (is_sequence && hours.length > 1)
        return `${zeroPad(hours.at(0)!.toString())}~${zeroPad(hours.at(-1)!.toString())}`;

    return hours.map((i) => zeroPad(i.toString())).join(',');
};
