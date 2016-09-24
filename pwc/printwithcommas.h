#ifndef PRINT_WITH_COMMAS_H
#define PRINT_WITH_COMMAS_H

namespace Util {

//=================================================================
int PrintWithCommas(char* buf, const unsigned long unsLong);
int PrintWithCommas2(char* buf, const unsigned long unsLong);

int PrintWithCommasRec(char* buf, const unsigned long unsLong);

//=================================================================
inline int
PrintWithCommas(char* buf, const long signLong)
{
    if (signLong >= 0) {
        const unsigned long ul = signLong;
        return PrintWithCommas(buf, ul);
    }

    long signLong1 = signLong + 1; // to avoid smallest (negative long) so negation works for sure
    signLong1 = -signLong1;
    unsigned long ul = signLong1;
    ++ul;
    buf[0] = '-';
    return 1 + PrintWithCommas(buf+1, ul);
}



//=================================================================
//=================================================================
inline int
PrintWithCommas(char* buf, const int signInt)
{
    const long signLong = signInt;
    return PrintWithCommas(buf, signLong);
}

//=================================================================
inline int
PrintWithCommas(char* buf, const short signShort)
{
    const long signLong = signShort;
    return PrintWithCommas(buf, signLong);
}

//=================================================================
inline int
PrintWithCommas(char* buf, const signed char signChar)
{
    const long signLong = signChar;
    return PrintWithCommas(buf, signLong);
}


//=================================================================
//=================================================================
inline int
PrintWithCommas(char* buf, const unsigned int unsInt)
{
    const unsigned long unsLong = unsInt;
    return PrintWithCommas(buf, unsLong);
}

//=================================================================
inline int
PrintWithCommas(char* buf, const unsigned short unsShort)
{
    const unsigned long unsLong = unsShort;
    return PrintWithCommas(buf, unsLong);
}

//=================================================================
inline int
PrintWithCommas(char* buf, const unsigned char unsChar)
{
    const unsigned long unsLong = unsChar;
    return PrintWithCommas(buf, unsLong);
}

}

#endif

