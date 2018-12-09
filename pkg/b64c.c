
#include "b64c.h"


int shuffle(char *a, int len) {
	if (len <= 1)
		return 0;
	srand(time(NULL));
	char j, i, tmp;
	for(i = len - 1; i > 0; i--) {
		j = rand() % (i + 1);
	    tmp = a[i];
		a[i] = a[j];
		a[j] = tmp;
	}
	return 1;
}

int base64_encode(const char* ms, int len, const char* cf, char* op){
	int f = 0, o = 0;
	char n, r, i, b1, b2, pad = *(cf+64);
	while (f < len){
		b1=0;
		b2=0;
		n = ms[f++];
        if (f < len) {
            r = ms[f++]; b1 = 1;
        }
        if (f < len) {
            i = ms[f++]; b2 = 1;
        }
        op[o++] = cf[n>>2];
        op[o++] = cf[((n&3)<<4)|(r>>4)];
        if (!b1) {
            op[o++] = pad;
            op[o++] = pad;
        } else if (!b2) {
            op[o++] = cf[((r&15)<<2)|(i>>6)];
            op[o++] = pad;
        } else {
            op[o++] = cf[((r&15)<<2)|(i>>6)];
            op[o++] = cf[i&63];
        }
	}
	return 1;
}


char index_of(const char* cipher, const char c){
	char i = 0;
	while(i<65)
		if(*(cipher+i++) == c)
			return i-1;
	return 64;
}


int base64_decode(const char* ms, int len, const char* cf, char* op){
	int f = 0, j = 0;
	char s, o, u, a;
	if(len % 4 != 0) return 0;
	while(f < len){
		s = index_of(cf, ms[f++]);
		o = index_of(cf, ms[f++]);
		u = index_of(cf, ms[f++]);
		a = index_of(cf, ms[f++]);
		op[j++] = (s<<2)|(o>>4);
		if(u != 64)
			op[j++] = ((o&15)<<4)|(u>>2);
		if(a != 64)
			op[j++] = ((u&3)<<6)|a;
	}
	return 1;
}