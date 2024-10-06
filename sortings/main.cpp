
// std
#include <iostream>
#include <cassert>
// projection
#include "bubble_sort.h"
#include "selection_sort.h"
#include "insertion_sort.h"
#include "merge_sort.h"
#include "quick_sort.h"

using namespace std;

int main( int argc, char *argv[] ){

    // init
    vector<int> array;
    const uint32_t arraySize = 20;
    const uint32_t leftRange = 0;
    const uint32_t rightRange = 100;

    cout << "// ------------------------" << endl;
    cout << "// quick sort" << endl;
    cout << "// ------------------------" << endl;
    QuickSort<int> qs;

    if( qs.initWithRandom( array, arraySize, leftRange, rightRange ) ){
        qs.printArray( array );

        qs.sort( array );
        assert( qs.checkSorted( array ) );

        qs.printArray( array );
    }

    cout << "// ------------------------" << endl;
    cout << "// merge sort (merge test)" << endl;
    cout << "// ------------------------" << endl;
    MergeSort<int> ms;

    vector<int> arrA = { 23, 47, 81, 95 };
    vector<int> arrB = { 7, 14, 39, 55, 62, 74 };
    vector<int> arrC;
    ms.printArray( arrA );
    ms.printArray( arrB );

    ms.mergeTest( arrA, arrB, arrC );
    assert( ms.checkSorted( arrC ) );

    ms.printArray( arrC );

    cout << "// ------------------------" << endl;
    cout << "// merge sort (sort test)" << endl;
    cout << "// ------------------------" << endl;
    if( ms.initWithRandom( array, arraySize, leftRange, rightRange ) ){
        ms.printArray( array );

        ms.sort( array );
        assert( ms.checkSorted( array ) );

        ms.printArray( array );
    }

    cout << "// ------------------------" << endl;
    cout << "// insertion sort" << endl;
    cout << "// ------------------------" << endl;
    InsertionSort<int> is;

    if( is.initWithRandom( array, arraySize, leftRange, rightRange ) ){
        is.printArray( array );

        is.sort( array );
        assert( is.checkSorted( array ) );

        is.printArray( array );
    }

    cout << "// ------------------------" << endl;
    cout << "// selection sort" << endl;
    cout << "// ------------------------" << endl;
    SelectionSort<int> ss;

    if( ss.initWithRandom( array, arraySize, leftRange, rightRange ) ){
        ss.printArray( array );

        ss.sort( array );
        assert( is.checkSorted( array ) );

        ss.printArray( array );
    }

    cout << "// ------------------------" << endl;
    cout << "// bubble sort" << endl;
    cout << "// ------------------------" << endl;
    BubbleSort<int> bs;

    if( bs.initWithRandom( array, arraySize, leftRange, rightRange ) ){
        bs.printArray( array );

        bs.sort( array );
        assert( is.checkSorted( array ) );

        bs.printArray( array );
    }

    return 0;
}
