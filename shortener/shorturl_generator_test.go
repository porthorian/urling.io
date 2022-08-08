package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://cuck.com/givotto12437213753274578345734862398427365782359239423894234"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.vultrcuck.com/12348123572345732875sfjsdfn182437325trnjsdfnsdjfnsdfjsdf"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://urling.io/thisismycucktherearemanylikeitbutthisoneismine"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)


	assert.Equal(t, "FKrg71k9", shortLink_1)
	assert.Equal(t, "CnF4J6Q2", shortLink_2)
	assert.Equal(t, "VqrEfDHx", shortLink_3)
}