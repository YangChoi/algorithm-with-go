/**
문장을 거꾸로 정렬한다.
단 스페이스 바 전 문장끼리만 역으로 정렬한다.
*/

func reverseWords(s string) string {
	// byte slice는 string slice에 비해 변형, 갱신이 가능하다
	// string은 새로 생성하는 수 밖에 없다
	// 그렇게 되면 가비지 컬렉터에 많은 부하가 추가되기 때문에
	// byte library 사용
	var res bytes.Buffer
	// byte 슬라이스
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			// space를 발견하면
			res.WriteByte(s[i])
			// 버퍼 끝에 s[i]를 byte로 쓴다
			i++
			continue
		}
		j := i
		for j < len(s) && s[j] != ' ' {
			// space가 아닌 경우(문자인 경우)
			j++
		}
		for k := j - 1; k >= i; k-- {
			res.WriteByte(s[k])
			// 거꾸로 가며 byte write
		}
		i = j
	}
	return res.String()
	// string으로 바꿔준후 반환
}