#pragma once

// project
#include "sort_machine.h"

// -------------------------
// complexity - O(n*log(n))
// -------------------------

template< typename T >
class MergeSort : public SortMachine<T>{

    using SortMachine<T>::swap;

public:

    MergeSort() {}
    ~MergeSort() {}

    void sort( std::vector<T> & _input ){

        m_temp.resize( _input.size() ); // недостаток данной сортировки

        recursiveSlicing( _input, 0, _input.size() - 1 );

        m_temp.clear();
    }

    void mergeTest( const std::vector<T> & _A, const std::vector<T> & _B, std::vector<T> & _C ){

        const uint32_t arrASize = _A.size();
        const uint32_t arrBSize = _B.size();
        uint32_t arrAIdx = 0;
        uint32_t arrBIdx = 0;
        uint32_t arrCIdx = 0;
        _C.resize( arrASize + arrBSize );

        while( arrAIdx < arrASize && arrBIdx < arrBSize ){

            if( _A[ arrAIdx ] < _B[ arrBIdx ] ){
                _C[ arrCIdx++ ] = _A[ arrAIdx++ ];
            }
            else{
                _C[ arrCIdx++ ] = _B[ arrBIdx++ ];
            }
        }
        while( arrAIdx < arrASize ){
            _C[ arrCIdx++ ] = _A[ arrAIdx++ ];
        }
        while( arrBIdx < arrBSize ){
            _C[ arrCIdx++ ] = _B[ arrBIdx++ ];
        }
    }

protected:

private:

    void recursiveSlicing( std::vector<T> & _result, const uint32_t _lowerBound, const uint32_t _upperBound ){

        if( _lowerBound == _upperBound ){
            return; // массив из 1го элемента уже отсортирован
        }
        else{
            const uint32_t middle = ( _lowerBound + _upperBound ) / 2;

            recursiveSlicing( _result, _lowerBound, middle );
            recursiveSlicing( _result, middle + 1, _upperBound );

            merge( _result, _lowerBound, middle + 1, _upperBound );
        }
    }

    void merge( std::vector<T> & _result, const uint32_t _beingFirstHalf, const uint32_t _beginSecondHalf, const uint32_t _endSecondHalf ){

        // центральная идея этой сортировки - слияние 2х предварительно отсортированных массивов
        uint32_t leftIdx = _beingFirstHalf;
        uint32_t rightIdx = _beginSecondHalf;
        uint32_t resultIdx = 0;

        while( leftIdx < _beginSecondHalf && rightIdx < (_endSecondHalf+1) ){

            if( _result[ leftIdx ] < _result[ rightIdx ] ){
                m_temp[ resultIdx++ ] = _result[ leftIdx++ ];
            }
            else{
                m_temp[ resultIdx++ ] = _result[ rightIdx++ ];
            }
        }
        while( leftIdx < _beginSecondHalf ){
            m_temp[ resultIdx++ ] = _result[ leftIdx++ ];
        }
        while( rightIdx < (_endSecondHalf+1) ){
            m_temp[ resultIdx++ ] = _result[ rightIdx++ ];
        }

        for( uint32_t i = _beingFirstHalf, j = 0; i <= _endSecondHalf; i++, j++ ){

            _result[ i ] = m_temp[ j ]; // отсортированный блок нужен будет для предыдущей рекурсии при раскрутке
        }
    }

    std::vector<T> m_temp;
};
















