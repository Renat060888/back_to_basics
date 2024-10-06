#pragma once

// MEMENTO: object state beyound from his scope

class Object;

// ----------------------------------
// State
// ----------------------------------
class ObjectState {
public:
    virtual ~ObjectState(){} // остальное извне это состояние могут ТОЛЬКО удалить

private:
    friend class Object; // только владелец состояния может ее менять
    ObjectState(){}

    int m_state;
};

// ----------------------------------
// Object
// ----------------------------------
class Object {
public:

    ObjectState * createSnapshot(){

        ObjectState * snapshot = new ObjectState();
        snapshot->m_state = 1;

        return snapshot;
    }

    void setSnapshot( ObjectState * _snapshot ){
        m_state = _snapshot->m_state;
    }

private:
    int m_state;
};
