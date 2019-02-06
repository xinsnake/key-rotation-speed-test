package main

import (
	"fmt"
	"testing"
)

var accountKeys = [...]string{
	"44989a20a35ef5167ebf9a92e047db55",
	"c352a202fa1aa99f76d35d272eb752c0",
	"540e7a1b59c9f0dbab018e045ba49685",
	"de10192fdcadc8a1448472cd37b633a0",
	"5a1578bd15d056a30dfe8dbc75ddf6ba",
	"2fdd589fb40fcb7ca89756e1a0ebdea1",
	"74791989c84ff3f46553ff5a715f7076",
	"ba41d70b6be2734c5e069374752496af",
	"f64e9c8b738c9d5902a7df46a62bc9c0",
	"b211cd9b403bea19b747bc108b42a04d",
	"ef54a5317b480268ee986bd829b2b17a",
	"596be46f66d5ff9324cce5e77f26c77c",
	"752a34c51d2fce762bc8a716c26a6607",
	"9da1a06f73703c8bf7ece1ed02f7694e",
	"452c5e335a290f6ce180f10e0c6632e6",
	"fe97788b887e3e9b9e567c7a42376487",
	"4d6fd632a819f225554edae95a4d3749",
	"3a9e9faa4f13f0434b358942533e6074",
	"7220672a40d651a5732294fbb4210d7b",
	"ba42f447585f0baff797240b8a18911c",
	"03458a497de0c6dea294dfcd4520f89e",
	"18be8237d82c63664181885bb2cb5864",
}

func BenchmarkFindAccount(b *testing.B) {
	var activeKey = "988d441fa6c532c583102360a7629691"
	for n := 1; n <= b.N; n++ {
		accountKey := accountKeys[n%len(accountKeys)]

		for m := 0; m <= 10000; m++ {
			secretKey := activeKey
			if m != 10000 {
				secretKey = secretKey + fmt.Sprintf("%d", m)
			}
			findAccount(secretKey, accountKey)
		}
	}
}
