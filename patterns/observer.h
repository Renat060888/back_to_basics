#pragma once

#include <vector>

// OBSERVER:

// ----------------------------------
// Event
// ----------------------------------
class ISwitchObserver {
public:
    virtual ~ISwitchObserver(){}
    virtual void switchSignal( const bool _turnOn ) = 0;
};

class Switch {
public:

    void addObserver( ISwitchObserver * _observer ){
        m_observsers.push_back( _observer );
    }
    void press(){
        for( ISwitchObserver * observer : m_observsers ){
            observer->switchSignal( (m_currentState = ~m_currentState) );
        }
    }

private:
    std::vector< ISwitchObserver * > m_observsers;
    bool m_currentState;
};

// ----------------------------------
// Observer
// ----------------------------------
class Light : ISwitchObserver {
public:

    virtual void switchSignal( const bool _turnOn ) override {
        m_on = _turnOn;
    }
private:
    bool m_on;
};
