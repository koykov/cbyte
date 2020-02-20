#ifndef CBYTE_EXPORT_H
#define CBYTE_EXPORT_H

/**
 * @file
 * Memory manipulation functions.
 */

#ifdef __cplusplus
extern "C" {
#endif

/* Determine the wordsize from the preprocessor defines. */
#if defined __x86_64__ && !defined __ILP32__
typedef unsigned long int uint64;
#else
__extension__ typedef unsigned long long int uint64;
#endif

/**
 * Initialize byte array with given capacity.
 *
 * @param cap Capacity of the array.
 * @return uint64
 */
uint64 cbyte_init(int cap);

/**
 * Initialize big byte array.
 * @see cbyte_init()
 */
uint64 cbyte_init64(uint64 cap);

/**
 * Initialize addresses array with given capacity.
 *
 * @param cap Capacity of the array.
 * @return uint64
 */
uint64 cbyte_init_set(int cap);

/**
 * Change capacity of the array using malloc().
 *
 * This function allows to reduce array's capacity as well.
 * @see http://www.cplusplus.com/reference/cstdlib/malloc
 * @see http://www.cplusplus.com/reference/cstring/memcpy
 * @see http://www.cplusplus.com/reference/cstdlib/free
 * @param addr  Address of reallocated array.
 * @param cap_o Old capacity.
 * @param cap_n New capacity of the array. May be less than old capacity.
 * @return uint64 address of first item of array in virtual memory.
 */
uint64 cbyte_grow_m(uint64 addr, int cap_o, int cap_n);

/**
 * Change capacity of the big array using malloc().
 * @see cbyte_grow_m()
 */
uint64 cbyte_grow64_m(uint64 addr, uint64 cap_o, uint64 cap_n);

/**
 * Change capacity of the array using realloc().
 *
 * This function allows to reduce array's capacity as well.
 * @see http://www.cplusplus.com/reference/cstdlib/realloc
 * @param addr Address of reallocated array.
 * @param cap  New capacity of the array. May be less than old capacity.
 * @return uint64 address of first item of array in virtual memory.
 */
uint64 cbyte_grow_r(uint64 addr, int cap);

/**
 * Change capacity of the big array using realloc().
 * @see cbyte_grow_r()
 */
uint64 cbyte_grow64_r(uint64 addr, uint64 cap);

/**
 * Release buffer memory.
 *
 * @param addr  Address of the array to release.
 */
void cbyte_release(uint64 addr);

#ifdef __cplusplus
}
#endif

#endif //CBYTE_EXPORT_H
