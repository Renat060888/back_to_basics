#pragma once

// PROTOTYPE: copy object without knowing which conrete object

class Car {
public:
    virtual ~Car(){}

    virtual Car * clone() = 0;
};

class Tesla : public Car {
public:
    virtual ~Tesla(){}

    virtual Car * clone() override {
        new Tesla(*this);
    }
};

class Ford : public Car {
public:
    virtual ~Ford(){}

    virtual Car * clone() override {
        new Ford(*this);
    }
};
