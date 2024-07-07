package common

import (
	"crypto"
	"crypto/rsa"
	"hash"
	"testing"
)

func TestSign(t *testing.T) {
	var privateKey, _ = ParsePrivateKey([]byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDokGGq7SlDoULN
PULY8lcb2uXJcrFKkJI/lSfPppIkGH4xPfQytZXRlonpXqgOvovflJT5VhRvoLe2
inJ/59kRF59KTerbCG5sG2IHhR/qCUGHervnZuPwgrjOOlnB19VCCUKY1tcplkZa
KIksUU3TVh09GB3lUngkuOeO15ihcFHMIknOiSpL+Q04+qQf0g++9CxdQUtNY5za
Z2Jdvch/4yFstR59qQu73ZCCYHFqXaVakyfOC3xOQkRB58jPOUvIab9zwo2hPukT
+6qkqfokqMhX979HhNshPAJEEUXp4szk0QtP+2n8hq8t3Dws+GY8ElAFvmeGHx5j
WzPYAcXvAgMBAAECggEBALIsu8caf/zCdc2MW8SelkJPCLG330DDVmjEO4YJlfl1
kmjjkE2xdSDn9q0GyjbRoZQf36rPWkTTmyyNEYAQ/urrcCybWY2J/h9xMz+TrIm/
oabMT13QJF5FqJTHe7DZTReUxKMYJixEZ433dHCxsbByT2BZM75X2pg32aBEaTl0
v9OfIMwfaJ5fSBmleJv4q/Lfd232/oOPzyr+EHfsMpTwOrgzQwPNoah1GvH+jBhz
goafi36vT8HVjJ+ZjOreH+Z2zVas683Le62rQaN/51jHS5vQGd2+z3qrI3kvu8KK
Wu0kIDQwCKtSFUT00MiKSaklE9JHf8rCNm3+en4kjfECgYEA/0Z2QaP4sOUt1+cZ
IOprsMkJOl2sLTTDx3MseD2BxUukLDT8P3HTWLtBN9AkGlL6XD8WSs1k9YlGCqf2
y9qC43Bgwsky8CH1ACk5K4PuWidGUQiqW4Oll/ris2vjagE+QfMHFFa7IalV5FvI
v9L01jMqSL4duoM5w/WlLyLu8s0CgYEA6TlppWuxUMbxhEL5jmCZSSvsyhmIOdRk
t4V9MuxwiRuItzywMo4+O7Hs6tjAxTnV/ROa33qyQtm4Olmd4Oa/TkmAJFX+mUrW
jUohDvm7Js2Y7/eeSRcRQLcgCjncpNe3AJoeVEvrGeaJMERYXnTwboUDKxsRyFFq
AyuZHfuh86sCgYAm3pjFF+2XKd5YIKUv4OHy8jmIfJjp7T3eUcg0qtDmtMTTwmGi
W3ed7C1bDUNiCr56a1S+oRW9WWCj4L1wft4tOYBSSIaMD++ZTa2Z1aXmblKDpjki
ZCJDyPzZ6xSeoH/VVOcADtDBqGIeumcP5lRHhVTr7J7kNnUGRJIZYk1WBQKBgQCl
LAIEI4cKnDrD3uL60LL+vVsPrpFp02AZETMf84+nqpZin1pyE4dDo7kUgbnUdCd2
+oF+sFi7O5Jb0MgdVY47FZbpJPYQ/o2AtvU+s+K1knozyPyS6wFPAeJxG5WGMTfr
9zpvnOy+BSU3x8+F5e+5df5OcvdfFTmtUR05vNJvzQKBgHUtziAeWo7H6vxknFcc
kVv7++a4IWF59eP+rpxlaHOtPTI43PLxJgSHEbw3epEzTUnCL9dpP8n48fuYuwM+
+vpAujDcaGjGffmxW40E6wuGjOYBNg1zjSfEyjxF2fY+D9WoICSPHnrWB0/BEAZB
aL9Lho8+BUEFergUMjxUdvAS
-----END PRIVATE KEY-----
`))

	type args struct {
		msg      []byte
		key      *rsa.PrivateKey
		hashType crypto.Hash
		hash     hash.Hash
	}
	tests := []struct {
		name     string
		args     args
		wantSign string
		wantErr  bool
	}{
		{
			name: "test1",
			args: args{
				msg: []byte(`800000010334001
017d6ae9ad6e
1720348216
Tn3G61IkxyHsyJfs0SFUtvSYYsJ9c3aT
{"req_time":"20210907150256","version":"3.0","out_org_code":"OP00000003","req_data":{"merchant_no":"822290070111135","term_no":"29034705","out_trade_no":"FD660E1FAA3A4470933CDEDAE1EC1D8E","auth_code":"135178236713755038","total_amount":"123","location_info":{"request_ip":"10.176.1.192","location":"+37.123456789,-121.123456789"},"out_order_no":"08F4542EEC6A4497BC419161747A92FA"}}
`),
				key:      privateKey,
				hashType: crypto.SHA256,
				hash:     crypto.SHA256.New(),
			},
			wantSign: "Mh7kp+YPtsJzH1n1fNg+xfD1auh2/iid/VpoXPhzzm65lghfDVaJVhFQLWu/hKWvglfm7K8mNPWjs0zjzuYKEY+LxvWEPr7fgGcrCjOhODmOsAZJWucyQsiqWAvE3C5FhpH9csQPaEY7kOXT183XQzR8383ZQiTlNTXgCTqmJTr+Y4ugzQr2yONYZ3nwi+mmrGcVSObb0YbzsE4wBR1v8Oei7VptUPtVti+F88+u2K5sXzjP1GtL2POUQgR2a73hT7AGaIVDDeThwjGYbYuBhQ9Ckzi+wctDkHTyHSb6UpliyauSN116mFYdY+6zCVwIMF6O7rj/tonX/BwFfMRCpw==",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSign, err := Sign(tt.args.msg, tt.args.key, tt.args.hashType, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSign != tt.wantSign {
				t.Errorf("Sign() gotSign = %v, want %v", gotSign, tt.wantSign)
			}
		})
	}
}
