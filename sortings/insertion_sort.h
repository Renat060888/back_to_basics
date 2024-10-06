#pragma once

// std
#include <cassert>
// project
#include "sort_machine.h"

// -------------------------
// complexity - O(n^2)
// -------------------------

template< typename T >
class InsertionSort : public SortMachine<T>{

    using SortMachine<T>::swap;

public:

    InsertionSort() {}
    ~InsertionSort() {}

    void sort( std::vector<T> & _array ){

        const uint32_t size = _array.size();
        // будет проверяться каждый элемент до конца
        for( uint32_t cursor = 1; cursor < size; cursor++ ){
            // в левую отсортированную часть будут вставляться элементы (1ый элемент считается отсортированным)
            for( uint32_t sortedPartIdx = 0; sortedPartIdx < cursor; sortedPartIdx++ ){

                if( _array[ cursor ] < _array[ sortedPartIdx ] ){

                    // 1
                    const T temp = _array[ cursor ]; // затрется сдвигом
                    // 2
                    shiftArray( _array, sortedPartIdx, cursor - sortedPartIdx );
                    // 3
                    _array[ sortedPartIdx ] = temp;
                    break;
                }
            }
        }
    }

protected:


private:

    void shiftArray( std::vector<T> & _array, const uint32_t _fromIdx, const uint32_t _count ){

        assert( (_fromIdx + _count) <= _array.size() );

        for( uint32_t i = _fromIdx + _count; i > _fromIdx; i-- ){
            _array[ i ] = _array[ i-1 ];
        }
    }
};

