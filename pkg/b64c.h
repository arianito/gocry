#ifndef B64C_H
#define B64C_H

#include <time.h>
#include <stdlib.h>

char index_of(const char* cipher, const char c);
int shuffle(char *a, int len);
int base64_encode(const char* ms, int len, const char* cf, char* op);
int base64_decode(const char* ms, int len, const char* cf, char* op);

#endif