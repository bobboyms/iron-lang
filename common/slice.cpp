export module Example;

#include <iostream>
#include <vector>

template <typename TT>
std::vector<TT> vector_slicing(std::vector<TT> &arr, int X, int Y)
{
    auto start = arr.begin() + X;
    auto end = arr.begin() + Y + 1;

    std::vector<TT> result(Y - X + 1);
    copy(start, end, result.begin());
    return result;
}
