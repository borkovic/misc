
#include "static_assert.h"

enum { X = 3, Y = 4};

#define Xval 3
#define Xcond X_must_be_3

#define Yval 5
#define Ycond Y_must_be_5

/* shortest message in gcc */
STATIC_ASSERT_T(X==Xval, Xcond);
STATIC_ASSERT_T(Y==Yval, Ycond);
/*
*/

STATIC_ASSERT_S(X==Xval, Xcond);
STATIC_ASSERT_S(Y==Yval, Ycond);
/*
*/

/* shortest message in clang */
STATIC_ASSERT_A(X==Xval, Xcond);
STATIC_ASSERT_A(Y==Yval, Ycond);
/*
*/


