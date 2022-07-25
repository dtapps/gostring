package gostring

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	str := "iPhone 11 Pro Max<iPhone12,5>"
	fmt.Printf("%d\n", strings.LastIndex(str, "<"))
	fmt.Printf("%d\n", strings.LastIndex(str, "("))
	fmt.Printf("%d\n", strings.LastIndex("iPad (6th generation, WiFi)<iPad7,5>", "<"))
	fmt.Printf("%d\n", strings.LastIndex("iPad (6th generation, WiFi)<iPad7,5>", "("))
}

func TestToInt64(t *testing.T) {
	t.Log(ToInt64("120"))
	t.Log(ToInt64("120.9"))
	t.Log(strings.Contains("120", ","))
	t.Log(strings.Contains("120,1", ","))
}

func TestString(t *testing.T) {
	str := "wx6566ef69e8738ad9"
	fmt.Println(strings.Contains(str, "wx"))
	myString := "www.5lmh.com"
	if strings.HasPrefix(myString, "www") {
		fmt.Println("Hello to you too")
	} else {
		fmt.Println("Goodbye")
	}
}

func TestSplit(t *testing.T) {
	t.Log(Split("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20", ","))
	t.Log(Split(",10,", ","))
	t.Log(len(Split(",10,", ",")))
	t.Log(Split(",10,", ",")[1 : len(Split(",1,", ","))-1])
	t.Log(len(Split(",10,", ",")[1 : len(Split(",1,", ","))-1]))
	t.Log(Contains("", ","))
	t.Log(len([]string{}))
	t.Log(Split("/pages/preferential_recharge/goods_details?goods_id=G02022062517457120", "="))
}

func TestContains(t *testing.T) {
	t.Log(Contains("1", ","))
	t.Log(Contains("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20", ","))
}

func TestToFloat64(t *testing.T) {
	t.Log(ToFloat64("120"))
	t.Log(ToFloat64("120.9"))
	t.Log(ToFloat64("100.100.100"))
	if "100.100.100" > "111.111.111" {
		log.Println("1")
	} else {
		log.Println("2")
	}
}

func TestReplace(t *testing.T) {
	t.Log(Replace("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20", ",", "|"))
	t.Log(Replace("/v1/a/{agent_user_id}/d/g", "{agent_user_id}", "A102DFB78FADE96F1E"))
}
