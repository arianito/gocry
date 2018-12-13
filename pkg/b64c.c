
#include "b64c.h"


int set_cipher(const char* c){
    for (int i = 0; i < 65; i++){
        cf[i] = c[i];
        dec[(unsigned char) c[i]] = i;
    }
    return 1;
}

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


int base64_encode(const char* ms, int len, char* op, int olen){
	int f = 0, o = 0;
	unsigned char a, b, c, pad = cf[64];
	unsigned int d = 0;
    for (int i = 0, j = 0; i < len;) {

        a = i < len ? ms[i++] : 0;
        b = i < len ? ms[i++] : 0;
        c = i < len ? ms[i++] : 0;

        d = (a << 0x10) + (b << 0x08) + c;

        op[j++] = cf[(d >> 3 * 6) & 0x3F];
        op[j++] = cf[(d >> 2 * 6) & 0x3F];
        op[j++] = cf[(d >> 1 * 6) & 0x3F];
        op[j++] = cf[(d >> 0 * 6) & 0x3F];
    }
    for (int i = 0; i < mod[len % 3]; i++)
        op[olen - 1 - i] = pad;
	return 1;
}

int base64_decode(const char* ms, int len, char* op, int olen){
	unsigned char pad = cf[64];
	if (ms[len-1] == pad)
		olen--;
	if (ms[len-2] == pad)
		olen--;
    if (len % 4 != 0) return 0;
	unsigned int a, b, c, d, e;
    for (int i = 0, j = 0; i < len;) {

        a = ms[i] == pad ? 0 & i++ : dec[ms[i++]];
        b = ms[i] == pad ? 0 & i++ : dec[ms[i++]];
        c = ms[i] == pad ? 0 & i++ : dec[ms[i++]];
        d = ms[i] == pad ? 0 & i++ : dec[ms[i++]];

        e = (a << 3 * 6)
        + (b << 2 * 6)
        + (c << 1 * 6)
        + (d << 0 * 6);

        if (j < olen) op[j++] = (e >> 2 * 8) & 0xFF;
        if (j < olen) op[j++] = (e >> 1 * 8) & 0xFF;
        if (j < olen) op[j++] = (e >> 0 * 8) & 0xFF;
    }

    return 1;
}