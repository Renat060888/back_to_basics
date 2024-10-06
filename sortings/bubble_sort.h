#pragma once

// project
#include "sort_machine.h"

// -------------------------
// complexity - O(n^2)
// -------------------------

template< typename T >
class BubbleSort : public SortMachine<T>{

    using SortMachine<T>::swap;

public:

    BubbleSort() {}
    ~BubbleSort() {}

    void sort( std::vector<T> & _array ){

        const uint32_t size = _array.size() - 1;

        for( uint32_t rightSortedEdge = size; rightSortedEdge > 0; rightSortedEdge-- ){ // сужение границы справа
            for( uint32_t cursor = 0; cursor < rightSortedEdge; cursor++ ){ // текущий элемент до правой границы

                if( _array[ cursor ] > _array[ cursor+1 ] ){
                    swap( _array, cursor, cursor+1 );
                }
            }
        }
    }

protected:


private:


};

