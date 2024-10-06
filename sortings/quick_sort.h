#pragma once

// project
#include "sort_machine.h"

// -------------------------
//
// -------------------------

template< typename T >
class QuickSort : public SortMachine<T>{

    using SortMachine<T>::swap;

public:

    QuickSort() {}
    ~QuickSort() {}

    void sort( std::vector<T> & _array ){

        const int32_t rightBound = _array.size() - 1;
        const int32_t leftBound = 0;

        recursiveSlicing( _array, leftBound, rightBound );
    }

protected:


private:

    // если в сортировке слиянием массив сначала разбивается потом обрабатывается (при раскрутке), то здесь наоборот
    void recursiveSlicing( std::vector<T> & _array, const int32_t _leftBound, const int32_t _rightBound ){

        if( (_rightBound - _leftBound) <= 0 ){
            return;
        }
        else{
            const T pivot = _array[ _rightBound ]; // можно и случайный

            const int32_t partitionIndex = partition( _array, _leftBound, _rightBound, pivot );

            // опорное значение не трогается, т.к. относительно него все сортировалось и оно на своем месте
            recursiveSlicing( _array, _leftBound, partitionIndex - 1);
            recursiveSlicing( _array, partitionIndex + 1, _rightBound );
        }
    }

    int32_t partition( std::vector<T> & _array, const int32_t _leftBound, const int32_t _rightBound, const T _pivot ){

        int32_t leftIdx = _leftBound - 1; // after ++
        int32_t rightIdx = _rightBound; // after --

        while( true ){

            while( _array[ ++leftIdx ] < _pivot ){
                // dummy
            }
            while( rightIdx > 0 && _array[ --rightIdx ] > _pivot ){
                // dummy
            }
            if( leftIdx >= rightIdx ){
                break;
            }
            swap( _array, leftIdx, rightIdx );
        }

        swap( _array, leftIdx, _rightBound ); // опорное значение (взятое справа) переносится в начало правого массива

        return leftIdx;
    }
};






