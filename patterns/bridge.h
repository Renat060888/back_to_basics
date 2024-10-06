#pragma once

#include <iostream>

// BRIDGE: bridge between Abstraction & Implementation

// ----------------------------------
// Implementation
// ----------------------------------
class Implementor{
public:
    Implementor(){}
    virtual ~Implementor() {}

    virtual void drawRect() = 0;
    virtual void drawText() = 0;
    virtual void drawButton() = 0;
    virtual void destroy() = 0;
};

class ImplementorWindows : public Implementor{
public:
    ImplementorWindows() {}
    virtual ~ImplementorWindows() {}

    virtual void drawRect() override {
        std::cout << "draw rect in windows" << std::endl;
    }
    virtual void drawText() override {
        std::cout << "draw text in windows" << std::endl;
    }
    virtual void drawButton() override {
        std::cout << "draw button in windows" << std::endl;
    }
    virtual void destroy() override {
        std::cout << "destroy in windows" << std::endl;
    }
};

class ImplementorLinux : public Implementor{
public:
    ImplementorLinux() {}
    virtual ~ImplementorLinux() {}

    virtual void drawRect() override {
        std::cout << "draw rect in linux" << std::endl;
    }
    virtual void drawText() override {
        std::cout << "draw text in linux" << std::endl;
    }
    virtual void drawButton() override {
        std::cout << "draw button in linux" << std::endl;
    }
    virtual void destroy() override {
        std::cout << "destroy in linux" << std::endl;
    }
};

// ----------------------------------
// Abstraction
// ----------------------------------
class AbsWindow{
public:
    AbsWindow( Implementor * _impl ) : m_impl(_impl){}
    virtual ~AbsWindow() {}

    virtual void open(){}

    virtual void close(){
        m_impl->destroy();
    }

protected:
    Implementor * m_impl;
private:
};

class ConreteAbsDialog : public AbsWindow{
public:
    ConreteAbsDialog( Implementor * _impl ) : AbsWindow(_impl) {}
    virtual ~ConreteAbsDialog() {}

    virtual void open() override {
        m_impl->drawRect();
        m_impl->drawButton();
    }
};

class ConreteAbsMessageBox : public AbsWindow{
public:
    ConreteAbsMessageBox( Implementor * _impl ) : AbsWindow(_impl) {}
    virtual ~ConreteAbsMessageBox() {}

    virtual void open() override {
        m_impl->drawRect();
        m_impl->drawText();
    }
};





