#pragma once

#include <vector>

// MEDIATOR:

class Widget;
class Button;
class CheckBox;

// ----------------------------------
// Mediator
// ----------------------------------
class DialogDirector {
public:

    DialogDirector(){}
    ~DialogDirector(){}

    virtual void widgetChanged( Widget * _widget );
    virtual void addWidget( Widget * _widget ){
        m_widgets.push_back( _widget );
    }

private:
    std::vector<Widget *> m_widgets;
};

// ----------------------------------
// Component
// ----------------------------------
class Widget {
public:

    Widget( DialogDirector * _director ) : m_director(_director){}
    virtual ~Widget(){}

    virtual void changed();

private:
    DialogDirector * m_director;
};

class Button : public Widget {
public:

    Button( DialogDirector * _director ) : Widget(_director){}
    virtual ~Button(){}

    void press();

private:
};

class CheckBox : public Widget {
public:

    CheckBox( DialogDirector * _director ) : Widget(_director){}
    virtual ~CheckBox(){}

    void check( const bool _checked );

private:
    bool m_checked;
};

