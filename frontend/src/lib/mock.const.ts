import type { Venue } from './interfaces/api';

export const venues: Venue[] = [
    {
        venue: 'studio',
        venueKor: '스튜디오',
        requirement: [
            '물품 보관 절대 금지 (제작 비품 및 미술 소품 포함)',
            '음식 반입 절대 금지',
            '음료 반입 가능 (본인 이름 기재 필수)',
            '허락 없는 기자재 사용 및 이동 금지 (책상, 의자 포함)'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'mastering1',
        venueKor: '마스터링룸 1',
        requirement: [
            '물품 보관 금지(제작 비품 및 미술 소품 일체 포함)',
            '음식 반입 절대 금지, 음료 반입 가능(단, 본인 이름 기재 필수)',
            '모든 장비나 라인은 그 자리에 둔 채로 공간을 이용할 것(책상 포함)',
            '임의 변경 발견시 전체 공간 72시간 사용 금지'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'mastering2',
        venueKor: '마스터링룸 2',
        requirement: [
            '물품 보관 금지(제작 비품 및 미술 소품 일체 포함)',
            '음식 반입 절대 금지, 음료 반입 가능(단, 본인 이름 기재 필수)',
            '모든 장비나 라인은 그 자리에 둔 채로 공간을 이용할 것(책상 포함)',
            '임의 변경 발견시 전체 공간 72시간 사용 금지'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'meeting',
        venueKor: '회의실',
        requirement: [
            '물품 보관시 이름표 부착 & 카페 게시판글 작성 필수',
            '음식/ 음료 반입 가능',
            '허락없는 기자재 이동 금지 (책상, 의자 포함)',
            '24시간 상시개방(단, 사전 사용 신청팀에게 우선권)'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'mixing',
        venueKor: '믹싱룸/ADR룸',
        requirement: [
            '물품 보관 절대 금지(제작 비품 및 미술 소품 일체 포함)',
            '음식 반입 절대 금지',
            '음료 반입 가능(단, 본인 이름 기재 필수)',
            '허락없는 기자재 사용 및 이동 금지(책상, 의자 포함)'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'editing',
        venueKor: '편집실',
        requirement: [
            '개인 물품 보관 금지(외장하드, 스크립북 포함)',
            '음식 반입 절대 금지',
            '음료 반입 가능(단, 본인 이름 기재 필수)',
            '허락없는 기자재 이동 및 도어락 건전지 빼기 금지(책상, 의자 포함)'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'lounge',
        venueKor: '과방',
        requirement: [
            '음식/ 음료 반입 가능',
            '허락없는 기자재 이동 금지 (책상, 의자 포함)',
            '24시간 상시개방(단, 사전 사용 신청팀에게 우선권)'
        ],
        approval_mode: 'auto'
    }
].map((v) => {
    return {
        ...v,
        requirement:
            `${v.venueKor}의 유의사항입니다. <br><ul>` +
            v.requirement?.map((i) => `<li>${i}</li>`).join('') +
            '</ul>',
        approval_mode: v.approval_mode as 'auto' | 'manual'
    };
});

export const purposes = ['수업', '워크샵', '과제', '소모임', '개인 작업', '기타'];
