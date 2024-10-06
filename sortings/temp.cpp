
#include <vector>
#include <stdint.h>

template< typename T >
void swap( std::vector<T> & _array, const uint32_t _element1Idx, const uint32_t _element2Idx ){

    const T temp = _array[ _element2Idx ];
    _array[ _element2Idx ] = _array[ _element1Idx ];
    _array[ _element1Idx ] = temp;
}

// ------------------------
// quick sort
// ------------------------


// ------------------------
// merge sort (merge test)
// ------------------------


// ------------------------
// merge sort (sort test)
// ------------------------


// ------------------------
// insertion sort
// ------------------------


// ------------------------
// selection sort
// ------------------------


// ------------------------
// bubble sort
// ------------------------
