#pragma once

// project
#include "sort_machine.h"

// -------------------------
// complexity - O(n^2)
// -------------------------

template< typename T >
class SelectionSort : public SortMachine<T>{

    using SortMachine<T>::swap;

public:

    SelectionSort() {}
    ~SelectionSort() {}

    void sort( std::vector<T> & _array ){

        const uint32_t size = _array.size() - 1;

        for( uint32_t leftSortedEdge = 0; leftSortedEdge < size; leftSortedEdge++ ){ // сужение левой отсортированной границы

            uint32_t currentMinIdx = leftSortedEdge;
            for( uint32_t cursor = leftSortedEdge; cursor <= size; cursor++ ){ // выбор минимального от левой границы до конца

                if( _array[ cursor ] < _array[ currentMinIdx ] ){
                    currentMinIdx = cursor;
                }
            }

            swap( _array, leftSortedEdge, currentMinIdx );
        }
    }

protected:


private:


};

