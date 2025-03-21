import type { Venue } from './interfaces/api';

export const venues: Venue[] = [
    {
        venue: 'apple',
        venueKor: '사과',
        requirement: '유의 사항.....',
        approval_mode: 'auto'
    },
    {
        venue: 'manual_banana',
        venueKor: '수동 바나나',
        requirement: '유의 사항.....',
        approval_mode: 'manual'
    },
    {
        venue: 'cherry',
        venueKor: '체리',
        requirement: '유의 사항.....',
        approval_mode: 'auto'
    },
    {
        venue: 'nothing',
        venueKor: '일정 없음',
        requirement: '유의 사항.....',
        approval_mode: 'auto'
    }
];

export const purposes = ['수업', '워크샵', '과제', '소모임', '개인 작업', '기타'];
