#pragma once

// project
#include "sort_machine.h"

// -------------------------
// complexity - O(n^2)
// -------------------------

template< typename T >
class ShellSort : public SortMachine<T>{

    using ShellMachine<T>::swap;

public:

    ShellSort() {}
    ~ShellSort() {}

    void sort( std::vector<T> & _array ){

        const uint32_t size = _array.size();

        int inner, outer;
        long temp;
        int h = 1;

        while( h <= size/3 ){
            h = h*3 + 1;
        }

        while( h > 0 ){ // пока не закончатся все приращения
            for( outer = h; outer < size; outer++ ){ // пока текущее приращение не дойдет до конца массиива

                temp = _array[ 0 ]; // TODO
            }
        }
    }

protected:


private:


};

