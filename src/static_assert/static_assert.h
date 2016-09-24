/*****************************************************************************/
#define CATSTR2(pre,post) pre##post
#define CATSTR(pre,post) CATSTR2(pre,post)

/*****************************************************************************/
#define STATIC_ASSERT_A(COND,MSG) typedef char \
    CATSTR(CATSTR(static_assert_line_,__LINE__),CATSTR(__,MSG))[2*(! !(COND))-1]

/*****************************************************************************/
#define STATIC_ASSERT(cond,msg) \
    typedef struct { int CATSTR(CATSTR(static_assert_line_,__LINE__), CATSTR(__,msg)) : ! !(cond); } CATSTR(static_assert_failed_line_,__LINE__)

#define STATIC_ASSERT_B(cond,msg) \
    struct CATSTR(static_assert_failed_line_,__LINE__) { int CATSTR(CATSTR(static_assert_line_,__LINE__), CATSTR(__,msg)) : ! !(cond); } 

