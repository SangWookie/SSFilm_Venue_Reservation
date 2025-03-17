export const splitArray = (arr: unknown[], size: number) => {
    const result = [];
    for (let i = 0; i < arr.length; i += size) {
        result.push(arr.slice(i, i + size));
    }
    return result;
};
export const zeroPad = (text: string) => text.padStart(2, '0');
