
// std
#include <iostream>
// projection
#include "mediator.h"

using namespace std;

// ----------------------------------
// Mediator
// ----------------------------------
void DialogDirector::widgetChanged( Widget * _widget ){

    for( Widget * widget : m_widgets ){
        if( widget == _widget ){
            // TODO ?
        }
    }
}

// ----------------------------------
// Component
// ----------------------------------
void Widget::changed(){
    m_director->widgetChanged( this );
}

void Button::press(){
    changed();
}

void CheckBox::check( const bool _checked ){
    m_checked = _checked;
    changed();
}
