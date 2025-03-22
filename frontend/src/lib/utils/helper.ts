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
