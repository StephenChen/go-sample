// hello.cpp

#include <iostream>

extern "C"
{
    #include "hello.h"
}

void SayHello4(const char* s)
{
    std::cout << s;
}