#pragma once

#include <string>

// STRATEGY: interchangeable algorithms

// ----------------------------------
// Interface & politics
// ----------------------------------
class Compositor {
public:
    virtual ~Compositor(){}
    virtual int compose( const std::string & _text ) = 0;
};

class SimpleCompositor : public Compositor {
public:
    virtual ~SimpleCompositor(){}
    virtual int compose( const std::string & _text ) override {
        // ...
    }
};

class TeXCompositor : public Compositor{
public:
    virtual ~TeXCompositor(){}
    virtual int compose( const std::string & _text ) override {
        // ...
    }
};

class ArrayCompositor : public Compositor{
public:
    virtual ~ArrayCompositor(){}
    virtual int compose( const std::string & _text ) override {
        // ...
    }
};

// ----------------------------------
// Context
// ----------------------------------
class Composition {
public:
    Composition( Compositor * _compositor ) : m_compositor(_compositor) {}
    ~Composition();

    void run( const std::string & _text ){
        // ...
        int breaks = m_compositor->compose( _text );
        // ...
    }

private:
    Compositor * m_compositor;
};
