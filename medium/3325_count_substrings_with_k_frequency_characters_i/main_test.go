package main

import (
	"reflect"
	"testing"
)

func TestNumberOfSubstrings(t *testing.T) {
	type testCase struct {
		S              string
		K              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			S:              "abc",
			K:              1,
			ExpectedResult: 6,
		},
		{
			S:              "abcd",
			K:              1,
			ExpectedResult: 10,
		},
		{
			S:              "abacb",
			K:              2,
			ExpectedResult: 4,
		},
		{
			S:              "abcde",
			K:              1,
			ExpectedResult: 15,
		},
		{
			S:              "hxccgfp",
			K:              1,
			ExpectedResult: 28,
		},
		{
			S:              "sykokhxgqerlhkmxrfgllmlwobiqhwbgjqhewprufmdalbxmlflsuxyzilbzeaozkkhqkdqmlltyyxtlbnudvvlnybfkrjctuxkpdcptfdffzzfdnhbfutfyibslngvatomwyavoceletpogvoxtxkiqyyvbepjnccmwafpeqsuammrvcuonerbztlrzhzfreefgkbdncgahpfnysfdfdqyrwkveilxeupfvgweqfyuvkqdgvasrhnmjvfytbycxspoawmoakrcfdchkezopikwgmsmztfexqgmdeireilxlrrsaizxaiewsipbyhasfuhlirezejflceyovwmlqcoixyctklaamngbjttfsocobikqawbmvgbdgyzsnxrjgqrgacekhsuncfeupcuvazrqkdsimnowmnxrjoijmpddryjnlgmgaldocrkroruesiakhkyalnwbdoesfiguvyqrnwvtzxrrrmhwumgebwuxkytui",
			K:              7,
			ExpectedResult: 92778,
		},
		{
			S:              "edfcqzcumqzgqeozaksvqnhuscetlquoqtiucdxfkudcgsswmhgqrtytojfirlvvtlhpwwmocszmpgnnvmxahmoqcmesugchgtgirladgmovajhouuydipmujuxvbuuodjsthhezpwbmltvdbievswyudrtkwfsqvnorjlpvhxpqqxuowneclqhgfdfbvgjgjyozcuuaeidkdvuczsdfehnzkhbadsltwnszuahhcggwgmnamaavgfyapkkswfyrfafwmawbzwmqikudqngfqfubgexwfqmlmcruenayoehfswltvfymchjbblmnqaqadevfbjkogasyjaiopnpakrpqopvukqg",
			K:              4,
			ExpectedResult: 53147,
		},
		{
			S:              "ghlihcggpzphokwwlvrtdhicjmqscafsvuqlxeftdpuoeqbznmjkphrkouijkujgcfabiwkdrqnozavgcdxebqvrkdwlvayqetpaccjkecoxfauctbufvtyfirnkyornuhfonsbrcgffjlwhklibbgjoenhoovvcprsrdghwmqifwpgfwafjamleufecivhfsjsfhtruzdwqrtbejpvrbwzfnnfkpmoymqbnybtugtnadscsfxrmpyvnqbvndlftubywjbatmlrsapjszwsvxumxyvcacfmhuchyciitmguszozojqvkstmnftnwiugxgbdtuccdiowlyultlckxuublhyhpyrpnbfsxjkwajixzzywqocapgtcrmfcktebivhihjkjlodqliobbsoyxgyyhfjwtmwdubvfmrbpecbcxxhvbnvwsdiabhgfocfeyddgovfieuxfcenejemvtobstfrrkpwmhfyysycdhuoifnxtfhrfeuqrulpoofczjozgvojslunsvpovidldebehgfaupufomygrrpvtwzigbnrdyphhrcruceveazebxdteoitdfuwbaveqxbhzeynqfqfvigntvnhfsfotpayrdhzsfinyfgsnhaucrragkzowmemopdudhwtshgexxglexeascyarkevek",
			K:              153,
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := numberOfSubstrings(cases[i].S, cases[i].K)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("s: %v, k: %v, expected: %v, got: %v", cases[i].S, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
