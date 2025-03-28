import type { Venue } from './interfaces/api';

export const venues: Venue[] = [
    {
        venue: 'mastering1',
        venueKor: '마스터링룸 1',
        requirement: [
            '물품 보관X 음식X 음료O *단, 본인 이름 기재 필수',
            '책상, 의자, 장비, 라인 이동 및 임의 변경 금지'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'mastering2',
        venueKor: '마스터링룸 2',
        requirement: [
            '물품 보관X 음식X 음료O *단, 본인 이름 기재 필수',
            '책상, 의자, 장비, 라인 이동 및 임의 변경 금지'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'studio',
        venueKor: '스튜디오',
        requirement: [
            '물품 보관X 음식X 음료O *단, 본인 이름 기재 필수',
            '책상, 의자, 장비, 라인 이동 및 임의 변경 금지'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'meeting',
        venueKor: '회의실',
        requirement: [
            '물품 보관O *이름표 부착 및 카페 게시판글 작성 필수',
            '음식O 음료O',
            '책상, 의자, 장비, 라인 이동 및 임의 변경 금지',
            '24시간 상시개방 *단, 사전 사용 신청팀에게 우선권'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'mixing',
        venueKor: '믹싱룸/ADR룸',
        requirement: [
            '<b style="color: red;">믹싱, ADR, 녹음, 노래연습, 건반 사용 이외의 목적은 승인 불가합니다.</b>',
            '<b>예약 신청 후 믹싱룸 관리자에게 문자를 남겨주셔야 신청이 완료됩니다.</b>',
            '<b>믹싱룸 관리자: 8기 김세연 010-9170-6176</b>',
            '<b>문자 내용: 기수/학번/이름/사용 날짜와 시간/연락처/동행인</b>',
            '물품 보관X 음식X 음료O *단, 본인 이름 기재 필수',
            '책상, 의자, 장비, 라인 이동 및 임의 변경 금지'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'lounge',
        venueKor: '과방',
        requirement: [
            '해당 공간은 공용 공간으로, 공간 부족(워크샵 오디션 등)시에만 승인됩니다.',
            '물품 보관X 음식O 음료O *단, 본인 이름 기재 필수',
            '책상, 의자 포함 모든 가구 무단 이동 금지.',
            '24시간 상시개방 *단, 사전 사용 신청팀에게 우선권'
        ],
        approval_mode: 'auto'
    },
    {
        venue: 'editing',
        venueKor: '편집실',
        requirement: [
            '해당 공간은 공용 공간으로, 공간 부족(워크샵 오디션 등)시에만 승인됩니다.',
            '개인 물품 보관 금지(외장하드, 스크립북 포함)',
            '물품 보관X 음식X 음료O *단, 본인 이름 기재 필수',
            '책상, 의자, 장비, 라인 이동 및 임의 변경 금지'
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
