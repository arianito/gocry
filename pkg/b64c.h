#ifndef B64C_H
#define B64C_H

#include <time.h>
#include <stdlib.h>


static char mod[] = {0, 2, 1};

static char cf[65];
static char dec[256];


int set_cipher(const char* c);

int shuffle(char *a, int len);
int base64_encode(const char* ms, int len, char* op, int olen);
int base64_decode(const char* ms, int len, char* op, int olen);

#endif