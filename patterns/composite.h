#pragma once

#include <vector>

// COMPOSITE: object at the same time Simple & Composite

class Graphic{ // одновременно и примитив и контейнер
public:

    Graphic(){}
    virtual ~Graphic(){}

    virtual void draw(){}
    virtual void add( Graphic * _component ){}
    virtual void remove( Graphic * _component ){}
    virtual void getChild( const int _idx ){}

};

// ----------------------------------
// Simple object
// ----------------------------------
class Line : public Graphic{
public:

    Line(){}
    ~Line(){}

    virtual void draw() override {
        // system call
    }

};

// ----------------------------------
// Composite object
// ----------------------------------
class Picture : public Graphic{
public:

    Picture(){}
    ~Picture(){}

    virtual void draw() override {
        for( Graphic * c : m_component ){
            c->draw();
        }
    }
    virtual void add( Graphic * _component ) override {
        m_component.push_back( _component );
    }
    virtual void remove( Graphic * _component ) override {
        // ...
    }
private:
    std::vector<Graphic *> m_component;
};

