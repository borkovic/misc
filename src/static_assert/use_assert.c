
#include "static_assert.h"

enum { X = 3, Y = 4};

#define Xval 3
#define Xcond X_must_be_3

#define Yval 5
#define Ycond Y_must_be_5

/*
*/
STATIC_ASSERT(X==Xval, Xcond);
STATIC_ASSERT(Y==Yval, Ycond);

STATIC_ASSERT_A(X==Xval, Xcond);
STATIC_ASSERT_A(Y==Yval, Ycond);

STATIC_ASSERT_B(X==Xval, Xcond);
STATIC_ASSERT_B(Y==Yval, Ycond);

