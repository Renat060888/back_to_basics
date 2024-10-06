#pragma once

// std
#include <vector>
#include <iostream>
#include <typeinfo>
#include <stdint.h>
#include <random>
#include <limits.h>

template< typename T >
class SortMachine{
public:

    SortMachine() : m_errorStr("no_errors") {}

    bool initWithRandom( std::vector<T> & _array, const uint32_t _requiredSize, const T _from = 0, const T _to = 0){

        const T to = (0 == _to) ? std::numeric_limits<T>::max() : _to;

        // will be used to obtain a seed for the random number engine
        std::random_device rd;
        // standard mersenne_twister_engine seeded with rd()
        std::mt19937 gen( rd() );

        _array.clear();
        _array.resize( _requiredSize );

        // genarate numbers for int, uint, float, double
        if( typeid(T) == typeid(int32_t) ){
            std::uniform_int_distribution<int32_t> dis( _from, to );

            for( uint64_t i = 0; i < _requiredSize; i++ ){
                _array[ i ] = dis( gen ); // functor
            }
        }
        else if( typeid(T) == typeid(uint32_t) ){
            std::uniform_int_distribution<uint32_t> dis( _from, to );

            for( uint64_t i = 0; i < _requiredSize; i++ ){
                _array[ i ] = dis( gen ); // functor
            }
        }
        else if( typeid(T) == typeid(float) ){
            std::uniform_real_distribution<float> dis( _from, to );

            for( uint64_t i = 0; i < _requiredSize; i++ ){
                _array[ i ] = dis( gen ); // functor
            }
        }
        else if( typeid(T) == typeid(double) ){
            std::uniform_real_distribution<double> dis( _from, to );

            for( uint64_t i = 0; i < _requiredSize; i++ ){
                _array[ i ] = dis( gen ); // functor
            }
        }
        else{
            m_errorStr = "type of array not INT32, UINT32, FLOAT or DOUBLE";
            _array.shrink_to_fit();
            return false;
        }

        return true;
    }

    bool checkSorted( const std::vector<T> & _array, const bool _ascendingOrder = true ){

        const uint32_t lastElementIdx = _array.size() - 1;

        if( _ascendingOrder ){
            for( uint32_t i = 0; i < lastElementIdx; i++ ){
                if( _array[ i ] > _array[ i+1 ] ){
                    return false;
                }
            }
        }
        else{
            for( uint32_t i = 0; i < lastElementIdx; i++ ){
                if( _array[ i ] < _array[ i+1 ] ){
                    return false;
                }
            }
        }

        return true;
    }

    void printArray( const std::vector<T> & _array ){

        std::cout << "-----------------------------" << std::endl;
        for( const T & element : _array ){
            std::cout << element << " ";
        }
        std::cout << std::endl;
        std::cout << "-----------------------------" << std::endl;
    }

    std::string getErrorStr(){ return m_errorStr; }

protected:

    void swap( std::vector<T> & _array, const uint32_t _element1Idx, const uint32_t _element2Idx ){

        const T temp = _array[ _element2Idx ];
        _array[ _element2Idx ] = _array[ _element1Idx ];
        _array[ _element1Idx ] = temp;
    }

    std::string m_errorStr;
};
