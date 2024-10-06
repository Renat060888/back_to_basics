#pragma once

// FACTORY METHOD:

enum class WarriorType_en {
    Archer,
    Knight,

};

// ----------------------------------
// Objects
// ----------------------------------
class Warrior{
public:
    Warrior(){}
    virtual ~Warrior(){}
};

class Archer : public Warrior{
public:
    Archer(){}
    virtual ~Archer(){}
};

class Knight : public Warrior{
public:
    Knight(){}
    virtual ~Knight(){}
};

// ----------------------------------
// Factory
// ----------------------------------
Warrior * factory( const WarriorType_en _type ){

    switch( _type ){
    case WarriorType_en::Archer : {
        return new Archer();
    }
    case WarriorType_en::Knight : {
        return new Knight();
    }
    default: {

    }
    }
}
