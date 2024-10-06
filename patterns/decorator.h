#pragma once

#include <iostream>

// DECORATOR: dynamic functional extend ( without subclass )

// ----------------------------------
// Component
// ----------------------------------
class VisualComponent{
public:

    VisualComponent() {}
    virtual ~VisualComponent() {}

    virtual void draw() = 0;
    virtual void resize() = 0;

private:

};

class TextWidget : public VisualComponent{
public:

    TextWidget(){}
    virtual ~TextWidget(){}

    virtual void draw() override {
        std::cout << "text widget draw" << std::endl;
    }
    virtual void resize() override {
        std::cout << "text widget resize" << std::endl;
    }

private:

};

// ----------------------------------
// Decorator
// ----------------------------------
class Decorator : public VisualComponent {
public:

    Decorator( VisualComponent * _visualComp ) : m_visualComp(_visualComp) {}
    virtual ~Decorator(){}

    virtual void draw() override {
        m_visualComp->draw();
    }
    virtual void resize() override {
        m_visualComp->resize();
    }

private:

    VisualComponent * m_visualComp; // является компонентом только на 1м уровне, дальше это декораторы
};

class BorderDecorator : public Decorator {
public:
    BorderDecorator( VisualComponent * _visualComp ) : Decorator(_visualComp) {}
    ~BorderDecorator(){}

    virtual void draw() override {
        Decorator::draw(); // под этим методом может скрываться как Декоратор, так и Компонент (за счет наследования декоратора и компонента от абстр. класса)
        drawBorder();
    }

private:

    void drawBorder(){
        std::cout << "border decorator draw" << std::endl;
    }
};

class ScrollDecorator : public Decorator {
public:
    ScrollDecorator( VisualComponent * _visualComp ) : Decorator(_visualComp) {}
    ~ScrollDecorator(){}

    virtual void draw() override {
        Decorator::draw();
        drawScroll();
    }

private:

    void drawScroll(){
        std::cout << "scroll decorator draw" << std::endl;
    }
};








