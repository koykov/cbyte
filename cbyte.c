#include <stdlib.h>
#include <string.h>
#include "cbyte.h"

uint64 cbyte_init(int cap) {
    return (uint64) malloc(cap);
}

uint64 cbyte_init_set(int cap) {
    return (uint64) malloc(cap * sizeof(uint64));
}

uint64 cbyte_grow_m(uint64 addr, int cap_o, int cap_n) {
    uint64 addr_n = (uint64) malloc(cap_n);
    if (addr_n == 0) {
        return addr_n;
    }
    if (cap_n < cap_o) {
        cap_o = cap_n;
    }
    memcpy((void*)addr_n, (void*)addr, cap_o);
    free((void*)addr);
    return addr_n;
}

uint64 cbyte_grow_r(uint64 addr, int cap) {
    return (uint64) realloc((void*)addr, cap);
}

void cbyte_release(uint64 addr) {
    if ((void*)addr != NULL) {
        free((void*)addr);
    }
}
