// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krpos

var (
	// NNG is 일반 명사
	NNG = POS{
		Pos:         "NNG",
		Category:    "체언",
		Description: "일반 명사",
	}

	// NNP is 고유 명사
	NNP = POS{
		Pos:         "NNP",
		Category:    "체언",
		Description: "고유 명사",
	}

	// NNB is 의존 명사
	NNB = POS{
		Pos:         "NNB",
		Category:    "체언",
		Description: "의존 명사",
	}

	// NP is 대명사
	NP = POS{
		Pos:         "NP",
		Category:    "체언",
		Description: "대명사",
	}

	// NR is 수사
	NR = POS{
		Pos:         "NR",
		Category:    "체언",
		Description: "수사",
	}

	// VV is 동사
	VV = POS{
		Pos:         "VV",
		Category:    "용언",
		Description: "동사",
	}

	// VA is 형용사
	VA = POS{
		Pos:         "VA",
		Category:    "용언",
		Description: "형용사",
	}

	// VX is 보조 용언
	VX = POS{
		Pos:         "VX",
		Category:    "용언",
		Description: "보조 용언",
	}

	// VCP is 금정 지정사
	VCP = POS{
		Pos:         "VCP",
		Category:    "용언",
		Description: "금정 지정사",
	}

	// VCN is 부정 지정사
	VCN = POS{
		Pos:         "VCN",
		Category:    "용언",
		Description: "부정 지정사",
	}

	// MM is 관형사
	MM = POS{
		Pos:         "MM",
		Category:    "수식언",
		Description: "관형사",
	}

	// MAG is 일반 부사
	MAG = POS{
		Pos:         "MAG",
		Category:    "수식언",
		Description: "일반 부사",
	}

	// MAJ is 접속 부사
	MAJ = POS{
		Pos:         "MAJ",
		Category:    "수식언",
		Description: "접속 부사",
	}

	// IC is 감탄산
	IC = POS{
		Pos:         "IC",
		Category:    "독립언",
		Description: "감탄산",
	}

	// JKS is 주격 조사
	JKS = POS{
		Pos:         "JKS",
		Category:    "관계언",
		Description: "주격 조사",
	}

	// JKC is 보격 조사
	JKC = POS{
		Pos:         "JKC",
		Category:    "관계언",
		Description: "보격 조사",
	}

	// JKG is 관현격 조사
	JKG = POS{
		Pos:         "JKG",
		Category:    "관계언",
		Description: "관현격 조사",
	}

	// JKO is 목적격 조사
	JKO = POS{
		Pos:         "JKO",
		Category:    "관계언",
		Description: "목적격 조사",
	}

	// JKB is 부사격 조사
	JKB = POS{
		Pos:         "JKB",
		Category:    "관계언",
		Description: "부사격 조사",
	}

	// JKV is 호격 조사
	JKV = POS{
		Pos:         "JKV",
		Category:    "관계언",
		Description: "호격 조사",
	}

	// JKQ is 인용격 조사
	JKQ = POS{
		Pos:         "JKQ",
		Category:    "관계언",
		Description: "인용격 조사",
	}

	// JX is 보조사
	JX = POS{
		Pos:         "JX",
		Category:    "관계언",
		Description: "보조사",
	}

	// JC is 접속 조사
	JC = POS{
		Pos:         "JC",
		Category:    "관계언",
		Description: "접속 조사",
	}

	// EP is 선어말 어미
	EP = POS{
		Pos:         "EP",
		Category:    "의존 형태",
		Description: "선어말 어미",
	}

	// EF is 종결 어미
	EF = POS{
		Pos:         "EF",
		Category:    "의존 형태",
		Description: "종결 어미",
	}

	// EC is 연결 어미
	EC = POS{
		Pos:         "EC",
		Category:    "의존 형태",
		Description: "연결 어미",
	}

	// ETN is 명사형 전성 어미
	ETN = POS{
		Pos:         "ETN",
		Category:    "의존 형태",
		Description: "명사형 전성 어미",
	}

	// ETM is 관형형 전성 어미
	ETM = POS{
		Pos:         "ETM",
		Category:    "의존 형태",
		Description: "관형형 전성 어미",
	}

	// XPN is 체언 접두사
	XPN = POS{
		Pos:         "XPN",
		Category:    "의존 형태",
		Description: "체언 접두사",
	}

	// XSN is 명사 파생 접미사
	XSN = POS{
		Pos:         "XSN",
		Category:    "의존 형태",
		Description: "명사 파생 접미사",
	}

	// XSV is 동사 파생 접미사
	XSV = POS{
		Pos:         "XSV",
		Category:    "의존 형태",
		Description: "동사 파생 접미사",
	}

	// XSA is 형용사 파생 접미사
	XSA = POS{
		Pos:         "XSA",
		Category:    "의존 형태",
		Description: "형용사 파생 접미사",
	}

	// XR is 어근
	XR = POS{
		Pos:         "XR",
		Category:    "의존 형태",
		Description: "어근",
	}

	// SF is 마침표, 물음표, 느낌표
	SF = POS{
		Pos:         "SF",
		Category:    "기호",
		Description: "마침표, 물음표, 느낌표",
	}

	// SP is 쉼표, 가운뎃점, 콜론, 빗금
	SP = POS{
		Pos:         "SP",
		Category:    "기호",
		Description: "쉼표, 가운뎃점, 콜론, 빗금",
	}

	// SS is 따옴표, 괄호표, 줄표
	SS = POS{
		Pos:         "SS",
		Category:    "기호",
		Description: "따옴표, 괄호표, 줄표",
	}

	// SE is 줄임표
	SE = POS{
		Pos:         "SE",
		Category:    "기호",
		Description: "줄임표",
	}

	// SO is 붙임표(물결, 숨김, 빠짐)
	SO = POS{
		Pos:         "SO",
		Category:    "기호",
		Description: "붙임표(물결, 숨김, 빠짐)",
	}

	// SL is 외국어
	SL = POS{
		Pos:         "SL",
		Category:    "기호",
		Description: "외국어",
	}

	// SH is 한자
	SH = POS{
		Pos:         "SH",
		Category:    "기호",
		Description: "한자",
	}

	// SW is 기타 기호(논리, 수학 기호, 화폐 기호 등)
	SW = POS{
		Pos:         "SW",
		Category:    "기호",
		Description: "기타 기호(논리, 수학 기호, 화폐 기호 등)",
	}

	// SWK is 한글 자소
	SWK = POS{
		Pos:         "SWK",
		Category:    "기호",
		Description: "한글 자소",
	}

	// SN is 숫자
	SN = POS{
		Pos:         "SN",
		Category:    "기호",
		Description: "숫자",
	}

	// ZN is 분석 불능(명사 추정)
	ZN = POS{
		Pos:         "ZN",
		Category:    "추정",
		Description: "분석 불능(명사 추정)",
	}

	// ZV is 분석 불능(용언 추정)
	ZV = POS{
		Pos:         "ZV",
		Category:    "추정",
		Description: "분석 불능(용언 추정)",
	}

	// ZZ is 분석 불능(기타)
	ZZ = POS{
		Pos:         "ZZ",
		Category:    "추정",
		Description: "분석 불능(기타)",
	}
)
